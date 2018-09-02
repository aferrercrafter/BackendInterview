using System;

namespace Domain.Eventos
{
    public class Events
    {
        public string Id { get; set; }
        public string Title { get; set; }
        public string EventImage { get; set; }
        public string Description { get; set; }
        public DateTime[] Dates { get; set; }
        public string Location { get; set; }
    }
}
