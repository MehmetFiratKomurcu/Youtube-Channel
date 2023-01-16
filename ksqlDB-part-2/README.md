Source file and commands for [Aggregate Kafka Messages | ksqlDB Part 2](https://youtu.be/yCYv1g0wzs8)
[![ksqlDB Part 2 Youtube Video Link](https://img.youtube.com/vi/yCYv1g0wzs8/0.jpg)](https://youtu.be/yCYv1g0wzs8)

## Docker Scripts
docker-compose.yml script

    docker-compose up
    
ksqlDB server script

    docker exec -it ksqldb-cli ksql http://ksqldb-server:8088

## ksqlDB Commands
- CREATE STREAM fulfillment_orders (id BIGINT, code varchar, productCount bigint, warehouseCode varchar, city varchar) WITH (kafka_topic='fulfillment_orders', value_format='json', partitions=1);
- SET 'auto.offset.reset' = 'earliest';
- create table warehouse_reports as select warehouseCode as warehouse, count(*) as orderCount, AVG(productCount) as productCount, LATEST_BY_OFFSET(city) as lastCity  from fulfillment_orders group by warehouseCode emit changes;

## Curl Commands
### Get Fulfillment Orders Stream
```json
curl -X "POST" "http://localhost:8088/query-stream" \
     -d $'{
  "sql": "SELECT * FROM fulfillment_orders;",
  "streamsProperties": {"auto.offset.reset": "earliest"}
}'
```

### Get Warehouse Reports Table
```json
curl -X "POST" "http://localhost:8088/query-stream" \
     -d $'{
  "sql": "SELECT * FROM warehouse_reports;",
  "streamsProperties": {"auto.offset.reset": "earliest"}
}'
```


## JSON Order Examples
### First Order
```json
{
    "id": 1,
    "code": "order-1",
    "productCount": 5,
    "warehouseCode": "Gotham",
    "city": "themyscira"
}
```
### Second Order
```json
{
    "id": 1,
    "code": "order-1",
    "productCount": 5,
    "warehouseCode": "Gotham",
    "city": "Coast City"
}
```