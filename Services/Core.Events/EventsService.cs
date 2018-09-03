using Infrastructure.CrossCutting;
using System.Collections.Generic;
using DataPersistence.Events.Respository;

namespace Core.Events
{
    public interface IEventsService : IServiceBase<Domain.Events.Events>
    {
        IEnumerable<Domain.Events.Events> GetFeatured();
    }

    public class EventsService : ServiceBase<Domain.Events.Events>, IEventsService
    {
        private IEventsRepository _eventsRepository;

        public EventsService(IEventsRepository eventsRepository) : base(eventsRepository)
        {
            _eventsRepository = eventsRepository;
        }

        public IEnumerable<Domain.Events.Events> GetFeatured()
        {
            return _eventsRepository.GetFeatured();
        }
    }
}
