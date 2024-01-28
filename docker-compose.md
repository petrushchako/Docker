# Docker Compose
> Documentation
> https://docs.docker.com/compose/gettingstarted/

|Command                |Description|
|---                    |---        |
|`docker—compose build` |Builds or rebuilds a service from the given Dockerfile|
|`docker—compose run`   |Allows user to run a one-off command in the service|
|`docker—compose up`    |Creates and runs service containers|
|`docker—compose down`  |Removes the containers, networks, images, and volumes related to the service|


### Basic example

```Dockerfile
# docker-compose.yaml
version: '2'

services:
    web:
        build: .
        # build from Dockerfile
        context: ./Path
        dockerfile: Dockerfile
        ports:
         - "5000:5000"
        volumes:
         - .:/code
    redis:
        image: redis
```

### Reference

#### Building
```Dockerfile
web:
    # build from Dockerfile
    build: .
```

```Dockerfile
# build from custom Dockerfile
build:
    context: ./dir
    dockerfile: Dockerfile.dev
```

```Dockerfile
# build from image
image: ubuntu
image: ubuntu:14.04
image: tutum/influxdb
image: example-registry:4000/postgresql
```




#### Ports




#### Commands




#### Environment variable




#### Dependencies




#### Other options




### Advanced features

#### Labels




#### DNS servers




#### Devices




#### External links





#### Hosts




#### services


