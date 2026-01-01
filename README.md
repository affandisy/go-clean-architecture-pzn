# belajar clean architecture dari pzn

```bash
go run cmd/web/main.go
```

### Membuat Migrasi

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Menjalankan Migrasi
```shell
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```