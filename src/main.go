package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	loveCmd *exec.Cmd
	// 添加防抖动相关变量
	debounceTimer *time.Timer
	debounceMutex sync.Mutex
	lastReload    time.Time
	// 防抖动时间配置（毫秒）
	debounceTime = getDebounceTime()
)

// 从环境变量获取防抖动时间，默认1000毫秒
func getDebounceTime() time.Duration {
	envTime := os.Getenv("LOVE2D_DEBOUNCE_TIME")
	if envTime == "" {
		return 3000 * time.Millisecond
	}
	
	timeMs, err := strconv.Atoi(envTime)
	if err != nil {
		log.Printf("Invalid debounce time in environment variable, using default 3000ms: %v", err)
		return 3000 * time.Millisecond
	}
	
	if timeMs < 100 {
		log.Printf("Debounce time too small, using minimum 100ms")
		return 100 * time.Millisecond
	}
	
	return time.Duration(timeMs) * time.Millisecond
}

func main() {
	// 创建文件监控器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}
	defer watcher.Close()

	// 设置要监控的目录（LÖVE2D 项目根目录）
	projectDir := "." // 当前目录，可改为具体路径
	err = filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Error walking directory:", err)
	}

	// 初次启动 LÖVE2D
	startLove()

	// 处理文件变化事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 忽略临时文件和非目标文件
				if shouldIgnore(event.Name) {
					continue
				}
				fmt.Println("Detected change in:", event.Name)
				// 使用防抖动机制重启
				debouncedRestartLove()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	// 保持程序运行
	select {}
}

// 启动 LÖVE2D
func startLove() {
	loveCmd = exec.Command("/Applications/love.app/Contents/MacOS/love", ".")
	loveCmd.Stdout = os.Stdout
	loveCmd.Stderr = os.Stderr
	if err := loveCmd.Start(); err != nil {
		log.Println("Error starting LÖVE2D:", err)
	}
	fmt.Println("LÖVE2D started")
}

// 重启 LÖVE2D
func restartLove() {
	if loveCmd != nil && loveCmd.Process != nil {
		loveCmd.Process.Kill() // 杀死当前进程
		loveCmd.Wait()         // 等待进程结束
	}
	// 短暂延迟以确保进程完全退出
	time.Sleep(100 * time.Millisecond)
	startLove()
}

// 防抖动重启函数
func debouncedRestartLove() {
	debounceMutex.Lock()
	defer debounceMutex.Unlock()

	// 如果已经有定时器在运行，重置它
	if debounceTimer != nil {
		debounceTimer.Stop()
	}

	// 设置新的定时器
	debounceTimer = time.AfterFunc(debounceTime, func() {
		debounceMutex.Lock()
		lastReload = time.Now()
		debounceMutex.Unlock()
		restartLove()
	})
}

// 忽略不相关的文件
func shouldIgnore(file string) bool {
	// 忽略临时文件、隐藏文件等
	ignoredExts := []string{".swp", ".tmp", ".log"}
	ignoredPrefixes := []string{".", "_"}
	ignoredDirs := []string{".git"}
	
	// 检查文件扩展名
	for _, ext := range ignoredExts {
		if filepath.Ext(file) == ext {
			return true
		}
	}
	
	// 检查文件前缀
	for _, prefix := range ignoredPrefixes {
		if filepath.Base(file)[0:1] == prefix {
			return true
		}
	}
	
	// 检查目录
	for _, dir := range ignoredDirs {
		if strings.Contains(file, "/"+dir+"/") || strings.HasPrefix(file, dir+"/") {
			return true
		}
	}
	
	return false
}