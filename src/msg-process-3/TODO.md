##Which enhancements could be realised ?

##ElasticSearch

This tool uses ElasticSearch (ES) to index all the data.
ES is really scallable, we can make a cluster of ES server with
multiple shards / nodes.
To make this application lightweight, we installed our ES server inside our
image ```sabmit/mrdrive```.
We can use Docker to deploy our ES cluster on multiple EC2 instances and then
configure our tool to use this cluster.

Also, our application does not use the ```script``` property to update the
document.
We should use this property if we want our application to handle large amount of
POST data.

##Go server
This server is written in Go, we can take advantage of its ```go``` instruction
to make a multi threaded server.

##Which extensions could be added ?
Discusing what are the possible extensions is quite difficult without knowing
the purpose of the application.
But here are a few possibilities.

###User Interface
We could add some charts to represent the data using GoogleChart.
We can add an admin area with login/password and create an 'UserManager'.

###Web Server
We could improve our search query to only find significant keywords.

###ES Server
