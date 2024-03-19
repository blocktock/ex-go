package app

import (
	"context"
	"errors"
	"ex-go/internal/config"
	"fmt"
	"github.com/blocktock/go-pkg/load"
	"github.com/blocktock/go-pkg/log"
	"github.com/blocktock/go-pkg/nobug"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"))

type App struct {
	Engine *gin.Engine
}

func Start(ctx context.Context) {

	load.FileConfig(config.C, "./conf/config.toml")

	//init logger
	log.InitLogger(config.C.Log.OutputPath)

	injector, injectorCleanFunc, err := InitServer()
	if err != nil {
		fmt.Printf("init failed, err: %v", err)
		return
	}

	cfg := config.C.HTTP

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: injector.Engine,
	}

	nobug.Nobug()

	quit := make(chan os.Signal, 1)
	// When received, the signal will be sent to the quit channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		injectorCleanFunc()

		ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)
		defer cancelCtx()

		if err = srv.Shutdown(ctx); err != nil {
			fmt.Println("shutdown error ", err)
		}

		<-ctx.Done()
		fmt.Println("Server exiting")
	}()

	// Start the server
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("listen: %s\n", err)
	}
}
