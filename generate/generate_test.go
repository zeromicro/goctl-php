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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pluginJson = "{\"Api\":{\"Info\":{\"Title\":\"\",\"Desc\":\"\",\"Version\":\"\",\"Author\":\"\",\"Email\":\"\"},\"Types\":[{\"Name\":\"User\",\"Annotations\":null,\"Members\":[{\"Annotations\":null,\"Name\":\"Name\",\"Type\":\"string\",\"Expr\":{\"StringExpr\":\"string\",\"Name\":\"string\"},\"Tag\":\"\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false}]},{\"Name\":\"request\",\"Annotations\":null,\"Members\":[{\"Annotations\":null,\"Name\":\"Name\",\"Type\":\"string\",\"Expr\":{\"StringExpr\":\"string\",\"Name\":\"string\"},\"Tag\":\"`json:\\\"name\\\"`\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false},{\"Annotations\":null,\"Name\":\"User\",\"Type\":\"User\",\"Expr\":{\"StringExpr\":\"User\"},\"Tag\":\"`json:\\\"user\\\"`\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false}]},{\"Name\":\"response\",\"Annotations\":null,\"Members\":[{\"Annotations\":null,\"Name\":\"Out\",\"Type\":\"string\",\"Expr\":{\"StringExpr\":\"string\",\"Name\":\"string\"},\"Tag\":\"\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false}]}],\"Service\":{\"Name\":\"template\",\"Groups\":[{\"Annotations\":[{\"Name\":\"server\",\"Properties\":{\"group\":\"template\",\"jwt\":\"Auth\"},\"Value\":\"\"}],\"Routes\":[{\"Annotations\":[{\"Name\":\"handler\",\"Properties\":null,\"Value\":\"handlerName\"}],\"Method\":\"get\",\"Path\":\"/users/id/:userId\",\"RequestType\":{\"Name\":\"request\",\"Annotations\":null,\"Members\":[{\"Annotations\":null,\"Name\":\"Name\",\"Type\":\"string\",\"Expr\":{\"StringExpr\":\"string\",\"Name\":\"string\"},\"Tag\":\"`json:\\\"name\\\"`\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false},{\"Annotations\":null,\"Name\":\"User\",\"Type\":\"User\",\"Expr\":{\"StringExpr\":\"User\"},\"Tag\":\"`json:\\\"user\\\"`\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false}]},\"ResponseType\":{\"Name\":\"response\",\"Annotations\":null,\"Members\":[{\"Annotations\":null,\"Name\":\"Out\",\"Type\":\"string\",\"Expr\":{\"StringExpr\":\"string\",\"Name\":\"string\"},\"Tag\":\"\",\"Comment\":\"\",\"Comments\":null,\"Docs\":null,\"IsInline\":false}]}}]}]}},\"ApiFilePath\":\"/Users/anqiansong/go/src/github.com/zeromicro/goctl-php/test.api\",\"Style\":\"\",\"Dir\":\"/Users/anqiansong/go/src/github.com/zeromicro/goctl-php/php\"}"

func TestDo(t *testing.T) {
	var plugin Plugin
	err := json.Unmarshal([]byte(pluginJson), &plugin)
	assert.Nil(t, err)

	plugin.Dir = t.TempDir()
	plugin.ParentPackage = "com.gozero"
	assert.Nil(t, Do(plugin))
}
