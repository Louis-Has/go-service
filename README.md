# go-service

```bash
protoc --go_out=./proto/. \
  --go-grpc_out=./proto/. \
  proto/hello.proto

```

```bash
goctl model mysql datasource \
  -url="root:development@tcp(localhost:3306)/testDB" -t="*" \
  -i created_at,updated_at \
  -d=./internal/model
```
```bash
goctl rpc protoc ./service/*.proto \
--go_out=./service/pb --go-grpc_out=./service/pb \
--zrpc_out=./service/article
```

```bash
 goctl api go -api  ./restful/*.api --dir ./restful/art
```

