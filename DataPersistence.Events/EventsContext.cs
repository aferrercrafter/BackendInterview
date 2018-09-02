using Microsoft.EntityFrameworkCore;
using Domain.Eventos;

namespace DataPersistence.Events
{
    public class EventsContext : DbContext
    {
        public EventsContext() { }
        public DbSet<Domain.Eventos.Events> Eventos { get; set; }
    }
}
