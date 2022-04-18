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
mockgen -package mockdb  -destination db/mock/store.go github.com/petrostrak/simple-bank/db/sqlc Store
```

###### NOTE: Due to recent changes in Go tooling, to make mockgen work, include an empty import of that package so it gets pulled. `import _ "github.com/golang/mock/mockgen/model"`

#### PASETO: Platform-Agnostic Security Tokens
PASETO (Platform-Agnostic SEcurity TOkens) is a specification and reference implementation for secure stateless tokens.

Unlike JSON Web Tokens (JWT), which gives developers more than enough rope with which to hang themselves, PASETO only allows secure operations. JWT gives you "algorithm agility", PASETO gives you "versioned protocols". It's incredibly unlikely that you'll be able to use PASETO in an [insecure way](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/).
###### Caution: Neither JWT nor PASETO were designed for stateless session management. PASETO is suitable for tamper-proof cookies, but cannot prevent replay attacks by itself.

To install Paseto
```
github.com/o1egl/paseto
```