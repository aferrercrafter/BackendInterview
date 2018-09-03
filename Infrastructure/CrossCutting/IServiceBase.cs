using Infrastructure.CrossCutting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Infrastructure.CrossCutting
{
    public interface IServiceBase<T> where T : class
    {
        IEnumerable<T> GetAll();
        T GetById(object Id);
        Task<T> GetByIdAsync(object Id);
        void Insert(T obj);
        Task InsertAsync(T obj);
        void Update(T obj);
        void Delete(object Id);
        void Save();
        Task<int> SaveAsync();
        IQueryable<T> Query();
    }
}
