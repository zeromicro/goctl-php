<?php

namespace {{.Namespace}};

class {{.MessageName}}
{
    {{range $k, $v := .Properties}}
    private ${{CamelCase $k false}};
    {{end}}

    {{range $k, $v := .Properties}}
    public function get{{CamelCase $k true}}() { return $this->{{CamelCase $k false}}; }
    public function set{{CamelCase $k true}}($v) { $this->{{CamelCase $k false}} = $v; return $this; }
    {{end}}

    public function toQueryString() {
        return http_build_query([
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k false}},
            {{end}}
        ]);
    }

    public function toJsonString() {
        return json_encode([
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k false}},
            {{end}}
        ], JSON_UNESCAPED_UNICODE);
    }

    public function toAssocArray() {
        return [
            {{range $k, $v := .Properties}}'{{$v}}' => $this->{{CamelCase $k false}},
            {{end}}
        ];
    }
}