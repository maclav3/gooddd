package domain

type EventProducer struct {
	events  []EventPayload
}

func (e *EventProducer) RecordThat(event EventPayload) {
	if event == nil {
		return
	}
	e.events = append(e.events, event)
}

func (e *EventProducer) PopEvents() []EventPayload {
	defer func() { e.events = nil }()
	return e.events
}
