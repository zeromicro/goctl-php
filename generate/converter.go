// MIT License
//
// Copyright (c) 2020 goctl-php
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

package generate

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/tal-tech/go-zero/core/collection"
	sx "github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	annotation "github.com/tal-tech/go-zero/tools/goctl/api/util"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
)

type (
	Plugin struct {
		Api           *spec.ApiSpec
		Style         string
		Dir           string
		ParentPackage string
	}
	Spec struct {
		Beans   []*Bean
		Service IService
	}

	Tag struct {
		tp   string
		name string
	}
	Type struct {
		Name string
	}
)

func (p *Plugin) SetParentPackage(parentPackage string) {
	p.ParentPackage = parentPackage
}

func (p *Plugin) Convert() (*Spec, error) {
	var (
		ret   Spec
		beans = make(map[string]*Bean)
	)

	for _, each := range p.Api.Types {
		list, err := getBean(p.ParentPackage, each)
		if err != nil {
			return nil, err
		}

		for _, bean := range list {
			beans[bean.Name.Lower()] = bean
		}
	}
	var r = make(map[string]spec.Route)
	for _, each := range p.Api.Service.Routes() {
		r[fmt.Sprintf("%v://%v", each.Method, each.Path)] = each
	}

	for _, each := range p.Api.Service.Groups {
		for _, each := range each.Routes {
			r[fmt.Sprintf("%v://%v", each.Method, each.Path)] = each
		}
	}
	var list []spec.Route
	for _, item := range r {
		list = append(list, item)
	}

	imports, routes := getRoute(list, beans)
	ret.Service = IService{
		ParentPackage: p.ParentPackage,
		Import:        strings.Join(trimList(imports), util.NL),
		Routes:        routes,
	}

	for _, each := range beans {
		ret.Beans = append(ret.Beans, each)
	}

	sort.Slice(ret.Beans, func(i, j int) bool {
		return ret.Beans[i].Name.Source() < ret.Beans[j].Name.Source()
	})

	sort.Slice(ret.Service.Routes, func(i, j int) bool {
		return ret.Service.Routes[i].Path < ret.Service.Routes[j].Path
	})

	return &ret, nil
}

func trimList(list []string) []string {
	var ret []string
	for _, each := range list {
		tmp := strings.TrimSpace(each)
		if len(tmp) == 0 {
			continue
		}
		ret = append(ret, tmp)
	}
	return ret
}

func getRoute(in []spec.Route, m map[string]*Bean) ([]string, []*Route) {
	var list []*Route
	var imports []string

	for _, each := range in {
		handlerName, ok := annotation.GetAnnotationValue(each.Annotations, "server", "handler")
		if !ok {
			continue
		}

		doc, _ := annotation.GetAnnotationValue(each.Annotations, "doc", "summary")
		if len(doc) > 0 {
			doc = strings.ReplaceAll(doc, "'", "")
			doc = strings.ReplaceAll(doc, "`", "")
			doc = strings.ReplaceAll(doc, `"`, "")
			doc = "// " + doc
		}

		path, ids, idsExpr := parsePath(each.Path)
		bean := m[strings.ToLower(each.RequestType.Name)]

		var queryId []string
		var queryExpr, pathIdExpr string
		var showRequestBody bool
		if bean != nil {
			imports = append(imports, bean.Import)
			for _, query := range bean.FormTag {
				queryId = append(queryId, "in.get"+stringx.From(query).ToCamel()+"()")
			}
			queryExpr = bean.GetQuery()
			showRequestBody = len(bean.JsonTag) > 0
			if showRequestBody {
				queryExpr = queryExpr + ", "
			}
			pathIdExpr = toRetrofitPath(ids, bean)
			if len(queryId) > 0 {
				pathIdExpr = pathIdExpr + ", "
			}
		}

		list = append(list, &Route{
			MethodName:       stringx.From(handlerName).Untitle(),
			Method:           strings.ToUpper(each.Method),
			Path:             path,
			RequestBeanName:  stringx.From(each.RequestType.Name).Title(),
			ResponseBeanName: stringx.From(each.ResponseType.Name).Title(),
			HasRequest:       len(each.RequestType.Name) > 0,
			ShowRequestBody:  showRequestBody,
			HasResponse:      len(each.ResponseType.Name) > 0,
			HavePath:         len(ids) > 0,
			PathId:           strings.Join(idsExpr, ","),
			PathIdExpr:       pathIdExpr,
			QueryId:          strings.Join(queryId, ","),
			HaveQuery:        len(queryId) > 0,
			QueryIdExpr:      queryExpr,
			Doc:              doc,
		})
	}
	return imports, list
}

func parsePath(path string) (string, []string, []string) {
	p := strings.Split(path, "/")
	var list, ids, idsExpr []string
	for _, each := range p {
		if strings.Contains(each, ":") {
			id := strings.ReplaceAll(each, ":", "")
			list = append(list, "{"+id+"}")
			ids = append(ids, id)
			idsExpr = append(idsExpr, "in.get"+stringx.From(id).ToCamel()+"()")
			continue
		}

		list = append(list, each)
	}
	return strings.Join(list, "/"), ids, idsExpr
}

func toRetrofitPath(ids []string, bean *Bean) string {
	if bean == nil {
		return ""
	}
	var list []string
	for _, each := range ids {
		m := bean.GetMember(each)
		if m == nil {
			continue
		}

		list = append(list, fmt.Sprintf(`@Path("%s") %s %s`, each, m.TypeName, each))
	}
	return strings.Join(list, ", ")
}

func getBean(parentPackage string, tp spec.Type) ([]*Bean, error) {
	var bean Bean
	var list []*Bean
	bean.Name = stringx.From(tp.Name)
	bean.ParentPackage = parentPackage

	for _, m := range tp.Members {
		externalBeans, err := getBeans(parentPackage, m, &bean)
		if err != nil {
			return nil, err
		}

		list = append(list, externalBeans...)
	}
	return list, nil
}

func getBeans(parentPackage string, member spec.Member, bean *Bean) ([]*Bean, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, member.Expr, member.Type == "interface{}")
	if err != nil {
		return nil, err
	}

	tag := NewTag(member.Tag)
	name := tag.GetTag()
	if tag.IsJson() {
		bean.JsonTag = append(bean.JsonTag, name)
	}
	if tag.IsPath() {
		bean.PathTag = append(bean.PathTag, name)
	}
	if tag.IsForm() {
		bean.FormTag = append(bean.FormTag, name)
	}

	bean.Import = strings.Join(imports, util.NL)
	comment := strings.Join(member.Comments, " ")
	doc := strings.Join(member.Docs, util.NL)
	if len(comment) > 0 {
		comment = "// " + comment
	}
	if len(doc) > 0 {
		doc = "// " + doc
	}
	bean.Members = append(bean.Members, &Member{
		Name:     stringx.From(name),
		TypeName: typeName,
		Comment:  comment,
		Doc:      doc,
	})
	beans = append(beans, bean)
	return beans, nil
}

func getTypeName(parentPackage string, expr interface{}, inter bool) ([]*Bean, []string, string, error) {
	set := collection.NewSet()
	switch v := expr.(type) {
	case map[string]interface{}:
		return getTypeName(parentPackage, unJsonMarshal(expr, inter), false)
	case spec.BasicType:
		imp, typeName := toPhpType(parentPackage, v.Name)
		set.AddStr(imp)
		return nil, set.KeysStr(), typeName, nil
	case spec.PointerType:
		return getTypeName(parentPackage, v.Star, false)
	case spec.MapType:
		set.AddStr("import java.util.HashMap;")
		beans, imports, typeName, err := toPhpMap(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		set.AddStr(imports...)
		return beans, set.KeysStr(), typeName, nil
	case spec.ArrayType:
		set.AddStr("import java.util.ArrayList;")
		beans, imports, typeName, err := toPhpArray(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		set.AddStr(imports...)
		return beans, set.KeysStr(), typeName, nil
	case spec.InterfaceType:
		return nil, nil, "Object", nil
	case spec.Type:
		beans, err := getBean(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		imp, typeName := toPhpType(parentPackage, v.Name)
		set.AddStr(imp)
		return beans, set.KeysStr(), typeName, nil
	case Type:
		return nil, nil, v.Name, nil
	default:
		return nil, nil, "", fmt.Errorf("unsupported type: %v", v)
	}
}

func unJsonMarshal(expr interface{}, inter bool) interface{} {
	m := expr.(map[string]interface{})
	data, err := json.Marshal(expr)
	if err != nil {
		return expr
	}

	var basicType spec.BasicType
	var pointerType spec.PointerType
	var mapType spec.MapType
	var arrayType spec.ArrayType
	var interfaceType spec.InterfaceType
	var tpType Type

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	if sx.Contains(keys, "StringExpr") && sx.Contains(keys, "Name") && len(keys) == 2 {
		_ = json.Unmarshal(data, &basicType)
		return basicType
	}
	if sx.Contains(keys, "StringExpr") && sx.Contains(keys, "Star") && len(keys) == 2 {
		_ = json.Unmarshal(data, &pointerType)
		return pointerType
	}
	if sx.Contains(keys, "StringExpr") && sx.Contains(keys, "Key") && sx.Contains(keys, "Value") && len(keys) == 3 {
		_ = json.Unmarshal(data, &mapType)
		return mapType
	}
	if sx.Contains(keys, "StringExpr") && sx.Contains(keys, "ArrayType") && len(keys) == 2 {
		_ = json.Unmarshal(data, &arrayType)
		return arrayType
	}
	if sx.Contains(keys, "StringExpr") && len(keys) == 1 {
		if inter {
			_ = json.Unmarshal(data, &interfaceType)
			return interfaceType
		}
		tpType.Name = fmt.Sprintf("%v", m[keys[0]])
		return tpType
	}

	return expr
}

func toPhpArray(parentPackage string, a spec.ArrayType) ([]*Bean, []string, string, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, a.ArrayType, false)
	if err != nil {
		typeName = ""
		return nil, nil, typeName, err
	}

	return beans, imports, fmt.Sprintf("$%s=array()", parentPackage), nil
}

func toPhpMap(parentPackage string, m spec.MapType) ([]*Bean, []string, string, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, m.Value, false)
	if err != nil {
		typeName = ""
		return nil, nil, typeName, err
	}

	return beans, imports, fmt.Sprintf("$%s=array()", parentPackage), nil
}

func toPhpType(parentPackage, goType string) (string, string) {
	switch goType {
	case "bool":
		return "", "bool"
	case "uint8", "uint16", "uint32", "int8", "int16", "int32", "int", "uint", "byte":
		return "", "int"
	case "uint64", "int64":
		return "", "float"
	case "float32":
		return "", "float"
	case "float64":
		return "", "float"
	case "string":
		return "", "String"
	default:
		return fmt.Sprintf("import %s.bean.%s;", parentPackage, goType), goType
	}
}

func NewTag(tagExpr string) *Tag {
	tagExpr = strings.ReplaceAll(tagExpr, "`", "")
	tagExpr = strings.ReplaceAll(tagExpr, `"`, "")
	commaIndex := strings.Index(tagExpr, ",")
	if commaIndex > 0 {
		tagExpr = tagExpr[:commaIndex]
	}
	splits := strings.Split(tagExpr, ":")
	var (
		tp, name string
	)
	if len(splits) == 2 {
		tp = splits[0]
		name = splits[1]
	}

	return &Tag{
		tp:   tp,
		name: name,
	}
}

func (t *Tag) IsJson() bool {
	return t.tp == "json"
}

func (t *Tag) IsPath() bool {
	return t.tp == "path"
}

func (t *Tag) IsForm() bool {
	return t.tp == "form"
}

func (t *Tag) GetTag() string {
	return t.name
}
