module github.com/huhaophp/eblog

go 1.13

require (
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.54.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.54.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/huhaophp/eblog/conf => /Users/huhao/goworkspace/eblog/conf
	github.com/huhaophp/eblog/controllers/admin => /Users/huhao/goworkspace/eblog/controllers/admin
	github.com/huhaophp/eblog/controllers/api => /Users/huhao/goworkspace/eblog/controllers/api
	github.com/huhaophp/eblog/middleware => /Users/huhao/goworkspace/eblog/middleware
	github.com/huhaophp/eblog/models => /Users/huhao/goworkspace/eblog/models
	github.com/huhaophp/eblog/pkg/setting => /Users/huhao/goworkspace/eblog/pkg/setting //pkg/setting
	github.com/huhaophp/eblog/pkg/util => /Users/huhao/goworkspace/eblog/pkg/util
	github.com/huhaophp/eblog/request => /Users/huhao/goworkspace/eblog/request
	github.com/huhaophp/eblog/routers => /Users/huhao/goworkspace/eblog/routers
)
