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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
)

func TestGenerateBean(t *testing.T) {
	err := generateBean(t.TempDir(), Bean{
		ParentPackage: "com.tal",
		Name:          stringx.From("user"),
		Members: []*Member{
			{
				Name:     stringx.From("id"),
				TypeName: "String",
			},
			{
				Name:     stringx.From("name"),
				TypeName: "String",
				Comment:  "// 姓名",
				Doc:      "// 姓名",
			},
			{
				Name:     stringx.From("age"),
				TypeName: "int",
				Doc:      "// 年龄",
			},
			{
				Name:     stringx.From("birthday"),
				TypeName: "String",
				Doc:      "// 生日",
			},
		},
		PathTag: []string{"id"},
		JsonTag: []string{"name", "age", "birthday"},
	})
	assert.Nil(t, err)
}
