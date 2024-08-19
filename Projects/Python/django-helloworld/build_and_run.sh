#!/bin/bash

# Build the Docker image
docker build -t django-helloworld .

# Run the Docker container
docker run -p 8000:8000 django-helloworld
