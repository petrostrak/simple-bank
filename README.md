#### Simple Bank
A backend application in GO that simulates simple bank functionality

##### To install golang migrate tool in debian
`https://github.com/golang-migrate/migrate/releases/tag/v4.15.1`

To create initial migrate files we run the following command

```
migrate create -ext sql -dir db/migration -seq init_schema
```

To run migrate up and make changes to the postgres DB run

```
migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

To run migrate down and make changes to the postgres DB run

```
migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
```

#### SQLC
SQLC generates Go-code from the SQL.
* You write SQL queries
* You run sqlc to generate code that presents type-safe interfaces to those queries
* You write application code that calls the methods sqlc generated.

To install sqlc

```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

To initialize sqlc

```
sqlc init
```

To generate Go-code from sql

```
sqlc generate
```

#### Viper
Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

To instal viper

```
go get github.com/spf13/viper
```

It supports:

* setting defaults
* reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
* live watching and re-reading of config files (optional)
* reading from environment variables
* reading from remote config systems (etcd or Consul), and watching changes
* reading from command line flags
* reading from buffer
* setting explicit values

#### Gomock
gomock is a mocking framework for the Go programming language.

To instal gomock

```
go get github.com/spf13/viper
```

To generate code
* first we choose the destination path 
* then, the path that the interface exists
specifying the name
```
mockgen -destination db/mock/store.go github.com/petrostrak/simple-bank/db/sqlc Store
```