# 设置 Go 编译器
GO := go

# 设置目标二进制文件名
BINARY_NAME := hot-reload.bin

# 设置构建目录
BUILD_DIR := build

# 设置源文件
SRC := src/main.go

# 默认目标
all: $(BUILD_DIR)/$(BINARY_NAME)

# 创建构建目录
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# 编译二进制文件
$(BUILD_DIR)/$(BINARY_NAME): $(SRC) | $(BUILD_DIR)
	$(GO) build -o $@ $(SRC)

# 清理构建文件
clean:
	rm -rf $(BUILD_DIR)

# 运行程序
run: $(BUILD_DIR)/$(BINARY_NAME)
	./$(BUILD_DIR)/$(BINARY_NAME)

# 声明伪目标
.PHONY: all clean run
