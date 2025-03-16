package utils

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func prompt(promt_msg string) bool {
    fmt.Print(promt_msg + "\n")
	fmt.Print("Press Enter to continue or any other key to quit...")

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    return input == "" 
}