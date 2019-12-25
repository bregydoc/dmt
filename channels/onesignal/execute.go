package onesignal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bregydoc/dmt"
)

type PushNotification struct {
	appID  string
	apiKey string
	done   bool
}

func (e *PushNotification) Type() dmt.WorkType {
	return "push-notification"
}

func (e *PushNotification) ExecuteTask(params dmt.TaskParams) error {
	e.done = false
	iSeg, err := params.Get("included_segments")
	if err != nil {
		return err
	}

	contents, err := params.Get("contents")
	if err != nil {
		return err
	}

	body := strings.NewReader(`{\"app_id\": \"` + e.AppID + `\",
	\"contents\": ` + contents.(string) + `,
	\"included_segments\": `+ iSeg.(string) +`}`)

	res, err := http.Post(NotificationsAPI, "application/json; charset=utf-8", body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid one signal api response, response code: %d", res.StatusCode)
	}

	e.done = true
	return nil
}

func (e *PushNotification) IsDone() bool {
	return e.done
}


