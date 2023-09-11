# Docker examples

This is project that contains the default golang server and Dockerfile.

## Usage

To run the server within the docker container clone the repo:

    git clone git@github.com:estacet/docker_examples.git

go to the project folder and create the docker image:

    docker build -t go_server:1.21 .

create and run the container based on your image:

    docker run --rm -p 8080:8080 go_server:1.21

