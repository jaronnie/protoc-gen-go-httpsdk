package interfaces

import "github.com/jaronnie/protoc-gen-go-httpsdk/internal/vars"

type Interface interface {
	GenGoMod() error
	GenRestFrame() error
	GenScopeClient(scopeResourceGws vars.ScopeResourceGateway) error
	GenResource(scopeResourceGws vars.ScopeResourceGateway) error
	GenClientSet(scopeResourceGws vars.ScopeResourceGateway) error
}
