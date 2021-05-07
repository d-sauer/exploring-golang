package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"time"
)

const MAX = 100

func main() {

	undo := initLogger()
	defer undo()

	fmt.Println("Start looping...")
	for i := 0; i < MAX; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Count: ", i, " of ", MAX)

		zap.S().Info("Test info message i:", i)
		zap.S().Debug("Test debug message i:", i)
	}

}

func initLogger() func() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		nanos := t.UnixNano()
		millis := nanos / int64(time.Millisecond)
		enc.AppendInt64(millis)
	}
	//encoder := zapcore.NewJSONEncoder(config)
	encoder := zapcore.NewConsoleEncoder(config)
	atom := zap.NewAtomicLevel()
	logr := zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atom))

	mux := http.NewServeMux()
	mux.Handle("/log_level", atom)
	go http.ListenAndServe(":1065", mux)

	return zap.ReplaceGlobals(logr)
}
