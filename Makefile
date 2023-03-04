DB_URL=postgresql://root:secret@localhost:5432/lpotl_go_dev?sslmode=disable

.PHONY: network
network:
	docker network create lpotl-network

.PHONY: postgres
postgres:
	docker run --name postgres-lpotl --network lpotl-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

.PHONY: createdb
createdb:
	docker exec -it postgres-lpotl createdb --username=root --owner=root lpotl_go_dev

.PHONY: dropdb
dropdb:
	docker exec -it postgres-lpotl dropdb lpotl_go_dev

.PHONY: migrateup
migrateup:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose down

.PHONY: db_schema
db_schema:
	dbml2sql --postgres -o postgres/schema.sql postgres/db.dbml

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: server
server:
	go run cmd/lpotl-go/main.go

.PHONY: mock
mock:
	mockgen -package mockdb -destination mock/store.go github.com/earlofurl/lpotl-go/db/sqlc Store

.PHONY: redis
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: seed
seed:
	go run postgres/seeder/seeder.go
