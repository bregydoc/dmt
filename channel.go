package dmt

type ChannelName string

type Channel interface {
	Name() ChannelName
	AddTask(task Task) error
	FlushAll() error
	Observe(func([]Work)) error
}
