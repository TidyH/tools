package utils

func printGreen(str string) {
	println("\x1b[32m" + str + "\x1b[0m\n")
}

func printRed(str string) {
	println("\x1b[31m" + str + "\x1b[0m\n")
}
