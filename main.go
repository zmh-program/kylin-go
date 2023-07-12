package main

import (
	"fmt"
	"kylin/interpret"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a file to run.")
		return
	}
	runtime := interpret.NewInterpreter(os.Args[1], nil)
	runtime.Run()
}
