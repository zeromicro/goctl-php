<?php

require __DIR__ . '/../vendor/autoload.php';

use Demo\GenLib\ApiException;
use Demo\GenLib\DemoApiClient;
use Demo\GenLib\Request;

echo 'start' . PHP_EOL;

try {
    $client = new DemoApiClient('127.0.0.1', 8888, 'http');
    $req = new Request();
    $req->getPath()->setName("me");
    $resp = $client->getFromName($req);
    echo $resp->getBody()->getMessage().PHP_EOL;
} catch (ApiException $e) {
    echo $e->getMessage().PHP_EOL;
} catch (Throwable $t) {
    echo $t->getMessage().PHP_EOL;
}
