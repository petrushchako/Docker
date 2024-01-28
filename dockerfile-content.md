# Dockerfile

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

<hr>

### Inheritance

```Dockerfile
FROM ruby:2.2.2
```

### Varaible

```Dockerfile
ENV APP_HOME /myapp
RUN mkdir $APP_HOME
```

### Initialize

```Dockerfile
RUN bundle install

WORKDIR /myapp

VOLUME ["/data"] #Specification for mount point

ADD file.xyz /file.xyz
COPY --chown=user:group host_file.xyz /path/container_file.xyz
```

### Onbuild

```Dockerfile
ONBUILD RUN bundle install
# when used with another file
```

### Commands

```Dockerfile
EXPOSE 5900
CMD ["bundle", "exec", "rails", "server"]
```
### Entrypoint

```Dockerfile
ENTRYPOINT ["executable", "param1", "param2"]
ENTRYPOINT command param1 param2
```
<br>_Configure a container that will run as executable_
```Dockerfile
ENTRYPOINT exec top -b
```
_This will use shell processing to substitute shell variable, and will ignore any **CMD** or **docker run** command line arguments._

### Metadata

```Dockerfile
LABEL version="1.0"
```

```Dockerfile
LABEL "com.example.vendor"="ACME incorporated"
LABEL com.example.label-with-value="foo"
```

```Dockerfile
LABEL description="This text illustrates \ 
that label-values can span mutiple lines."
```

### See also
- https://docs.docker.com/engine/reference/builder/