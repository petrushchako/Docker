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
|Start a new container from image|**`docker run`** ___IMAGE___<br>Example: <br>`docker run nginx`|
|... and assign it a name|**`docker run --name`** ___NAME IMAGE___<br>Example: <br>`docker run --name web nginx`|
|... and map a port|**`docker run -p`** ___HOSTPORT:CONTAINER_PORT IMAGE___<br>Example:<br> `docker run -p 8080:80 nginx`|
|... and map all ports|**`docker run -P`** ___IMAGE___<br>Example: <br>`docker run -P nginx`|
|... and start container in background|**`docker run -d `** ___IMAGE___<br>Example:<br> `docker run -d nginx`|
|... and assign it a hostname|**`docker run --hostname`** ___HOSTNAME IMAGE___<br>Example: <br>`docker run --hostname srv nginx`|
|... and add a dns entry|**`docker run --add-host`** ___HOSTNAME:IP IMAGE___|
|... and map a local dir into a container|**`docker run -v `** ___HOST_DIR:TARGET_DIR IMAGE___<br>Example:<br> `docker run -v ~/:/usr/share/nginx/html nginx`|
|... but change the entrypoint|**`docker run -it --entrypoint`** ___EXECUTABLE IMAGE___<br>Example: <br>`docker run -it --entrypoint bash nginx`|
