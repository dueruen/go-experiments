# Go-kit and DDD experiment
A experiment to see how go-kit works and and how it handles gRPC. And an experiment to show a project could be organised, inspired by DDD.

## Inspired by
https://github.com/shijuvar/gokit-examples  
https://www.youtube.com/watch?v=oL6JBUk6tj0&list=PLUn4JBwKdy9tEjKW1pNIyw6klSXwz1bgk&index=3  
https://github.com/katzien/go-structure-examples/tree/master/domain-hex-actor  

## Run the experiment
Start the server:
```bash
go run chat/cmd/chat-server/main.go
```
I have not made a client so a cli client can bu used, like [Evans](https://github.com/ktr0731/evans)
```bash
evans chat/pkg/transport/pb/chat.proto
```

## Auto-Generate using the proto file
```bach
cd chat/pkg/transport/pb
protoc chat.proto --go_out=plugins=grpc:.
```