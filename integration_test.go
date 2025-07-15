package main

import (
	"os"
	"strings"
	"testing"
)

// 集成测试：测试主程序的参数处理
func TestMain_NoArgs(t *testing.T) {
	// 备份原始参数
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// 设置测试参数（只有程序名）
	os.Args = []string{"xgit"}

	output := captureOutput(func() {
		main()
	})

	// 检查是否显示了使用说明
	if !strings.Contains(output, "xgit - 中文拼音首字母的Git命令工具") {
		t.Error("没有参数时应该显示使用说明")
	}
}

func TestMain_HelpCommand(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	testCases := []string{"bz", "help"}

	for _, helpCmd := range testCases {
		t.Run("help_command_"+helpCmd, func(t *testing.T) {
			os.Args = []string{"xgit", helpCmd}

			output := captureOutput(func() {
				main()
			})

			// 检查是否显示了命令列表
			if !strings.Contains(output, "xgit 命令列表:") {
				t.Errorf("命令 %s 应该显示帮助列表", helpCmd)
			}
		})
	}
}

func TestMain_SpecificHelpCommand(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"xgit", "bz", "kl"}

	output := captureOutput(func() {
		main()
	})

	expectedElements := []string{
		"命令: kl",
		"说明: 克隆仓库",
		"用法示例:",
	}

	for _, element := range expectedElements {
		if !strings.Contains(output, element) {
			t.Errorf("特定帮助命令输出中缺少元素: %s", element)
		}
	}
}

func TestMain_GitEquivalentCommand(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"xgit", "bz", "--git", "kl"}

	output := captureOutput(func() {
		main()
	})

	if !strings.Contains(output, "kl → git clone") {
		t.Error("--git 选项应该显示git等价命令")
	}
}

// 测试完整的命令流程（不实际执行git）
func TestCompleteCommandFlow(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		contains    []string
	}{
		{
			name:        "显示基本帮助",
			args:        []string{"xgit"},
			expectError: false,
			contains:    []string{"中文拼音首字母的Git命令工具", "用法:"},
		},
		{
			name:        "显示命令列表",
			args:        []string{"xgit", "bz"},
			expectError: false,
			contains:    []string{"命令列表:", "【仓库操作】", "【文件操作】"},
		},
		{
			name:        "显示特定命令帮助",
			args:        []string{"xgit", "bz", "tj"},
			expectError: false,
			contains:    []string{"命令: tj", "提交更改", "用法示例:"},
		},
		{
			name:        "显示git等价命令",
			args:        []string{"xgit", "bz", "--git", "kstj"},
			expectError: false,
			contains:    []string{"复合命令:", "git add .", "git commit -m", "git push"},
		},
		{
			name:        "未知命令帮助",
			args:        []string{"xgit", "bz", "unknown"},
			expectError: false,
			contains:    []string{"未知命令: unknown"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args

			output := captureOutput(func() {
				main()
			})

			for _, expected := range tt.contains {
				if !strings.Contains(output, expected) {
					t.Errorf("测试 %s: 输出中缺少期望内容 '%s'\n实际输出:\n%s",
						tt.name, expected, output)
				}
			}
		})
	}
}

// 测试命令映射的完整性
func TestCommandMappingCompleteness(t *testing.T) {
	// 检查所有在分类中的命令都有对应的映射或帮助
	for category, commands := range commandCategories {
		for _, cmd := range commands {
			t.Run("mapping_"+category+"_"+cmd, func(t *testing.T) {
				hasMapping := false

				// 检查基本命令映射
				if _, exists := commandMap[cmd]; exists {
					hasMapping = true
				}

				// 检查复合命令映射
				if _, exists := compositeCommands[cmd]; exists {
					hasMapping = true
				}

				if !hasMapping {
					t.Errorf("分类 %s 中的命令 %s 没有对应的映射", category, cmd)
				}

				// 检查是否有帮助信息
				if _, exists := commandHelp[cmd]; !exists {
					t.Errorf("分类 %s 中的命令 %s 没有帮助信息", category, cmd)
				}
			})
		}
	}
}

// 测试命令一致性
func TestCommandConsistency(t *testing.T) {
	// 所有commandMap中的命令都应该在commandHelp中
	for cmd := range commandMap {
		if _, exists := commandHelp[cmd]; !exists {
			t.Errorf("命令 %s 在commandMap中但不在commandHelp中", cmd)
		}
	}

	// 所有compositeCommands中的命令都应该在commandHelp中
	for cmd := range compositeCommands {
		if _, exists := commandHelp[cmd]; !exists {
			t.Errorf("复合命令 %s 在compositeCommands中但不在commandHelp中", cmd)
		}
	}

	// 所有在commandCategories中的命令都应该存在于映射中
	allCategorizedCommands := make(map[string]bool)
	for _, commands := range commandCategories {
		for _, cmd := range commands {
			allCategorizedCommands[cmd] = true
		}
	}

	// 检查基本命令
	for cmd := range commandMap {
		if !allCategorizedCommands[cmd] {
			t.Errorf("基本命令 %s 没有被分类", cmd)
		}
	}

	// 检查复合命令
	for cmd := range compositeCommands {
		if !allCategorizedCommands[cmd] {
			t.Errorf("复合命令 %s 没有被分类", cmd)
		}
	}
}

// 性能测试：测试命令查找性能
func BenchmarkMainCommandLookup(b *testing.B) {
	testCommands := []string{"kl", "tj", "ts", "lq", "kstj", "bz"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := testCommands[i%len(testCommands)]

		// 模拟main函数中的命令查找逻辑
		switch cmd {
		case "bz", "help":
			// 帮助命令
		case "git":
			// git命令
		default:
			// 拼音命令查找
			_, compositeExists := compositeCommands[cmd]
			_, basicExists := commandMap[cmd]
			gitExists := isGitCommand(cmd)
			_ = compositeExists || basicExists || gitExists
		}
	}
}

// 测试内存使用情况
func TestMemoryUsage(t *testing.T) {
	// 检查映射表的大小是否合理
	totalCommands := len(commandMap) + len(compositeCommands)
	totalHelp := len(commandHelp)
	totalCategories := 0
	for _, commands := range commandCategories {
		totalCategories += len(commands)
	}

	t.Logf("命令映射数量: %d", len(commandMap))
	t.Logf("复合命令数量: %d", len(compositeCommands))
	t.Logf("总命令数量: %d", totalCommands)
	t.Logf("帮助信息数量: %d", totalHelp)
	t.Logf("分类命令总数: %d", totalCategories)

	// 基本的一致性检查
	if totalHelp < totalCommands {
		t.Errorf("帮助信息数量 (%d) 少于命令数量 (%d)", totalHelp, totalCommands)
	}

	if totalCategories < totalCommands {
		t.Errorf("分类命令数量 (%d) 少于总命令数量 (%d)", totalCategories, totalCommands)
	}
}
