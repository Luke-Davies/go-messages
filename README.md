# go-messages
A simple messaging REST service implemented in go using the gorilla mux router.
The application allows you to store and retrieve messages via http.
Messages are stored in memory and are lost when the server is stopped.

### Swagger

A rough API definition can be found on swaggerhub [here](https://app.swaggerhub.com/apis/Luke-Davies/Messages_API/1.0.0).
Note that, for this implementation, some error responses may not be implemented.

### Running the server:

- `go get github.com/luke-davies/go-messages`
- `go install github.com/luke-davies/go-messages`
- `$ go-messages`
  or
  `$ $GOPATH/bin/go-messages`
This will launch the server on localhost, port `3000`
   
### Calling the messages service

The following examples demonstrate calling the service on localhost:

Add a message:
```
$ curl localhost:3000/messages -d 'test message'
{"id":1000}
```
View the message:
```
$ curl localhost:3000/messages/1000
test message
```
See all messages:
```
$ curl localhost:3000/messages
[{"id":1000,"text":"test message"}]
```
Clear all messages:
```
$ curl -X DELETE localhost:3000/messages

```

---
### Node.js Version
Want to see the same thing implemented in Node.js? [Click here](https://github.com/Luke-Davies/messages-express)
