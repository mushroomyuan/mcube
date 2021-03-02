package generater

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	routerPackage = protogen.GoImportPath("github.com/infraboard/mcube/http/router")
)

// NewGenerater todo
func NewGenerater(gen *protogen.Plugin, file *protogen.File) *Generater {
	filename := file.GeneratedFilenamePrefix + "_http.pb.go"
	return &Generater{
		gen:  gen,
		file: file,
		g:    gen.NewGeneratedFile(filename, file.GoImportPath),
	}
}

// Generater todo
type Generater struct {
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile
}

// GenerateFile generates a _grpc.pb.go file containing gRPC service definitions.
func (m *Generater) GenerateFile() {
	if len(m.file.Services) == 0 {
		return
	}

	m.g.P("// Code generated by protoc-gen-go-http. DO NOT EDIT.")
	m.g.P()
	m.g.P("package ", m.file.GoPackageName)
	m.g.P()
	m.generateFileContent()
	return
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func (m *Generater) generateFileContent() {
	m.generateHTTPEntry()
}

// generateHTTPEntry generates the gRPC service definitions, excluding the package statement.
func (m *Generater) generateHTTPEntry() {
	pn := m.file.Proto.GetPackage()

	g := m.g
	g.P("// HttpEntry todo")
	g.P("func ", "HttpEntry", "()", " *", routerPackage.Ident("EntrySet"), " {")
	g.P("set := &", routerPackage.Ident("EntrySet"), "{")
	g.P("Items: []*", routerPackage.Ident("Entry"), "{")
	for _, service := range m.file.Proto.Service {
		for _, method := range service.GetMethod() {
			opt := GetServiceMethodRestAPIOption(method)
			fullPath := fmt.Sprintf("/%s.%s/%s", pn, service.GetName(), method.GetName())
			g.P("{")
			g.P(`Path: "`, fullPath, `",`)
			g.P(`Method: "`, opt.Method, `",`)
			g.P(`FunctionName: "`, method.GetName(), `",`)
			g.P(`Resource: "`, opt.Resource, `",`)
			g.P("AuthEnable: ", opt.AuthEnable, ",")
			g.P("PermissionEnable: ", opt.PermissionEnable, ",")
			g.P("Labels: ", m.genLable(opt.Labels), ",")
			g.P("},")
		}
	}
	g.P("},")
	g.P("}")
	g.P("return set")
	g.P("}")
	g.P()
}

func (m *Generater) genLable(data map[string]string) string {
	kv := []string{}
	for k, v := range data {
		kv = append(kv, fmt.Sprintf(`"%s": "%s"`, k, v))
	}
	return fmt.Sprintf(`map[string]string{%s}`, strings.Join(kv, ","))
}
