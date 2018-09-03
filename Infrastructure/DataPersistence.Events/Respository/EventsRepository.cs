using Infrastructure.CrossCutting;
using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;

namespace DataPersistence.Events.Respository
{
    public interface IEventsRepository : IRepository<Domain.Events.Events>
    {
        IEnumerable<Domain.Events.Events> GetFeatured();
    }

    public class EventsRepository : Repository<Domain.Events.Events>, IEventsRepository
    {
        public EventsRepository (EventsContext context) : base(context)
        {

        }

        public IEnumerable<Domain.Events.Events> GetFeatured()
        {
            return DbSet().Where(x => x.Dates.Where(d => d.Date.Day == System.DateTime.Today.Day).Any()).ToList();
        }
    }
}
