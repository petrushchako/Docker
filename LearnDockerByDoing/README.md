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

### Solution

**Installing Docker**

1. Install the Docker prerequisites:

    `sudo yum install -y yum-utils device-mapper-persistent-data lvm2`

2. Using yum-config-manager, add the CentOS-specific Docker repo:

    `sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo`

3. Install Docker:

    `sudo yum -y install docker-ce`

**Enable the Docker Daemon**

1. Enable the Docker daemon:

    `sudo systemctl enable --now docker`

**Configure User Permissions**

1. Add the lab user to the docker group. This will allow `cloud_user` to use docker commands without `sudo`:

    `sudo usermod -aG docker cloud_user`

<br>

*Verify configuration after restart by checking groups user is assigned to*:

    `groups`

> Note: You will need to exit the server for the change to take effect.

**Run a Test Image**

1. Using docker, run the hello-world image to verify that the environment is set up properly:

`docker run hello-world`


<br><br><br>

## Working With Prebuilt Docker Images

**ABOUT THIS LAB**

After installation, the best way to familiarize yourself with Docker, is to run containers from a few prebuilt images. In this lab, you will explore Docker Hub for images that will run a website. Once you find suitable images, you will get them into your development environment, and begin experimenting. You will run, stop, and delete containers from those images. You will also learn how to use existing data in containers.



### Solution

**Explore Docker Hub**
- Sign in to Docker Hub.
- At the top of the page, search for "httpd".
- In the left-hand menu, filter for **Application Infrastructure**, and **Official Images**.
- Select the `httpd` project.
- At the top of the page, click the **Tags** tab.
- Under *latest*, select **linux/amd64**.
- Back in the list of available images, select nginx
- Review the *How to use this image* section.

<br>

**Get and View httpd**
- In the Docker Instance, verify that docker is installed:

    `docker ps`

- Using docker, pull the httpd:2.4 image:

    `docker pull httpd:2.4`

- Run the image:

    `docker run --name httpd -p 8080:80 -d httpd:2.4`

- Check the status of the container:

    `docker ps`
    
- In a web browser, test connectivity to the container:

    `<PUBLIC_IP_ADDRESS>:8080`

<br>

**Run a Copy of the Website in httpd**
- Clone the Widget Factory Inc repository:

    `git clone https://github.com/linuxacademy/content-widget-factory-inc`

- Change to the content-widget-factory-inc directory:

    `cd content-widget-factory-inc`

- Check the files:

    `ll`

- Move to the web directory:

    `cd web`

- Check the files:

    `ll`

- Stop the httpd container:

    `docker stop httpd`

- Remove the httpd container:

    `docker rm httpd`

- Verify that the container has been removed:

    `docker ps -a`

- Run the container with the website data:

    `docker run --name httpd -p 8080:80 -v $(pwd):/usr/local/apache2/htdocs:ro -d httpd:2.4`

- Check the status of the container:

    `docker ps`

- In a web browser, check connectivity to the container:

    `<PUBLIC_IP_ADDRESS>:8080`

<br>

**Get and View Nginx**
- Using docker, pull the latest version of nginx:

    `docker pull nginx`

- Verify that the image was pulled successfully:

    `docker images`

- Run the container using the nginx image:

    `docker run --name nginx -p 8081:80 -d nginx `

- Check the status of the container:

    `docker ps`

- Verify connectivity to the nginx container:

    `<PUBLIC_IP_ADDRESS>:8081`

<br>

**Run a Copy of the Website in Nginx**
- Stop the nginx container:

    `docker stop nginx`

- Remove the nginx container:

    `docker rm nginx`

- Verify that the container has been removed:

    `docker ps -a`

- Run the nginx container, and mount the website data:

    `docker run --name nginx -v $(pwd):/usr/share/nginx/html:ro -p 8081:80 -d nginx`

- Check the status of the container:

    `docker ps`

- In a web browser, verify connectivity to the container:

    `<PUBLIC_IP_ADDRESS>:8081`

- Stop the nginx container:

    `docker stop nginx`

- Remove the nginx container:

    `docker rm nginx`

- Verify that the container has been removed:

    `docker ps -a`


<br><br><br>

## Handcrafting a Container image

**ABOUT THIS LAB**

If you run your website from a pre-built base image, it will require a manual process to set up the container each time it runs. For repeatability and scalability, the container, and your website code should be made into an image. In this lab, you will start with a base webserver image, modify settings in the container for your website, and then create images from the container. You'll demonstrate the importance of small changes to your container, and how they affect your image. Lastly, you will use your new images to create containers to see your hard work in action.



