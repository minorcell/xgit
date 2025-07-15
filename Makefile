# xgit Makefile

BINARY_NAME=xgit
MAIN_PACKAGE=.

.PHONY: build clean install test help

# 默认目标
all: build

# 构建二进制文件
build: check-config
	@echo "构建 xgit..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "构建完成: ./$(BINARY_NAME)"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -f $(BINARY_NAME)
	go clean

# 检查JSON配置文件
check-config:
	@echo "检查配置文件..."
	@if [ ! -f commands.json ]; then \
		echo "错误: commands.json 配置文件不存在"; \
		exit 1; \
	fi
	@echo "配置文件检查通过"

# 安装到系统PATH（需要管理员权限）
install: build
	@echo "安装 xgit 到 /usr/local/bin..."
	sudo cp $(BINARY_NAME) /usr/local/bin/
	sudo cp commands.json /usr/local/bin/
	@echo "安装完成，现在可以在任意位置使用 xgit 命令"

# 运行测试
test:
	@echo "运行测试..."
	go test -v ./...

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./...

# 检查代码
vet:
	@echo "检查代码..."
	go vet ./...

# 开发模式：构建并运行帮助
dev: build
	@echo "开发模式测试："
	./$(BINARY_NAME) bz

# 显示帮助
help:
	@echo "xgit 构建工具"
	@echo ""
	@echo "可用命令:"
	@echo "  make build        - 构建 xgit 二进制文件"
	@echo "  make clean        - 清理构建文件"
	@echo "  make install      - 安装到系统 PATH"
	@echo "  make test         - 运行测试"
	@echo "  make fmt          - 格式化代码"
	@echo "  make vet          - 检查代码"
	@echo "  make dev          - 开发模式测试"
	@echo "  make check-config - 检查JSON配置文件"
	@echo "  make help         - 显示此帮助信息" 