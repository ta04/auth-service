module github.com/SleepingNext/auth-service

go 1.13

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.1.0
	github.com/G0tYou/user-service v0.0.0-20200406111304-e1323d4a0a54
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.4
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
)

// +heroku goVersion go1.13.8
