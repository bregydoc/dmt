package onesignal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bregydoc/dmt"
)

type PushNotification struct {
	appID    string
	apiKey   string
	done     bool
	Contents map[Language]string
}

func (e *PushNotification) Type() dmt.WorkType {
	return "push-notification"
}

func (e *PushNotification) State() dmt.WorkState {
	if e.done {
		return dmt.WorkDone
	}
	return dmt.WorkPending
}

func (e *PushNotification) ExecuteTask() error {
	if e.IsDone() {
		return errors.New("work done")
	}

	iSeg := `["Subscribed Users"]`

	c, err := json.Marshal(e.Contents)
	if err != nil {
		return err
	}

	contents := string(c)

	body := strings.NewReader(`{"app_id": "` + e.appID + `", "contents": ` + contents + `, "included_segments": ` + iSeg + `}`)

	req, err := http.NewRequest(http.MethodPost, notificationsAPI, body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic "+e.apiKey)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("invalid one signal api response. code: %d. body: %s", res.StatusCode, string(body))
	}

	e.done = true
	return nil
}

func (e *PushNotification) IsDone() bool {
	return e.done
}
