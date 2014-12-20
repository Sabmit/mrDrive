#Tests for mrDrive

##Crawling-2

The crawler uses ```Tor``` to make hidden HTTP request.
Also, it uses ```Redis``` to stock data.

```sh
sudo docker run --name redis-server -d redis
sudo docker run -it --name mrDrive --link redis-server:db sabmit/mrdrive
```

Inside the container please make sure ```Tor``` is running (use the ```startup.sh``` script).

```sh
root@e605cd33c90b:/apps/crawling-2# ls -l
total 24
-rw-rw-r-- 1 root root  217 Dec 18 21:40 proxyConfiguration.ini
-rw-rw-r-- 1 root root 8281 Dec 19 04:15 proxyConnector.class.php
-rwxr-xr-x 1 root root  459 Dec 19 04:17 readData.php
-rwxr-xr-x 1 root root 1770 Dec 19 04:17 start.php
root@e605cd33c90b:/apps/crawling-2#
```

Use ```./start.php``` to get the data.
Use ```./readData.php``` to make sure the data is insite the redis db
