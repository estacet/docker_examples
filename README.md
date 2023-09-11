# Docker examples

This is project that contains the default golang server and Dockerfile.

## Go

To run the server within the docker container clone the repo:

    git clone git@github.com:estacet/docker_examples.git

go to the go project folder:

    cd go

and create the docker image:

    docker build -t go_server:1.21 .

create and run the container based on your image:

    docker run --rm -p 8080:8080 go_server:1.21


## JS

To run the server within the docker container clone the repo:

    git clone git@github.com:estacet/docker_examples.git

go to the js project folder:

    cd js

and create the docker image:

    docker build -t node_server:20 .

create and run the container based on your image:

    docker run --rm -p 3000:3000 node_server:20