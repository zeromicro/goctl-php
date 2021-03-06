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
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/zeromicro/goctl-php/template"
)

func generateService(dir string, data IService) error {
	filename := filepath.Join(dir, "Service.php")
	base := filepath.Dir(filename)
	err := util.MkdirIfNotExist(base)
	if err != nil {
		return err
	}

	return util.With("service").Parse(template.Service).SaveTo(data, filename, true)
}
