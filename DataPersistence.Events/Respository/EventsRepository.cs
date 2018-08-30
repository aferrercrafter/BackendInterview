using Infrastructure.CrossCutting;
using Domain.Eventos;
using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;

namespace DataPersistence.Events.Respository
{
    public interface IEventsRepository : IRepository<Eventos>
    {
        IEnumerable<Eventos> GetFeatured();
    }

    public class EventsRepository : IRepository<Eventos>
    {
        public readonly EventsContext _eventsContext;

        public DbSet<Eventos> DbSet()
        {
            throw new System.NotImplementedException();
        }

        public void Delete(object Id)
        {
            throw new System.NotImplementedException();
        }

        public IEnumerable<Eventos> GetAll()
        {
            throw new System.NotImplementedException();
        }

        public Eventos GetById(object Id)
        {
            throw new System.NotImplementedException();
        }

        public void Insert(Eventos obj)
        {
            throw new System.NotImplementedException();
        }

        public IQueryable<Eventos> Query()
        {
            throw new System.NotImplementedException();
        }

        public void Save()
        {
            throw new System.NotImplementedException();
        }

        public void Update(Eventos obj)
        {
            throw new System.NotImplementedException();
        }
    }
}
