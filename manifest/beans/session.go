package beans

import (
    "github.com/go-session/redis"
    "github.com/go-session/session"
    "github.com/mix-go/bean"
    "github.com/mix-go/dotenv"
    "time"
)

func Session() {
    Beans = append(Beans,
        bean.BeanDefinition{
            Name:            "session",
            Reflect:         bean.NewReflect(session.NewManager),
            Scope:           bean.SINGLETON,
            ConstructorArgs: bean.ConstructorArgs{bean.NewReference("session-option")},
        },
        bean.BeanDefinition{
            Name:    "session-option",
            Reflect: bean.NewReflect(session.SetStore),
            ConstructorArgs: bean.ConstructorArgs{
                redis.NewRedisStore(&redis.Options{
                    Addr:        dotenv.Getenv("REDIS_ADDR").String(),
                    Password:    dotenv.Getenv("REDIS_PASSWORD").String(),
                    DB:          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
                    DialTimeout: time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
                }),
            },
        },
    )
}
