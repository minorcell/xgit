package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	args := os.Args[1:]
	command := args[0]

	switch command {
	case "bz", "help":
		showHelp(args[1:])
	case "git":
		// 直接执行git命令
		executeGitCommand(args[1:])
	default:
		// 处理拼音命令
		handlePinyinCommand(command, args[1:])
	}
}
