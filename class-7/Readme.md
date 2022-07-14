```go get github.com/dgrijalva/jwt-go

openssl genrsa -out app.rsa 1024 // certificado privado

openssl rsa -in app.rsa -pubout > app.rsa.pub // publico
```

