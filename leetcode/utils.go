package utils

func printGreen(str string) string {
	return "\x1b[32m" + str + "\x1b[0m\n"
}

func printRed(str string) string {
	return "\x1b[31m" + str + "\x1b[0m\n"
}
