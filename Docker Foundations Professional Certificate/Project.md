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


<br><br><br>

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


<br><br><br>

## Building a Docker Image from a Dockerfile  
To create a Docker image, we use the `docker build` command. Below is a step-by-step guide on how to build an image using a Dockerfile.  

> Branch > `02_02`

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

<br><br><br>

### Searching for Images in Docker Hub  
When creating a **Dockerfile**, you often need a **base image**. You can find suitable images using the `docker search` command, which searches the **Docker Hub Registry** for available images.  

> Branch > `03_01`

### **Basic Search Command**  
To search for a Python image:  
```sh
docker search python
```
This returns a list of images with:  
- **Name** – Image name  
- **Stars** – Number of users who liked the image  
- **Official** – Whether it’s an official image from Docker  
- **Automated** – Indicates if it’s automatically built  

### **Filtering Search Results**  

| Filter | Command Example | Description |
|---------|----------------|-------------|
| **Official Images Only** | `docker search --filter=is-official=true python` | Show only **official** images. |
| **Minimum Stars** | `docker search --filter=stars=100 python` | Show images with at least **100 stars**. |
| **Automated Builds** | `docker search --filter=is-automated=true python` | Show images built automatically from GitHub. |
| **Limit Results** | `docker search --limit=4 python` | Display only **top 4** results. |
| **No Truncate** | `docker search --no-trunc python` | Show full descriptions without truncation. |

### **Formatting Search Results**  
To customize output using **Go templates**:  
```sh
docker search --format "{{.Name}}\t{{.Description}}" python
```
- `\t` adds tab spacing for table formatting.  
- Available columns: `Name`, `IsOfficial`, `IsAutomated`, `Description`, `StarCount`.  

### **Pulling an Image from Docker Hub**  
Once you've identified an image, pull it using:  
```sh
docker pull python
```
- This pulls the **latest** version by default.  

To pull a **specific version** (recommended for stability):  
```sh
docker pull python:3.10
```
- Versions (tags) can be found on the **Docker Hub website**.  

### **Why Avoid `latest` in Dockerfiles?**  
Using `latest` can lead to **unexpected errors** due to automatic updates.  
- Instead, **explicitly specify** a version to ensure **stability**.  

<br><br><br>


