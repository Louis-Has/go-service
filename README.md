# go-service

```bash
protoc --go_out=./proto/. \
  --go-grpc_out=./proto/. \
  proto/hello.proto

```

```bash
 goctl api go -api  ./restful/*.api --dir ./restful/art
```

```bash
goctl model mysql datasource \
  -url="root:development@tcp(localhost:3306)/testDB" -t="*" \
  -d=./internal/model

```