package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// --- 1. 定义 lumberjack 配置的命令行参数 ---
	filename := flag.String("filename", "app.log", "Log file name.")
	maxSize := flag.Int("max-size", 1000, "Maximum size in megabytes of the log file before it gets rotated.")
	maxAge := flag.Int("max-age", 30, "Maximum number of days to retain old log files.")
	maxBackups := flag.Int("max-backups", 10, "Maximum number of old log files to retain.")
	localTime := flag.Bool("local-time", true, "Use local time for formatting timestamps in backup files.")
	compress := flag.Bool("compress", false, "Compress rolled-over files using gzip.")

	flag.Parse()

	// --- 2. 配置 lumberjack Logger ---
	lumberjackLogger := &lumberjack.Logger{
		Filename:   *filename,
		MaxSize:    *maxSize,
		MaxAge:     *maxAge,
		MaxBackups: *maxBackups,
		LocalTime:  *localTime,
		Compress:   *compress,
	}

	// --- 3. 从 stdin 逐行读取并使用 fmt 写入日志 ---
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 直接使用 fmt.Fprintln 写入到 lumberjackLogger
		// 它返回写入的字节数和错误，这里我们暂时忽略
		_, _ = fmt.Fprintln(lumberjackLogger, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// 如果读取 stdin 出错，打印到标准错误并退出
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}
}
