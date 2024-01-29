# Docker

## Builder main commands

|**Command**                    |**Description**|
|---                            |---|
|**`FROM`** ___image___         | base image for the build|
|**`MAINTAINER`** ___email___   | Name of maintainer (metadata)|
|**`COPY`** ___path dest___     | copy ___path___ from the context into the container at the location ___dst___|
|**`ADD`** ___src dst___        | same as **`COPY`** but untar archives and accepts http urls|
|**`RUN`** ___args...___        | run an arbitrary command inside the container |
|**`USER`** ___name___          | set the default username|
|**`WORKDIR`** ___path___       | set the default working directory|
|**`CMD`** ___args...___        | set the default command|
|**`ENV`** ___name value___     | set an environment variable(s)|


## Run a new Container
|**Description**|**Command**|
|---|---|
|Start a new container from image<br><br><br>|**`docker run`** ___IMAGE___<br>Example: <br>`docker run nginx`|
|... and assign it a name<br><br><br>|**`docker run --name`** ___NAME IMAGE___<br>Example: <br>`docker run --name web nginx`|
|... and map a port<br><br><br>|**`docker run -p`** ___HOSTPORT:CONTAINER_PORT IMAGE___<br>Example:<br> `docker run -p 8080:80 nginx`|
|... and map all ports<br><br><br>|**`docker run -P`** ___IMAGE___<br>Example: <br>`docker run -P nginx`|
|... and start container in background<br><br><br>|**`docker run -d `** ___IMAGE___<br>Example:<br> `docker run -d nginx`|
|... and assign it a hostname<br><br><br>|**`docker run --hostname`** ___HOSTNAME IMAGE___<br>Example: <br>`docker run --hostname srv nginx`|
|... and add a dns entry|**`docker run --add-host`** ___HOSTNAME:IP IMAGE___|
|... and map a local dir into a container<br><br><br>|**`docker run -v `** ___HOST_DIR:TARGET_DIR IMAGE___<br>Example:<br> `docker run -v ~/:/usr/share/nginx/html nginx`|
|... but change the entrypoint<br><br><br>|**`docker run -it --entrypoint`** ___EXECUTABLE IMAGE___<br>Example: <br>`docker run -it --entrypoint bash nginx`|


## Manage Containers

|**Description**|**Command**|
|---|---|
|Show a list of running containers|**`docker ps`**|
|Show a list of all containers|**`docker ps -a`**|
|Delete container<br><br><br>|**`docker rm CONTAINER`**<br>Exameple:<br>`docker rm web`|
|Delete stopped containers|**`docker container prune`**|
|Stop a running container<br><br><br>|**`docker stop CONTAINER`**<br>Example:<br>`docker stop web`|
|Start stopped container<br><br><br>|**`docker start CONTAINER`**<br>Example<br>`docker start web`|
|Copy file from container to the host<br><br><br>|**`docker cp CONTAINER:SOURCE TARGET`**<br>Example:<br>`docker cp web:/index.html index.html`|
|Copy file from the host to container<br><br><br>|**`docker cp TARGET CONTAINER:SOURCE`**<br>Example:<br>`docker cp index.html web:/index.html`|
|Start shell inside of running container<br><br><br>|**`docker exec -it CONTAINER EXECUTABLE`**<br>Example:<br>`docker exec -it web bash`|
|Rename container<br><br><br>|**`docker rename OLD_NAME NEW_NAME`**<br>Example:<br>`docker rename 096 web`|
|Create image out of container<br><br><br>|**`docker commit CONTAINER`**<br>Example:<br>`docker commit web`|

## Manage Images

|**Description**|**Command**|
|---|---|
|Download an image<br><br><br>|**`docker pull IMAGE[:TAG]`**<br>Example:<br>`docker pull nginx `|
|Upload an image to repository<br><br><br>|**`docker push IMAGE`**<br>Example:<br>`docker push my-image:1.0`|
|Delete image|**`docker rmi IMAGE`**|
|Show list of local images|**`docker images`**|
|Delete dangling images|**`docker image prune`**|
|Detele all unused images|**`docker image prune -a`**|
|Build image from Dockerfile<br><br><br>|**`docker build DIRECTORY`**<br>Example:<br>`docker build .`|
|Tag an image<br><br><br>|**`docker tag IMAGE NEW_IMAGE`**<br>Example:<br>`docker tag ubuntu ubuntu:18.04`|
|Build and tag an image from a Docker<br><br><br>|**`docker build -t IMAGE DIRECTORY`**<br>Example:<br>`docker build -t my-image .`|
|Save image to tar file<br><br><br>|**`docker save IMAGE > FILE`**<br>Example:<br>`docker save nginx > nginx.tar`|
|Load image from a tar file<br><br><br>|**`docker load -i TARFILE`**<br>Example:<br>`docker load -i nginx.tar`|



## Info & Stats

|**Description**|**Command**|
|---|---|
|Show the log of a container<br><br><br>|**`docker logs CONTAINER`**<br>Example:<br>`docker logs web`|
|Show stats of running container|**`docker stats`**|
|Show processes of container<br><br><br>|**`docker top CONTAINER`**<br>Example:<br>`docker top web`|
|Show installed docker version|**`docker version`**|
|Get detailed info about an object<br><br><br>|**`docker inspect NAME`**<br>Example:<br>`docker inspect nginx`|
|Show all modified files in container<br><br><br>|**`docker diff CONTAINER`**<br>Example:<br>`docker diff web`|
|Show mapped ports of a container<br><br><br>|**`docker port CONTAINER`**<br>Example:<br>`docker port web`|




## Container management commands

|**Command**|**Description**|
|---|---|
|**`docker create`** _image_ [_command_]<br>**`docker run`** _image_ [_command_]| create container<br>= create + start |
|**`docker start`** _CONTAINER..._<br>**`docker stop`** _CONTAINER..._<br>**`docker kill`** _CONTAINER..._<br>**`docker restart`** _CONTAINER..._|start the container<br>* graceful stop<br>kill (SIGKILL) the container<br>= stop + start|
|**`docker pause`** _CONTAINER..._<br>**`docker unpause`** _CONTAINER..._|suspend the container<br>resume the container|
|**`docker rm`** [`-f`]** _CONTAINER..._|destroy the container|

`*` - send SIGTERM to the main process + SIGKILL 10 seconds later<br>
`**` - `-f` allows removing running containers (= `docker kill` + `docker rm`)


## Interacting with the container


|**Command**|**Description**|
|---|---|
|**`docker attach`** _CONTAINER_| - attach to running container (stdin/stdout/stderr) |
|**`docker cp`** _CONTAINER:PATH HOSTPATH_<br>**`docker cp`** _HOSTPATH CONTAINER:PATH| - copy files fomr container<br> - copy files into the container |
|**`docker export`** _CONTAINER_| - export the content of the container (tar archive) |
|**`docker exec`** _CONTAINER ARGS..._| - run a command in an existing container (usefull for debugging) |
|**`docker wait`** _CONTAINER_| - wait until the container terminates and return the exit code |
|**`docker commit`** _CONTAINER IMAGE_| - commit a new docker image (snapshot of the container) |


## Image transfer commands

|**Command**|**Description**|
|---|---|
|**`docker pull`** _REPO_[_:TAG_]<br>**`docker push`** _REPO_[_:TAG_]<br>**`docker search`** _TEXT_| - pull an image/repo from a registry<br>- push an image/repo from a registry<br>- search and image on the official registry |
|**`docker login`** ...<br>**`docker logout`** ...|- login to a registry<br>- logout from a registry|
|**`docker save`** _REPO_[_:TAG_]<br>**`docker load`**| - export an image/repo as a tarbal<br>- load images from a tarball |
|**`docker-ssh`**<br>| - proposed script to transfer images between two daemons over ssh |


## Examples
### docker exec
```shell
docker exec [options] CONTAINER COMMAND
    -d, --detach        # run in beackground
    -i, --interactive   # stdin
    -t, --tty           # interactive

docker exec CONTAINER tail logs/development.log
docker exec -t -i CONTAINER rails c
```

### create container from an image
```shell
docker create --name app_redis_1 \
--expose 6379 \
redis:3.0.2
```

