# Learn Docker By Doing

**About the course**

You learn faster and better when you learn by doing. With that in mind, this course has been designed to teach core Docker fundamentals and features through a 100% hands-on experience. To accomplish this, our Training Architects have hand-selected a set of the best Docker Hands-on Labs we have to offer.

Everything in this course will be inside of a real Linux environment provisioned with whatever needed through our hands-on lab and Cloud Playground platform.

**Topics**

- **Docker Basics**
    - Running your first docker container
    - Deploying a static website to the container (Docker)
    - Build container images
    - Dockerize a node.js application
- **Docker Optimization**
    - Optimizing Docker builds with onBuild
    - Ignoring files during `docker build`
- **Storing Data and Networking in Docker**
    - Creating Data Containers
    - Container Networking with Links
    - Container Netowrking with Networks
    - Persist Data Volumes
- **Doing More with Docker**
    - Container Logging
    - Updating Container with Watchtower
    - Adding Metadata and Labels
    - Load Balancing Containers
    - Building Services with Docker Compose
- **Monitoring with Prometheus**
    - Monitoring Containers with Prometheus
    - Using Grafana with Prometheus for Alerting and Monitoring
- **Working with Docker Swarm**
    - Setting Up a Docker Swarm
    - Backing Up and Restoring a Docker Swarm
    - Scaling a Docker Swarm Service
- **Container Orchestration with Kubernetes**
    - Setting up a Kubernetes Cluster with Docker
    - Scaling Pods in Kubernetes
    - Creating a Helm Chart

<br><br><br>

## Initializing the Docker Environment 

**ABOUT THIS LAB**

Docker is the leading containerization platform. If you are using containers, you are likely using Docker. In order to work with Docker, you must have the Docker daemon, and CLI available. This lab teaches you how to set up your environment, so you can get started working with Docker today!

**Solution**

- Log in to the server using the credentials provided:

    - `ssh cloud_user@<PUBLIC_IP_ADDRESS>`

- **Installing Docker**

    1. Install the Docker prerequisites:

        `sudo yum install -y yum-utils device-mapper-persistent-data lvm2`

    2. Using yum-config-manager, add the CentOS-specific Docker repo:

        `sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo`

    3. Install Docker:

        `sudo yum -y install docker-ce`

- **Enable the Docker Daemon**

    1. Enable the Docker daemon:

        `sudo systemctl enable --now docker`

- **Configure User Permissions**

    1. Add the lab user to the docker group:

    `sudo usermod -aG docker cloud_user`

    > Note: You will need to exit the server for the change to take effect.

- **Run a Test Image**

    1. Using docker, run the hello-world image to verify that the environment is set up properly:

    `docker run hello-world`