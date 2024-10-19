package logger

import "go.uber.org/zap"

func (cfg *Config) InitializeGlobally() *zap.Logger {
	var lg *zap.Logger

	if cfg.DevMode {
		lg, _ = zap.NewDevelopment()
	} else {
		lg, _ = zap.NewProduction()
	}

	zap.ReplaceGlobals(lg)

	return lg
}
