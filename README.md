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
