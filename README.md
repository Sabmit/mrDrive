Tests for mrDrive
=================

#Setup
First, you will need to download and run two docker images by executing these commands :
```bash
sudo docker run --name redis-server -d redis
sudo docker run -it --name mrDrive --link redis-server:db sabmit/mrdrive
```
The ```mrDrive``` image is now running and linked to a redis database (which is used in the Test 2).

Inside the image please use the ```startup.sh``` script to start every services.

#Tests
##Graph-1
Inside the directory ```graph-1``` you will find two python scripts : ```main_networkx.py``` and ```main_scipy.py```, the later one is a simple implementation using scipy module, but it uses a compressed sparse graph so we can not find a list of different paths with the same weight.

The script using networkx is to use as follow :
```bash
Usage: python main_networkx.py [options] Source Target
Options:
  -h, --help         show this help message and exit
  --all              Print the set of minimal distance path from S to T
  -f FILE, --filename=FILE
                        Read the graph from FILE in csv format
```
By default it uses ```graph.csv``` to read each edges properties.
```bash
root@x:/apps/graph-1# python main_networkx.py --all 0 6
```
```bash
Shortest path from 0 to 6 = [0, 2, 7, 3, 6]
This path takes 5 steps and costs 1.51
The list of all path with weight 1.51 :
[[0, 2, 7, 3, 6]]
```
##Crawling-2
The crawler uses ```Tor``` to make hidden HTTP request.
Also, it uses ```Redis``` to stock data.
These scrips are written in PHP which could be executed in an HHVM environment.

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
Please make sure ```Tor``` is running by executing the ```startup.sh``` script.

##Msg-processing-3
For this Test, I decided to write a Rest API in golang.
First, build the project using the makefile.
Here are the steps :
```bash
root@x:/apps/msg-process-3# make vendor_get && make run &
```
The server is now running.
You can add some test data inside the database by executing this command :

```bash
root@x:/apps/msg-process-3# curl -XPUT 127.0.0.1:9200/mrdrive/keywords/_bulk?pretty --data-binary @datasetM.json;
```

You may notice that the index has been created during the building of the server.
You will now need to retreive the Ip of the running mrDrive image
Here is the command line to execute on your machine :
```bash
sudo docker inspect --format '{{ .NetworkSettings.IPAddress }}' mrDrive
```
On my machine, mrDrive is using this ip ```172.17.0.84```
I will now use my browse to see the dashboard at this address : ```http://172.17.0.84:8080```
The dasboard is written in JS using angularJS / Bootsrap.

You can check the ```TODO.md``` file for further informations.

##Classification-4
The classification is made in Python using a stochastic gradient descent (SGD) training algorithm.
The classifier uses the file ```data/train_test.csv``` to train (85% of the file) and test (15% of the file)
Simply run this command to classify products from ```data/data.csv```.
```bash
root@x:/apps/classification-4# python classifier.py 
``` 

