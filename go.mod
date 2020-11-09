module github.com/mix-go/web-skeleton

go 1.13

replace (
	github.com/mix-go/bean => ../mix/src/bean
	github.com/mix-go/console => ../mix/src/console
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/event => ../mix/src/event
	github.com/mix-go/gin => ../mix/src/gin
	github.com/mix-go/logrus => ../mix/src/logrus
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/go-session/redis v3.0.1+incompatible
	github.com/go-session/session v3.1.2+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/configor v1.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mix-go/bean v1.0.21
	github.com/mix-go/console v1.0.21
	github.com/mix-go/dotenv v1.0.21
	github.com/mix-go/event v1.0.21
	github.com/mix-go/gin v1.0.21
	github.com/mix-go/logrus v1.0.21
	github.com/sirupsen/logrus v1.6.0
	github.com/ugorji/go v1.1.13 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	gorm.io/gorm v0.2.34 // indirect
)
