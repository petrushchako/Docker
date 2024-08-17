#!/bin/bash

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker build -t react-helloworld .

# Step 2: Run the Docker container
echo "Running the Docker container..."
docker run -d -p 3000:3000 react-helloworld

echo "React application is running on http://localhost:3000"
