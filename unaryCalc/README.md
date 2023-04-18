# First model
## simple greet

Simple go gRPC project to greet server and client.

---

### gRPC commands

#### for build the protobuf
~~~
protoc -I unaryCalc/proto --go_out=. --go_opt=module=github.com/jhowilbur/grpc-project --go-grpc_out=. --go-grpc_opt=module=github.com/jhowilbur/grpc-project unaryCalc/proto/sum.proto
~~~

#### to build cert
~~~
openssl genrsa -out server.key 2048
openssl req -new -sha256 -key server.key -out server.csr -config config.conf
openssl x509 -req -days 3650 -in server.csr -out server.crt -signkey server.key -extensions ext -extfile config.conf
~~~

---

