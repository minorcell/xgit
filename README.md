# xgit - 中文拼音首字母的 Git 命令工具

xgit 是基于 git 的扩展。使用中文拼音首字母作为命令，让中文用户更直观地使用 Git，同时简化了一些繁杂的 Git 命令。

### 基础命令映射

```bash
# 仓库操作
xgit kl <url>           # 克隆 (ke long) → git clone
xgit csh                # 初始化 (chu shi hua) → git init

# 文件操作
xgit tj [file]          # 添加 (tian jia) → git add
xgit tj -m "msg"        # 提交 (ti jiao) → git commit -m
xgit ch <file>          # 撤回 (che hui) → git checkout -- <file>

# 分支操作
xgit fz                 # 分支列表 (fen zhi) → git branch
xgit cj <branch>        # 创建分支 (chuang jian) → git checkout -b
xgit qh <branch>        # 切换分支 (qie huan) → git checkout

# 远程操作
xgit ts                 # 推送 (tui song) → git push
xgit lq                 # 拉取 (la qu) → git pull
xgit hq                 # 获取 (huo qu) → git fetch

# 高级操作
xgit hb <branch>        # 合并 (he bing) → git merge
xgit zf <branch>        # 整合 (zheng he) → git rebase
xgit ht                 # 回退 (hui tui) → git reset
```

### 复合命令简化

```bash
xgit kstj "msg"         # 快速提交 → git add . && git commit -m && git push
xgit tbfz <branch>      # 同步分支 → git fetch && git checkout && git pull
```

### 原生支持

```bash
xgit git <any-git-cmd>  # 原生git命令支持
xgit commit -m "xxx"    # 直接使用git原命令
```

### 帮助系统

```bash
xgit bz                 # 帮助 (bang zhu) - 显示所有命令
xgit bz <cmd>          # 显示具体命令用法
xgit bz --git <cmd>    # 显示对应的git命令
xgit bz --list         # 显示完整映射表
```

## 📄 许可证

[MIT License](LICENSE)
