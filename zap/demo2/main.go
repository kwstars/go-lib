package main

import "go.uber.org/zap"

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)

	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}
