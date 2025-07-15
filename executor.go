package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 处理拼音命令
func handlePinyinCommand(command string, args []string) {
	// 检查是否是复合命令
	if composite, exists := compositeCommands[command]; exists {
		executeCompositeCommand(command, composite, args)
		return
	}

	// 检查是否是基本命令
	if gitCmd, exists := commandMap[command]; exists {
		fullArgs := append(gitCmd, args...)
		executeGitCommand(fullArgs)
		return
	}

	// 检查是否是原生git命令
	if isGitCommand(command) {
		fullArgs := append([]string{command}, args...)
		executeGitCommand(fullArgs)
		return
	}

	fmt.Printf("未知命令: %s\n", command)
	fmt.Println("运行 'xgit bz' 查看所有可用命令")
	os.Exit(1)
}

// 执行复合命令
func executeCompositeCommand(cmdName string, commands [][]string, args []string) {
	fmt.Printf("执行复合命令: %s\n", cmdName)

	switch cmdName {
	case "kstj": // 快速提交
		executeQuickCommit(args)
	case "ycsh": // 远程设置
		executeRemoteSetup(args)
	default:
		fmt.Printf("未实现的复合命令: %s\n", cmdName)
	}
}

// 执行快速提交
func executeQuickCommit(args []string) {
	if len(args) == 0 {
		fmt.Println("错误: 需要提供提交信息")
		fmt.Println("用法: xgit kstj \"提交信息\"")
		return
	}

	message := args[0]

	// 1. git add .
	fmt.Println("→ 添加所有文件...")
	if err := executeGitCommandWithError([]string{"add", "."}); err != nil {
		fmt.Printf("添加文件失败: %v\n", err)
		return
	}

	// 2. git commit -m
	fmt.Printf("→ 提交更改: %s\n", message)
	if err := executeGitCommandWithError([]string{"commit", "-m", message}); err != nil {
		fmt.Printf("提交失败: %v\n", err)
		return
	}

	// 3. git push
	fmt.Println("→ 推送到远程...")
	if err := executeGitCommandWithError([]string{"push"}); err != nil {
		fmt.Printf("推送失败: %v\n", err)
		return
	}

	fmt.Println("✅ 快速提交完成！")
}

// 执行远程设置
func executeRemoteSetup(args []string) {
	if len(args) == 0 {
		fmt.Println("错误: 需要提供远程仓库URL")
		fmt.Println("用法: xgit ycsh <远程仓库URL>")
		return
	}

	url := args[0]

	// 1. git remote add origin <url>
	fmt.Printf("→ 添加远程仓库: %s\n", url)
	if err := executeGitCommandWithError([]string{"remote", "add", "origin", url}); err != nil {
		fmt.Printf("添加远程仓库失败: %v\n", err)
		return
	}

	// 2. git push -u origin main (或当前分支)
	fmt.Println("→ 推送并设置上游分支...")
	branch := "main" // 默认使用main分支
	if len(args) > 1 {
		branch = args[1] // 如果提供了分支名，使用指定分支
	}

	if err := executeGitCommandWithError([]string{"push", "-u", "origin", branch}); err != nil {
		fmt.Printf("推送失败: %v\n", err)
		return
	}

	fmt.Println("✅ 远程设置完成！")
}

// 执行git命令
func executeGitCommand(args []string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Printf("执行git命令时出错: %v\n", err)
		os.Exit(1)
	}
}

// 执行git命令并返回错误（用于复合命令的错误处理）
func executeGitCommandWithError(args []string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
