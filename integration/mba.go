package integration

import (
	"log"
)

// Mba struct sys
type Mba struct {
	Code string
	Data interface{}
}

// NewMbaService for new mba service
func NewMbaService(code string, data interface{}) Vendor {
	return &Mba{Code: code, Data: data}
}

// Do for execution
func (mba *Mba) Do() error {
	listMenu, listFn := MbaEvent(mba)
	err := Routing(mba.Code, listMenu, listFn)

	return err
}

// PublishPulsa func for refil pulsa xl 25 rb
func (mba *Mba) PublishPulsa(req *Request) error {
	log.Printf("isi pulsa mba code: %s data: %s publishCode: %s", mba.Code, mba.Data, req.PublishCode)
	return nil
}
