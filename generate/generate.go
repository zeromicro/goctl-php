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
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util/console"
)

func Do(in Plugin) error {
	log := console.NewColorConsole()
	spec, err := in.Convert()
	if err != nil {
		return err
	}
	// mkdir
	dir, err := mkDir(in.Dir)
	if err != nil {
		return err
	}

	// generate bean
	for _, each := range spec.Beans {
		err = generateBean(dir[dirBean], *each)
		if err != nil {
			return err
		}
	}

	for _, item := range spec.Service.Routes {
		if item.HasRequest || item.HasResponse {
			if strings.TrimSpace(spec.Service.Import) != "" {
				spec.Service.Import = fmt.Sprintf("use %s\\Bean.*;", spec.Service.ParentPackage) + "\n" + spec.Service.Import
				continue
			}
			spec.Service.Import = fmt.Sprintf("use %s\\Bean.*;", spec.Service.ParentPackage)
			break
		}
	}

	// generate interface
	err = generateIService(dir[dirService], spec.Service)
	if err != nil {
		return err
	}

	// generate implement
	err = generateService(dir[dirService], spec.Service)
	if err != nil {
		return err
	}

	log.MarkDone()
	return nil
}
