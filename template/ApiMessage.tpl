<?php

namespace {{.Namespace}};

class {{.MessageName}}
{
    {{range $k, $v := .Properties}}
    private ${{$k}};
    {{end}}

    public function __construct(){
        {{range $k, $v := .Properties}}
        $this->{{$k}} = new {{$v}}();
        {{end}}
    }

    {{range $k, $v := .Properties}}
    public function get{{PascalCase $k}}() { return $this->{{$k}}; }
    {{end}}
}