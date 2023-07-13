package main

import (
	"kylin/include"
	"kylin/interpret"
	"os"
)

func main() {
	var file string
	if len(os.Args) < 2 {
		file = "main.ky"
	} else {
		file = os.Args[1]
	}

	runtime := interpret.NewInterpreter(
		file,
		include.NewGlobalScope(),
	)
	runtime.Run()
}
