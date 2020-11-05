package manifest

import (
    "github.com/mix-go/web-skeleton/manifest/beans"
    "github.com/mix-go/web-skeleton/manifest/commands"
    "github.com/mix-go/console"
    "github.com/mix-go/dotenv"
)

var (
    ApplicationDefinition console.ApplicationDefinition
)

func Init() {
    beans.Init()

    ApplicationDefinition = console.ApplicationDefinition{
        Name:    "app",
        Version: "1.0.0-alpha",
        Debug:   dotenv.Getenv("APP_DEBUG").Bool(),
        Beans:      beans.Beans,
        Commands:   commands.Commands,
    }
}
