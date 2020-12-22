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
)

const (
	dirBean    = "bean"
	dirService = "service"
)

func mkDir(target string) (map[string]string, error) {
	m := map[string]string{
		dirBean:    filepath.Join(target, "bean"),
		dirService: filepath.Join(target, "service"),
	}

	for _, each := range m {
		abs, err := filepath.Abs(each)
		if err != nil {
			return nil, err
		}

		err = util.MkdirIfNotExist(abs)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
