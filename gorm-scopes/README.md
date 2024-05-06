
Source file and commands for [Creating a new endpoint with GORM Scopes | Golang API Part 4](https://youtu.be/FJ3pqRhOAA0)
[![#Creating a new endpoint with GORM Scopes | Golang API Part 4 Youtube Video Link](https://img.youtube.com/vi/FJ3pqRhOAA0/0.jpg)](https://youtu.be/m_8CjSQ0ZR8)


## Commands

### Run Postgres Docker Container
```json
$ docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

### Update Swagger
```json 
swag init
swag fmt
```
