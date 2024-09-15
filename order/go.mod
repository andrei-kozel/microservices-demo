module github.com/andrei-kozel/microservices-demo/order

go 1.23.1

require (
	github.com/andrei-kozel/microservices-demo-proto/golang/order v1.0.12
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
)

require github.com/andrei-kozel/microservices-demo-proto/golang/payment v1.0.12

require github.com/joho/godotenv v1.5.1

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.66.2
	google.golang.org/protobuf v1.34.2 // indirect
)
