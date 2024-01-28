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
|**Description**|**Command**|**Example**|
|---|---|---|
|Start a new container from image|**`docker run`** ___IMAGE___|
|... and assign it a name|**`docker run --name`** ___NAME IMAGE___|
|... and map a port|**`docker run -p`** ___HOSTPORT:CONTAINER_PORT IMAGE___|
|... and map all ports|**`docker run -P`** ___IMAGE___|
|... and start container in background|**`docker run -d `** ___IMAGE___|
|... and assign it a hostname|**`docker run --hostname`** ___HOSTNAME IMAGE___|
|... and add a dns entry|**`docker run --add-host`** ___HOSTNAME:IP IMAGE___|
|... and map a local dir into a container|**`docker run -v `** ___HOST_DIR:TARGET_DIR IMAGE___|
|... but change the entrypoint|**`docker run -it --entrypoint`** ___EXECUTABLE IMAGE___|
