package cli

import (
	"kylin/i18n"
	"kylin/lexer"
	"kylin/lib"
	"os"
	"strings"
)

func ParseCommand(data []string) []string {
	for i := 0; i < len(data); i++ {
		data[i] = strings.TrimSpace(data[i])
	}
	return data[1:]
}

func ExecCommand() {
	command := ParseCommand(os.Args)
	if len(command) == 0 {
		command = []string{"main.ky"}
	}
	switch command[0] {
	default:
		ExecFile(command[0])
	}
}

func ExecFile(file string) {
	//runtime := interpret.NewRuntime(
	//	file,
	//	include.NewGlobalScope(),
	//	i18n.NewManager(),
	//)
	//runtime.Run()

	parser := lexer.NewParser(lib.ReadKylinFile(file), i18n.NewManager())
	parser.FpWrite(file)
}
