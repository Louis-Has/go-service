# go-service

```bash
protoc --go_out=./proto/. \
  --go-grpc_out=./proto/. \
  proto/hello.proto

```

model datasource

```bash
goctl model mysql datasource \
  -url="root:development@tcp(localhost:3306)/testDB" \
  -t="order,order_detail,order_receive_mes,product,product_category,user_mes" \
  -i created_at,updated_at \
  -d=./internal/model -c
```

rpc & gateway : user_mes

```bash
goctl rpc protoc ./service/user_mes.proto \
--go_out=./internal/pb --go-grpc_out=./internal/pb \
--zrpc_out=./service/user_mes -m

protoc --descriptor_set_out=gateway/user_mes.pb service/user_mes.proto
```

rpc & gateway : product

```bash
goctl rpc protoc ./service/product.proto \
--go_out=./internal/pb --go-grpc_out=./internal/pb \
--zrpc_out=./service/product -m

protoc --descriptor_set_out=gateway/product.pb service/product.proto
```

rpc & gateway : order

```bash
goctl rpc protoc ./service/order.proto \
--go_out=./internal/pb --go-grpc_out=./internal/pb \
--zrpc_out=./service/order -m

protoc --descriptor_set_out=gateway/order.pb service/order.proto
```


api

```bash
 goctl api go -api  ./restful/*.api --dir ./restful/art
```

