# fauna-go
Sample application that demonstrates basics operations we can do perform in FQL
by using a Fauna client for Go

### Running the application locally
```
> go run main.go
```

### Verification
```
> curl -H "Content-Type: application/json" 
-X POST http://localhost:8080/product
-d '{"name":"aa","categoriesrefs":["288228455849394693"]}'
```

### View swagger documentation
```
> go run main.go
```
You can see the swagger UI at http://localhost:8080/swagger/index.html#/
