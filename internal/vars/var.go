package vars

type Scope string

type Resource string

type ScopeResourceGateway map[Scope]map[Resource][]*Gateway

type ServiceGateway map[Resource][]*Gateway

type Gateway struct {
	ProtoRequestBody ProtoRequestBody
	HttpRequestBody  HttpRequestBody
	HttpResponseBody HttpResponseBody

	IsStreamClient   bool // 是否是 stream client
	IsStreamServer   bool // 是否是 stream server
	ProtoServiceName string
	FuncName         string
	HttpMethod       string
	Url              string
	PathParams       []*PathParam
	QueryParams      []*QueryParam
}

type ProtoRequestBody struct {
	Name         string
	GoImportPath string // git.hyperchain.cn/bfsdk/pb/corev1
	RootPath     string // corev1
}

type HttpRequestBody struct {
	BodyName   string // * or others
	GoBodyName string
}

type HttpResponseBody struct {
	Name         string
	GoImportPath string
	RootPath     string
}

type PathParam struct {
	Index  int
	Name   string
	GoName string
}

type QueryParam struct {
	GoName string
	Name   string
}
