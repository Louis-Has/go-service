# go-service

```bash
protoc --go_out=./proto/. \
  --go-grpc_out=./proto/. \
  proto/hello.proto

```

model

```bash
goctl model mysql datasource \
  -url="root:development@tcp(localhost:3306)/testDB" -t="*" \
  -i created_at,updated_at \
  -d=./internal/model
```

rpc

```bash
goctl rpc protoc ./service/*.proto \
--go_out=./service/pb --go-grpc_out=./service/pb \
--zrpc_out=./service/article -m
```

gateway

```bash
protoc --descriptor_set_out=gateway/article.pb service/article.proto
```


api

```bash
 goctl api go -api  ./restful/*.api --dir ./restful/art
```

