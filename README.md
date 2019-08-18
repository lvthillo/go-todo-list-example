# go-todo-list-example
An example project to create your own todo list written in Go.

It supports:
- Adding notes (POST)
- Listing notes (GET)
- Completing notes (PUT)
- Deleting notes (DELETE)


Run the application:
```
$ go run main.go
```

Add a note to the todo list (POST):
```
$ curl localhost:3000/todo -d '{"message": "I still have to cook dinner."}'
{"id":"blcnvkkllhcm5r8n5ctg"}
```

List notes (GET):
```
$ curl localhost:3000/todo
[{"id":"blcnvkkllhcm5r8n5ctg","message":"I still have to cook dinner.","complete":false}]
```

Complete note (PUT):
```
$ curl -X PUT -H "Content-Type: application/json" -d '{"id":"blcnvkkllhcm5r8n5ctg"}' localhost:3000/todo
```


Delete note (DELETE):
```
$ curl -X DELETE localhost:3000/todo/blcnvkkllhcm5r8n5ctg
```
