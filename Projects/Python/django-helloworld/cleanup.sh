#!/bin/bash

# Step 1: Stop the running container
container_id=$(docker ps -q --filter ancestor=django-helloworld)

if [ -n "$container_id" ]; then
    docker stop $container_id
    echo "Stopped container $container_id"
else
    echo "No running container found for django-helloworld"
fi

# Step 2: Remove the container
docker rm $container_id
echo "Removed container $container_id"

# Step 3: Remove the Docker image
docker rmi django-helloworld
echo "Removed Docker image django-helloworld"
