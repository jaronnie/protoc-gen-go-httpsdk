package fake

type FakeScopeClientData struct {
	GoModule       string   // github.com/jaronnie/autosdk
	ScopeVersion   string   // corev1
	UpScopeVersion string   // Corev1
	UpResources    []string // [Credential, Machine, Chain]
}

var FakeScopeClientTpl = `
// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"{{.GoModule}}/rest"
	"{{$.GoModule}}/typed/{{.ScopeVersion}}"
)

type Fake{{.UpScopeVersion}} struct {}

func (f *Fake{{.UpScopeVersion}}) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

{{range $v := .UpResources}}func (f *Fake{{$.UpScopeVersion}}) {{$v}}() {{$.ScopeVersion}}.{{$v}}Interface {
	return &Fake{{$v}}{Fake: f}
}

{{end}}
`
