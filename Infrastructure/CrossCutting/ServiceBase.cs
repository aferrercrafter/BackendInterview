using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Infrastructure.CrossCutting
{
    public class ServiceBase<T> : IServiceBase<T> where T : class
    {
        private readonly IRepository<T> _repository;

        public ServiceBase(IRepository<T> repository)
        {
            _repository = repository;
        }

        public IEnumerable<T> GetAll()
        {
            return _repository.GetAll();
        }

        public T GetById(object Id)
        {
            return _repository.GetById(Id);
        }

        public Task<T> GetByIdAsync(object Id)
        {
            return _repository.GetByIdAsync(Id);
        }

        public IQueryable<T> Query()
        {
            return _repository.Query();
        }

        public void Insert(T obj)
        {
            _repository.Insert(obj);
        }

        public Task InsertAsync(T obj)
        {
            return _repository.InsertAsync(obj);
        }

        public void Update(T obj)
        {
            _repository.Update(obj);
        }

        public void Delete(Object Id)
        {
            _repository.Delete(Id);
        }

        public void Save()
        {
            _repository.Save();
        }

        public Task<int> SaveAsync()
        {
            return _repository.SaveAsync();
        }
    }
}
