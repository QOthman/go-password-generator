package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	includeSpecialChar := false
	includeDigits := false
	includeUppercase := false
	passwordLength := 0

	if len(args) == 0 {
		fmt.Println("Usage: program_name [-s] [-d] [-u] [-l password_length]")
		return
	}

	for i := 0; i < len(args); i++ {
		if args[i][0] == '-' {
			switch args[i][1] {
			case 'h':
				printUsage()
				return
			case 's':
				includeSpecialChar = true
			case 'd':
				includeDigits = true
			case 'a':
				includeUppercase = true
			case 'l':
				if len(args) > i+1 {
					passwordLength, _ = strconv.Atoi(args[i+1])
					i++
				} else {
					fmt.Println("Password length is missing")
					return
				}
			default:
				fmt.Println("Invalid option")
				return
			}
		} else {
			fmt.Println("Invalid option")
			return
		}
	}

	if !includeSpecialChar && !includeDigits && !includeUppercase {
		fmt.Println("At least one option (-s, -d, -a) must be selected")
		return
	}

	charSet := ""

	if includeSpecialChar {
		charSet += "!&%$£(){}-_><?=+*,;."
	}

	if includeDigits {
		charSet += "0123456789"
	}

	if includeUppercase {
		charSet += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	password := generatePassword(charSet, passwordLength)
	fmt.Println(password)
}

func generatePassword(charSet string, length int) string {
	password := ""
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSet))
		password += string(charSet[randomIndex])
	}
	return password
}

func printUsage() {
	fmt.Println(`This program generates random passwords based on user-defined criteria.
Usage: 
go run main.go [options]
Options:
-s         Include special characters in the password.
-d         Include digits (0-9) in the password.
-a         Include uppercase letters in the password.
-l length  Specify the length of the password.`)
}
