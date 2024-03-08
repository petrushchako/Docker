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

<br><br><br><br>

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


<br><br><br><br>
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


<br><br><br><br>

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


<br><br><br><br>

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

<br>

#### Create a Docker Volume
    
- Create a Docker volume:

    `docker volume create website`
    
- Verify that the volume was created successfully:

    `docker volume ls`
    
- Copy the widget-factory-inc data to the website container:

    `sudo cp -r /home/cloud_user/widget-factory-inc/web/* /var/lib/docker/volumes/website/_data/`
    
- List the copied data:

    `sudo ls -l /var/lib/docker/volumes/website/_data/`

<br>


#### Use the Website Volume with Containers  
- Run a docker container with the website volume:

    `docker run -d --name web1 -p 80:80 -v website:/usr/local/apache2/htdocs:ro httpd:2.4`
    
- Check the status of the container:

    `docker ps`
    
- In a web browser, verify connectivity to the container:

    `<PUBLIC_IP_ADDRESS>:80`
    
- Run a second container with the --rm flag:

    `docker run -d --name webTmp --rm -v website:/usr/local/apache2/htdocs:ro httpd:2.4`
    
- Check the status of the containers:

    `docker ps`
    
- Stop the webTmp container:

    `docker stop webTmp`
    
- Check the status of the containers:

    `docker ps -a`
    
- Verify that the website can still be accessed through a web browser:

    `<PUBLIC_IP_ADDRESS>`

<br>

#### Clean Up Unused Volumes
    
- Clean up the unused volumes:

    `docker volume prune`
    
- Check the currently running containers:

    `docker ps -a `   

- Remove the db2 container:

    `docker rm db2`
    
- Clean up the unused volumes again:

    `docker volume prune`
    
- List the current volumes:

    `docker volume ls`

<br>

#### Back Up and Restore the Docker Volume  
- Switch to the root user:

    `sudo -i`
    
- Find where the website volume data is stored:

    `docker volume inspect website`

- Copy the **Mountpoint**.

- Back up the volume:

    `tar czf /tmp/website_$(date +%Y-%m-%d-%H%M).tgz -C /var/lib/docker/volumes/website/_data .`

- Verify that the data was backed up properly:

    `ls -l /tmp/website_*.tgz`

- List the contents of the tgz file:

    `tar tf <BACKUP_FILE_NAME>.tgz`

- Exit out of root:
    
    `exit`

- Run a new container using the website volume, and create a backup:

    `docker run -it --rm -v website:/website -v /tmp:/backup bash tar czf /backup/website_$(date +%Y-%m-%d-%H-%M).tgz -C /website .`

- Verify that the data was backed up properly:

    `ls -l /tmp/website_*.tgz`

- Switch to the root user:

    `sudo -i`

- Change to the /docker/volumes/ directory:

    `cd /var/lib/docker/volumes/`

- List the volumes:

    `ls -l`

- Change to the website/_data directory:

    `cd website/_data/`

- Remove the contents of the directory:

    `rm * -rf`

- Verify that the backups are still available:

    `ls -l /tmp/website_*.tgz`

- Restore the data to the current directory:

    `tar xf /tmp/<BACKUP_FILE_NAME>.tgz .`

- Verify that the data was restored successfully:

    `ls -l`

<br><br><br><br>

## Storing Container Data in AWS S3

### Introduction

Using Docker volumes is the preferred method of storing container data locally. Volume support is built directly into Docker, making it an easy tool to use for storage, as well as more portable. However, storing container data in Docker volumes still requires you to back up the data in those volumes on your own.

There is another option - storing your container data in the cloud. It's not a solution for every problem, but after this lab, you'll have another tool at your disposal.

This lab will show you how to mount an **S3 bucket** onto your local system as a directory. You will then mount that directory into your Docker container. We will use an **httpd** container to serve the contents of that bucket as a webpage, but you can use it to share any common data between containers.

This will demonstrate how flexible Docker can be. You can make changes to your bucket and all of your containers using the S3 bucket will near-instantly have access to the content.

<br>

### Solution

#### Configuration and Installation
- Install the `awscli`, while checking if there are any versions currently installed, and not stopping any user processes:

    `pip install --upgrade --user awscli`

- Configure the CLI:

    `aws configure`

- Enter the following:
    - AWS Access Key ID: <ACCESS_KEY_ID> (AKIAW3BU6B7QT624QTYM)
    - AWS Secret Access Key: <SECRET_ACCESS_KEY> (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)
    - Default region name: us-east-1
    - Default output format: json

- Copy the CLI configuration to the `root` user:

    `sudo cp -r ~/.aws /root`

- Install the `s3fs` package:
    
    `sudo yum install s3fs-fuse -y`

<br>

#### Prepare the Bucket

- Create a mount point for the s3 bucket:

    `sudo mkdir /mnt/widget-factory`

- Export the bucket name:

    `export BUCKET=<S3_BUCKET_NAME>`

- Mount the S3 bucket:

    `sudo s3fs $BUCKET /mnt/widget-factory -o allow_other -o use_cache=/tmp/s3fs`

- Verify that the bucket was mounted successfully:

    `ll /mnt/widget-factory`

- Copy the website files to the s3 bucket:

    `cp -r ~/widget-factory-inc/web/* /mnt/widget-factory`

- Verify the files are in the folder:

    `ll /mnt/widget-factory`

- Verify the files are in the s3 bucket:

    `aws s3 ls s3://$BUCKET`

<br>

#### Use the S3 Bucket Files in a Docker Container

- Run an `httpd` container using the S3 bucket:
    `docker run -d --name web1 -p 80:80 --mount type=bind,source=/mnt/widget-factory,target=/usr/local/apache2/htdocs,readonly httpd:2.4`

- In a web browser, verify connectivity to the container:

    `<SERVER_PUBLIC_IP_ADDRESS>`

- Check the bucket cache:

- Change to the /mnt/widget-factory/ directory:

    `cd /mnt/widget-factory`

- Create a new page within the bucket:

    `cp index.html newPage.html`

- In a web browser, verify that the new page is accessible:

    `<SERVER_PUBLIC_IP_ADDRESS>/newPage.html`

- Verify that the page was added to the bucket:

    `aws s3 ls $BUCKET`


<br><br><br><br>

## Storing Container Data in Google Cloud Storage

#### Configuration and Installation
- Export the `projnum` variable:

    `export projnum=$(curl http://metadata.google.internal/computeMetadata/v1/project/numeric-project-id -sH "Metadata-Flavor: Google")`

- Verify that the variable was set successfully:

    `echo $projnum`
    
- Export the BUCKET variable:
  
    `export BUCKET="widgetfactory-${projnum}"`

<br>

#### Prepare the Cloud Storage Bucket

- Using gsutil, create a new bucket:

    `gsutil mb -l us-central1 -c standard gs://$BUCKET`

- Verify that the gcsfuse repo is available on the server:

    `cat /etc/yum.repos.d/gcsfuse.repo`

- Install gcsfuse (Note: please wait a few minutes before continuing to allow the lab's startup scripts to release the yum lock).

    `sudo yum install -y gcsfuse`

- Update the fuse.conf file to allow the user to mount the bucket properly:

    `sudo sed -ri 's/# user_allow_other/user_allow_other/' /etc/fuse.conf`

- Configure the directories needed to mount the bucket:

    `sudo mkdir /mnt/widget-factory /tmp/gcs`

- Change ownership of the directories to the cloud_user:

    `sudo chown cloud_user: /mnt/widget-factory/ /tmp/gcs`

- Mount the bucket:

    `gcsfuse -o allow_other --temp-dir=/tmp/gcs $BUCKET /mnt/widget-factory/`

- Copy the website files into the bucket:

    `cp -r /home/cloud_user/widget-factory-inc/web/* /mnt/widget-factory/`

- List the contents of the bucket:

    `gsutil ls gs://$BUCKET`

<br>

#### Use the GCS Bucket in a Container

- Mount the directory into the Docker container:

    `sudo docker run -d --name web1 --mount type=bind,source=/mnt/widget-factory,target=/usr/local/apache2/htdocs,readonly -p 80:80 httpd:2.4`

- Using a web browser, verify connectivity to the container:

    `<SERVER_PUBLIC_IP_ADDRESS>`


<br><br><br><br>

## Storing Container Data in Azure Blob Storage

### Introduction

This lab shows how to mount a Blob Storage container onto our local system as a directory. We will then mount that directory into our Docker container. We will use an httpd container to serve the contents of that bucket as a webpage, but we can use it to share any common data between containers. This will demonstrate how flexible Docker can be. We can make changes to our bucket, and all our containers using the Blob Storage container will near-instantly have access to the content.

### Solution

#### Configuration and Installation

- Obtain the Azure login credentials:

    `az login`

- Copy the code provided by the command.
- Open a browser and navigate to https://microsoft.com/devicelogin.
- Enter the code copied in a previous step and click Next.
- Use the login credentials from the lab page to finish logging in.
- Switch back to the terminal and wait for the confirmation.

<br>

#### Prepare the Storage
- Find the name of the Storage account:

    `az storage account list | grep name | head -1`

- Copy the name of the Storage account to the clipboard.
- Export the Storage account name:

    `export AZURE_STORAGE_ACCOUNT=<COPIED_STORAGE_ACCOUNT_NAME>`

- Retrieve the Storage access key:

    `az storage account keys list --account-name=$AZURE_STORAGE_ACCOUNT`

- Copy the key1 "value" for later use.
- Export the key value:
  
    `export AZURE_STORAGE_ACCESS_KEY=<KEY1_VALUE>`

- Install blobfuse:

    `sudo rpm -Uvh https://packages.microsoft.com/config/rhel/7/packages-microsoft-prod.rpm`
    `sudo yum install blobfuse fuse -y`

- Modify the `fuse.conf` configuration file:

    `sudo sed -ri 's/# user_allow_other/user_allow_other/' /etc/fuse.conf`

<br>

#### Use the Azure Blob Storage Container

- Create necessary directories:

    `sudo mkdir -p /mnt/widget-factory /mnt/blobfusetmp`

- Change ownership of the directories:

    `sudo chown cloud_user /mnt/widget-factory/ /mnt/blobfusetmp/`

- Mount the Blob Storage from Azure:

    `blobfuse /mnt/widget-factory --container-name=website --tmp-path=/mnt/blobfusetmp -o allow_other`

- Copy website files into the Blob Storage container:

    `cp -r ~/widget-factory-inc/web/* /mnt/widget-factory/`

- Verify the copy worked:

    `ll /mnt/widget-factory/`

- Verify the files made it to Azure Blob Storage:

    `az storage blob list -c website --output table`

- Run a Docker container:

    `docker run -d --name web1 -p 80:80 --mount type=bind,source=/mnt/widget-factory,target=/usr/local/apache2/htdocs,readonly httpd:2.4`

- Once the command is complete, open a web browser and navigate to the public IP address of the server.

- Verify the website is up and running.


<br><br><br><br>

## Building Container Images Using Dockerfiles

**Introduction**

Creating a container image by hand is possible, but it requires manual processes. There has to be a more automatic way to build images. Manual processes do not scale and are not easily version controlled. Docker provides a solution to this problem - the Dockerfile. In this lab, you will create a Dockerfile to build an image, and host a static website.

<br>

### Solution

#### Build a First Version

- Change to the widget-factory-inc directory:

    `cd widget-factory-inc`

- Create a Dockerfile that uses httpd:2.4 as the base image:

    `vim Dockerfile`

- In the new file, insert the following:

    ```yml
    FROM httpd:2.4
    RUN apt update -y && apt upgrade -y && apt autoremove -y && apt clean && rm -rf /var/lib/apt/lists*
    ```

- Save the file.
- Verify that the file was saved successfully:

    `cat Dockerfile`

- Build the 0.1 version of the widgetfactory image using the Dockerfile:

    `docker build -t widgetfactory:0.1 .`

- Set variables to examine the image's size and layers:

    `export showLayers='{{ range .RootFS.Layers }}{{ println . }}{{end}}'`
    
    `export showSize='{{ .Size }}'`

- Compare the httpd and widgetfactory images:

    `docker images`

- Show the widgetfactory image's size:

    `docker inspect -f "$showSize" widgetfactory:0.1`

- Show the layers:

    `docker inspect -f "$showLayers" widgetfactory:0.1`

- Show the layers of the httpd:2.4 image:

    `docker inspect -f "$showLayers" httpd:2.4`

- Compare the layers. Are they the same?

<br>

#### Load the Website into the Container

- Open the Dockerfile:

    `vim Dockerfile`

- Remove the Apache welcome page from the image by adding the following:

    `RUN rm -f /usr/local/apache2/htdocs/index.html`

- Save the file.

- Build version 0.2 of the widgetfactory image:

    `docker build -t widgetfactory:0.2 .`

- Inspect both versions of the widgetfactory image to see the differences in size and layers:

    `docker images`

- Show the widgetfactory:0.1 image's size:

    `docker inspect -f "$showSize" widgetfactory:0.1`

- Compare it to the image size for widgetfactory:0.2:

    `docker inspect -f "$showSize" widgetfactory:0.2`

- Using an interactive terminal, check the htdocs folder for widgetfactory:0.2. Are the website files in the folder?:

    `docker run --rm -it widgetfactory:0.2 bash`
    
    `ls htdocs`

- Exit the container:

    `exit`

- Show the layers for the widgetfactory:0.1 image:

    `docker inspect -f "$showLayers" widgetfactory:0.1`

- Show the layers for the widgetfactory:0.2 image and compare the two:

    `docker inspect -f "$showLayers" widgetfactory:0.2`

- Open the Dockerfile:

    `vim Dockerfile`

- Add the website data to the container by adding the following to the end of the file:

    `WORKDIR /usr/local/apache2/htdocs`

    `COPY ./web .`

- Save the file.

- Build version 0.3 of the widgetfactory image:

    `docker build -t widgetfactory:0.3 .`

- Inspect versions 0.2 and 0.3 to see the differences in size and layers:

    `docker images`

- Show the widgetfactory:0.2 image's size:

    `docker inspect -f "$showSize" widgetfactory:0.2`

- Compare it to the image size for widgetfactory:0.3:

    `docker inspect -f "$showSize" widgetfactory:0.3`

- Show the layers for the widgetfactory:0.2 image:

    `docker inspect -f "$showLayers" widgetfactory:0.2`

- Show the layers for the widgetfactory:0.3 image and compare the two:

    `docker inspect -f "$showLayers" widgetfactory:0.3`

- Using an interactive terminal, check the htdocs folder for widgetfactory:0.3:

    `docker run --rm -it widgetfactory:0.3 bash`

- Are the website files in the folder?:

    `ls -l`

<br>

#### Run a Container from the Image

- Run a container from the widgetfactory:0.3 image. What commmand does it use to run the server? Remember to publish the web server port:

    `docker run --name web1 -p 80:80 widgetfactory:0.3`

- Exit the container:

    `CTRL+C`

- Check the status of the container:

    `docker ps -a`

- Start the container:

    `docker start web1`

- Using docker exec connect to the web1 container:

    `docker exec -it web1 bash`

- View the website files in the container:

    `ls -l`

- Exit the container:

    `exit`

- Retrieve the main website page from the container:

    `wget localhost`

- Compare it to the copy on the server:

    `diff index.html web/index.html`

<br><br><br><br>

## Docker Container Networking

### Introduction

Each container should serve a single purpose, such as running one application like a web server. Containers can be powerful by themselves, but when connected together, they are far more useful.

For example, a web server container can be connected to a database container to provide application storage. Docker provides multiple options for networking containers.

In this lab, you'll explore a few of the common types of networks that Docker supports, and learn how containers within those networks interact.

<br>


### Solution

#### Explore the Default Network

- List the default networks:

    `docker network ls`

- Run an httpd container named web1, without specifying a network, and see which network it uses:

    `docker run -d --name web1 httpd:2.4`
    `docker inspect web1`
    
- Take note of the IPAddress.

- Run a container using the busybox image, and see if you can connect to the web1 server:

    `docker run --rm -it busybox`

- Check the container's networking, and verify it is in the same IP range as web1:

    `ip addr`

- Ping the web1 container using the IP address:
    
    `ping <WEB1_IP_ADDRESS>` (ping 172.17.0.2)

- Attempt to ping the web1 container by name:

    `ping web1`

    > In bridge network, you can reach octher contaitners by IP address, but not by container's name.

- Attempt to access web1 using wget:

    `wget <WEB1_IP_ADDRESS>`

    > This works without publishing the port on the host, since you are on the same network.

- Exit the container:

    `exit`

<br>

#### Explore Bridge Networks

- Create a new bridge network named test_application:

    `docker network create test_application`

- Run an httpd container named web2, in the test_application network:

    `docker run -d --name web2 --network test_application httpd:2.4`

- Check the status of the container:

    `docker ps -a`

- Verify that web2 was added to the test_application network:

    `docker inspect web2`

- Run a container using the busybox image, and see if you can connect to the web2 server, within the test_application network:

    `docker run --rm -it --network test_application busybox`

- Check the container's networking, and verify it is in the same IP range as web2:

    `ip addr`

- Ping the web2 container using the IP address:

    `ping <WEB2_IP_ADDRESS>`

    > In non-dafult bridge network, you can still connect via IP address, but the added benefit is that you can also connect via container name. 

- Attempt to ping the web2 container by name:

    `ping web2`

- Using wget, attempt to access web2 with the hostname:

    `wget web2`

- Attempt to ping web1:

    `ping <WEB1_IP_ADDRESS>`

    > This will not work since web1 and web2 conatier are assigned to a different networks. 

- Attempt to access web1 using wget:

    `wget <WEB1_IP_ADDRESS>`

    > Does not work for the same reason as command above

- Exit the container:

    `exit`

<br>

#### Explore the Host Network

- Run an httpd container named web3 on the host network:

    `docker run -d --name web3 --network host httpd:2.4`

- Check the status of the container:

    `docker ps -a`

- Attempt to connect to web3 directly from the server:

    `wget localhost`

    > This command execution will return the index.html file

- Stop web3:

    `docker stop web3`

- Attempt to connect to web3 directly from the server again:

    `wget localhost`

- Start web3:

    `docker start web3`

- Run a container using the busybox image, and see if you can connect to the web3 server:
    `docker run --rm -it --network host busybox`
    `ping web3`

    > The `ping web3` will not be able to connect to web3 conatiner, since web3 is not published name on the host network. 

- Using wget, attempt to access localhost within the busybox image:

    `wget localhost`

- Attempt to ping web2:

    `ping <WEB2_IP_ADDRESS>`

    > `ping 172.18.0.2, as well as command below, should be reachable since we are pinging from the local host (host) network. 

- Attempt to ping web1:
  
    `ping <WEB1_IP_ADDRESS>`


<br><br><br><br>

## Dockerize a Flask Application

### Introduction
Migrating static content into containers was a great way to learn the basics of Docker, but real-world uses are usually dynamic, including full applications.

This lab will show you the process to dockerize a Flask application. Flask is a lightweight Python WSGI micro web framework, however, you won't need to know any Python to complete this lab.


#### Create the Build Files

- Change to the notes directory:

    `cd notes`

- List the files in the directory:

    `ls -la`

- Inspect the config.py file:

    `cat config.py`

- Crate the .dockerignore file:

    `vim .dockerignore`

- In the file, paste the following:
    ```yaml
    .dockerignore
    Dockerfile
    .gitignore
    Pipfile.lock
    migrations/
    ```
- Create the Dockerfile:

    `vim Dockerfile`

- In the file, paste the following:
    ```yaml
    FROM python:3
    ENV PYBASE /pybase
    ENV PYTHONUSERBASE $PYBASE
    ENV PATH $PYBASE/bin:$PATH
    RUN pip install pipenv

    WORKDIR /tmp
    COPY Pipfile .
    RUN pipenv lock
    RUN PIP_USER=1 PIP_IGNORE_INSTALLED=1 pipenv install -d --system --ignore-pipfile

    COPY . /app/notes
    WORKDIR /app/notes
    EXPOSE 80
    CMD ["flask", "run", "--port=80", "--host=0.0.0.0"]
    ```

<br>

#### Build and Setup Environment

- Build the notesapp image:

    `docker build -t notesapp:0.1 .`

- Check the status of the image:

    `docker images`

- List the containers:

    `docker ps -a`

- List the docker networks:

    `docker network ls`

- Run a container using the notesapp image and mount the migrations directory:

    `docker run --rm -it --network notes -v /home/cloud_user/notes/migrations:/app/notes/migrations notesapp:0.1 bash`

- Once connected to the container, enable SQLAlchemy:

    `flask db init`

- Check the migrations folder:

    `ls -l migrations`

- Create the files needed to configure the database:

    `flask db migrate`

- Apply the files:

    `flask db upgrade`

<br>

#### Run, Evaluate, and Upgrade

- Run a container using the `notesapp:0.1` image:

    `docker run --rm -it --network notes -p 80:80 notesapp:0.1`

- Using a web browser, navigate to the public IP address for the server.

- Sign up for a new account using an email address and password.

- Once you are signed up, log in to your account.

- Create your first note.

- Verify that you can edit the note.

- Back in the terminal, disable `Debug mode` by editing the `.env` file:

    `vim .env`

- Remove the `export FLASK_ENV='development'` line.

- Build the image again:

    `docker build -t notesapp:0.2 .`

- Run a container using the updated image:

    `docker run --rm -it --network notes -p 80:80 notesapp:0.2`

- In a web browser, navigate to the public IP address for the server, and log in to your account.

- Verify that you can add a second note.

- In the terminal, stop the container:

    `CTRL+C`

<br>

#### Upgrade to Gunicorn

- Check the Pipfile:

    `cat Pipfile`

- Run a container and modify the pip file:

    `docker run --rm -it -v /home/cloud_user/notes/Pipfile:/tmp/Pipfile notesapp:0.2 bash`

- Once connected, change to the /tmp directory:

    `cd /tmp`

- Add gunicorn to the list of dependencies:

    `pipenv install gunicorn`

- Exit the container:

    `exit`

- Verify that gunicorn was added under [packages]:

    `cat Pipfile`

- Modify the init.py script:

    `vim __init__.py`

- Beneath the import section, add the following:
    ```python
    from dotenv import load_dotenv, find_dotenv
    load_dotenv(find_dotenv())
    ```
- Open the Dockerfile:

    `vim Dockerfile`

- To the bottom of the file, make the following changes:
    ```yaml
    COPY . /app/notes
    WORKDIR /app
    EXPOSE 80
    CMD ["gunicorn", "-b 0.0.0.0:80", "notes:create_app()"]
    ```
<br>






<br><br><br><br>

## 
























<br><br><br><br>

## Container Logging

**Introduction**

In this lab, we are going to configure syslog on a Docker instance, configure Docker to use syslog instead of the JSON logging driver, and spin up two containers to test our configuration.

Let's get started. Open your terminal application, and log in to the live environment using the credentials provided on the lab instructions page.

### Solution

####  Configure Docker to Use Syslog
- Open the rsyslog.conf file.

    `vim /etc/rsyslog.conf`

- In the file editor, uncomment the two lines under Provides UDP syslog reception by removing `#`.
    ```yaml
    $ModLoad imudp
    $UDPServerRun 514
    ```

- Then, start the syslog service.

    `systemctl start rsyslog`

- Now that syslog is running, let's configure Docker to use syslog as the default logging driver. We'll do this by creating a file called daemon.json.
    ```yml
    mkdir /etc/docker
    vi /etc/docker/daemon.json
    ```

- In the vi editor, enter the following, making sure to replace <PRIVATE_IP> with the private IP of your cloud server:
    ```json
    {
        "log-driver": "syslog",
        "log-opts": {
            "syslog-address": "udp://<PRIVATE_IP>:514"
        }
    }
    ```

- Next, save and quit. Then, start the Docker service.

    `systemctl start docker`

- The next step is to see if there are any logs coming in from Docker.

    `tail /var/log/messages`

- The output of this command tells us that there are. Now let's test our setup.

<br>

#### Verify configuration

- We're going to create two new containers using the httpd image. The first one will be called syslog-logging and will use none for the log driver. The second will be called json-logging and will use the JSON file for the log driver.

- Let's create our first container.

    `docker container run -d --name syslog-logging --log-driver none httpd`

- Then run `docker ps` to make sure our container is up and running correctly.

- Next, execute the `docker logs` command to see what logs we have.

    `docker logs syslog-logging`

- To confirm that we are logging to syslog, we can check the content of /var/log/messages. Verify that the syslog-logging container is sending its logs to syslog by running tail on the message log file.

    `tail /var/log/messages`

- The output shows us the logs that are being input to syslog.

- Now let's create our second test container. This time, we'll specify the log driver as the JSON file.

    `docker container run -d --name json-logging --log-driver json-file httpd`

- Run the `docker ps` command again. This time, we should see two containers running.

- Next, verify that the json-logging container is sending its logs to the JSON file. Execute the `docker logs` command on the json-logging container.

    `docker logs json-logging`

    > This time, the logs do not appear in /var/log/messages because they are being sent to a JSON file instead.