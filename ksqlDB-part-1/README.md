Source file and commands for [Filter your kafka topics on the fly | ksqlDB Part 1](https://youtu.be/STV2u2n0ULU)
[![ksqlDB Part 1 Youtube Video Link](https://img.youtube.com/vi/STV2u2n0ULU/0.jpg)](https://www.youtube.com/watch?v=STV2u2n0ULU&ab_channel=MFKinTech)

## Docker Scripts
docker-compose.yml script

    docker-compose up
    
ksqlDB server script

    docker exec -it ksqldb-cli ksql http://ksqldb-server:8088

## ksqlDB Commands
- CREATE STREAM cities (id BIGINT, name VARCHAR, population BIGINT) WITH (kafka_topic='cities', value_format='json', partitions=1);
- INSERT INTO cities (id, name, population) VALUES (1, 'istanbul', 20000000);
- CREATE STREAM cities (id BIGINT, name VARCHAR, population BIGINT, districts ARRAY<STRUCT<id BIGINT, name VARCHAR, population BIGINT>>) WITH (kafka_topic='cities', value_format='json', partitions=1);
- CREATE STREAM big_cities WITH (kafka_topic='big_cities', partitions=1) AS SELECT * FROM cities WHERE population > 5000000  EMIT CHANGES;
