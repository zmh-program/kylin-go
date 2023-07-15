package main

import (
	"kylin/i18n"
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

	runtime := interpret.NewRuntime(
		file,
		include.NewGlobalScope(),
		i18n.NewManager(),
	)
	runtime.Run()
}
