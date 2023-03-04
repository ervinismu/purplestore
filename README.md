# Purplestore (WIP)

## Deps :
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [golangci-lint](https://golangci-lint.run/)
- [migrate](https://github.com/golang-migrate/migrate)
- [gotenv](https://github.com/subosito/gotenv)
- postgres

## Pre
- install golangci-lint, follow the instruction in [this link here!](https://golangci-lint.run/usage/install/)
- Running linters `golangci-lint`

## Run
- Run application `go run main.go`, will running in port `8080`
- Run linter `golangci-lint run`

### About DB Migration
This project using [migrate](https://github.com/golang-migrate/migrate). For now, you need to install it locally and running for trigger `create migration`, `up migration` and `down migration` :
- set environment variable : `export POSTGRESQL_URL='postgres://youruser:yourpassword@localhost:5433/purpledb?sslmode=disable'`
- create migration : `migrate create -ext sql -dir db/migrations -seq migration_name`
- run migration up : `migrate -database ${POSTGRESQL_URL} -path db/migrations up`
- run migration down : `migrate -database ${POSTGRESQL_URL} -path db/migrations down`

## References
- https://www.calhoun.io/using-mvc-to-structure-go-web-applications/
- https://go.dev/doc/tutorial/web-service-gin
- https://peter.bourgon.org/go-in-production/
