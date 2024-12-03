package logger

import "go.uber.org/zap"

func (cfg *Config) InitializeGlobally() *zap.Logger {
	var (
		lg  *zap.Logger
		err error
	)

	if cfg.DevMode {
		lg, err = zap.NewDevelopment()
	} else {
		lg, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(lg)

	return lg
}
