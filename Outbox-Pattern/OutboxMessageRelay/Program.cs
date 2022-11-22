using System.Text;
using Dapper;
using Microsoft.Data.SqlClient;
using OutboxMessageRelay.Entities;
using OutboxMessageRelay.Enums;
using Quartz;
using Quartz.Impl;
using RabbitMQ.Client;

var factory = new StdSchedulerFactory();
var scheduler = await factory.GetScheduler();
await scheduler.Start();

var job = CreateJob();
var trigger = TriggerJob();
await scheduler.ScheduleJob(job, trigger);

Console.ReadKey();

IJobDetail CreateJob()
{
    return JobBuilder.Create<OutboxJob>()
        .WithIdentity("job1", "group1")
        .Build();
}

ITrigger TriggerJob()
{
    return TriggerBuilder.Create()
        .WithIdentity("trigger1", "group1")
        .StartNow()
        .WithSimpleSchedule(x => x
            .WithIntervalInSeconds(5)
            .RepeatForever())
        .Build();
}

[DisallowConcurrentExecution]
public class OutboxJob : IJob
{
    public async Task Execute(IJobExecutionContext context)
    {
        await using var dbConnection = new SqlConnection(GetConnectionString());
        using var rabbitConnection = GetRabbitMqConnectionFactory().CreateConnection();
        using var channel = rabbitConnection.CreateModel();

        var sql = $"select * from Outbox where StatusId = {(int)OutboxStatus.Started}";
        var outboxRows = await dbConnection.QueryAsync<Outbox>(sql);

        foreach (var outboxRow in outboxRows)
        {
            PublishMessage(outboxRow, channel);

            var updateSql = $"UPDATE Outbox SET StatusId = {(int)OutboxStatus.Done} WHERE Id = {outboxRow.Id}";
            await dbConnection.ExecuteAsync(updateSql);
        }
    }

    private static void PublishMessage(Outbox outboxRow, IModel channel)
    {
        var body = Encoding.UTF8.GetBytes(outboxRow.Body);

        channel.BasicPublish(exchange: "",
            routingKey: "user-suspended",
            basicProperties: null,
            body: body);
    }

    private static ConnectionFactory GetRabbitMqConnectionFactory()
    {
        var factory = new ConnectionFactory
            { HostName = "localhost", Port = 5672, UserName = "guest", Password = "guest" };
        return factory;
    }

    private static string GetConnectionString()
    {
        const string connectionString =
            "data source=localhost;Initial Catalog=master;User Id=sa;Password=MFKinTech12*;Application Name=OutboxPattern;TrustServerCertificate=True;";
        return connectionString;
    }
}