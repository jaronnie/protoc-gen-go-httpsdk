package tpl

type ClientSetData struct {
	GoModule      string            //github.com/jaronnie/autosdk
	RootModule    string            // bfsdk
	ScopeVersions map[string]string // ["corev1":"CoreV1", "oauthv1":"OauthV1"]
}

var ClientSetTpl = `
// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package {{.RootModule}}

import (
	"{{.GoModule}}/rest"
	"{{.GoModule}}/typed"

	{{range $k, $v := .ScopeVersions}}{{$k}} "{{$.GoModule}}/typed/{{$k}}"
	{{end}}
)

type Interface interface {
	Direct() typed.DirectInterface
	
	{{range $k, $v := .ScopeVersions}}{{$v}}() {{$k}}.{{$v}}Interface
	{{end}}
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	{{range $k, $v := .ScopeVersions}}{{$k}} *{{$k}}.{{$v}}Client
	{{end}}
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

{{range $k, $v := .ScopeVersions}}func (x *Clientset) {{$v}}() {{$k}}.{{$v}}Interface {
	return x.{{$k}}
}

{{end}}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.direct, err = typed.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	{{range $k, $v := .ScopeVersions}}cs.{{$k}}, err = {{$k}}.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	{{end}}
	return &cs, nil
}
`
