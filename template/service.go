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

var Service = `<?php
namespace {{.ParentPackage}}\service;

{{.Import}}
use com.alibaba.fastjson.JSON;
use okhttp3.MediaType;
use okhttp3.RequestBody;
use retrofit2.Call;
use retrofit2.Callback;
use retrofit2.Retrofit;
use retrofit2.converter.gson.GsonConverterFactory;

public class Service {
    private static final String MEDIA_TYPE_JSON = "application/json; charset=utf-8";
    private static final String BASE_RUL = "http://localhost:8888/";// TODO replace to your host and delete this comment
    private static Service instance;
    private static IService service;

    private Service() {
        Retrofit retrofit = new Retrofit.Builder()
                ->baseUrl(BASE_RUL)
                ->addConverterFactory(GsonConverterFactory.create())
                ->build();
        service = retrofit->create(IService.class);
    }

    public static Service getInstance() {
        if (instance == null) {
            instance = new Service();
        }
        return instance;
    }

    private RequestBody buildJSONBody(Object obj) {
        String s = JSON.toJSONString(obj);
        return RequestBody.create(s, MediaType.parse(MEDIA_TYPE_JSON));
    }
	{{range $index,$item := .Routes}}{{$item.Doc}}
    public void {{$item.MethodName}}({{if $item.HasRequest}}{{$item.RequestBeanName}} in, {{end}}Callback{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} callback) {
        Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} call = service.{{$item.MethodName}}({{if $item.HavePath}}{{$item.PathId}}{{end}}{{if $item.HaveQuery}}{{$item.QueryId}}{{end}}{{if $item.ShowRequestBody}}buildJSONBody(in){{end}});
        call.enqueue(callback);
    }
	{{end}}
}
`
