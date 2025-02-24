package template

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/samber/lo"
)

//go:embed ApiBaseClient.tpl
var ApiBaseClient string

//go:embed ApiClient.tpl
var ApiClient string

//go:embed ApiException.tpl
var ApiException string

//go:embed ApiMessage.tpl
var ApiMessage string

//go:embed ApiSubMessage.tpl
var ApiSubMessage string

//go:embed ApiBody.tpl
var ApiBody string

type PhpTemplateData struct {
	Namespace string
}

type PhpApiClientRouteTemplateData struct {
	HttpMethod            string
	UrlPath               string
	Prefix                string
	ActionName            string
	ActionPrefix          string
	RequestType           *string
	RequestHasPathParams  bool
	RequestHasQueryString bool
	RequestHasBody        bool
	RequestHasHeaders     bool
	ResponseType          *string
	ResponseBody          map[string]string
	ResponseHeaders       map[string]string
}

type PhpApiClientTemplateData struct {
	PhpTemplateData
	ClientName string
	Routes     []PhpApiClientRouteTemplateData
}

type PhpApiMessageTemplateData struct {
	PhpTemplateData
	MessageName string
	Properties  map[string]string
}

type PhpApiSubMessageTemplateData struct {
	PhpTemplateData
	MessageName string
	Properties  map[string]string
}

func WriteFile[T any](dir string, name string, tpl string, data T) error {
	tmpl, err := template.New(name).
		Funcs(template.FuncMap{
			"CamelCase":  lo.CamelCase,
			"PascalCase": lo.PascalCase,
		}).
		Parse(tpl)
	if err != nil {
		return err
	}
	p := filepath.Join(dir, fmt.Sprintf("%s.php", name))
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
