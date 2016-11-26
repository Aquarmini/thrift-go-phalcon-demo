<?php

error_reporting(E_ALL);

define('BASE_PATH', dirname(__DIR__));
define('APP_PATH', BASE_PATH . '/app');

use Phalcon\Mvc\Application;

try {

    include APP_PATH . '/bootstrap.php';
    /**
     * Handle the request
     */
    $application = new Application($di);
    $application->useImplicitView(false);

    echo $application->handle()->getContent();
} catch (\Exception $e) {
    $error = $e->getMessage();
    echo $error;
    logger($error, 'error');
}
