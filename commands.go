package main

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
	"fz":   {"branch"},         // 分支列表
	"fzxq": {"branch", "-v"},   // 分支详情(详细)
	"ycfz": {"branch", "-r"},   // 远程分支列表
	"cjfz": {"checkout", "-b"}, // 创建分支
	"qhfz": {"checkout"},       // 切换分支

	// 远程操作
	"ts": {"push"},  // 推送
	"lq": {"pull"},  // 拉取
	"hq": {"fetch"}, // 获取

	// 远程仓库管理
	"ycck": {"remote", "-v"},      // 远程查看
	"yctz": {"remote", "add"},     // 远程添加
	"ycsc": {"remote", "remove"},  // 远程删除
	"yczm": {"remote", "rename"},  // 远程重命名
	"ycxg": {"remote", "set-url"}, // 远程修改URL
	"ycxq": {"remote", "show"},    // 远程详情

	// 高级操作
	"hb": {"merge"},  // 合并
	"zf": {"rebase"}, // 整合
	"ht": {"reset"},  // 回退

	// 日志操作
	"rz":   {"log"},              // 日志
	"yhrz": {"log", "--oneline"}, // 一行日志

	// 状态
	"zt":   {"status"},       // 状态
	"ztxq": {"status", "-s"}, // 状态详情

	// 标签
	"bq":   {"tag"},       // 标签列表
	"cjbq": {"tag", "-a"}, // 创建标签
	"bqxq": {"tag", "-l"}, // 标签详情
}

// 复合命令映射
var compositeCommands = map[string][][]string{
	"kstj": { // 快速提交
		{"add", "."},
		{"commit", "-m"},
		{"push"},
	},
	"ycsh": { // 远程设置 - 添加origin并推送
		{"remote", "add", "origin"},
		{"push", "-u", "origin", "main"},
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
	"fzxq": "分支详情 (fen zhi xiang qing) → git branch -v",
	"ycfz": "远程分支 (yuan cheng fen zhi) → git branch -r",
	"cjfz": "创建分支 (chuang jian fen zhi) → git checkout -b <branch>",
	"qhfz": "切换分支 (qie huan fen zhi) → git checkout <branch>",
	"ts":   "推送代码 (tui song) → git push",
	"lq":   "拉取代码 (la qu) → git pull",
	"hq":   "获取更新 (huo qu) → git fetch",
	"ycck": "查看远程仓库 (yuan cheng cha kan) → git remote -v",
	"yctz": "添加远程仓库 (yuan cheng tian jia) → git remote add <name> <url>",
	"ycsc": "删除远程仓库 (yuan cheng shan chu) → git remote remove <name>",
	"yczm": "重命名远程仓库 (yuan cheng zhong ming) → git remote rename <old> <new>",
	"ycxg": "修改远程URL (yuan cheng xiu gai) → git remote set-url <name> <url>",
	"ycxq": "远程仓库详情 (yuan cheng xiang qing) → git remote show <name>",
	"hb":   "合并分支 (he bing) → git merge <branch>",
	"zf":   "整合分支 (zheng he) → git rebase <branch>",
	"ht":   "回退版本 (hui tui) → git reset",
	"rz":   "查看日志 (ri zhi) → git log",
	"yhrz": "一行日志 (yi hang ri zhi) → git log --oneline",
	"zt":   "状态 (zhuang tai) → git status",
	"ztxq": "状态详情 (zhuang tai xiang qing) → git status -s",
	"bq":   "标签列表 (biao qian) → git tag",
	"cjbq": "创建标签 (chuang jian biao qian) → git tag -a <tag> -m <message>",
	"bqxq": "标签详情 (biao qian xiang qing) → git tag -l",
	"kstj": "快速提交 (kuai su ti jiao) → git add . && git commit -m && git push",
	"ycsh": "远程设置 (yuan cheng she zhi) → git remote add origin <url> && git push -u origin main",
}

// 命令分类
var commandCategories = map[string][]string{
	"仓库操作": {"kl", "csh"},
	"文件操作": {"tja", "tj", "ch"},
	"分支操作": {"fz", "fzxq", "ycfz", "cjfz", "qhfz"},
	"远程操作": {"ts", "lq", "hq", "ycck", "yctz", "ycsc", "yczm", "ycxg", "ycxq"},
	"高级操作": {"hb", "zf", "ht"},
	"日志操作": {"rz", "yhrz"},
	"状态操作": {"zt", "ztxq"},
	"标签操作": {"bq", "cjbq", "bqxq"},
	"复合命令": {"kstj", "ycsh"},
}

// 检查是否是标准git命令
func isGitCommand(command string) bool {
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
