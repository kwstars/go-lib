package demo10

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

var SugarLogger *zap.SugaredLogger

//func main() {
//	InitLogger()
//	defer SugarLogger.Sync()
//	SimpleHttpGet("www.baidu.com")
//	SimpleHttpGet("http://www.baidu.com")
//}

func InitLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	SugarLogger = logger.Sugar()
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = SyslogTimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeLevel = CustomLevelEncoder
	encoderConfig.FunctionKey = "function"
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: "./test.log",
		//MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     24,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func SimpleHttpGet(url string) {
	SugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
		//SugarLogger.Errorw("Error fetching URL",
		//	"uid", 123455,
		//	"test", "test",
		//	"asdf", 12312,
		//	)
		zap.S().Errorw("Error fetching URL",
			"uid", 123455,
			"test", "test",
			"asdf", 12312,
			)
	} else {
		SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
