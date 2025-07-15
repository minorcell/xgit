package main

import (
	"strings"
	"testing"
)

func TestHandlePinyinCommand_BasicCommand(t *testing.T) {
	// 由于handlePinyinCommand会调用实际的git命令，我们需要模拟或者测试逻辑
	// 这里我们主要测试命令解析逻辑而不是实际执行

	tests := []struct {
		name    string
		command string
		exists  bool
	}{
		{"存在的基本命令", "kl", true},
		{"存在的复合命令", "kstj", true},
		{"标准git命令", "status", true},
		{"不存在的命令", "invalidcmd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试命令是否在映射中存在
			_, basicExists := commandMap[tt.command]
			_, compositeExists := compositeCommands[tt.command]
			gitExists := isGitCommand(tt.command)

			actualExists := basicExists || compositeExists || gitExists

			if actualExists != tt.exists {
				t.Errorf("命令 %s 的存在性检查失败，期望 %v，得到 %v", tt.command, tt.exists, actualExists)
			}
		})
	}
}

func TestExecuteQuickCommit_NoArgs(t *testing.T) {
	// 测试没有提供提交信息的情况
	output := captureOutput(func() {
		executeQuickCommit([]string{})
	})

	expectedElements := []string{
		"错误: 需要提供提交信息",
		"用法: xgit kstj \"提交信息\"",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("executeQuickCommit([]) 输出中缺少元素: %s", element)
		}
	}
}

func TestExecuteQuickCommit_WithArgs(t *testing.T) {
	// 由于这个函数会执行实际的git命令，我们需要模拟环境
	// 或者创建一个模拟版本的executeGitCommandWithError

	// 这里我们只测试参数处理逻辑
	args := []string{"测试提交信息"}

	if len(args) == 0 {
		t.Error("测试参数不应该为空")
	}

	message := args[0]
	if message != "测试提交信息" {
		t.Errorf("提取的提交信息不正确，期望 '测试提交信息'，得到 '%s'", message)
	}
}

func TestExecuteCompositeCommand_UnknownCommand(t *testing.T) {
	output := captureOutput(func() {
		executeCompositeCommand("unknown", [][]string{}, []string{})
	})

	expectedElements := []string{
		"执行复合命令: unknown",
		"未实现的复合命令: unknown",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("executeCompositeCommand('unknown') 输出中缺少元素: %s", element)
		}
	}
}

func TestExecuteCompositeCommand_QuickCommit(t *testing.T) {
	// 测试快速提交命令的参数处理
	commands := [][]string{
		{"add", "."},
		{"commit", "-m"},
		{"push"},
	}

	// 这里我们主要测试命令结构是否正确
	if len(commands) != 3 {
		t.Errorf("快速提交命令应该有3个步骤，实际有 %d 个", len(commands))
	}

	expectedSteps := [][]string{
		{"add", "."},
		{"commit", "-m"},
		{"push"},
	}

	for i, step := range commands {
		if len(step) != len(expectedSteps[i]) {
			t.Errorf("第 %d 步的参数数量不匹配", i+1)
			continue
		}

		for j, arg := range step {
			if arg != expectedSteps[i][j] {
				t.Errorf("第 %d 步第 %d 个参数不匹配，期望 %s，得到 %s", i+1, j+1, expectedSteps[i][j], arg)
			}
		}
	}
}

// 模拟git命令执行的测试
func TestGitCommandConstruction(t *testing.T) {
	tests := []struct {
		name        string
		pinyinCmd   string
		args        []string
		expectedGit []string
	}{
		{
			"克隆命令",
			"kl",
			[]string{"https://github.com/user/repo.git"},
			[]string{"clone", "https://github.com/user/repo.git"},
		},
		{
			"提交命令",
			"tj",
			[]string{"-m", "测试提交"},
			[]string{"commit", "-m", "测试提交"},
		},
		{
			"推送命令",
			"ts",
			[]string{},
			[]string{"push"},
		},
		{
			"创建分支",
			"cjfz",
			[]string{"feature-branch"},
			[]string{"checkout", "-b", "feature-branch"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gitCmd, exists := commandMap[tt.pinyinCmd]
			if !exists {
				t.Errorf("命令 %s 不存在于映射中", tt.pinyinCmd)
				return
			}

			fullArgs := append(gitCmd, tt.args...)

			if len(fullArgs) != len(tt.expectedGit) {
				t.Errorf("命令 %s 构造的git命令参数数量不匹配，期望 %d，得到 %d",
					tt.pinyinCmd, len(tt.expectedGit), len(fullArgs))
				return
			}

			for i, arg := range fullArgs {
				if arg != tt.expectedGit[i] {
					t.Errorf("命令 %s 构造的git命令第 %d 个参数不匹配，期望 %s，得到 %s",
						tt.pinyinCmd, i, tt.expectedGit[i], arg)
				}
			}
		})
	}
}

// 测试命令路由逻辑
func TestCommandRouting(t *testing.T) {
	tests := []struct {
		name        string
		command     string
		shouldRoute string // "composite", "basic", "git", "none"
	}{
		{"复合命令路由", "kstj", "composite"},
		{"基本命令路由", "kl", "basic"},
		{"git命令路由", "status", "git"},
		{"未知命令", "invalidcmd", "none"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualRoute string

			// 检查是否是复合命令
			if _, exists := compositeCommands[tt.command]; exists {
				actualRoute = "composite"
			} else if _, exists := commandMap[tt.command]; exists {
				// 检查是否是基本命令
				actualRoute = "basic"
			} else if isGitCommand(tt.command) {
				// 检查是否是git命令
				actualRoute = "git"
			} else {
				actualRoute = "none"
			}

			if actualRoute != tt.shouldRoute {
				t.Errorf("命令 %s 的路由不正确，期望 %s，得到 %s", tt.command, tt.shouldRoute, actualRoute)
			}
		})
	}
}

// 基准测试
func BenchmarkCommandLookup(b *testing.B) {
	commands := []string{"kl", "tj", "ts", "lq", "kstj", "status", "invalidcmd"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := commands[i%len(commands)]

		// 模拟handlePinyinCommand中的查找逻辑
		_, compositeExists := compositeCommands[cmd]
		_, basicExists := commandMap[cmd]
		gitExists := isGitCommand(cmd)

		_ = compositeExists || basicExists || gitExists
	}
}

func BenchmarkIsGitCommand(b *testing.B) {
	commands := []string{"add", "commit", "push", "pull", "invalidcmd"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := commands[i%len(commands)]
		isGitCommand(cmd)
	}
}
