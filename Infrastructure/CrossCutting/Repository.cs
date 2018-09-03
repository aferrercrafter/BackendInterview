using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Infrastructure.CrossCutting
{
    public class Repository<T> : IRepository<T> where T : class
    {
        private DbContext db;
        private DbSet<T> dbSet;

        public Repository(DbContext context)
        {
            db = context;
            dbSet = db.Set<T>();
        }
        public IEnumerable<T> GetAll()
        {
            return dbSet.ToList();
        }
        public T GetById(object Id)
        {
            return dbSet.Find(Id);
        }
        public Task<T> GetByIdAsync(object Id)
        {
            return dbSet.FindAsync(Id);
        }
        public IQueryable<T> Query()
        {
            return dbSet.AsQueryable<T>();
        }
        public void Insert(T obj)
        {
            dbSet.Add(obj);
        }
        public Task InsertAsync(T obj)
        {
            return dbSet.AddAsync(obj);
        }
        public void Update(T obj)
        {
            db.Entry(obj).State = EntityState.Modified;
        }
        public void Delete(object Id)
        {
            T getObjById = dbSet.Find(Id);
            dbSet.Remove(getObjById);
        }
        public DbSet<T> DbSet()
        {
            return dbSet;
        }
        public void Save()
        {
            db.SaveChanges();
        }
        public Task<int> SaveAsync()
        {
            return db.SaveChangesAsync();
        }
        protected virtual void Dispose(bool disposing)
        {
            if (disposing)
            {
                if (this.db != null)
                {
                    this.db.Dispose();
                    this.db = null;
                }
            }
        }


    }
}