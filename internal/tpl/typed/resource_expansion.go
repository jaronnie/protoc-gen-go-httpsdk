package typed

type ResourceExpansionData struct {
	ScopeVersion string // corev1
	UpResource   string // Credential
}

var ResourceExpansionTpl = `
// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package {{.ScopeVersion}}

type {{.UpResource}}Expansion interface {}
`
