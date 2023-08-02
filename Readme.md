Run the server and test the CRUD operations.
`go run main.go`

You can use tools like curl or Postman to interact with the endpoints:

Create a new user:
`curl -X POST -d '{"name":"John Doe"}' http://localhost:8080/create
`

Read all users:
`curl http://localhost:8080/read`

Update a user:
`curl -X PUT -d '{"id":1,"name":"Updated Name"}' http://localhost:8080/update
`

Delete a user:
`curl -X DELETE -d '{"id":1}' http://localhost:8080/delete`