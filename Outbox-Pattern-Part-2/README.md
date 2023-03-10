Source file and commands for [Polling Publisher vs Log Tailing | Debezium Outbox Router | Outbox Pattern Part 2](https://youtu.be/8HynE1SgozQ)
[![Polling Publisher vs Log Tailing | Debezium Outbox Router | Outbox Pattern Part 2 Youtube Video Link](https://img.youtube.com/vi/8HynE1SgozQ/0.jpg)](https://youtu.be/8HynE1SgozQ)

## Docker Scripts
Note: Extract debezium folder and then run docker compose

docker-compose.yml script

    docker-compose up

## Curl Commands
### Create Outbox Connector
```json
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d '{
    "name": "order-outbox-connector", 
    "config": {
        "connector.class": "io.debezium.connector.sqlserver.SqlServerConnector", 
        "database.hostname": "sqlserver", 
        "database.port": "1433", 
        "database.user": "sa", 
        "database.password": "Password!", 
        "database.names": "testDB", 
        "database.encrypt": "false",
        "topic.prefix": "fullfillment", 
        "table.include.list": "dbo.order_outbox", 
        "schema.history.internal.kafka.bootstrap.servers": "kafka:9092", 
        "schema.history.internal.kafka.topic": "schemahistory.fullfillment.outbox",
        "transforms": "outbox",
        "transforms.outbox.type": "io.debezium.transforms.outbox.EventRouter",
        "transforms.outbox.table.fields.additional.placement": "topic:envelope,eventType:envelope"
    }
  }'
```

## Resources
- [https://microservices.io/patterns/data/polling-publisher.html](https://microservices.io/patterns/data/polling-publisher.html)
- [https://microservices.io/patterns/data/transaction-log-tailing.html](https://microservices.io/patterns/data/transaction-log-tailing.html)
- [https://debezium.io/documentation/reference/2.1/transformations/outbox-event-router.html](https://debezium.io/documentation/reference/2.1/transformations/outbox-event-router.html)
- [https://github.com/debezium/debezium-examples/blob/main/tutorial/docker-compose-sqlserver.yaml](https://github.com/debezium/debezium-examples/blob/main/tutorial/docker-compose-sqlserver.yaml)