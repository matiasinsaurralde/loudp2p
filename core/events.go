package loudp2p

// EventType is an event identifier for the enum-like thing.
type EventType int

const (
	_                       = iota
	SayHello      EventType = 1
	HelloPeerAddr           = 2
	HelloRPCPort            = 3
)

// EventHandler is the event handler base data structure.
type EventHandler struct {
	Listeners []*chan (*Event)
}

// Event is a data structure for events.
type Event struct {
	Type     EventType
	Metadata interface{}
}

// NewEventHandler returns a brand new EventHandler{}.
func NewEventHandler() (eventHandler EventHandler) {
	return eventHandler
}

// AddListener initializes a listener channel and returns a pointer to it.
func (e *EventHandler) AddListener() *chan (*Event) {
	var eventChan chan (*Event)
	eventChan = make(chan *Event)
	e.Listeners = append(e.Listeners, &eventChan)
	return &eventChan
}

// Emit broadcasts the emitted event to available listeners.
func (e *EventHandler) Emit(event Event) {
	for _, listener := range e.Listeners {
		go func(event *Event, l *chan (*Event)) {
			*l <- event
		}(&event, listener)
	}
}
