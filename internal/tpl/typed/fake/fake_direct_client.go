package fake

type FakeDirectClientData struct {
	GoModule string // github.com/jaronnie/autosdk
}

var FakeDirectClientTpl = `
// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"{{.GoModule}}/rest"
)

type FakeDirect struct {}

func (f *FakeDirect) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
`
