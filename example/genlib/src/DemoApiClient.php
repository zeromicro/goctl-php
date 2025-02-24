<?php

namespace Demo\GenLib;

class DemoApiClient extends ApiBaseClient
{
    
    public function getFromName(
        $request, 
        $body=null
    ) {
        $result = $this->request(
            '/from/:name',
            'get',
            $request->getPath(),
            null,
            null,
            $body
        );
        
        $response = new Response();
        
        
        $response->getBody()
            ->setMessage($result['body']['message'])
            ;
        
        return $response;
        
    }
    
}