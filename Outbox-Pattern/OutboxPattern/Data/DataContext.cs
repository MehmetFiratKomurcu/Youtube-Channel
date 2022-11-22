using Microsoft.EntityFrameworkCore;
using OutboxPattern.Entities;

namespace OutboxPattern.Data
{
    public class DataContext : DbContext
    {
        public DataContext(DbContextOptions<DataContext> options) : base(options)
        {
        }

        public virtual DbSet<Outbox> Outbox { get; set; }
        public virtual DbSet<User> Users { get; set; }
    }
}