#!/bin/bash

# Build the Docker image
docker build -t django_helloworld .

# Run the Docker container
docker run -p 8000:8000 django_helloworld
