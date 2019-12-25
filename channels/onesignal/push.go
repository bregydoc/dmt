package onesignal

import "encoding/json"


type PushNotificationForAll struct {
	Contents map[Language]string `json:"contents"`
}

func (p PushNotificationForAll) Get(name string) (interface{}, error) {
	if name == "contents" {
		c, err := json.Marshal(p.Contents)
		if err != nil {
			return nil, err
		}
		return string(c), nil
	}else if name == "included_segments" {
		return `[\"Subscribed Users\"]`, nil
	}
	return nil, nil
}


