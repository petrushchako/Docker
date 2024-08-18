#!/bin/bash

# Step 1: Stop the running container
container_id=$(docker ps -q --filter ancestor=angular-hello-world)

if [ -n "$container_id" ]; then
    docker stop $container_id
    echo "Stopped container $container_id"
else
    echo "No running container found for angular-hello-world"
fi

# Step 2: Remove the container
docker rm $container_id
echo "Removed container $container_id"

# Step 3: Remove the Docker image
docker rmi angular-hello-world
echo "Removed Docker image angular-hello-world"
