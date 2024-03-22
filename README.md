# Go Password Generator

This Go program generates random passwords based on user-defined criteria.

## Installation

1. Clone this repository.
2. Navigate to the directory containing the code.
3. Run `go build` to build the executable.
4. Run the executable with appropriate command-line options.

## Usage

go run main.go [options]


### Options

- `-s`: Include special characters in the password.
- `-d`: Include digits (0-9) in the password.
- `-a`: Include uppercase letters in the password.
- `-l length`: Specify the length of the password.

## Example

To generate a password with special characters, digits, and a length of 12:

go run main.go -s -d -a -l 12

