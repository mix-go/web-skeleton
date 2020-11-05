package beans

import (
    "github.com/go-redis/redis/v8"
    "github.com/mix-go/bean"
    "github.com/mix-go/dotenv"
    "time"
)

func Redis() {
    Beans = append(Beans,
        bean.BeanDefinition{
            Name:            "redis",
            Reflect:         bean.NewReflect(redis.NewClient),
            Scope:           bean.SINGLETON,
            ConstructorArgs: bean.ConstructorArgs{bean.NewReference("redis-options")},
        },
        bean.BeanDefinition{
            Name:    "redis-options",
            Reflect: bean.NewReflect(redis.Options{}),
            Fields: bean.Fields{
                "Addr":        dotenv.Getenv("REDIS_ADDR").String(),
                "Password":    dotenv.Getenv("REDIS_PASSWORD").String(),
                "DB":          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
                "DialTimeout": time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
            },
        },
    )
}
