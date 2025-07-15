package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 拼音首字母到git命令的映射
var commandMap = map[string][]string{
	// 仓库操作
	"kl":  {"clone"}, // 克隆
	"csh": {"init"},  // 初始化

	// 文件操作
	"tja": {"add"},            // 添加
	"tj":  {"commit"},         // 提交
	"ch":  {"checkout", "--"}, // 撤回文件

	// 分支操作
	"fz": {"branch"},         // 分支列表
	"cj": {"checkout", "-b"}, // 创建分支
	"qh": {"checkout"},       // 切换分支

	// 远程操作
	"ts": {"push"},  // 推送
	"lq": {"pull"},  // 拉取
	"hq": {"fetch"}, // 获取

	// 高级操作
	"hb": {"merge"},  // 合并
	"zf": {"rebase"}, // 整合
	"ht": {"reset"},  // 回退
}

// 复合命令映射
var compositeCommands = map[string][][]string{
	"kstj": { // 快速提交
		{"add", "."},
		{"commit", "-m"},
		{"push"},
	},
}

// 命令说明
var commandHelp = map[string]string{
	"kl":   "克隆仓库 (ke long) → git clone <url>",
	"csh":  "初始化仓库 (chu shi hua) → git init",
	"tja":  "添加文件 (tian jia) → git add <file>",
	"tj":   "提交更改 (ti jiao) → git commit -m <message>",
	"ch":   "撤回文件 (che hui) → git checkout -- <file>",
	"fz":   "查看分支 (fen zhi) → git branch",
	"cj":   "创建分支 (chuang jian) → git checkout -b <branch>",
	"qh":   "切换分支 (qie huan) → git checkout <branch>",
	"ts":   "推送代码 (tui song) → git push",
	"lq":   "拉取代码 (la qu) → git pull",
	"hq":   "获取更新 (huo qu) → git fetch",
	"hb":   "合并分支 (he bing) → git merge <branch>",
	"zf":   "整合分支 (zheng he) → git rebase <branch>",
	"ht":   "回退版本 (hui tui) → git reset",
	"kstj": "快速提交 (kuai su ti jiao) → git add . && git commit -m && git push",
}

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

func showHelp(args []string) {
	if len(args) == 0 {
		fmt.Println("xgit 命令列表:")
		fmt.Println()

		categories := map[string][]string{
			"仓库操作": {"kl", "csh"},
			"文件操作": {"tja", "tj", "ch"},
			"分支操作": {"fz", "cj", "qh"},
			"远程操作": {"ts", "lq", "hq"},
			"高级操作": {"hb", "zf", "ht"},
			"复合命令": {"kstj"},
		}

		for category, commands := range categories {
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
		switch targetCmd {
		case "kl":
			fmt.Println("用法示例:")
			fmt.Println("  xgit kl https://github.com/user/repo.git")
		case "tj":
			fmt.Println("用法示例:")
			fmt.Println("  xgit tj -m \"提交信息\"")
		case "kstj":
			fmt.Println("用法示例:")
			fmt.Println("  xgit kstj \"快速提交信息\"")
		}
	} else {
		fmt.Printf("未知命令: %s\n", targetCmd)
		fmt.Println("运行 'xgit bz' 查看所有可用命令")
	}
}

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

func executeCompositeCommand(cmdName string, commands [][]string, args []string) {
	fmt.Printf("执行复合命令: %s\n", cmdName)

	switch cmdName {
	case "kstj": // 快速提交
		if len(args) == 0 {
			fmt.Println("错误: 需要提供提交信息")
			fmt.Println("用法: xgit kstj \"提交信息\"")
			return
		}

		message := args[0]

		// 1. git add .
		fmt.Println("→ 添加所有文件...")
		executeGitCommand([]string{"add", "."})

		// 2. git commit -m
		fmt.Printf("→ 提交更改: %s\n", message)
		executeGitCommand([]string{"commit", "-m", message})

		// 3. git push
		fmt.Println("→ 推送到远程...")
		executeGitCommand([]string{"push"})
	}
}

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

func isGitCommand(command string) bool {
	// 检查是否是标准git命令
	gitCommands := []string{
		"add", "commit", "push", "pull", "clone", "init", "status",
		"branch", "checkout", "merge", "rebase", "reset", "log",
		"fetch", "remote", "tag", "diff", "stash",
	}

	for _, gitCmd := range gitCommands {
		if command == gitCmd {
			return true
		}
	}
	return false
}
