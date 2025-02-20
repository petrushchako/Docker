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
    - Working with a private image repository
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


## Removing Docker Images
Cleaning up unused images helps maintain system efficiency by freeing up disk space.
### **1. Remove a Specific Image by Name and Tag**
```sh
docker rmi <image_name>:<tag>
```
Example:
```sh
docker rmi myapp:v2
```

### **2. Remove an Image by ID**
To delete all tags associated with an image:
```sh
docker rmi <image_id>
```
Example:
```sh
docker rmi 123456789abc
```
If the image is referenced in multiple repositories, you may get an error:
```
unable to delete (must be forced) - image is referenced in multiple repositories.
```
Use the `-f` (force) flag to override:
```sh
docker rmi -f <image_id>
```

### **3. Remove an Image by Digest**
List image digests:
```sh
docker images --digests
```
Remove an image using its digest:
```sh
docker rmi <repository>@<digest>
```
Example:
```sh
docker rmi myapp@sha256:abc123...
```

### **4. Remove Images Using VS Code**
- Open the **Docker extension** in VS Code.
- Right-click an image in the **Images panel**.
- Select **Remove**.

### **5. Remove Images from Docker Hub**
- Navigate to **Docker Hub**.
- Open your **repository** and go to the **Tags** section.
- Select the tags you want to delete.
- Click **Delete**.


<br><br><br>

## Running a container
### **Key Commands for Running Containers**
1. **Creating and Running a Container**
   ```sh
   docker run [OPTIONS] IMAGE[:TAG] [COMMAND] [ARG...]
   ```
   - Creates and starts a new container.
   - If the image does not exist locally, Docker pulls it from the registry.
2. **Starting an Existing Stopped Container**
   ```sh
   docker start CONTAINER_ID_or_NAME
   ```
   - Useful if you've already created the container but stopped it.
3. **Stopping vs. Killing a Container**
   ```sh
   docker stop CONTAINER_ID_or_NAME
   ```
   - Sends **SIGTERM** (graceful shutdown).
   
   ```sh
   docker kill CONTAINER_ID_or_NAME
   ```
   - Sends **SIGKILL** (forceful termination, immediate stop).

### **Common `docker run` Options**
| Option | Description |
|---|---|
| `-it` | Interactive mode + TTY (useful for debugging) |
| `-d` | Detached mode (runs in the background) |
| `-p HOST_PORT:CONTAINER_PORT` | Maps container port to host |
| `-v HOST_DIR:CONTAINER_DIR` | Mounts a volume for persistent data |

### **Example: Running a Flask App in a Container**
If your app listens on port **5000**, you can run:
```sh
docker run -d -p 5000:5000 my-flask-app
```
Then, open **http://localhost:5000** in your browser.

<br><br><br>


##  Listing and Filtering Containers in Docker
### **Basic Command:**
```sh
docker ps
```
- Lists only **running** containers.
- Displays **Container ID, Image, Status, Ports, and Name**.

### **Showing All Containers (Including Stopped)**
```sh
docker ps -a
```
- Includes stopped containers.

### **Filtering Containers**
| Option | Description |
|---|---|
| `-n <number>` | Shows the last `<number>` created containers (running or stopped). |
| `-q` | Displays only container IDs (useful for scripting). |
| `-s` | Shows the total file size of all containers. |
| `-l` | Shows the **last created** container. |
| `--filter "status=exited"` | Lists all containers that have **stopped**. |
| `--filter "ancestor=IMAGE_NAME"` | Lists containers based on a specific image. |
| `--filter "label=key=value"` | Filters based on a specific label. |
| `--format "{{json .}}"` | Outputs in JSON format. |
| `--format "table {{.ID}}\t{{.Image}}\t{{.Status}}"` | Displays selected fields in a table. |

### **Checking Exit Codes**
To find failed containers:
```sh
docker ps -a --filter "status=exited"
```
- **Exit code `137`** → Container was **killed** (possibly due to memory limits).
- **Exit code `0`** → Container exited **normally**.
- **Exit code `1`** → General error inside the container.

### **Example: Show Only Names and Status**
```sh
docker ps --format "table {{.Names}}\t{{.Status}}"
```

<br><br><br>

## Inspecting Docker Containers
### **Basic Command**
```sh
docker inspect <container_id_or_name>
```
- Returns detailed **JSON output** about the container.

### **Key Fields in `docker inspect` Output**
| Field | Description |
|---|---|
| `.Id` | The **Container ID**. |
| `.State` | Shows the container **status** (running, stopped, etc.). Also includes **PID** (process ID). |
| `.Image` | The **Image ID** used to create the container. |
| `.Name` | The **user-defined name** of the container. |
| `.RestartCount` | Number of times the container has been **restarted**. |
| `.HostConfig` | Contains CPU, memory limits, network settings, and mount paths. |
| `.LogPath` | The **log file path** on the host system. |
| `.Config` | Shows **environment variables, commands, and exposed ports**. |

### **Filtering Specific Information**
Instead of viewing the full JSON, you can extract only the required details:

| Command | Description |
|---|---|
| `docker inspect -f '{{.Id}}' <container>` | Show **only** the container ID. |
| `docker inspect -f '{{.Config.Image}}' <container>` | Show the **image name** used for the container. |
| `docker inspect -f '{{.State.Status}}' <container>` | Display the **container status** (running, exited, etc.). |
| `docker inspect -f '{{.LogPath}}' <container>` | Get the **log file path** of the container. |
| `docker inspect -f '{{json .HostConfig}}' <container>` | Show **HostConfig** details in JSON format. |

### **Example: Inspect a Running Container's Image**
```sh
docker inspect -f '{{.Config.Image}}' my-container
```

<br><br><br>


## **Reviewing container log files**
### **Basic Command**
```sh
docker logs <container_id_or_name>
```
- Displays **STDOUT** and **STDERR** logs of the running container.

### **Common Options**
| Option | Description |
|----|---|
| `-f` (follow) | Continuously streams logs in real time. |
| `--tail <num>` | Shows the last `<num>` lines of logs (e.g., `--tail 50`). |
| `--tail all` | Shows **all logs** from the container. |
| `--details` | Displays additional metadata in logs. |
| `--timestamps` | Adds **timestamps** to log output. |
| `--since <time>` | Shows logs **after** a specific time (`--since 10m` for last 10 min). |
| `--until <time>` | Shows logs **before** a specific time. |

### **Examples**
1. **View Logs with Live Updates**
```sh
docker logs -f my-container
```
- Streams new logs as they appear.
2. **Show Last 20 Log Entries**
```sh
docker logs --tail 20 my-container
```
3. **Show Logs from the Last 5 Minutes**
```sh
docker logs --since 5m my-container
```
4. **Show Logs Between Specific Time Range**
```sh
docker logs --since "2024-02-11T14:00:00Z" --until "2024-02-11T15:00:00Z" my-container
```
1. **Show Logs with Timestamps**
```sh
docker logs --timestamps my-container
```

<br><br><br>


## **Working with volumes and mounts**

### **Volumes**
- **Definition**: A volume is a type of data storage provided by Docker, which exists outside the container's filesystem but is accessible to it.
- **Use Case**: Best for persisting data that needs to survive container restarts or deletion. Commonly used for databases or any data you want to retain across container recreation.
  
#### **Commands for Volumes**
- **Create a volume**:
  ```sh
  docker volume create <volume_name>
  ```
- **List all volumes**:
  ```sh
  docker volume ls
  ```
- **Inspect a volume**:
  ```sh
  docker volume inspect <volume_name>
  ```
  This provides details like the mount point (the path inside the container).
- **Attach a volume to a container**:
  ```sh
  docker run -v <volume_name>:<container_path> <image_name>
  ```
- **Remove a volume** (only if it's not in use):
  ```sh
  docker volume rm <volume_name>
  ```

### **Bind Mounts**
- **Definition**: A bind mount is a way to link a directory on the host machine to a directory in a container. It is commonly used for development environments where you want to modify code on the host and immediately reflect those changes in the container.
- **Use Case**: Useful for sharing files, code, or configuration between the host and container.
- **Key Difference**: Data inside bind mounts is not persistent like volumes. If the container is deleted, data in bind mounts could be lost.

#### **Commands for Bind Mounts**
- **Create a bind mount**:
  ```sh
  docker run -v <host_directory>:<container_directory> <image_name>
  ```
  If the directory doesn't exist on the host, the `-v` flag will create it.
- **Verify Bind Mount**:
  1. Use `docker ps` to get the container ID.
  2. Access the container shell:
     ```sh
     docker exec -it <container_id> sh
     ```
  3. List files inside the container to verify the bind mount.


### **Key Differences between Volumes and Bind Mounts**

| **Feature**| **Volume**| **Bind Mount**|
|---|---|---|
| **Persistence**| Data persists even after container deletion | Data is lost when the container is removed |
| **Use Case**| Best for persisting application or database data | Best for sharing files (code/config) between host and container |
| **Storage Location** | Managed by Docker (stored outside container) | Stored directly on the host filesystem |
| **Ease of Use**| Docker handles management| Requires manual management of host directory |

<br><br><br>


### **Daily Docker Workflow and Cleanup**

When working with Docker daily, it is important to clean up unused objects such as images, containers, and volumes to free up system resources. Docker provides several commands to make this process efficient.

#### **1. Cleaning Up Images**
- **Remove dangling images (untagged and unreferenced by any container):**  
  ```sh
  docker image prune
  ```
  - Prompts before deletion (use `-f` to force).
- **Remove all images not used by existing containers:**  
  ```sh
  docker image prune -a
  ```
- **Bypass prompt:**  
  ```sh
  docker image prune -a -f
  ```
- **Filter images for selective pruning:**  
  ```sh
  docker image prune --filter "until=24h"
  ```
  (Removes images older than 24 hours.)

#### **2. Cleaning Up Containers**
- **Remove all stopped containers:**  
  ```sh
  docker container prune
  ```
- **Force removal without prompt:**  
  ```sh
  docker container prune -f
  ```
- **Remove containers stopped before a specific time:**  
  ```sh
  docker container prune --filter "until=24h"
  ```

#### **3. Cleaning Up Volumes**
- **Remove all unused volumes:**  
  ```sh
  docker volume prune
  ```
  - Volumes can only be filtered by labels, not timestamps.
  