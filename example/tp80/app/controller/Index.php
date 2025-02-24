<?php

namespace app\controller;

use app\BaseController;
use think\facade\View;
use Demo\GenLib\ApiException;
use Demo\GenLib\DemoApiClient;
use Demo\GenLib\Request;
use Throwable;

class Index extends BaseController
{
    public function index()
    {
        return View::fetch('index');
    }

    public function callDemo()
    {
        $message = [];
        try {
            $client = new DemoApiClient('127.0.0.1', 8888, 'http');
            $req = new Request();
            $req->getPath()->setName("me");
            $resp = $client->getFromName($req);
            $message = $resp->getBody()->toAssocArray();
        } catch (ApiException $e) {
            $message['error'] = $e->getMessage();
        } catch (Throwable $t) {
            $message['error'] = $t->getMessage();
        }

        return json($message);
    }
}
