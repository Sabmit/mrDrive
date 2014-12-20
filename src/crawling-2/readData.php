#!/usr/bin/env php
<?php

require_once('../vendor/predis/autoload.php');

try {
  $client = new Predis\Client($_ENV["DB_PORT"]);
  $client->connect();
}
catch (Exception $e) {
  echo $e->getMessage() . "\n";
  echo "Redis database seems to be unavailable\n";
  return 1;
}

$products = $client->keys("product:*");
if (empty($products)) {
  echo "No data received from " . $_ENV["DB_PORT"] . "\n";
  return 0;
}

foreach ($products as $product) {

  $productContent = $client->hmget($product, "label", "price");

  var_dump(trim($productContent[0]));
  var_dump(floatval($productContent[1]));

}
