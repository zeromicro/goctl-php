<?php

namespace Demo\GenLib;

class ResponseBody
{
    
    private $message;
    

    
    public function getMessage() { return $this->message; }
    public function setMessage($v) { $this->message = $v; return $this; }
    

    public function toQueryString() {
        return http_build_query([
            'message' => $this->message,
            
        ]);
    }

    public function toJsonString() {
        return json_encode([
            'message' => $this->message,
            
        ], JSON_UNESCAPED_UNICODE);
    }

    public function toAssocArray() {
        return [
            'message' => $this->message,
            
        ];
    }
}