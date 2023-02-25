# go-clean-arch

## Changelog

- **v1**: checkout to the [v1 branch](https://github.com/bxcodec/go-clean-arch/tree/v1) <br>
  Proposed on 2017, archived to v1 branch on 2018 <br>
  Desc: Initial proposal by me. The story can be read here: https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

- **v2**: checkout to the [v2 branch](https://github.com/bxcodec/go-clean-arch/tree/v2) <br>
  Proposed on 2018, archived to v2 branch on 2020 <br>
  Desc: Improvement from v1. The story can be read here: https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf

- **v3**: master branch <br>
  Proposed on 2019, merged to master on 2020. <br>
  Desc: Introducing Domain package, the details can be seen on this PR [#21](https://github.com/bxcodec/go-clean-arch/pull/21)

## How to run
- go mod init
- go mod tidy
- go mod vendor
- go run app/main.go

## Use migration
### Create Migration
- migrate create -ext sql -dir db/migrations -seq create_users_table
### Run Migration
- migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
example up:
- migrate -database "mysql://root:#d4esUqzQpS9XZNv@tcp(localhost:3306)/article" -path db/migrations up
example down:
- migrate -database "mysql://root:#d4esUqzQpS9XZNv@tcp(localhost:3306)/article" -path db/migrations down
### Detail YOUR_DATABASE_URL
["mysql://root:secret@tcp(localhost:3306)/simple_bank"]
- We’re using mysql, so the driver name is mysql.
- Then the username is root
- The password is secret
- The address is localhost, port 3306.
- And the database name is simple_bank.

## Use Mockery
### Install
- go install github.com/vektra/mockery/v2@v2.20.0
### Run Mockery
- cd domain
- mockery --all --case=camel