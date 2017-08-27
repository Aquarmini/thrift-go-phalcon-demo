<?php

namespace App\Tasks\Test;

use App\Tasks\Task;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Protocol\TMultiplexedProtocol;
use Thrift\Transport\TBufferedTransport;

class UserTask extends Task
{

    public function mainAction()
    {

    }

    public function getAction()
    {
        $thrift = di('thrift');

        $thrift->setHost('127.0.0.1');
        $thrift->setPort('10086');
        $socket = $thrift->socket();

        // $transport = new TFramedTransport($socket, 1024, 1024);
        $transport = new TBufferedTransport($socket, 1024, 1024);
        $protocol = new TBinaryProtocol($transport);

        $app_protocol = new TMultiplexedProtocol($protocol, "app");
        $user_protocal = new TMultiplexedProtocol($protocol, "user");

        $client = new \MicroService\UserClient($user_protocal);

        $transport->open();

        $res = $client->getByUserId(100);
        print_r($res);
        echo PHP_EOL;

        $transport->close();
    }

}

