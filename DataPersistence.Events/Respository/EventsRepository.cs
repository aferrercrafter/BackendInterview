using Infrastructure.CrossCutting;
using Domain.Eventos;
using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;

namespace DataPersistence.Events.Respository
{
    public interface IEventsRepository : IRepository<Domain.Eventos.Events>
    {
        IEnumerable<Domain.Eventos.Events> GetFeatured();
    }

    public class EventsRepository : Repository<Domain.Eventos.Events>, IRepository<Domain.Eventos.Events>
    {
        public EventsRepository (EventsContext context) : base(context)
        {

        }

        public IEnumerable<Domain.Eventos.Events> GetFeatured()
        {
            return DbSet().Where(x => x.Dates.Where(d => d.Day == System.DateTime.Today.Day).Any()).ToList();
        }
    }
}
