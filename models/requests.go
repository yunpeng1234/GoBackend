package models

type RequestRegister struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}
type RequestSuspend struct {
	Student string `json:"student"`
}

type RequestNotification struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}
