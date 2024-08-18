#!/bin/bash

# Step 1: Build the Docker image
docker build -t angular-hello-world .

# Step 2: Run the Docker container
docker run -d -p 4201:4200 angular-hello-world

# Step 3: Output the running status
echo "Application is running on http://localhost:4201"
