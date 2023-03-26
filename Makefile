createschema:
	migrate create -ext sql -dir db/migration -seq init_schema_db
migrateup:
	migrate -path db/migration -database "postgres://spsuserone:@Th1nkIT@localhost:5432/sps1?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://spsuserone:@Th1nkIT@localhost:5432/sps1?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/codeninjaug/authexample/db/sqlc Store

PHONY:	createdb migrateup migratedown sqlc test mock