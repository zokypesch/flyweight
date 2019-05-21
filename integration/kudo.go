package integration

import "log"

// Kudo struct sys
type Kudo struct {
	Code string
	Data interface{}
	// ListOrder map[string]interface{}
}

// NewKudoService for new mba service
func NewKudoService(code string, data interface{}) Vendor {
	return &Kudo{Code: code, Data: data}
}

// Do for execution
func (kudo *Kudo) Do() error {
	listMenu, listFn := KudoEvent(kudo)
	err := Routing(kudo.Code, listMenu, listFn)

	return err
}

// PublishPulsa purchase pulsa
func (kudo *Kudo) PublishPulsa(req *Request) error {
	log.Printf("isi pulsa kudo code: %s data: %s publishCode: %s", kudo.Code, kudo.Data, req.PublishCode)
	return nil
}
