<?php

namespace Demo\GenLib;

class Request
{
    
    private $path;
    

    public function __construct(){
        
        $this->path = new RequestPath();
        
    }

    
    public function getPath() { return $this->path; }
    
}