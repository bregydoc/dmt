package onesignal

type Fabric struct {
	AppID string
	APIKey string
}

func NewFabric(AppID, APIKey string) (*Fabric, error) {
	return &Fabric{
		AppID:  AppID,
		APIKey: APIKey,
	}, nil
}

func (os *Fabric) NewPushNotificationWork() (*PushNotification, error) {
	return &PushNotification{
		appID:  os.AppID,
		apiKey: os.APIKey,
		done:   false,
	}, nil
}