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

package action

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zeromicro/goctl-php/generate"
)

func Php(ctx *cli.Context) error {
	pkg := ctx.String("package")
	std, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var plugin Plugin
	plugin.ParentPackage = pkg
	err = json.Unmarshal(std, &plugin)
	if err != nil {
		return err
	}

	return Do(plugin)
}
