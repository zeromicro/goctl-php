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

package template

var IService = `<?php
namespace {{.ParentPackage}};

{{.Import}}
use okhttp3.RequestBody;
use retrofit2.Call;
use retrofit2.http.*;

public interface IService {
    {{range $index,$item := .Routes}}{{$item.Doc}}
	@{{$item.Method}}("{{$item.Path}}")
	Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} {{$item.MethodName}}({{if $item.HavePath}}{{$item.PathIdExpr}}{{end}}{{if $item.HaveQuery}}{{$item.QueryIdExpr}}{{end}}{{if $item.ShowRequestBody}}@Body RequestBody req{{end}});
	{{end}}
}`
