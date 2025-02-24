<?php

namespace Demo\GenLib;

class RequestPath
{
    
    private $name;
    

    
    public function getName() { return $this->name; }
    public function setName($v) { $this->name = $v; return $this; }
    

    public function toQueryString() {
        return http_build_query([
            'name' => $this->name,
            
        ]);
    }

    public function toJsonString() {
        return json_encode([
            'name' => $this->name,
            
        ], JSON_UNESCAPED_UNICODE);
    }

    public function toAssocArray() {
        return [
            'name' => $this->name,
            
        ];
    }
}