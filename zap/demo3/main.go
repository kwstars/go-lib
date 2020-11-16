package main

import (
	"go.uber.org/zap"
)

func test (main *zap.Logger)  {
	// Additional calls to Named create a period-separated path.
	main.Named("subpackage2").Info("sub-logger2")
}

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	// By default, Loggers are unnamed.
	logger.Info("no name")

	// The first call to Named sets the Logger name.
	main := logger.Named("main")
	main.Info("main logger")

	// Additional calls to Named create a period-separated path.
	main.Named("subpackage").Info("sub-logger")

	test(main)
}
