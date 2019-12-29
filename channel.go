package dmt

type ChannelName string

type Channel interface {
	Name() ChannelName
	AddTask(task Task) error
	FlushAll() error
	Observe(func([]Work)) error
}

func (e *Engine) flushAllChannels() error {
	for _, ch := range e.channels {
		if err := ch.FlushAll(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) AddNewChannel(ch Channel) error {
	if e.channels == nil {
		e.channels = make([]Channel, 0)
	}
	e.channels = append(e.channels, ch)
	return nil
}