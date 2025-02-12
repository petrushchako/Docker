# Docker: Your First Project
Docker Foundations Professional Certificate

### Contents
- The Project
    - Application using Docker
    - Setting up your development environment
    - Reviewing a project on GitHub
- Creating a First Docker Project
    - Writing a Dockerfile
    - Building a Docker image from Dockerfile
- Pulling and Pushing Docker Images
    - Searching for images in Docker Hub
    - Working with custom images
    - Tagging and labeling images
    -  Working with a private image repository
    - Inspecting images
    - Removing images
- Running Docker Containers
    - Running a container
    - Listing containers
    - Inspecting containers
    - Reviewing container log files
    - Working with volumes and mounts
    - Reviewing your daily workflow


## Project Requirements: Big Star Collectibles Website Containerization

### Overview
Big Star Collectibles, a company specializing in high-quality, exclusive, and unique items, is in the process of containerizing its existing website. The objective is to modernize its DevOps processes, reduce errors, and enhance productivity by leveraging Docker for containerization.

### Goals
- Enable developers to configure their local development environments quickly.
- Streamline and standardize the deployment process across all environments.
- Eliminate inconsistencies in setup, especially for contractors working with different machine configurations.
- Reduce manual installations and improve automation to minimize errors.

### Technical Stack
- **Programming Language:** Python
- **Framework:** Flask (lightweight web application framework)
- **Containerization Tool:** Docker
- **Version Control & Repository Management:** Docker Image Repository

### Deliverables
The final project will include:
- Website files, including code and assets.
- A **Dockerfile** containing the necessary instructions to build a Docker image.
- A **Docker image** with the required software and dependencies to run the application.
- A **Docker image repository** for securely storing and managing the Docker image.
- A **Docker command cheat sheet** for quick reference and streamlined execution.

### Expected Benefits
- Developers can start working on the project quickly without manually installing dependencies.
- Consistent and reliable deployment environments for both full-time employees and contractors.
- Automated and error-free DevOps processes, reducing time spent on manual configurations.
- Improved efficiency in testing, development, and deployment workflows.

<br>

## Setting up your development environment
To set up your local development environment, you will need Docker Desktop, Visual Studio Code (VS Code), and the Docker extension for VS Code.

### Install Docker Desktop:
- Navigate to the Docker website and download Docker Desktop.
- It's available for Windows, Mac, and Linux.
- Once installed, open Docker Desktop and sign in with your Docker Hub credentials.
- If you don’t have a Docker Hub account, you can create one during the sign-in process.

### Install Visual Studio Code (VS Code):
- Download and install VS Code from code.visualstudio.com.
- It is available for Windows, Mac, and Linux.
- Install the Docker Extension for VS Code:
- Open VS Code and navigate to the Extensions tab.
- Search for "Docker" and install the Docker extension.
- This extension allows you to run Docker commands and manage images, containers, networks, and other resources within VS Code.

<br>

## Reviewing a Project on GitHub    
### **Accessing the Repository**  
- The GitHub repository contains branches for each section of the course:
    - https://github.com/LinkedInLearning/docker-your-first-project-4485003
- Each branch is named according to the chapter and movie number (e.g., `02_01` for Chapter 2, Movie 1).  

### **Navigating the Project**  
- To switch to a specific branch:  
  ```sh
  git checkout <branch_name>
  ```  
  - Example:  
    ```sh
    git checkout 02_01
    ```  
- To list all available branches:  
  ```sh
  git branch
  ```  


<br>

## Building a Docker Image from a Dockerfile  
To create a Docker image, we use the `docker build` command. Below is a step-by-step guide on how to build an image using a Dockerfile.  

> branch `02_02`

### **Understanding the Docker Build Command**  
The basic syntax for building a Docker image:  
```sh
docker build [OPTIONS] <PATH> | <URL>
```
- **Path**: Local directory containing the **Dockerfile** (most common).  
- **URL**: A remote repository (useful for CI/CD pipelines).  

### **Building an Image Using a Local Path**  
- The most common command:  
  ```sh
  docker build .
  ```
  - The `.` (dot) sets the **current directory** as the build context.  
  - Docker looks for a **Dockerfile** in this directory by default.  

### **Specifying a Custom Dockerfile**  
- If your **Dockerfile** is not in the root directory, specify its location:  
  ```sh
  docker build -f path/to/Dockerfile .
  ```
  - Useful when managing different Dockerfiles for **QA, staging, and production** environments.  

### **Optimizing Builds with Additional Options**  

|**Option**|**Description**|
|---|---|
|`--force-rm`|Removes intermediate containers after a failed build.|
|`--rm=true`|Keeps intermediate containers for debugging (default: true).|
|`--no-cache`|Forces a **clean build**, bypassing cached layers.|

Example:  
```sh
docker build --no-cache --force-rm .
```
- This ensures **fresh** layers and removes failed build artifacts.  

### **Viewing Available Build Options**  
To see more build options, use:  
```sh
docker build --help
```
- This displays all available flags and their descriptions.  

<br>

