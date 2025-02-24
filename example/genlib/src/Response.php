<?php

namespace Demo\GenLib;

class Response
{
    
    private $body;
    

    public function __construct(){
        
        $this->body = new ResponseBody();
        
    }

    
    public function getBody() { return $this->body; }
    
}