#!/bin/bash

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker build -t flask-helloworld .

# Step 2: Run the Docker container
echo "Running the Docker container..."
docker run -d -p 5000:5000 flask-helloworld

echo "Flask application is running on http://localhost:5000"
