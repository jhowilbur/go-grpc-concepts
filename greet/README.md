# First model
## simple greet

Simple go gRPC project to greet server and client.

---

### gRPC commands
#### to see details about the OS
~~~
make about
~~~

#### to build protobuf
~~~
make greet
~~~

---

#### for the dummy proto
~~~
protoc -I greet/proto/dummy --go_out=. --go_opt=module=github.com/jhowilbur/grpc-project --go-grpc_out=. --go-grpc_opt=module=github.com/jhowilbur/grpc-project  greet/proto/dummy/dummy.proto
~~~