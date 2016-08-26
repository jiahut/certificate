## https和双向认证


> openssl genrsa -out ca.key 2048
> openssl req -x509 -new -nodes -key ca.key -subj "/CN=zhangzhijia.io" -days 5000 -out ca.crt

> openssl genrsa -out server.key 2048
> openssl req -new -key server.key -subj "/CN=server.zhangzhijia.io"  -out server.csr
> openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

> cd https
> go run server.go
> go run client.go
> curl --cacert ca.crt https://server.zhangzhijia.io:8081


> openssl genrsa -out client.key 2048
> openssl req -new -key client.key -subj "/CN=client.zhangzhijia.io" -out client.csr
> echo extendedKeyUsage=clientAuth > extfile.conf
> openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile extfile.conf -out client.crt -days 5000

> cd two-way
> go run server.go
> go run client.go

