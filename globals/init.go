package globals

import (
    "fmt"
    "github.com/go-session/redis"
    "github.com/go-session/session"
    "github.com/mix-go/console"
    "github.com/mix-go/dotenv"
    "github.com/mix-go/logrus"
    logrus2 "github.com/sirupsen/logrus"
    "io"
    "os"
    "time"
)

func Init() {
    // logger
    logger := Logger()
    file := logrus.NewFileWriter(fmt.Sprintf("%s/../runtime/logs/mix.log", console.App.BasePath), 7, 0)
    writer := io.MultiWriter(os.Stdout, file)
    logger.SetOutput(writer)
    if console.App.Debug {
        logger.SetLevel(logrus2.DebugLevel)
    }

    // db
    db := DB()
    db.SetLogger(logger)
    db.LogMode(console.App.Debug)

    // session
    session.InitManager(
        session.SetStore(redis.NewRedisStore(&redis.Options{
            Addr:        dotenv.Getenv("REDIS_ADDR").String(),
            Password:    dotenv.Getenv("REDIS_PASSWORD").String(),
            DB:          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
            DialTimeout: time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
        })),
    )
}
