module github.com/huhaophp/eblog

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.6.1
	github.com/go-ini/ini v1.54.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.8 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.3 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/tools v0.0.0-20200325010219-a49f79bcc224 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.54.0 // indirect
)

replace (
	github.com/huhaophp/eblog/conf => /Users/huhao/goworkspace/eblog/conf
	github.com/huhaophp/eblog/controllers/admin => /Users/huhao/goworkspace/eblog/controllers/admin
	github.com/huhaophp/eblog/controllers/api => /Users/huhao/goworkspace/eblog/controllers/api
	github.com/huhaophp/eblog/docs => /Users/huhao/goworkspace/eblog/docs
	github.com/huhaophp/eblog/middleware/jwt => /Users/huhao/goworkspace/eblog/middleware/jwt
	github.com/huhaophp/eblog/models => /Users/huhao/goworkspace/eblog/models
	github.com/huhaophp/eblog/pkg/setting => /Users/huhao/goworkspace/eblog/pkg/setting //pkg/setting
	github.com/huhaophp/eblog/pkg/util => /Users/huhao/goworkspace/eblog/pkg/util
	github.com/huhaophp/eblog/request => /Users/huhao/goworkspace/eblog/request
	github.com/huhaophp/eblog/routers => /Users/huhao/goworkspace/eblog/routers
)
