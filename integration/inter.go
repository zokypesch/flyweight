package integration

import "fmt"

// Vendor needed
type Vendor interface {
	Do() error
	PublishPulsa(req *Request) error
}

// Request request from http
type Request struct {
	Code        string
	PublishCode string
	Fn          string
}

//FnMapping mapping function
var FnMapping map[string]interface{}

// Routing for routing your request
func Routing(code string, list []*Request, listFn map[string]interface{}) error {

	var exist bool
	var req *Request

	for _, v := range list {
		if v.Code == code {
			exist = true
			req = v

			break
		}
	}

	if !exist {
		return fmt.Errorf("request not found")
	}

	v, ok := listFn[req.Fn]

	if !ok {
		return fmt.Errorf("routing funct not found")
	}

	return v.(func(*Request) error)(req)

}
