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
  exit;
}

for ($i = 0; $i < 517; ++$i) {
  $label = $client->hget("product_$i", "label");
  $price = $client->hget("product_$i", "price");

  var_dump(trim($label));
  var_dump(floatval($price));

}
