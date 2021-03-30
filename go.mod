module github.com/mix-go/web-skeleton

go 1.13

replace (
	github.com/mix-go/dotenv => ../mix/src/dotenv
	github.com/mix-go/xcli => ../mix/src/xcli
	github.com/mix-go/xdi => ../mix/src/xdi
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-session/redis v3.0.1+incompatible
	github.com/go-session/session v3.1.2+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/configor v1.2.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mix-go/dotenv v1.0.22
	github.com/mix-go/xcli v1.1.2
	github.com/mix-go/xdi v1.1.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/ugorji/go v1.1.13 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
