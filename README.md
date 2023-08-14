<p align="center">
  <br>
  <img width="300" src="./assets/images/logo.png" alt="repo logo">
  <br>
</p>

# gotes
disclaimer: pet project

gotes - it's like a notes on Go

For begin you can create `.env.local`  to set your configure at `configs/.`

In order to `up` migration run command:
```
go run cmd/migration/main.go -c "up"
```

For `down`:
```
go run cmd/migration/main.go -c "down"
```

For `create` migration with `any_name` run:
```
go run cmd/migration/main.go -c "create" -args "any_namne sql"
```

Run tests:
```
make test
```

Run build:
```
make
```