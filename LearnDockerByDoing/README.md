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

<br>

### Solution

#### Get and Run the Base Image
- Retrieve the httpd image:

    `docker pull httpd:2.4`

- Run the image:

    `docker run --name webtemplate -d httpd:2.4`

- Check the status of the webtemplate container:

    `docker ps`

<br>

#### Install Tools and Code in the Container
- Log in to the container:

    `docker exec -it webtemplate bash`

- Run `apt update` and `install git`

    `apt update && apt install git -y`

- Clone the website code from GitHub:

    `git clone  https://github.com/linuxacademy/content-widget-factory-inc.git /tmp/widget-factory-inc`

- Verify that the code was cloned successfully:

    `ls -l /tmp/widget-factory-inc/`

- List the files in the htdocs/ directory:

    `ls -l htdocs/`

- Remove the index.html file:
    
    `rm htdocs/index.html`

- Copy the webcode from /tmp/ to the htdocs/ folder:

    `cp -r /tmp/widget-factory-inc/web/* htdocs/`

- Verify that they were copied over successfully:

    `ls -l htdocs/`

- Exit the container:

    `exit`

<br>

**Create an Image from the Container**
- Copy the Container ID:

    `docker ps`

- Create an image from the container:
    
    `docker commit <CONTAINER_ID> example/widgetfactory:v1`

- Verify that the image was created successfully:

    `docker images`

- Take note of the image size.

<br>

**Clean up the Template for a Second Version**
- Log in to the container:

    `docker exec -it webtemplate bash`

- Remove the cloned code from the /tmp/ directory:

    `rm -rf /tmp/widget-factory-inc/`

- Use apt to uninstall git and clean the cache:

    `apt remove git -y && apt autoremove -y && apt clean `

- Exit the container:

    `exit`

- Check the status of the container:

    `docker ps`

- Create an image from the updated container:

    `docker commit <CONTAINER_ID> example/widgetfactory:v2`

- Verify that both images are now running:

    `docker images`

- Delete the v1 image:

    `docker rmi example/widgetfactory:v1`

<br>

**Run Multiple Containers from the Image**
- Run multiple containers using the new image:
    
    `docker run -d --name web1 -p 8081:80 example/widgetfactory:v2`

    `docker run -d --name web2 -p 8082:80 example/widgetfactory:v2`

    `docker run -d --name web3 -p 8083:80 example/widgetfactory:v2`

- Check the status of the containers:

    `docker ps`

- Stop the base webtemplate image:

    `docker stop webtemplate`

- Verify that only the created containers are running:

    `docker ps`

- Using a web browser, verify that the containers are running successfully:
    
    `<SERVER_PUBLIC_IP_ADDRESS>:8081`

    `<SERVER_PUBLIC_IP_ADDRESS>:8082`

    `<SERVER_PUBLIC_IP_ADDRESS>:8083`


<br><br><br>

## Storing Container Data In Docker Volumes

**ABOUT THIS LAB**

Storing data within a container image is one option for automating a container with data, but it requires a copy of the data to be in each container you run. For static files, this can be a waste of resources. Each file might not amount to more than a few megabytes, but once the containers are scaled up to handle a production load, that few megabytes can turn into gigabytes of waste. Instead, you can store one copy of the static files in a Docker volume for easy sharing between containers. In this lab, you will learn how Docker volumes interact with containers. You will do this by creating new volumes and attaching them to containers. You'll then clean up space left by anonymous volumes created automatically by the containers. Finally, you'll learn about backup strategies for your volumes.

### Solution

#### Discover Anonymous Docker Volumes

- Check the docker images:

    `docker images`

- Run a container using the `postgres:12.1` repository:

    `docker run -d --name db1 postgres:12.1`
    
- Run a second container using the same image:

    `docker run -d --name db2 postgres:12.1`
    
- Check the status of the containers:

    `docker ps`
    
- List the anonymous volumes:

    `docker volume ls`
    
- Inspect the db1 container:

    `docker inspect db1 -f '{{ json .Mounts }}' | python -m json.tool`
    
- Inspect the db2 container:

    `docker inspect db2 -f '{{ json .Mounts}}' | python -m json.tool`
    
- Create a third container using the --rm flag:

    `docker run -d --rm --name dbTmp postgres:12.1`
    
- Check the status of the container:

    `docker ps -a`
    
- List the anonymous volumes:

    `docker volume ls`
    
- Stop the db2 and dbTmp containers:

    `docker stop db2 dbTmp`
    
- List the anonymous volumes:

    `docker volume ls`
    
- Check the status of the containers:

    `docker ps -a`

