# Docker

## Builder main commands

|Command                    |Description|
|---                        |---|
|**`FROM`** _image_         | base image for the build|
|**`MAINTAINER`** _email_   | Name of maintainer (metadata)|
|**`COPY`** _path_ _dest_   | copy ___path___ from the context into the container at the location ___dst___|
|**`ADD`** _src_ _dst_      | same as **`COPY`** but untar archives and accepts http urls|
|**`RUN`** _args...         | run an arbitrary command inside the container |
|**`USER`** _name_          | set the default username|
|**`WORKDIR`** _path_       | set the default working directory|
|**`CMD`** _args..._        | set the default command|
|**`ENV`** _name_ _value_   | set an environment variable(s)|
