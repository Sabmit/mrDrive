#What can be done to make it more powerfull ?

##ElasticSearch

This tool uses ElasticSearch (ES) to index all the data.
ES is really scallable, we can make a cluster of ES server with multiple shards / nodes.
To make this application lightweight, we installed our ES server inside our image ```sabmit/mrdrive```.
We can use Docker to deploy our ES cluster on multiple EC2 instances and then configure our tool to use this cluster.


##Go server

This server is written in golang, we can take advantage of its powerfull ```go``` instruction to make a multi threaded server where each request/search can be served by multiple thread.
