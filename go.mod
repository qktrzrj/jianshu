module github.com/shyptr/hello-world-web

go 1.13

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.3.0
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/micro/go-micro/v2 v2.4.0
	github.com/rs/zerolog v1.18.0
	github.com/shyptr/graphql v1.0.0-beta.2
	github.com/shyptr/plugins v0.0.0-20200409063656-665bdab483ce
	github.com/shyptr/sqlex v1.4.11
	github.com/sony/sonyflake v1.0.0
	go.opencensus.io v0.22.3 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/tools v0.0.0-20200410194907-79a7a3126eef // indirect
	google.golang.org/grpc v1.28.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

replace github.com/shyptr/graphql => /Users/yan/GolandProjects/graphql

replace github.com/shyptr/sqlex => /Users/yan/GolandProjects/sqlex

replace github.com/shyptr/plugins => /Users/yan/GolandProjects/plugins
