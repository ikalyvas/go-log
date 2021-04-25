package utils

import "go.uber.org/zap"

func CustomLogger() *zap.SugaredLogger {

	logger,_ := zap.NewDevelopment()
	defer logger.Sync()
	slog := logger.Sugar()
	return slog

}
