package main

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

func worker(wg *sync.WaitGroup, logger *log.Logger, workerID int) {
	defer wg.Done()

	if workerID%2 == 0 {
		for {
			logger.Println("foobar std")
		}
	} else {
		for {
			if ce := zap.L().Check(zap.DebugLevel, "foobar zap"); ce != nil {
				ce.Write(zap.Int("intfield", 200), zap.String("onefield", "string"))
				//ce.Write(zap.String("onefield", "string"))
			}
		}
	}
}

func main() {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.OutputPaths = []string{"stdout"}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(zapLogger)

	logger, err := zap.NewStdLogAt(
		zap.L().With(zap.String("package", "foobar")),
		zap.DebugLevel,
	)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 63; i++ {
		wg.Add(1)
		go worker(&wg, logger, i)
	}

	wg.Wait()
}
