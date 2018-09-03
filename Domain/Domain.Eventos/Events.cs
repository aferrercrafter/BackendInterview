using System;

namespace Domain.Events
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
}
