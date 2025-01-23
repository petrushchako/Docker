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


<br>

### Install Docker on a Mac with Docker Desktop


<br>

### Install Docker on Windows with Docker Desktop


<br>

### Install Docker on Linux




## Using Docker
### Exploring the Docker CLI


<br>

### Create a Docker container


<br>

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


