package generate

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/goctl-php/template"
	"github.com/zeromicro/goctl-php/util"
)

const (
	formTagKey   = "form"
	pathTagKey   = "path"
	headerTagKey = "header"
	bodyTagKey   = "json"
)

var (
	tagKeys = []string{pathTagKey, formTagKey, headerTagKey, bodyTagKey}
)

func tagToSubName(tagKey string) string {
	suffix := tagKey
	switch tagKey {
	case "json":
		suffix = "body"
	case "form":
		suffix = "query"
	}
	return suffix
}

func getMessageName(tn string, tagKey string, isPascal bool) string {
	suffix := tagToSubName(tagKey)
	return util.CamelCase(fmt.Sprintf("%s-%s", tn, suffix), isPascal)
}

func hasTagMembers(t spec.Type, tagKey string) bool {
	definedType, ok := t.(spec.DefineStruct)
	if !ok {
		return false
	}
	ms := definedType.GetTagMembers(tagKey)
	return len(ms) > 0
}

func genMessages(dir string, ns string, api *spec.ApiSpec) error {
	for _, t := range api.Types {
		tn := t.Name()
		definedType, ok := t.(spec.DefineStruct)
		if !ok {
			return fmt.Errorf("type %s not supported", tn)
		}

		data := template.PhpApiMessageTemplateData{
			PhpTemplateData: template.PhpTemplateData{Namespace: ns},
			MessageName:     util.CamelCase(tn, true),
			Properties:      map[string]string{},
		}

		// 子类型
		for _, tagKey := range tagKeys {
			// 获取字段
			ms := definedType.GetTagMembers(tagKey)
			if len(ms) <= 0 {
				continue
			}
			cn := getMessageName(tn, tagKey, true)
			data.Properties[tagToSubName(tagKey)] = cn

			// 写入
			if err := writeSubMessage(dir, ns, cn, ms); err != nil {
				return err
			}
		}

		// 主类型
		if err := template.WriteFile(dir, data.MessageName, template.ApiMessage, data); err != nil {
			return nil
		}
	}

	return nil
}

func writeSubMessage(dir string, ns string, cn string, ms []spec.Member) error {
	data := template.PhpApiSubMessageTemplateData{
		PhpTemplateData: template.PhpTemplateData{Namespace: ns},
		MessageName:     cn,
		Properties:      map[string]string{},
	}
	// 字段
	for _, m := range ms {
		tags := m.Tags()
		n := util.CamelCase(m.Name, false)
		k := ""
		if len(tags) > 0 {
			k = tags[0].Name
		} else {
			k = n
		}
		data.Properties[m.Name] = k
	}

	return template.WriteFile(dir, cn, template.ApiSubMessage, data)
}
