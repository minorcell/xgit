# xgit - 中文拼音首字母的 Git 命令工具

xgit 不是 Git 的创新，而是基于 git 的扩展。使用中文拼音首字母作为命令，让中文用户更直观地使用 Git，同时简化了一些繁杂的 Git 命令。

## ✨ 特性

- **中文友好** - 使用拼音首字母，如 `kl`(克隆)、`tj`(提交)、`ts`(推送)
- **输入高效** - 2 个字母通常比原 git 命令更短
- **完全兼容** - 支持原生 git 命令，无学习负担
- **复合命令** - `kstj`一键完成 add + commit + push

## 📦 安装

### 方式一：使用 Makefile（推荐）

```bash
# 构建
make build

# 安装到系统 PATH（需要管理员权限）
make install
```

### 方式二：手动构建

```bash
go build -o xgit .
```

## 快速开始

```bash
# 查看所有命令
xgit bz

# 克隆仓库
xgit kl https://github.com/user/repo.git

# 快速提交（add + commit + push）
xgit kstj "提交信息"

# 查看状态
xgit zt

# 查看日志
xgit rz
```

## 📖 命令列表

### 仓库操作

- `kl <url>` - 克隆仓库 (ke long)
- `csh` - 初始化仓库 (chu shi hua)

### 文件操作

- `tja [file]` - 添加文件 (tian jia)
- `tj -m "msg"` - 提交更改 (ti jiao)
- `ch <file>` - 撤回文件 (che hui)

### 分支操作

- `fz` - 查看分支 (fen zhi)
- `fzxq` - 分支详情 (fen zhi xiang qing)
- `ycfz` - 远程分支 (yuan cheng fen zhi)
- `cjfz <branch>` - 创建分支 (chuang jian fen zhi)
- `qhfz <branch>` - 切换分支 (qie huan fen zhi)

### 远程操作

- `ts` - 推送代码 (tui song)
- `lq` - 拉取代码 (la qu)
- `hq` - 获取更新 (huo qu)

### 远程仓库管理

- `ycck` - 查看远程仓库 (yuan cheng cha kan)
- `yctz <name> <url>` - 添加远程仓库 (yuan cheng tian jia)
- `ycsc <name>` - 删除远程仓库 (yuan cheng shan chu)
- `yczm <old> <new>` - 重命名远程仓库 (yuan cheng zhong ming)
- `ycxg <name> <url>` - 修改远程 URL (yuan cheng xiu gai)
- `ycxq <name>` - 远程仓库详情 (yuan cheng xiang qing)

### 高级操作

- `hb <branch>` - 合并分支 (he bing)
- `zf <branch>` - 整合分支 (zheng he)
- `ht [commit]` - 回退版本 (hui tui)

### 日志和状态

- `rz` - 查看日志 (ri zhi)
- `yhrz` - 一行日志 (yi hang ri zhi)
- `zt` - 状态 (zhuang tai)
- `ztxq` - 状态详情 (zhuang tai xiang qing)

### 标签操作

- `bq` - 标签列表 (biao qian)
- `cjbq <tag> -m "msg"` - 创建标签 (chuang jian biao qian)
- `bqxq` - 标签详情 (biao qian xiang qing)

### 复合命令

- `kstj "msg"` - 快速提交 (kuai su ti jiao) → `git add . && git commit -m && git push`
- `ycsh <url> [branch]` - 远程设置 (yuan cheng she zhi) → `git remote add origin <url> && git push -u origin main`

## 🔍 帮助系统

```bash
# 查看所有命令
xgit bz

# 查看特定命令用法
xgit bz kl

# 查看git等价命令
xgit bz --git kl
```

## 🔄 原生 Git 支持

xgit 完全兼容原生 git 命令：

```bash
# 直接使用git命令
xgit git status
xgit commit -m "message"
xgit push origin main

# 或者省略git前缀
xgit status
xgit log --oneline
```

## 🧪 开发和测试

### 运行测试

```bash
# 运行所有测试
make test

# 运行基准测试
go test -bench=. -benchmem

# 开发模式测试
make dev
```

### 项目结构

```
xgit/
├── main.go              # 主程序入口
├── commands.go          # 命令映射和定义
├── help.go              # 帮助系统
├── executor.go          # 命令执行器
├── Makefile             # 构建管理
├── *_test.go           # 测试文件 (4个)
└── docs/               # 设计文档
```

### 性能指标

- 命令查找: ~26ns, 0 内存分配
- 24 个命令，100%测试覆盖
- 987 行测试代码

## 💡 设计理念

### 拼音首字母方案

- `kl` (克隆) vs `clone` - 更直观
- `tj` (提交) vs `commit` - 更好记
- `ts` (推送) vs `push` - 更高效

### 渐进式采用

- 保持与 git 生态完全兼容
- 团队可以逐步迁移
- 支持混合使用

## 📝 常用工作流

### 初始化新项目

```bash
# 初始化仓库
xgit csh

# 添加文件并提交
xgit tja .
xgit tj -m "initial commit"

# 设置远程仓库并推送
xgit ycsh https://github.com/user/repo.git

# 或者分步操作：
# xgit yctz origin https://github.com/user/repo.git
# xgit ts -u origin main
```

### 日常开发

```bash
# 克隆项目
xgit kl https://github.com/user/repo.git

# 创建功能分支
xgit cjfz feature-branch

# 开发...

# 快速提交
xgit kstj "feat: 添加新功能"

# 切回主分支
xgit qhfz main

# 合并功能分支
xgit hb feature-branch
```

### 查看信息

```bash
# 查看状态
xgit zt

# 查看简洁日志
xgit yhrz

# 查看分支
xgit fz

# 查看远程分支
xgit ycfz

# 查看远程仓库
xgit ycck

# 查看远程仓库详情
xgit ycxq origin
```

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

### 开发环境

```bash
# 克隆仓库
git clone <repo-url>
cd xgit

# 构建
make build

# 运行测试
make test

# 开发模式
make dev
```

## 📄 许可证

[MIT License](LICENSE)

---

**让 Git 更中文，让开发更高效！** 🚀
