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
	"fmt"
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/goctl-php/template"
)

type (
	Bean struct {
		ParentPackage string
		Name          stringx.String
		Import        string
		Members       []*Member
		PathTag       []string
		FormTag       []string
		JsonTag       []string
	}
	Member struct {
		Name     stringx.String
		TypeName string
		Comment  string
		Doc      string
	}
)

func (b *Bean) IsJsonRequest() bool {
	return len(b.JsonTag) > 0
}

func (b *Bean) IsFormRequest() bool {
	return len(b.FormTag) > 0
}

func (b *Bean) HavePath() bool {
	return len(b.PathTag) > 0
}

func (b *Bean) GetQuery() string {
	var query []string
	for _, item := range b.FormTag {
		m := b.GetMember(item)
		if m == nil {
			continue
		}
		query = append(query, fmt.Sprintf(`@Query("%s") %s %s`, item, m.TypeName, m.Name.Untitle()))
	}
	return strings.Join(query, ", ")
}

func (b *Bean) GetMember(name string) *Member {
	for _, item := range b.Members {
		if item.Name.Lower() == strings.ToLower(name) {
			return item
		}
	}
	return nil
}

func generateBean(dir string, bean Bean) error {
	filename := filepath.Join(dir, bean.Name.ToCamel()+".php")
	base := filepath.Dir(filename)
	err := util.MkdirIfNotExist(base)
	if err != nil {
		return err
	}

	return util.With("bean").Parse(template.Bean).SaveTo(bean, filename, true)
}
