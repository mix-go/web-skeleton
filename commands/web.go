package commands

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mix-go/dotenv"
	"github.com/mix-go/web-skeleton/globals"
	"github.com/mix-go/web-skeleton/routes"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xcli/flag"
	"github.com/mix-go/xcli/process"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type WebCommand struct {
}

func (t *WebCommand) Main() {
	if flag.Match("d", "daemon").Bool() {
		process.Daemon()
	}

	logger := globals.Logrus()
	addr := dotenv.Getenv("GIN_ADDR").String(":8080")
	mode := dotenv.Getenv("GIN_MODE").String(gin.ReleaseMode)

	// server
	gin.SetMode(mode)
	router := gin.New()
	routes.SetRoutes(router)
	srv := &http.Server{
		Addr:    flag.Match("a", "addr").String(addr),
		Handler: router,
	}
	globals.Server = srv

	// signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("Server shutdown")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf("Server shutdown error: %s", err)
		}
	}()

	// logger
	if mode != gin.ReleaseMode {
		handlerFunc := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s|%s|%d|%s",
					params.Method,
					params.Path,
					params.StatusCode,
					params.ClientIP,
				)
			},
			Output: logger.Out,
		})
		router.Use(handlerFunc)
	}

	// templates
	router.LoadHTMLGlob(fmt.Sprintf("%s/../templates/*", xcli.App().BasePath))

	// static file
	router.Static("/static", fmt.Sprintf("%s/../public/static", xcli.App().BasePath))
	router.StaticFile("/favicon.ico", fmt.Sprintf("%s/../public/favicon.ico", xcli.App().BasePath))

	// run
	welcome()
	logger.Infof("Server start at %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
		panic(err)
	}
}
