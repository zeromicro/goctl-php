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
)

func TestGenerateService(t *testing.T) {
	err := generateService(t.TempDir(), IService{
		ParentPackage: "com.tal",
		Import:        "import com.tal.bean.*;",
		Routes: []*Route{
			{
				MethodName:      "register",
				Method:          "POST",
				Path:            "/api/user/register",
				RequestBeanName: "RegisterRequest",
				HasRequest:      true,
				ShowRequestBody: true,
				Doc:             "// 注册",
			},
			{
				MethodName:       "login",
				Method:           "POST",
				Path:             "/api/user/login",
				RequestBeanName:  "LoginRequest",
				ResponseBeanName: "LoginRespinse",
				HasRequest:       true,
				ShowRequestBody:  true,
				HasResponse:      true,
				Doc:              "// 登录",
			},
			{
				MethodName:       "getUserInfo",
				Method:           "GET",
				Path:             "/api/user/info/{id}",
				RequestBeanName:  "UserInfoRequest",
				ResponseBeanName: "UserInfoResponse",
				HasRequest:       true,
				ShowRequestBody:  false,
				HasResponse:      true,
				HavePath:         true,
				PathIdExpr:       `@Path("id")int id`,
				PathId:           "in.getId()",
				Doc:              "// 获取用户信息",
			},
			{
				MethodName:       "searchUser",
				Method:           "GET",
				Path:             "/api/user/search",
				RequestBeanName:  "UserInfoRequest",
				ResponseBeanName: "UserInfoResponse",
				HasRequest:       true,
				ShowRequestBody:  false,
				HasResponse:      true,
				HavePath:         false,
				HaveQuery:        true,
				QueryId:          `in.getKeyWord()`,
				QueryIdExpr:      `@Query("keyword")String keyword`,
				Doc:              "// 搜索用户信息",
			},
		},
	})
	assert.Nil(t, err)
}
