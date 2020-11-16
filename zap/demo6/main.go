package main

import "go.uber.org/zap"

func main() {
	logger1 := zap.NewExample()
	defer logger1.Sync()

	logger2, _ := zap.NewDevelopment()
	defer logger2.Sync()

	logger3, _ := zap.NewProduction()
	defer logger3.Sync()


	logger1.Info("hello world1")
	logger2.Info("hello world2")
	logger3.Info("hello world3")
}
