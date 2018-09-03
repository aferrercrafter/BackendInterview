using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using DataPersistence.Events;
using Core.Events;
using AutoMapper;
using Api.Eventos.dto.events;

namespace Api.Eventos.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class EventsController : ControllerBase
    {
        private readonly IEventsService _eventsService;
        private readonly IMapper _mapper;

        public EventsController(IEventsService eventsService, IMapper mapper)
        {
            _eventsService = eventsService;
            _mapper = mapper;
        }

        // GET: api/Events
        [HttpGet]
        public IEnumerable<Events> GetEventos()
        {
            return _mapper.Map<IEnumerable<Domain.Events.Events>,IEnumerable<Events>>(_eventsService.GetAll());
        }

        // GET: api/Events/Featured
        [Route("api/[controller]/featured")]
        [HttpGet]
        public IEnumerable<Events> GetFeatured()
        {
            return _mapper.Map<IEnumerable<Domain.Events.Events>,IEnumerable<Events>>(_eventsService.GetFeatured());
        }

        // GET: api/Events/5
        [HttpGet("{id}")]
        public async Task<IActionResult> GetEvents([FromRoute] string id)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            var events = _mapper.Map<Domain.Events.Events, Events>(await _eventsService.GetByIdAsync(id));

            if (events == null)
            {
                return NotFound();
            }

            return Ok(events);
        }

        // PUT: api/Events/5
        [HttpPut("{id}")]
        public async Task<IActionResult> PutEvents([FromRoute] string id, [FromBody] Events events)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != events.Id)
            {
                return BadRequest();
            }

            _eventsService.Update(_mapper.Map<Events, Domain.Events.Events>(events));

            try
            {
                await _eventsService.SaveAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (_eventsService.GetById(id) == null)
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        // POST: api/Events
        [HttpPost]
        public async Task<IActionResult> PostEvents([FromBody] dto.events.Events events)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _eventsService.Insert(_mapper.Map<dto.events.Events, Domain.Events.Events>(events));
            await _eventsService.SaveAsync();

            return CreatedAtAction("GetEvents", new { id = events.Id }, events);
        }

        // DELETE: api/Events/5
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteEvents([FromRoute] string id)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            var events = await _eventsService.GetByIdAsync(id);
            if (events == null)
            {
                return NotFound();
            }

            _eventsService.Delete(events);
            await _eventsService.SaveAsync();

            return Ok(events);
        }
    }
}