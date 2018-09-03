using System;

namespace Api.Eventos.dto.events
{
    public class Events
    {
        public string Id { get; set; }
        public string Title { get; set; }
        public string EventImage { get; set; }
        public string Description { get; set; }
        public DatePrice[] Dates { get; set; }
        public string Location { get; set; }
    }

    public class DatePrice
    {
        public int Id { get; set; }
        public DateTime Date { get; set; }
        public float Price { get; set; }
    }
}
