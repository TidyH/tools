#!/usr/bin/env python
# Bootstrap an AWS account with some basics

# std library imports
import pip
import sys
import json

# foreign libs
# boto3
try:
    import boto3
except ImportError:
    pip.main(['install', boto3])
    import boto3

def create_admin_role_for_ec2():
    role_name = 'admin'
    iam = boto3.client('iam')

    # Define the trust policy for the role
    trust_policy = {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "ec2.amazonaws.com"  # EC2 service principal
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }

    # Create IAM role with the trust policy
    role_response = iam.create_role(
        RoleName=role_name,
        AssumeRolePolicyDocument=json.dumps(trust_policy)
    )

    # Attach AdministratorAccess policy to the role
    iam.attach_role_policy(
        RoleName=role_name,
        PolicyArn='arn:aws:iam::aws:policy/AdministratorAccess'
    )

    # Print the ARN of the created role
    print("ARN of the created role:", role_response['Role']['Arn'])
    return role_name

# Create/Delete iam instance profile
def create_admin_instance_profile(profile_name, role_name):
    iam_client = boto3.client('iam')

    # Create IAM instance profile
    response_create_profile = iam_client.create_instance_profile(
        InstanceProfileName=profile_name
    )

    # Get ARN of the instance profile
    instance_profile_arn = response_create_profile['InstanceProfile']['Arn']

    # Attach AdministratorAccess policy to the instance profile
    response_attach_policy = iam_client.add_role_to_instance_profile(
        InstanceProfileName=profile_name,
        RoleName=role_name  # Assuming 'AdministratorAccess' is the name of the policy
    )

    print("Admin IAM instance profile created successfully.")
    print("Instance Profile ARN:", instance_profile_arn)

def delete_instance_profile(profile_name):
    iam_client = boto3.client('iam')

    # Delete IAM instance profile
    iam_client.delete_instance_profile(InstanceProfileName=profile_name)

    print("Instance profile '{}' deleted successfully.".format(profile_name))

# create codecommit repo
def create_codecommit_repo(profile_name):
    codecommit_client = boto3.client('codecommit')

    # Create CodeCommit repository
    response = codecommit_client.create_repository(
        repositoryName=profile_name
    )

    print("CodeCommit repository '{}' created successfully.".format(profile_name))
    print("Repository ARN:", response['repositoryMetadata']['Arn'])

# create ec2 linked to codecommit with iam role profile
def create_ec2_instance(instance_profile_arn, ami_id, instance_type, security_group_ids, subnet_id):
    ec2_client = boto3.client('ec2')

    # Specify IAM instance profile
    iam_profile = {
        'Arn': instance_profile_arn
    }

    # Launch EC2 instance with specified IAM role
    response = ec2_client.run_instances(
        ImageId=ami_id,
        InstanceType=instance_type,
        SecurityGroupIds=security_group_ids,
        SubnetId=subnet_id,
        IamInstanceProfile=iam_profile,
        MaxCount=1,
        MinCount=1,
        UserData='#!/bin/bash\nyum update -y\n',  # Example UserData script
        TagSpecifications=[
            {
                'ResourceType': 'instance',
                'Tags': [
                    {
                        'Key': 'Name',
                        'Value': 'AdminCC'
                    },
                ]
            },
        ]
    )

    instance_id = response['Instances'][0]['InstanceId']
    print("EC2 instance '{}' created successfully.".format(instance_id))

# filter for latest ami
def get_latest_amazon_linux_2023_ami():
    # Create SSM client
    ssm_client = boto3.client('ssm')
    parameter_name = '/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-6.1-x86_64'

    try:
        # Get the parameter value
        response = ssm_client.get_parameter(
            Name=parameter_name,
            WithDecryption=True  # Set to True if the parameter value is encrypted
        )

        # Extract the parameter value from the response
        parameter_value = response['Parameter']['Value']

        return parameter_value

    except ssm_client.exceptions.ParameterNotFound:
        print(f"Parameter '{parameter_name}' not found.")
        return None

    except Exception as e:
        print(f"An error occurred: {str(e)}")
        return None

# filter for 'default' security group
def get_default_security_group_id():
    ec2_client = boto3.client('ec2')

    # Describe security groups
    response = ec2_client.describe_security_groups(
        Filters=[
            {'Name': 'group-name', 'Values': ['default']}
        ]
    )

    if response['SecurityGroups']:
        return [response['SecurityGroups'][0]['GroupId']]
    else:
        raise Exception("Could not find the default security group.")
    
# filter for VPC
def get_vpc_id_by_name(vpc_name):
    ec2_client = boto3.client('ec2')

    # Describe VPCs
    response = ec2_client.describe_vpcs(
        Filters=[
            {'Name': 'tag:Name', 'Values': [vpc_name]}
        ]
    )

    if response['Vpcs']:
        return response['Vpcs'][0]['VpcId']
    else:
        raise Exception("Could not find the VPC with the specified name.")

def get_subnet_id_by_name():
    ec2_client = boto3.client('ec2')

    vpc_id = get_vpc_id_by_name('aws-controltower-VPC')

    # Describe subnets in the specified VPC
    response = ec2_client.describe_subnets(
        Filters=[
            {'Name': 'vpc-id', 'Values': [vpc_id]}
        ]
    )

    if response['Subnets']:
        return response['Subnets'][0]['SubnetId']
    else:
        raise Exception("Could not find any subnets in the specified VPC.")

if __name__ == "__main__":
    boto3.setup_default_session(region_name='us-east-1')
    if len(sys.argv) < 3:
        print("""Usage: python script_name.py <function_name> <profile_name>
              create-iam-role
              delete-iam-role
              create-codecommit-repo
              create-admin-ec2
              """)
        sys.exit(1)

    function_name = sys.argv[1]
    profile_name = sys.argv[2]

    if function_name == 'create-iam-role':
        role_name = create_admin_role_for_ec2()
        create_admin_instance_profile(profile_name, role_name)
    elif function_name == 'delete-iam-role':
        delete_instance_profile(profile_name)
    elif function_name == 'create-codecommit-repo':
        create_codecommit_repo(profile_name)
    elif function_name == 'create-admin-ec2':
        instance_profile_arn = profile_name # pass in arn not name ie: arn:aws:iam::654654138626:instance-profile/admin
        ami_id = get_latest_amazon_linux_2023_ami()
        instance_type = "t3.small"
        security_group_ids = get_default_security_group_id()
        subnet_id = get_subnet_id_by_name()

        create_ec2_instance(instance_profile_arn, ami_id, instance_type, security_group_ids, subnet_id)
    elif function_name == 'create-all':
        '''Piece all individual pieces together for a single run all
        '''
    else:
        print("Invalid function name. Supported functions: create-iam-role, delete-iam-role, create-codecommit-repo, create-admin-ec2")
        sys.exit(1)