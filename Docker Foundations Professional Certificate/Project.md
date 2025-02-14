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



## Working with Custom Images
When you pull images from an image repository, the image is saved locally. You can manage and inspect these images using various Docker commands.

### Listing Local Images

To view your local images, use the following command:

```sh
docker image ls
```

By default, this command hides intermediate images. Below are some useful options:

- `-a`, `--all`: Show all images, including intermediate images.
- `--no-trunc`: Prevents truncation of output.
- `-q`, `--quiet`: Returns only image IDs, suppressing other columns.
- `--digests`: Displays the image digest (a unique hash that identifies an image).

### Filtering Images

The `docker image ls` command allows filtering using the `--filter` option. Some common filters include:

- `reference=<name>`: Isolates images with a specific name or tag.
- `before=<image>`: Lists images created before a given image (cannot specify a date/time).
- `since=<image>`: Lists images created after a given image.
- `dangling=true`: Shows untagged images (also known as dangling images).
- `label=<key>=<value>`: Filters images based on labels.

### Formatting Output

To format output, use the `--format` option with Go templates. Common placeholders:

- `.ID`: Image ID  
- `.Repository`: Image repository name  
- `.Tag`: Image tag  
- `.Size`: Image size  
- `.CreatedSince`: Time since image creation  
- `.CreatedAt`: Exact timestamp of image creation  
- `.Digest`: Image digest  

Example: Display only image names and tags in table format:

```sh
docker image ls --format "table {{.Repository}}\t{{.Tag}}"
```

### Viewing Images in VS Code

You can also view and manage your Docker images using the **Docker Extension** for VS Code. This provides a graphical interface to explore images, containers, and networks.


<br><br><br>

## Tagging and Labeling Images
Tags and labels help organize Docker images.

### **Tags**
Tags are used to identify distinct versions of an image. They are commonly used to mark releases so users can select between different versions.

- In a **Dockerfile**, the base image is specified with a tag:
  ```Dockerfile
  FROM python:3.9-alpine
  ```
  - Here, `python` is the image name.
  - `3.9-alpine` is the tag, specifying both version and operating system.
  
- **Tagging Conventions:**
  - Industry standard: `<version>-<OS>` (e.g., `3.9-alpine`).
  - Helps differentiate versions when multiple exist.

- **Tagging an Image During Build:**
  ```sh
  docker build -t myapp:1.0 .
  ```
  - `-t` specifies a tag.
  - If no tag is provided, `latest` is used by default.

- **Best Practices for Tags:**
  - Always specify the image version explicitly.
  - Ensure `latest` represents the most stable release.
  - Example of building multiple tags:
    ```sh
    docker build -t myapp:latest -t myapp:1.0 --no-cache .
    ```
  - The `--no-cache` option ensures images are rebuilt without using cached layers.

- **Viewing Image Tags:**
  ```sh
  docker image ls
  ```
  - Displays images, including repository, tag, and image ID.
  - Multiple tags can exist for the same image ID.

- **Applying Tags to an Existing Image:**
  ```sh
  docker tag myapp:1.0 myapp:v1
  ```
  - The same image gets a new tag without re-building.

- **Avoiding Overwriting Images:**
  - Always build new versions with a fresh image ID using:
    ```sh
    docker build -t myapp:2.0 .
    ```
  - This prevents accidental overwrites.


### **Labels**
Labels allow adding metadata to images using key-value pairs.

- **Adding Labels in a Dockerfile:**
  ```Dockerfile
  LABEL vendor="My Company" \
        version="1.0" \
        description="My application"
  ```
  - Uses key-value pairs.
  - `\` allows better readability in multiple lines.

- **Filtering Images by Labels:**
  ```sh
  docker image ls --filter "label=vendor=My Company"
  ```
  - Searches images based on label metadata.

### **Summary**
- **Tags** organize versions of an image.
- **Labels** provide metadata for filtering and categorization.
- **Use both** to keep images well-structured, especially in repositories.

<br><br><br>


## Working with a Private Image Repository

Docker Hub provides a platform for managing private image repositories, ensuring secure storage while allowing controlled access.

### **Creating a Private Repository**
1. Navigate to the **Repositories** section in Docker Hub.
2. Click **Create a New Repository** and provide:
   - A **name** for the repository.
   - A **description** (optional).
3. Configure repository settings:
   - **Visibility**: Choose **Public** or **Private**.
   - If set to **Private**, you can invite specific users by entering their Docker Hub usernames.

### **Logging into a Private Repository**
To authenticate and interact with your repository:

```sh
docker login
```
- Enter your Docker Hub **username** and **password** when prompted.
- After login, you can push and pull images from your private repository.

### **Pushing an Image to a Private Repository**
1. **Tag the image** to match the repository name:
   ```sh
   docker tag my-local-image my-dockerhub-username/my-repo-name:tag
   ```
   Example:
   ```sh
   docker tag big-star-collectibles sbenhoff/big-star-collectibles-repo:v1
   ```
2. **Push the image** to Docker Hub:
   ```sh
   docker push my-dockerhub-username/my-repo-name:tag
   ```
   Example:
   ```sh
   docker push sbenhoff/big-star-collectibles-repo:v1
   ```
3. **Verify the upload**:
   - Refresh the **Registries panel** in VS Code.
   - Check your repository in **Docker Hub**.


### **Pulling an Image from a Private Repository**
To retrieve an image from your private repository:
```sh
docker pull my-dockerhub-username/my-repo-name:tag
```
Example:
```sh
docker pull sbenhoff/big-star-collectibles-repo:v1
```
- If the image already exists locally, Docker will indicate that it is **up to date**.

<br><br><br>



## Inspecting Docker Images
Inspecting Docker images provides key details for troubleshooting and verifying configurations.

### **Using VS Code Docker Extension**
- **Right-click** any image tag.
- Select **Inspect** to view image details.

### **Using CLI to Inspect Images**
#### **1. Get the Image ID**
List all images and retrieve the **image ID**:
```sh
docker image ls
```
#### **2. Inspect the Image**
Use the `docker image inspect` command with the image ID:
```sh
docker image inspect <image_id>
```
Example:
```sh
docker image inspect 123456789abc
```
This displays detailed JSON output, including:
- **Full Image ID**
- **RepoTags**: All associated tags.
- **Creation Date**: Shown in UTC format.
- **Container Info**: If the image is linked to a container.
- **Config Section**: Includes environment variables, command execution, and labels.

#### **Key Information in Image Inspection**
- **RepoTags**: Lists all tags linked to the image.
- **Created**: Displays when the image was created.
- **Container Info**: If attached, it shows container details.
- **Config Section**:
  - **Environment Variables**: Verify paths, languages, and dependencies.
  - **Command Configuration**: The command that runs when a container starts (from the `Dockerfile`).
  - **Labels**: Metadata defined in the `Dockerfile`.

### **Filtering Image Inspection Results**
To extract specific details, format output using JSON.

#### **View Labels Only**
```sh
docker image inspect --format '{{ json .Config.Labels }}' <image_id>
```

#### **Extract Environment Variables**
```sh
docker image inspect --format '{{ json .Config.Env }}' <image_id>
```

This is useful when dealing with large images and only needing specific details.
By effectively inspecting images, you can troubleshoot and ensure configurations are correctly applied.

<br><br><br>



