Source Code for [How to Deliver Event Messages Successfully | Outbox Pattern](https://youtu.be/vM-WNyEiNk8)
[![Options Pattern Youtube Video Link](https://img.youtube.com/vi/vM-WNyEiNk8/0.jpg)](https://www.youtube.com/watch?v=vM-WNyEiNk8)

## Docker Scripts
MSSQL Script

    docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=MFKinTech12*" -p 1433:1433 -d mcr.microsoft.com/mssql/server:2022-latest

RabbitMQ Script

    docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 rabbitmq:3-management
