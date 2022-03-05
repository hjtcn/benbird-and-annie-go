module github.com/hjtcn/benbird-and-annie-go

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/spf13/cobra v1.3.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/hjtcn/go-base-homework@v0.0.2 => ./module_three/practice/refactorFatRate
