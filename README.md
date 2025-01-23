folder organization
https://gist.github.com/ayoubzulfiqar/9f1a34049332711fddd4d4b2bfd46096

Migrate database
```sh
migrate -path internal/database/migration -database "postgres://postgres:example@0.0.0.0:5432/jubawink?sslmode=disable" -verbose up
```

Create new migration file
```sh
migrate create -ext sql -dir internal/database/migration -seq create_users_table
```

generate swagger docs
```sh
swag init -g ./cmd/jubawink/main.go -o ./docs
```

build project
```sh
make build
```

genarate swagger docs and run project
```sh
make
```

