package main

import (
	"testing"
)

func TestCommandMap(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected []string
	}{
		{"克隆命令", "kl", []string{"clone"}},
		{"初始化命令", "csh", []string{"init"}},
		{"添加命令", "tja", []string{"add"}},
		{"提交命令", "tj", []string{"commit"}},
		{"推送命令", "ts", []string{"push"}},
		{"拉取命令", "lq", []string{"pull"}},
		{"分支列表命令", "fz", []string{"branch"}},
		{"创建分支命令", "cjfz", []string{"checkout", "-b"}},
		{"切换分支命令", "qhfz", []string{"checkout"}},
		{"日志命令", "rz", []string{"log"}},
		{"一行日志命令", "yhrz", []string{"log", "--oneline"}},
		{"状态命令", "zt", []string{"status"}},
		{"标签命令", "bq", []string{"tag"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, exists := commandMap[tt.command]
			if !exists {
				t.Errorf("命令 %s 不存在于 commandMap 中", tt.command)
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("命令 %s 的参数数量不匹配，期望 %d，得到 %d", tt.command, len(tt.expected), len(result))
				return
			}

			for i, arg := range result {
				if arg != tt.expected[i] {
					t.Errorf("命令 %s 的第 %d 个参数不匹配，期望 %s，得到 %s", tt.command, i, tt.expected[i], arg)
				}
			}
		})
	}
}

func TestCompositeCommands(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected [][]string
	}{
		{
			"快速提交命令",
			"kstj",
			[][]string{
				{"add", "."},
				{"commit", "-m"},
				{"push"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, exists := compositeCommands[tt.command]
			if !exists {
				t.Errorf("复合命令 %s 不存在于 compositeCommands 中", tt.command)
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("复合命令 %s 的步骤数量不匹配，期望 %d，得到 %d", tt.command, len(tt.expected), len(result))
				return
			}

			for i, step := range result {
				if len(step) != len(tt.expected[i]) {
					t.Errorf("复合命令 %s 的第 %d 步参数数量不匹配", tt.command, i)
					continue
				}
				for j, arg := range step {
					if arg != tt.expected[i][j] {
						t.Errorf("复合命令 %s 的第 %d 步第 %d 个参数不匹配，期望 %s，得到 %s",
							tt.command, i, j, tt.expected[i][j], arg)
					}
				}
			}
		})
	}
}

func TestCommandHelp(t *testing.T) {
	// 测试所有命令都有对应的帮助信息
	for cmd := range commandMap {
		if _, exists := commandHelp[cmd]; !exists {
			t.Errorf("命令 %s 缺少帮助信息", cmd)
		}
	}

	for cmd := range compositeCommands {
		if _, exists := commandHelp[cmd]; !exists {
			t.Errorf("复合命令 %s 缺少帮助信息", cmd)
		}
	}
}

func TestCommandCategories(t *testing.T) {
	// 测试命令分类是否包含所有命令
	allCategorizedCommands := make(map[string]bool)
	for _, commands := range commandCategories {
		for _, cmd := range commands {
			allCategorizedCommands[cmd] = true
		}
	}

	// 检查基本命令是否都被分类
	for cmd := range commandMap {
		if !allCategorizedCommands[cmd] {
			t.Errorf("命令 %s 没有被分类", cmd)
		}
	}

	// 检查复合命令是否都被分类
	for cmd := range compositeCommands {
		if !allCategorizedCommands[cmd] {
			t.Errorf("复合命令 %s 没有被分类", cmd)
		}
	}
}

func TestIsGitCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected bool
	}{
		{"标准git命令-add", "add", true},
		{"标准git命令-commit", "commit", true},
		{"标准git命令-push", "push", true},
		{"标准git命令-pull", "pull", true},
		{"标准git命令-clone", "clone", true},
		{"标准git命令-init", "init", true},
		{"标准git命令-status", "status", true},
		{"标准git命令-branch", "branch", true},
		{"标准git命令-checkout", "checkout", true},
		{"标准git命令-merge", "merge", true},
		{"标准git命令-rebase", "rebase", true},
		{"标准git命令-reset", "reset", true},
		{"标准git命令-log", "log", true},
		{"标准git命令-fetch", "fetch", true},
		{"标准git命令-remote", "remote", true},
		{"标准git命令-tag", "tag", true},
		{"标准git命令-diff", "diff", true},
		{"标准git命令-stash", "stash", true},
		{"非git命令", "invalidcommand", false},
		{"拼音命令", "kl", false},
		{"空字符串", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isGitCommand(tt.command)
			if result != tt.expected {
				t.Errorf("isGitCommand(%s) = %v，期望 %v", tt.command, result, tt.expected)
			}
		})
	}
}

func TestCommandMapConsistency(t *testing.T) {
	// 测试命令映射的一致性
	for cmd, gitCmd := range commandMap {
		if len(gitCmd) == 0 {
			t.Errorf("命令 %s 映射的git命令为空", cmd)
		}

		// 基本的git命令应该是有效的
		if len(gitCmd) > 0 && !isGitCommand(gitCmd[0]) {
			// 允许一些特殊情况，比如带有 "--" 的命令
			if gitCmd[0] != "--" {
				t.Errorf("命令 %s 映射到无效的git命令: %s", cmd, gitCmd[0])
			}
		}
	}
}
