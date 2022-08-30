package logger

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"sync"
)

type initData struct {
	Dir      string `validate:"required"`
	File     string `validate:"required"`
	Encoding string `validate:"required"`
}

type Logger struct {
	*zap.SugaredLogger
}

var (
	once      sync.Once
	zapLogger *zap.Logger
	inD       *initData
)

func Init(dir, file, encoding string) *Logger {
	once.Do(func() {
		if err := validateDefaults(dir, file, encoding); err != nil {
			panic(fmt.Sprintf("validation of init params not passed. Err - %v", err))
		}
		err := os.MkdirAll(inD.Dir, 0744)
		if err != nil {
			panic(err)
		}
		_, err = os.OpenFile(getLogFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			panic(err)
		}
		devConf := zap.NewDevelopmentConfig()
		devConf.OutputPaths = []string{getLogFilePath()}
		devConf.ErrorOutputPaths = devConf.OutputPaths
		devConf.Encoding = inD.Encoding
		zapLogger, err = devConf.Build()
		if err != nil {
			panic(fmt.Sprintf("cant initialize logger. Error - %v", err))
		}
	})
	defer zapLogger.Sync()
	return &Logger{zapLogger.Sugar()}
}

func getLogFilePath() string {
	return fmt.Sprintf("%v/%v", inD.Dir, inD.File)
}

func validateDefaults(dir, file, encoding string) error {
	inD = &initData{
		Dir:      dir,
		File:     file,
		Encoding: encoding,
	}
	vv := validator.New()
	if err := vv.Struct(inD); err != nil {
		return err
	}
	return nil
}
