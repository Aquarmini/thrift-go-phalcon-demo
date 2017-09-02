<?php

namespace App\Logics\Thrift\Clients;

use App\Logics\Thrift\Client;

class UserClient extends Client
{
    protected $host = '127.0.0.1';

    protected $port = '10086';

    protected $service = 'user';

    protected $clientName = \MicroService\UserClient::class;

}

