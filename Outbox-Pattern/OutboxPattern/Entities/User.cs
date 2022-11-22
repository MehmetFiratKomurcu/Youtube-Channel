namespace OutboxPattern.Entities
{
    public class User
    {
        public long Id { get; set; }
        public string? Username { get; set; }
        public bool IsSuspended { get; set; }
        public DateTime CreatedDate { get; set; }
    }
}