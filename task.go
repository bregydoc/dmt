package dmt

type TaskParams map[string]interface{}

type Task struct {
	// ID      string     `json:"id" yaml:"id"`
	Channel ChannelName `json:"channel" yaml:"channel"`
	Type    WorkType    `json:"type" yaml:"type"`
	Params  TaskParams  `json:"params" yaml:"params"`
}
