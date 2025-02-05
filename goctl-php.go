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

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
	"github.com/zeromicro/goctl-php/action"
)

var (
	version  = "20250205"
	commands = []*cli.Command{
		{
			Name:   "php",
			Usage:  "generates http client for php",
			Action: action.Php,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "namespace",
					Usage: "the namespace of php",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate http client code for php."
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-php: %+v\n", err)
	}
}
