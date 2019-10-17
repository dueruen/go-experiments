# Go and GraphQL experiment
## Inspired by
https://www.youtube.com/watch?v=640sSD-1CJQ&list=PLUn4JBwKdy9tEjKW1pNIyw6klSXwz1bgk&index=8&t=0s  
https://www.youtube.com/watch?v=ldpnqwszE_I&list=PLUn4JBwKdy9tEjKW1pNIyw6klSXwz1bgk&index=9&t=0s

## Run the experiment
Create the go files from proto files
```bash
make protoAll
```
Start the services
```bash
go run service/user/cmd/main.go
go run service/house/cmd/main.go
go run service/api_gateway/cmd/main.go
```
## GraphQL Playground
The GraphQL Playground is on `localhost:8081`
Some query:
```json
mutation{createUser(input:{firstname:"WEBTEST", lastname:"LASTNAMEWEB", age:42}){firstname, lastname, age, id}}
query{users {firstname, lastname, age, id}}
query{houses {address, age, id, ownerId}}
```
