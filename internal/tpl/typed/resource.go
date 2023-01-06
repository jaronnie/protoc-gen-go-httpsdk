package typed

import "github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"

type ResourceData struct {
	Gateways []*vars.Gateway

	IsWarpHttpResponse                bool     // is warped code,data,message
	IsResourceExpansionCreateOrUpdate bool     // is to create or update resource expansion
	GoModule                          string   // github.com/jaronnie/autosdk
	GoImportPaths                     []string // pb import path [github.com/jaronnie/autosdk/pb/corev1]
	ScopeVersion                      string   // corev1
	UpScopeVersion                    string   // Corev1
	Resource                          string   // credential
	UpResource                        string   // Credential
}

var ResourceTpl = `
{{define "methodDefine"}}{{.FuncName}}({{if or .IsStreamServer .IsStreamClient}}{{else}}ctx context.Context,{{end}} param *{{.ProtoRequestBody.RootPath}}.{{.ProtoRequestBody.Name}}) ({{if or .IsStreamServer .IsStreamClient}}*rest.Request{{else}}*{{.HttpResponseBody.RootPath}}.{{.HttpResponseBody.Name}}{{end}}, error){{end}}
// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package {{.ScopeVersion}}

import (
	"context"

	{{range $v := .GoImportPaths | uniq}}"{{$v}}"
	{{end}}

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	
	"{{.GoModule}}/rest"
)

type {{.UpResource}}Getter interface {
	{{.UpResource}}() {{.UpResource}}Interface
}

type {{.UpResource}}Interface interface {
	{{range $k, $v := .Gateways}}{{template "methodDefine" $v}}
	{{end}}
	{{if .IsResourceExpansionCreateOrUpdate}}{{.UpResource}}Expansion{{end}}
}

type {{.Resource}}Client struct {
	client rest.Interface
}

func new{{.UpResource}}Client(c *{{.UpScopeVersion}}Client) *{{.Resource}}Client {
	return &{{.Resource}}Client{
		client: c.RESTClient(),
	}
}

{{range $k, $v := .Gateways}}func (x *{{$.Resource}}Client) {{template "methodDefine" $v}} {
	{{if or $v.IsStreamServer $v.IsStreamClient}}request := x.client.Verb("{{$v.HttpMethod}}").
		SubPath(
			"{{$v.Url}}",{{range $p := $v.PathParams}}
			rest.PathParam{Name: "{{$p.Name}}", Value: param.{{$p.GoName}}},{{end}}
		).
		Params({{if eq $v.HttpRequestBody.BodyName "*"}}{{else}}{{range $q := $v.QueryParams}}
			rest.QueryParam{Name: "{{$q.Name}}", Value: param.{{$q.GoName}}},{{end}}{{end}}
		).
		Body({{if eq $v.HttpRequestBody.BodyName ""}}nil{{else if eq $v.HttpRequestBody.BodyName "*"}}param{{else if ne $v.HttpMethod "GET"}}param.{{$v.HttpRequestBody.GoBodyName}}{{else}}nil{{end}})

	return request, nil{{else}}var resp {{$v.HttpResponseBody.RootPath}}.{{$v.HttpResponseBody.Name}}
		response, err := x.client.Verb("{{$v.HttpMethod}}").
		SubPath(
			"{{$v.Url}}",{{range $p := $v.PathParams}}
			rest.PathParam{Name: "{{$p.Name}}", Value: param.{{$p.GoName}}},{{end}}
		).
		Params({{if eq $v.HttpRequestBody.BodyName "*"}}{{else}}{{range $q := $v.QueryParams}}
			rest.QueryParam{Name: "{{$q.Name}}", Value: param.{{$q.GoName}}},{{end}}{{end}}
		).
		Body({{if eq $v.HttpRequestBody.BodyName ""}}nil{{else if eq $v.HttpRequestBody.BodyName "*"}}param{{else if ne $v.HttpMethod "GET"}}param.{{$v.HttpRequestBody.GoBodyName}}{{else}}nil{{end}}).
		Do(ctx).
		{{if $.IsWarpHttpResponse}}TransformResponse(){{else}}RawResponse(){{end}}

	if err != nil {
		return nil, err
	}

	jsonb := new(runtime.JSONPb)
	err = jsonb.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil{{end}}
}
{{end}}
`
