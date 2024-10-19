package http_transport

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func New(config *Config, router http.Handler, withoutCORS bool) func() error {
	if !withoutCORS {
		router = initializedCors.Handler(router)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		zap.L().Info(fmt.Sprintf("Server started on address: %s", addr))
		return http.ListenAndServe(addr, router)
	})
	g.Go(func() error {
		<-gCtx.Done()
		zap.L().Info("Graceful shutdown triggered")
		//return http.Shutdown(context.Background())
		//TODO добавить грейсфул закрытие
		return nil
	})

	if err := g.Wait(); err != nil {
		zap.L().Error(fmt.Sprintf("exit reason: %s \n", err))
	}

	return nil
}
