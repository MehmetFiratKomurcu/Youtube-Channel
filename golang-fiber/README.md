
Source file and commands for [Unlocking Golang Fiber: A Practical Approach to Endpoints, Middlewares and Validations](https://youtu.be/dtgrKy-cjFU)
[![Unlocking Golang Fiber: A Practical Approach to Endpoints, Middlewares and Validations Youtube Video Link](https://img.youtube.com/vi/dtgrKy-cjFU/0.jpg)](https://youtu.be/dtgrKy-cjFU)


## Go Get Commands
### Get Fiber
```json
go get github.com/gofiber/fiber/v2
```

### Get Validator
```json
go get github.com/go-playground/validator/v10
```


## Curl Commands
### POST Order
```json
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"shipmentNumber": "55555"}' \
  http://localhost:3000/orders
```

### POST Order with CountryCode
```json
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"shipmentNumber": "55555", "countryCode": "TR"}' \
  http://localhost:3000/orders
```

### POST Order with CountryCode and Age
```json
curl --header "Content-Type: application/json" \
--request POST \
--data '{"shipmentNumber": "55555", "countryCode": "TR", "age": 45}' \
http://localhost:3000/orders
```
