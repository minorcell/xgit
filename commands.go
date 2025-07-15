package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// JSON配置结构体
type Command struct {
	Args        []string `json:"args"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
}

type CompositeCommand struct {
	Steps       [][]string `json:"steps"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
}

type CommandConfig struct {
	Commands          map[string]Command          `json:"commands"`
	CompositeCommands map[string]CompositeCommand `json:"composite_commands"`
	GitCommands       []string                    `json:"git_commands"`
}

// 全局变量
var (
	commandMap        map[string][]string
	compositeCommands map[string][][]string
	commandHelp       map[string]string
	commandCategories map[string][]string
	gitCommands       []string
	config            *CommandConfig
)

// 初始化函数，读取JSON配置
func init() {
	loadConfig()
}

// 加载配置文件
func loadConfig() {
	// 获取执行文件所在目录
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("错误：无法获取执行路径: %v\n", err)
		os.Exit(1)
	}

	// 配置文件路径
	configPath := filepath.Join(filepath.Dir(execPath), "commands.json")

	// 如果执行文件目录没有配置文件，尝试当前工作目录
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = "commands.json"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("错误：无法读取配置文件 %s: %v\n", configPath, err)
		os.Exit(1)
	}

	// 解析JSON
	config = &CommandConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		fmt.Printf("错误：无法解析配置文件: %v\n", err)
		os.Exit(1)
	}

	// 生成映射
	generateMappings()
}

// 根据配置生成映射
func generateMappings() {
	// 初始化映射
	commandMap = make(map[string][]string)
	compositeCommands = make(map[string][][]string)
	commandHelp = make(map[string]string)
	commandCategories = make(map[string][]string)

	// 处理基本命令
	for key, cmd := range config.Commands {
		commandMap[key] = cmd.Args
		commandHelp[key] = cmd.Description

		// 添加到分类
		if commandCategories[cmd.Category] == nil {
			commandCategories[cmd.Category] = []string{}
		}
		commandCategories[cmd.Category] = append(commandCategories[cmd.Category], key)
	}

	// 处理复合命令
	for key, cmd := range config.CompositeCommands {
		compositeCommands[key] = cmd.Steps
		commandHelp[key] = cmd.Description

		// 添加到分类
		if commandCategories[cmd.Category] == nil {
			commandCategories[cmd.Category] = []string{}
		}
		commandCategories[cmd.Category] = append(commandCategories[cmd.Category], key)
	}

	// 设置Git命令列表
	gitCommands = config.GitCommands
}

// 检查是否是标准git命令
func isGitCommand(command string) bool {
	for _, gitCmd := range gitCommands {
		if command == gitCmd {
			return true
		}
	}
	return false
}
