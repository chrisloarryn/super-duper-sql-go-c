# Testing with go

|                     | Definition                                                                                                                    |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------|
| Unit testing        | test functionality separately                                                                                                 |
| Integration testing | - test functionality together and testing to some external services<br/> - only is integration if the service is being in use |


## Unit Testing

```go
func multiply(a, b int) int {
    return a * b
}
```

```go
func multiply(a, b int) int {
	r := a * b
	fmt.Println("Hello World")
	return r
}
```


## Integration Testing

```go
func ControllerSave(db databaseInterface, name string) {
    db.Save(name)
}
```
`if connect to db, it will test integration test`

if I simulate the db, it will test unit test