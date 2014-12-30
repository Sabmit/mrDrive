#!/bin/sh

tor --quiet;
/elasticsearch/bin/elasticsearch -d;
chmod 755 /apps/crawling-2/start.php /apps/crawling-2/readData.php;
