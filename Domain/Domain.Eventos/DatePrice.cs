using System;

namespace Domain.Events
{
    public class DatePrice
    {
        public int Id { get; set; }
        public DateTime Date { get; set; }
        public float Price { get; set; }
    }
}
