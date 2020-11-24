package commands

import (
    "context"
    "fmt"
    "github.com/mix-go/console"
    "github.com/mix-go/console/flag"
    "github.com/mix-go/dotenv"
    "github.com/mix-go/gin"
    "github.com/mix-go/web-skeleton/globals"
    "github.com/mix-go/web-skeleton/routes"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
)

const Addr = ":8080"

type WebCommand struct {
}

func (t *WebCommand) Main() {
    logger := globals.Logger()
    mode := dotenv.Getenv("GIN_MODE").String(gin.ReleaseMode)

    // server
    gin.SetMode(mode)
    router := gin.New(routes.WebDefinition())
    srv := &http.Server{
        Addr:    flag.Match("a", "addr").String(Addr),
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
        router.Use(gin.LoggerWithFormatter(logger, func(params gin.LogFormatterParams) string {
            return fmt.Sprintf("%s|%s|%d|%s",
                params.Method,
                params.Path,
                params.StatusCode,
                params.ClientIP,
            )
        }))
    }

    // templates
    router.LoadHTMLGlob(fmt.Sprintf("%s/../templates/*", console.App.BasePath))

    // static file
    router.Static("/static", fmt.Sprintf("%s/../public/static", console.App.BasePath))
    router.StaticFile("/favicon.ico", fmt.Sprintf("%s/../public/favicon.ico", console.App.BasePath))

    // run
    welcome()
    logger.Infof("Server start at %s", srv.Addr)
    if err := srv.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
        panic(err)
    }
}
