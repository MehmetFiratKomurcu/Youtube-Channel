create database TestDB;
use TestDB;
exec sp_cdc_enable_db;

create table order_outbox
(
    id            uniqueidentifier not null primary key,
    aggregatetype nvarchar(255)    not null,
    aggregateid   nvarchar(255)    not null,
    topic         nvarchar(255)    not null,
    eventType     nvarchar(255)    not null,
    payload       nvarchar(max)
)

insert into order_outbox (id, aggregatetype, aggregateid, topic, eventType, payload)
values ('23f15631-22d4-38ea-866f-32359dd1002d', 'Order', '81552721', 'fulfillment.order.created', 'OrderCreated', '{
  "id": 169108165,
  "code": "123adf5",
}');

EXEC sys.sp_cdc_enable_table @source_schema = 'dbo', @source_name = 'order_outbox', @role_name = NULL,
     @supports_net_changes = 0;

insert into order_outbox (id, aggregatetype, aggregateid, topic, eventType, payload)
values ('23f15631-32d4-38ea-866f-32359dd1902c', 'Order', '81552721', 'fulfillment.order.created', 'OrderCreated', '{
  "id": 10,
  "code": "Ben10",
}');

select * from order_outbox;