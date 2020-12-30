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

var Bean = `<?php
namespace {{.ParentPackage}}\Bean;
{{.Import}}
public class {{.Name.ToCamel}} {
	{{range $index,$item :=  .Members}}{{$item.Doc}}
	private {{$item.TypeName}} ${{$item.Name.Untitle}}; {{$item.Comment}}
	{{end}}{{range $index,$item :=  .Members}}
	public {{$item.TypeName}} get{{$item.Name.ToCamel}}() {
		return $this->{{$item.Name.Untitle}};
	}

	public void set{{$item.Name.ToCamel}}({{$item.TypeName}} {{$item.Name.Untitle}}) {
		$this->{{$item.Name.Untitle}} = {{$item.Name.Untitle}};
	}
	{{end}}
}`
