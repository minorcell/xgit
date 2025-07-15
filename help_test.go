package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// 捕获输出的帮助函数
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestShowUsage(t *testing.T) {
	output := captureOutput(func() {
		showUsage()
	})

	// 检查基本元素是否存在
	expectedElements := []string{
		"xgit - 中文拼音首字母的Git命令工具",
		"用法:",
		"xgit <拼音命令> [参数...]",
		"xgit git <git命令> [参数...]",
		"xgit bz [命令]",
		"常用命令:",
		"xgit kl <url>",
		"xgit tja .",
		"xgit tj -m",
		"xgit ts",
		"xgit lq",
		"运行 'xgit bz' 查看完整命令列表",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("showUsage() 输出中缺少元素: %s", element)
		}
	}
}

func TestShowHelp_EmptyArgs(t *testing.T) {
	output := captureOutput(func() {
		showHelp([]string{})
	})

	// 检查帮助列表的基本元素
	expectedElements := []string{
		"xgit 命令列表:",
		"【仓库操作】",
		"【文件操作】",
		"【分支操作】",
		"【远程操作】",
		"【高级操作】",
		"【日志操作】",
		"【状态操作】",
		"【标签操作】",
		"【复合命令】",
		"使用 'xgit bz <命令>' 查看具体命令用法",
		"使用 'xgit bz --git <命令>' 查看对应的git命令",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("showHelp([]) 输出中缺少元素: %s", element)
		}
	}

	// 检查是否包含主要命令
	mainCommands := []string{"kl", "tj", "ts", "lq", "kstj"}
	for _, cmd := range mainCommands {
		if !strings.Contains(output, cmd) {
			t.Errorf("showHelp([]) 输出中缺少命令: %s", cmd)
		}
	}
}

func TestShowHelp_SpecificCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected []string
	}{
		{
			"克隆命令帮助",
			"kl",
			[]string{
				"命令: kl",
				"说明: 克隆仓库 (ke long) → git clone <url>",
				"用法示例:",
				"xgit kl https://github.com/user/repo.git",
			},
		},
		{
			"提交命令帮助",
			"tj",
			[]string{
				"命令: tj",
				"说明: 提交更改 (ti jiao) → git commit -m <message>",
				"用法示例:",
				"xgit tj -m \"提交信息\"",
			},
		},
		{
			"快速提交命令帮助",
			"kstj",
			[]string{
				"命令: kstj",
				"说明: 快速提交 (kuai su ti jiao) → git add . && git commit -m && git push",
				"用法示例:",
				"xgit kstj \"快速提交信息\"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				showHelp([]string{tt.command})
			})

			for _, expected := range tt.expected {
				if !strings.Contains(output, expected) {
					t.Errorf("showHelp([%s]) 输出中缺少元素: %s\n实际输出:\n%s", tt.command, expected, output)
				}
			}
		})
	}
}

func TestShowHelp_UnknownCommand(t *testing.T) {
	output := captureOutput(func() {
		showHelp([]string{"invalidcommand"})
	})

	expectedElements := []string{
		"未知命令: invalidcommand",
		"运行 'xgit bz' 查看所有可用命令",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("showHelp([invalidcommand]) 输出中缺少元素: %s", element)
		}
	}
}

func TestShowGitEquivalent(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected []string
	}{
		{
			"基本命令映射",
			"kl",
			[]string{"kl → git clone"},
		},
		{
			"带参数的命令映射",
			"cjfz",
			[]string{"cjfz → git checkout -b"},
		},
		{
			"复合命令映射",
			"kstj",
			[]string{
				"kstj → 复合命令:",
				"1. git add .",
				"2. git commit -m",
				"3. git push",
			},
		},
		{
			"未知命令",
			"invalidcommand",
			[]string{"未知命令: invalidcommand"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				showGitEquivalent(tt.command)
			})

			for _, expected := range tt.expected {
				if !strings.Contains(output, expected) {
					t.Errorf("showGitEquivalent(%s) 输出中缺少元素: %s\n实际输出:\n%s",
						tt.command, expected, output)
				}
			}
		})
	}
}

func TestShowHelp_GitEquivalent(t *testing.T) {
	output := captureOutput(func() {
		showHelp([]string{"--git", "kl"})
	})

	if !strings.Contains(output, "kl → git clone") {
		t.Errorf("showHelp([--git, kl]) 应该显示git等价命令，实际输出:\n%s", output)
	}
}

func TestShowUsageExamples(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected []string
	}{
		{
			"克隆命令示例",
			"kl",
			[]string{
				"用法示例:",
				"xgit kl https://github.com/user/repo.git",
				"xgit kl https://github.com/user/repo.git my-folder",
			},
		},
		{
			"提交命令示例",
			"tj",
			[]string{
				"用法示例:",
				"xgit tj -m \"提交信息\"",
				"xgit tj --amend",
			},
		},
		{
			"创建分支示例",
			"cjfz",
			[]string{
				"用法示例:",
				"xgit cjfz feature-branch",
				"xgit cjfz hotfix/bug-123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				showUsageExamples(tt.command)
			})

			for _, expected := range tt.expected {
				if !strings.Contains(output, expected) {
					t.Errorf("showUsageExamples(%s) 输出中缺少元素: %s\n实际输出:\n%s",
						tt.command, expected, output)
				}
			}
		})
	}
}

func TestShowUsageExamples_NoExample(t *testing.T) {
	// 测试没有特定示例的命令
	output := captureOutput(func() {
		showUsageExamples("ts")
	})

	// 对于没有特定示例的命令，不应该有任何输出
	if strings.TrimSpace(output) != "" {
		t.Errorf("showUsageExamples(ts) 应该没有输出，但得到: %s", output)
	}
}
