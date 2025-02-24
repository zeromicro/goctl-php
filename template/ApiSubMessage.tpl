<?php

namespace {{.Namespace}};

class {{.MessageName}}
{
    {{range $k, $v := .Properties}}
    private ${{CamelCase $k}};
    {{end}}

    {{range $k, $v := .Properties}}
    public function get{{PascalCase $k}}() { return $this->{{CamelCase $k}}; }
    public function set{{PascalCase $k}}($v) { $this->{{CamelCase $k}} = $v; return $this; }
    {{end}}

    public function toQueryString() {
        return http_build_query([
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k}},
            {{end}}
        ]);
    }

    public function toJsonString() {
        return json_encode([
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k}},
            {{end}}
        ], JSON_UNESCAPED_UNICODE);
    }

    public function toAssocArray() {
        return [
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k}},
            {{end}}
        ];
    }
}