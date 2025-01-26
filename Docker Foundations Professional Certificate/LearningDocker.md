# Learning Docker
#### **By:** Carlos Nunez
Containers help developers pack up their applications, ensuring they run smoothly anywhere. Software developers need to understand these technologies to deliver software as a team. Whether you are new to Docker or looking to review the fundamentals, these courses can help. Tune in, pass the final exam, and earn your certificate.

### Learning objectives
- Define Docker and explain common use cases.
- Understand how containers functionally and operationally differ from virtual machines.
- Explore three key technologies that make Docker different: layered containers, the Dockerfile, and the Docker API.
- Learn how to create and manage containers using the Docker CLI.
- Understand how to create custom container images using Dockerfiles.
- Learn how to push Docker images to the Docker registry and manage them.
- Troubleshoot common container issues using Docker CLI commands.
- Understand common best practices and problems when working with Docker containers and images.

<br>

### Table of contents
- **Docker Explained**
  - What is Docker?
  - Containers vs. virtual machines
  - The anatomy of a container
  - The Docker difference
- **Installing Docker**
  - Docker Desktop
  - Install Docker on a Mac with Docker Desktop
  - Install Docker on Windows with Docker Desktop
  - Install Docker on Linux
- **Using Docker**
  - Exploring the Docker CLI
  - Create a Docker container
  - Create a Docker container: The short way
  - Create a Docker container from Dockerfiles, part 1
  - Create a Docker container from Dockerfiles, part 2
  - Interact with your container
  - Stopping and removing the container
  - Binding ports to your container
  - Saving data from containers
  - Introducing the Docker Hub
  - Pushing images to the Docker registry
  - Checking your images in Docker Hub
  - Challenge: Starting NGINX
  - Solution: Starting NGINX
- **When Things Go Wrong**
  - Help! I can't seem to create more containers
  - Help! My container is really slow
  - Challenge: Fix a broken container
  - Solution: Fix a broken container
- **Additional Docker Resources**
  - Docker best practices
  - Taking it to the next level with Docker Compose
  - Level up even more with Kubernetes

<br><br><br>

## Docker Explained
### What is Docker?
#### Previous Solutions
- **Configuration managment tools (Chef, Puppet, Ansible)**<br>Require knowledge about hardware and operating system
- **Virtual machines as code (Vagrant)**<br>Heavy, slowish, require inconvenient configuration
#### Docker
Docker uses **images** and **containers** to allow apps to run anywhere, consistently.

<br>

### Containers vs. virtual machines
Containers and virtual machines are often compared, but they differ significantly. Here's a breakdown:

#### Key Differences

| **Area** | **VMs** | **Containers** |
|---|---|---|
|**Virtualization Layer**|Virtualize hardware via a hypervisor | Virtualize the operating system kernel via container runtimes.|
|**Operation**|Operate using a hypervisor that emulates hardware (memory, processors, disks, etc.).| Utilize the host operating system and hardware directly through container runtimes, avoiding emulation.|
|**Setup and Booting**|Require separate installation of operating systems and configuration for each VM.|Share the host's operating system and do not require a boot-up process, leading to faster application startups.|
|**Resource Usage**|Require virtual memory and disks, consuming significant disk space.|Do not need virtual memory or disks, enabling more applications to run simultaneously and using less space.|
|**Application Hosting**|Can host multiple apps simultaneously within one VM.|Designed to run only one app per container.|
|**Security and Isolation**|Apps running on VMs are isolated and cannot interact with the host directly.|Share the host's operating system, allowing visibility and potential modification of the host (a resolved security concern in most cases).|

#### Summary of Advantages
- **Virtual Machines**: Flexibility, isolation, and the ability to emulate a complete computer system.
- **Containers**: Faster startup times, lower resource consumption, and higher density of applications on the same hardware.

<br>

### The anatomy of a container
Containers rely on specific Linux kernel features to function. They consist of two primary components: **Linux Namespaces** and **Linux Control Groups (cgroups)**. Here's a breakdown:

#### Components of a Container

1. **Linux Namespaces**
   - **Purpose**: Provide isolation by exposing different "views" of the system to applications running within the container.
   - **Functionality**: Allow an application to perceive it has access to a full system while limiting actual access.
   - **Types of Namespaces**:
     |Namespace|Description|
     |---|---|
     |**USERNS**| Manages user views and permissions |
     |**MOUNT**| Controls file system access |
     |**NET**| Manages network communication capabilities |
     |**IPC**| Handles interprocess communication |
     |**TIME**| Controls the ability to modify time settings (not used by Docker containers) |
     |**PID**| Manages process IDs |
     |**CGROUP**| Controls creation and listing of control groups |
     |**UTS**| Manages host and domain names |
   - **Docker Limitation**: Docker uses all namespaces except for the TIME namespace.

2. **Linux Control Groups (cgroups)**
   - **Purpose**: Restrict and monitor hardware resource usage per process.
   - **Functions**:
     - Monitor and limit **CPU usage** (CPU time per container).
     - Restrict **network and disk bandwidth**.
     - Control **memory consumption**, preventing resource overuse by busy containers.
   - **Limitations**:
     - Cannot assign disk quotas directly; container-native storage solutions address this.

#### Key Characteristics of Containers

- **Lightweight Isolation**: Instead of virtualizing entire hardware components (as virtual machines do), containers use namespaces and cgroups for efficient resource management and isolation.
- **Linux Dependency**: 
  - Containers rely on Linux kernel features, making Docker natively compatible only with Linux (and newer versions of Windows).
- **Kernel Compatibility**: 
  - Containers are tied to the kernel they were created from.
  - **Linux Containers** run only on Linux.
  - **Windows Containers** run only on Windows.

#### Limitations and Workarounds
- Containers can't natively run across different kernels (e.g., Linux containers on Windows), but established workarounds exist to address these limitations.


<br>

### The Docker difference
#### A Brief History of Containers
- Containers have existed in various forms long before Docker. Here's a timeline of their evolution:
  - **1979: `chroot`**  
    - Introduced by Bill Joy, popularized in 1982 with 4.2 BSD.  
    - **Purpose**: Isolates file systems by changing the root directory for applications.  
    - **Limitation**: Applications in chroots could still interact with other applications on the host if they had the necessary libraries.
   
      ![](img/chroot.png)
      
  - **1999: BSD Jails**  
    - Extended the idea of chroots by creating isolated virtual environments for applications.  
    - Provided process and resource restrictions similar to modern containers.  
  - **2004: Solaris Zones**  
    - Introduced by Solaris, offering similar capabilities to BSD Jails.
   
      ![](img/BSD-Solaris.png)
  - **2007: Linux Containers (LXC)**  
    - Leveraged **control groups (cgroups)** and **namespaces** to isolate workloads.  
    - Supported restricting and isolating processes, networking, and resources.  

#### Why Docker Stands Out
Docker revolutionized the container ecosystem by making containerization more accessible to developers. Here are its key advantages:

1. **Ease of Configuration**  
   - **Dockerfiles**: Developers can define container environments in simple configuration files.  
   - Flexible and powerful configuration options.  
   - Automates the packaging of applications and their dependencies into Docker images.  

2. **Seamless Sharing**  
   - **Docker Hub**: A global repository maintained by Docker for sharing container images.  
   - Supports private and custom registries for image distribution.  

3. **Simplified Workflow**  
   - **Command-line Interface (CLI)**:  
     - Easy to create and start containers with commands like `docker run`.  
     - Simplifies UID mappings, network setups, and volume management.  
   - Developers don't need to handle the complexity of LXC configurations.  

<br><br><br>

## Installing Docker
### Docker Desktop
#### Containers on Non-Linux Systems
- **Background**:  
  Containers rely on Linux-specific features like **namespaces** and **control groups**, making Docker natively Linux-based. However, most developers use macOS or Windows, posing a challenge for Docker to support these platforms.

#### Early Solution: Docker Machine
- **What It Was**:  
  A tool that used Oracle’s **VirtualBox** to create a small Linux virtual machine (VM) solely for running the Docker engine. Developers had to run scripts to connect their Docker CLI to the VM.  

- **Limitations**:  
  1. **Complexity**:  
     - Users needed knowledge of VirtualBox and its CLI tool, `VBoxManage`, for tasks like:  
       - Exposing network ports.  
       - Mounting directories.  
     - This complexity slowed adoption.  
  2. **Performance Issues**:  
     - **Slow disk performance** when using mounted volumes.  
     - **Slow networking** when exposing ports.  
     - These issues stemmed from VirtualBox dependencies, which Docker couldn't control.  

#### Modern Solution: Docker Desktop
- **Introduced in 2016**:  
  A more efficient alternative to Docker Machine, improving usability and performance for Mac and Windows users.  

- **Key Features**:  
  1. **Performance**:  
     - Uses a smaller, faster VM:  
       - **Mac**: Apple’s native **Hypervisor.framework** (Virtual Kit).  
       - **Windows**: Microsoft’s **Hyper-V**.  
  2. **Ease of Use**:  
     - Automatic handling of:  
       - Mounting volumes.  
       - Exposing network ports.  
  3. **User-Friendly Interface**:  
     - Includes a GUI for:  
       - Configuring the VM.  
       - Managing Kubernetes clusters.  
       - Performing common Docker tasks.  

#### Licensing Changes and Alternatives
- **2021 License Update**:  
  Mirantis, Docker’s parent company, required companies with:  
  - **>250 employees** or  
  - **> $10 million revenue**  
  to purchase Docker subscriptions for Docker Desktop use.  

- **Emergence of Alternatives**:  
  - Due to licensing changes, new tools emerged as alternatives to Docker Desktop.  
  - Some developers reverted to using Docker Machine.  

<br>

### Install Docker on a Mac with Docker Desktop
#### Prerequisites
- **macOS Version**: 10.15 or newer.
- **Memory**: At least 4 GB.
- **Check System Information**:  
  1. Click the Apple icon in the upper left corner of your screen.  
  2. Select **About This Mac** to view your macOS version and memory.

#### Installation via Docker Website
1. **Visit Docker's Website**:
   - Go to [docker.com](https://www.docker.com).
   - Click the blue **Download** button for macOS.
   - If using an M1 or M2 chip, select the **Apple Chip** option.

2. **Download and Install**:
   - Open the downloaded `.dmg` file.
   - Drag the Docker icon into the **Applications** folder.

3. **Launch Docker**:
   - Press `Command + Space`, type **Docker**, and hit Enter.
   - macOS will prompt you to open Docker since it was downloaded from the internet. Click **Open**.
   - Enter your password when prompted to install backend components.
   - Accept the Docker license agreement.

4. **Verify Installation**:
   - Look for the Docker whale icon in the taskbar. When the boxes on the whale stop moving, Docker is ready.


#### Installation via Homebrew
1. **Install Homebrew**:
   - Visit [Homebrew's website](https://brew.sh).
   - Copy the Shell command provided.
   - Open **Terminal** (`Command + Space`, type **Terminal**, and press Enter).
   - Paste the copied command and press Enter.
   - Enter your password when prompted.

2. **Install Docker**:
   - In Terminal, type:
     ```bash
     brew install docker --cask
     ```
   - Homebrew will download, install, and configure Docker automatically.

3. **Launch Docker**:
   - Open Docker using `Command + Space`, search **Docker**, and press Enter.

<br>

#### Verify Docker is Working
1. Open **Terminal**.
2. Run the following command:
   ```bash
   docker run --rm hello-world
   ```
3. Docker will:
   - Pull the `hello-world` image.
   - Create a container from it.
   - Print a "Hello, world!" message.
   - Remove the container due to the `--rm` flag.

<br>

### Install Docker on Windows with Docker Desktop
#### Prerequisites
- **Windows Version**: Windows 10 or 11 (Docker requires Windows Pro, Enterprise, or Education editions for Hyper-V).
- **Requirements**: Verify all minimum requirements on Docker's website, such as:
  - Hardware virtualization enabled in BIOS.
  - At least 4 GB of memory.

#### Installation Steps
1. **Download Docker Desktop**:
   - Open your browser and visit [docker.com](https://www.docker.com).
   - Click the **Download** button for Windows.
   - Save the `.exe` file to your computer.
2. **Run the Installer**:
   - Locate and double-click the downloaded `.exe` file.
   - If prompted, ensure the publisher is verified as **Docker Inc.** and click **Yes** to continue.
3. **Select Backend**:
   - Choose the **WSL 2** (Windows Subsystem for Linux 2) backend:
     - Recommended by Docker as it performs nearly as well as a native Linux install.
     - Check the box for WSL 2 if not already selected.
   - Alternatively, you can choose **Hyper-V** to run Docker in a virtual machine.
4. **Shortcut Option**:
   - If you want a shortcut on your desktop, leave the box checked. Otherwise, uncheck it.
   - Click **OK** to start the installation.
5. **Complete Installation**:
   - Wait for the installation to complete.
   - If prompted, restart your computer.
6. **Launch Docker Desktop**:
   - Click the **Start** button, type **Docker Desktop**, and select it.
   - Accept the Docker Subscription Service Agreement.

#### Verify Docker Installation
1. **Check Initialization**:
   - Look at the **status bar** in the bottom left corner of Docker Desktop:
     - Orange = Initializing.
     - Green = Ready.
   - The Docker whale icon in the taskbar will have moving boxes while initializing. Hover over it to see the status.
2. **Address WSL 2 Update Errors**:
   - If prompted, install the latest WSL 2 update using the provided link.
   - Restart Docker Desktop after completing the update.
3. **Optional Tutorial**:
   - After initialization, Docker Desktop will present a tutorial wizard. You can start the tutorial or skip it.

<br>

### Install Docker on Linux
#### Overview
Docker works natively with Linux, so you don’t need Docker Desktop—only the **Docker Engine** and the **Docker CLI**. This guide uses Ubuntu, but the steps are similar for other distributions. For distribution-specific instructions, visit the [Docker documentation](https://docs.docker.com).

#### Step-by-Step Installation
1. **Install `curl`**:
   - Open a terminal and ensure `curl` is installed:
     ```bash
     sudo apt update
     sudo apt install curl
     ```
   - Enter your password if prompted.
2. **Download Docker's Installation Script**:
   - Use `curl` to download Docker's official installation script to the `/tmp` directory:
     ```bash
     curl -o /tmp/get-docker.sh https://get.docker.com
     ```
   - *(Optional)* Review the script to ensure it meets your security standards:
     ```bash
     less /tmp/get-docker.sh
     ```
3. **Run the Installation Script**:
   - Execute the script to install Docker Engine:
     ```bash
     sh /tmp/get-docker.sh
     ```
   - The script will download and install Docker components. This might take a few minutes depending on your system.
4. **Test Docker Installation**:
   - Verify that Docker is installed by running a test container:
     ```bash
     sudo docker run hello-world
     ```
   - If successful, you’ll see a "Hello from Docker!" message.
5. **Add Your User to the Docker Group**:
   - To avoid using `sudo` every time you run Docker, add your user to the `docker` group:
     ```bash
     sudo usermod -aG docker $USER
     ```
   - Replace `$USER` with your username if necessary. To check your username:
     ```bash
     whoami
     ```
6. **Apply Group Changes**:
   - Normally, you’d need to log out and back in for the group changes to take effect. Alternatively, you can apply them without logging out:
     ```bash
     newgrp docker
     ```
   - Confirm your membership in the `docker` group:
     ```bash
     groups
     ```
7. **Run Docker Without `sudo`**:
   - Test Docker without `sudo` to confirm everything is set up:
     ```bash
     docker run hello-world
     ```
   - You should see the same "Hello from Docker!" message.

<br><br><br>

## Using Docker
### Exploring the Docker CLI
### Exploring the Docker CLI

#### Overview
The **Docker Command Line Interface (CLI)** is the primary tool for interacting with Docker containers. It’s straightforward to use and provides a variety of commands to manage containers, images, networks, and more.

#### Key Features of the Docker CLI
1. **Top-Level Commands**:
   - Commands like `run`, `pull`, `network`, and `image` allow you to perform high-level operations. For example:
     - `docker run` starts a container.
     - `docker pull` downloads an image.
     - `docker network` manages Docker networks.
2. **Subcommands**:
   - Some top-level commands have **subcommands**. For example:
     ```bash
     docker network --help
     ```
   - This will list subcommands like `connect`, `create`, and `disconnect`.
3. **Options and Flags**:
   - Many commands support additional **options** to customize behavior. For example:
     ```bash
     docker run --name my_container -d nginx
     ```
     - `--name my_container`: Names the container "my_container."
     - `-d`: Runs the container in detached mode (background).

#### Using `--help`
The `--help` flag is available for every Docker command and subcommand. It provides:
- A **description** of the command.
- A **usage example**.
- A list of **supported options**.

Examples:
- View help for `docker network`:
  ```bash
  docker network --help
  ```
- View help for a subcommand, such as `docker network create`:
  ```bash
  docker network create --help
  ```
If you try running a command incorrectly or without required arguments, Docker will often display a shorter usage message to guide you.

<br>

### Create a Docker container
#### Overview
Docker containers are instances of Docker images. While images contain the app, its environment, and configuration, containers are the running (or stopped) versions of those images. 

#### **The Long Way to Create a Docker Container**
1. **Pulling the Image**:
   - Containers are created from Docker images. If an image is not available locally, Docker automatically pulls it from Docker Hub or another registry.
   - Use the `docker container create` command to specify the image:
     ```bash
     docker container create hello-world:linux
     ```
     - `hello-world` is the image name.
     - `:linux` is the image tag (to specify the Linux-specific version).
2. **Checking the Container**:
   - After creating the container, Docker assigns it a unique **ID**. The container will be in the "created" state but not running.
   - Verify its status using:
     ```bash
     docker ps --all
     ```
     - The `--all` flag lists all containers, including stopped ones.
3. **Starting the Container**:
   - Start the created container:
     ```bash
     docker container start <container-ID>
     ```
     - Replace `<container-ID>` with the actual container ID or even the first few characters of it.
4. **Viewing Logs**:
   - If the container exited quickly (as with the `hello-world` image), view its output using:
     ```bash
     docker logs <container-ID>
     ```
     - You’ll see the "Hello from Docker!" message.
5. **Attaching to the Container**:
   - To attach your terminal to the container and view its output live:
     ```bash
     docker container start --attach <container-ID>
     ```

<br>

#### **The Short Way to Create and Run a Docker Container**
The long process can be simplified using `docker run`, which combines image creation, container creation, and starting it into a single command:

```bash
docker run hello-world:linux
```

- **What Happens**:
  - Docker pulls the image (if not already available locally).
  - Creates a container from the image.
  - Starts the container.
  - Attaches your terminal to the container’s output, so you see the friendly "Hello from Docker!" message immediately.


### Create a Docker container: The short way


<br>

### Create a Docker container from Dockerfiles, part 1


<br>

### Create a Docker container from Dockerfiles, part 2


<br>

### Interact with your container


<br>

### Stopping and removing the container


<br>

### Binding ports to your container


<br>

### Saving data from containers


<br>

### Introducing the Docker Hub


<br>

### Pushing images to the Docker registry


<br>

### Checking your images in Docker Hub


<br>

### Challenge: Starting NGINX


<br>

### Solution: Starting NGINX


<br><br><br>

## When Things Go Wrong
### Help! I can't seem to create more containers


<br>

### Help! My container is really slow


<br>

### Challenge: Fix a broken container


<br>

### Solution: Fix a broken container


<br><br><br>

## Additional Docker Resources
### Docker best practices


<br>

### Taking it to the next level with Docker Compose


<br>

### Level up even more with Kubernetes


