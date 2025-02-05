package generate

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/goctl-php/template"
)

func genClient(dir string, ns string, api *spec.ApiSpec) error {
	data := template.PhpTemplateData{
		Namespace: ns,
	}

	if err := template.WriteFile(dir, "ApiBaseClient", template.ApiBaseClient, data); err != nil {
		return err
	}
	if err := template.WriteFile(dir, "ApiException", template.ApiException, data); err != nil {
		return err
	}
	if err := template.WriteFile(dir, "ApiBody", template.ApiBody, data); err != nil {
		return err
	}
	return writeClient(dir, ns, api)
}

func writeClient(dir string, ns string, api *spec.ApiSpec) error {
	name := camelCase(api.Service.Name, true)

	data := template.PhpApiClientTemplateData{
		PhpTemplateData: template.PhpTemplateData{Namespace: ns},
		ClientName:      name,
		Routes:          []template.PhpApiClientRouteTemplateData{},
	}

	for _, g := range api.Service.Groups {
		prefix := g.GetAnnotation("prefix")

		// 路由
		for _, r := range g.Routes {
			route := template.PhpApiClientRouteTemplateData{
				HttpMethod:   strings.ToLower(r.Method),
				UrlPath:      r.Path,
				Prefix:       prefix,
				ActionPrefix: camelCase(prefix, true),
				ActionName:   camelCase(r.Path, true),
			}

			if r.RequestType != nil {
				requestType := camelCase(r.RequestType.Name(), true)
				route.RequestType = &requestType
				route.RequestHasPathParams = hasTagMembers(r.RequestType, pathTagKey)
				route.RequestHasQueryString = hasTagMembers(r.RequestType, formTagKey)
				route.RequestHasHeaders = hasTagMembers(r.RequestType, headerTagKey)
				route.RequestHasBody = hasTagMembers(r.RequestType, bodyTagKey)
			}

			if r.ResponseType != nil {
				responseType := camelCase(r.ResponseType.Name(), true)
				route.ResponseType = &responseType

				definedType, ok := r.ResponseType.(spec.DefineStruct)
				if !ok {
					return fmt.Errorf("type %s not supported", responseType)
				}
				if rh, err := enumResponseSubMessageKey(&definedType, headerTagKey); err != nil {
					return err
				} else {
					route.ResponseHeaders = rh
				}
				if rb, err := enumResponseSubMessageKey(&definedType, bodyTagKey); err != nil {
					return err
				} else {
					route.ResponseBody = rb
				}
			}

			data.Routes = append(data.Routes, route)
		}
	}

	return template.WriteFile(dir, fmt.Sprintf("%sClient", name), template.ApiClient, data)
}

func enumResponseSubMessageKey(definedType *spec.DefineStruct, tag string) (map[string]string, error) {
	// 获取字段
	ms := definedType.GetTagMembers(tag)
	if len(ms) <= 0 {
		return nil, nil
	}

	result := map[string]string{}

	for _, m := range ms {
		tags := m.Tags()
		k := ""
		if len(tags) > 0 {
			k = tags[0].Name
		} else {
			k = m.Name
		}
		result[camelCase(m.Name, true)] = k
	}
	return result, nil
}
