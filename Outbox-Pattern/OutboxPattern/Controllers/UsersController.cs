using System.Text.Json;
using Microsoft.AspNetCore.Mvc;
using OutboxPattern.Data;
using OutboxPattern.Entities;
using OutboxPattern.Enums;
using OutboxPattern.Models;

namespace OutboxPattern.Controllers;

[ApiController]
[Route("[controller]")]
public class UsersController : ControllerBase
{
    private readonly DataContext _dataContext;

    public UsersController(DataContext dataContext)
    {
        _dataContext = dataContext;
    }

    [HttpPost, Route("{id}/suspend")]
    public async Task<IActionResult> SuspendUser(long id)
    {
        var user = _dataContext.Users.FirstOrDefault(x => x.Id == id);
        if (user is null)
        {
            return NotFound("User Not Found!");
        }

        await using (var transaction = await _dataContext.Database.BeginTransactionAsync())
        {
            try
            {
                user.IsSuspended = true;
                
                await _dataContext.SaveChangesAsync();

                var userSuspended = CreateUserSuspended(id);
                
                /*
                 * _messagePublisher.Publish(userSuspended);
                 */
                
                var outbox = CreateOutbox(userSuspended, user);

                await _dataContext.Outbox.AddAsync(outbox);
                await _dataContext.SaveChangesAsync();

                await transaction.CommitAsync();
            }
            catch (Exception e)
            {
                await transaction.RollbackAsync();
                Console.WriteLine(e);
                throw;
            }
        }

        return Ok(user);
    }

    private static Outbox CreateOutbox(UserSuspended userSuspended, User user)
    {
        return new Outbox
        {
            Body = JsonSerializer.Serialize(userSuspended),
            CreatedDate = DateTime.Now,
            StatusId = (int)OutboxStatus.Started,
            EventId = user.Id.ToString()
        };
    }

    private static UserSuspended CreateUserSuspended(long id)
    {
        return new UserSuspended
        {
            Id = id
        };
    }
}