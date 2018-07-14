package internal

import (
	"fmt"
	"os"
)

func PrintErr(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31mslices: "+msg+"\n", args...)
}

func PrintSuccess(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[32mslices: "+msg+"\n", args...)
}

func PrintInfo(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[36mslices: "+msg+"\n", args...)
}
