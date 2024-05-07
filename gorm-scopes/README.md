Source file and commands for [Creating a new endpoint with GORM Scopes | Golang API Part 4](https://www.youtube.com/watch?v=pzhjUgwP-f8)
[![#Creating a new endpoint with GORM Scopes | Golang API Part 4 Youtube Video Link](https://img.youtube.com/vi/pzhjUgwP-f8/0.jpg)](https://www.youtube.com/watch?v=pzhjUgwP-f8)


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

### Cargo Table DDL
```json
create table cargos
(
    id          serial
        primary key,
    code        varchar(100) not null,
    description varchar(100) not null,
    created_at  timestamp default CURRENT_TIMESTAMP
);

alter table cargos
    owner to postgres;
```
