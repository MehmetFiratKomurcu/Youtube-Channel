
Source file and commands for [Crafting RESTful APIs with Go: Leveraging GORM & PostgreSQL](https://youtu.be/FJ3pqRhOAA0)
[![#Crafting RESTful APIs with Go: Leveraging GORM & PostgreSQL Youtube Video Link](https://img.youtube.com/vi/FJ3pqRhOAA0/0.jpg)](https://youtu.be/m_8CjSQ0ZR8)


## Commands

### Run Postgres Docker Container
```json
$ docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

### Get Gorm and Postgres Driver
```json 
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgre
```

### Update Swagger
```json 
swag init
swag fmt
```

