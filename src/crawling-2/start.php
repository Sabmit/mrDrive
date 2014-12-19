#!/usr/bin/env php
<?php

include("./proxyConnector.class.php");
require_once("../vendor/simple_html_dom.php");
require_once('../vendor/predis/autoload.php');

function process_data($html, $redisEnabled, $client) {
  $id = 0;

  foreach ($html->find('.item') as $item) {
    $label = $item->find(".nom h3", 0)->plaintext;
    $price = preg_replace(["/[^0-9,\.]/", "/,/"], ["", "."], $item->find(".prix h3", 0)->plaintext);

    if ($redisEnabled) {
      $client->hmset("product_$id", "label", $label, "price", $price);
    }
    echo "Product $id => Name = " . $label . "\n";
    echo "Product $id => Price = " . $price . "\n";
    $id++;
  }
}

$redisEnabled = true;

if (!empty($argv[1]) && $argv[1] === "-r")
  $redisEnabled = false;

$connection = proxyConnector::getIstance();
$dataPostJson = [["t" => ["zoneid" => "idZoneListProduits"]]];

echo "1 - Loading.\n";
$connection->launch('http://www.coursesu.com/legrandquevilly/AccueilMagasin/_mag', ""); //getting the cookie for this store
echo "2 - Getting data...\n";
$connection->launch('http://www.coursesu.com/listeproduits.vignettes.navtop.showall?t:ac=136246', "", null, json_encode($dataPostJson));

$page = $connection->getProxyData();
$html = new simple_html_dom();
$html->load(json_decode($page["return"])->content);

if (empty($html)) {
  echo "Bad data received\n";
  var_dump($page);
  exit;
}

if ($redisEnabled) {
  try {
    echo "2.5 - Connecting to redis server on " . $_ENV['DB_PORT'] . "...\n";
    $client = new Predis\Client($_ENV["DB_PORT"]);
    $client->connect();
  }
  catch (Exception $e) {
    echo $e->getMessage() . "\n";
    echo "Redis database seems to be unavailable\n";
    $redisEnabled = false;
  }
}

echo "3 - Processing data.\n";
process_data($html, $redisEnabled, $client);