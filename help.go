package main

import (
	"fmt"
	"strings"
)

// 显示基本使用说明
func showUsage() {
	fmt.Println("xgit - 中文拼音首字母的Git命令工具")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  xgit <拼音命令> [参数...]     # 使用拼音首字母命令")
	fmt.Println("  xgit git <git命令> [参数...]  # 直接执行git命令")
	fmt.Println("  xgit bz [命令]               # 查看帮助")
	fmt.Println()
	fmt.Println("常用命令:")
	fmt.Println("  xgit kl <url>      # 克隆仓库")
	fmt.Println("  xgit tja .         # 添加所有文件")
	fmt.Println("  xgit tj -m 'msg'   # 提交更改")
	fmt.Println("  xgit ts            # 推送代码")
	fmt.Println("  xgit lq            # 拉取代码")
	fmt.Println()
	fmt.Println("运行 'xgit bz' 查看完整命令列表")
}

// 显示帮助信息
func showHelp(args []string) {
	if len(args) == 0 {
		fmt.Println("xgit 命令列表:")
		fmt.Println()

		for category, commands := range commandCategories {
			fmt.Printf("【%s】\n", category)
			for _, cmd := range commands {
				if help, exists := commandHelp[cmd]; exists {
					fmt.Printf("  %-6s %s\n", cmd, help)
				}
			}
			fmt.Println()
		}

		fmt.Println("使用 'xgit bz <命令>' 查看具体命令用法")
		fmt.Println("使用 'xgit bz --git <命令>' 查看对应的git命令")
		return
	}

	targetCmd := args[0]
	if len(args) > 1 && args[0] == "--git" {
		targetCmd = args[1]
		showGitEquivalent(targetCmd)
		return
	}

	if help, exists := commandHelp[targetCmd]; exists {
		fmt.Printf("命令: %s\n", targetCmd)
		fmt.Printf("说明: %s\n", help)
		fmt.Println()

		// 显示用法示例
		showUsageExamples(targetCmd)
	} else {
		fmt.Printf("未知命令: %s\n", targetCmd)
		fmt.Println("运行 'xgit bz' 查看所有可用命令")
	}
}

// 显示git等价命令
func showGitEquivalent(command string) {
	if gitCmd, exists := commandMap[command]; exists {
		fmt.Printf("%s → git %s\n", command, strings.Join(gitCmd, " "))
	} else if _, exists := compositeCommands[command]; exists {
		fmt.Printf("%s → 复合命令:\n", command)
		for i, cmd := range compositeCommands[command] {
			fmt.Printf("  %d. git %s\n", i+1, strings.Join(cmd, " "))
		}
	} else {
		fmt.Printf("未知命令: %s\n", command)
	}
}

// 显示用法示例
func showUsageExamples(command string) {
	switch command {
	case "kl":
		fmt.Println("用法示例:")
		fmt.Println("  xgit kl https://github.com/user/repo.git")
		fmt.Println("  xgit kl https://github.com/user/repo.git my-folder")
	case "tj":
		fmt.Println("用法示例:")
		fmt.Println("  xgit tj -m \"提交信息\"")
		fmt.Println("  xgit tj --amend")
	case "kstj":
		fmt.Println("用法示例:")
		fmt.Println("  xgit kstj \"快速提交信息\"")
	case "cjfz":
		fmt.Println("用法示例:")
		fmt.Println("  xgit cjfz feature-branch")
		fmt.Println("  xgit cjfz hotfix/bug-123")
	case "qhfz":
		fmt.Println("用法示例:")
		fmt.Println("  xgit qhfz main")
		fmt.Println("  xgit qhfz feature-branch")
	case "hb":
		fmt.Println("用法示例:")
		fmt.Println("  xgit hb feature-branch")
		fmt.Println("  xgit hb --no-ff feature-branch")
	case "zf":
		fmt.Println("用法示例:")
		fmt.Println("  xgit zf main")
		fmt.Println("  xgit zf origin/main")
	case "cjbq":
		fmt.Println("用法示例:")
		fmt.Println("  xgit cjbq v1.0.0 -m \"Release version 1.0.0\"")
	case "ch":
		fmt.Println("用法示例:")
		fmt.Println("  xgit ch file.txt")
		fmt.Println("  xgit ch .")
	case "ht":
		fmt.Println("用法示例:")
		fmt.Println("  xgit ht HEAD~1")
		fmt.Println("  xgit ht --hard HEAD~2")
	}
}
