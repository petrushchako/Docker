#!/bin/bash

# Step 1: Clean and package the application using Maven
#echo "Building the Maven project..."
#mvn clean package -DskipTests

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker build -t helloworld-spring .

# Step 2: Run the Docker container
echo "Running the Docker container..."
docker run -d -p 8080:8080 helloworld-spring

echo "Application is running on http://localhost:8080"
