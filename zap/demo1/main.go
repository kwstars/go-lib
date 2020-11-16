package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	// 使用zap的预设构造函数是最简单的方法封装，它不允许太多自定义。
	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://example.org/api"

	// 在大多数情况下，请使用SugaredLogger。它比大多数其他结构化日志记录包快4-10倍，并且具有熟悉的松散类型的API。
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", 2*time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	//在每微秒都很重要的异常情况下，请使用Logger记录器。它比SugaredLogger更快，但仅支持结构化日志记录。
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
