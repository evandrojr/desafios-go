package event

import "time"

type NewOrderListed struct {
	Name    string
	Payload interface{}
}

func NewNewOrderListed() *NewOrderListed {
	return &NewOrderListed{
		Name: "NewOrderListed",
	}
}

func (e *NewOrderListed) GetName() string {
	return e.Name
}

func (e *NewOrderListed) GetPayload() interface{} {
	return e.Payload
}

func (e *NewOrderListed) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *NewOrderListed) GetDateTime() time.Time {
	return time.Now()
}
