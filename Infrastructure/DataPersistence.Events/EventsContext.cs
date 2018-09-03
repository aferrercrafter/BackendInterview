using Microsoft.EntityFrameworkCore;

namespace DataPersistence.Events
{
    public class EventsContext : DbContext
    {
        public EventsContext(DbContextOptions<EventsContext> dbContextOptions) : base(dbContextOptions) { }
        public DbSet<Domain.Events.Events> Eventos { get; set; }
    }
}
