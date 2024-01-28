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

```Dockerfile
ports:
    - "3000"
    - "8000:80" # guest:host
```

```Dockerfile
# expose ports to linked services (not to host)
expose: ["3000"]
```

#### Commands

```Dockerfile
# command to execute
command: bundle exec thin -p 3000
command: [bundle, exec, thin, -p, 3000]
```

```Dockerfile
# over the entrypoint
entrypoint: /app/start.sh
entrypoint: [php, -d, vendor/bin/phpunit]
```



#### Environment variable

```Dockerfile
# environment vars
environment:
    RACK_ENV: development

environment:
    - RACK_ENV=development
```

```Dockerfile
# environment vars from file
env_file: .env
env_file: [.env, .development.env]
```



#### Dependencies

```Dockerfile
# makes the 'db' service available as the hostname 'database'
# (implies depends_on)
links:
    - db:database
    - redis
```

```Dockerfile
# make sure 'db' is available before starting
depends_on:
    - db
```



#### Other options

```Dockerfile
# make this service extend another
extends:
    file: common.yaml # optional
    service: webapp
```

```Dockerfile
volumes:
    - /var/lib/mysql
    - ./_data:/var/lib/mysql
```

### Advanced features

#### Labels

```Dockerfile
services:
    web:
        labels:
            com.example.description: "Accounting web app"
```

#### DNS servers

```Dockerfile
services:
    web:
        dns: 8.8.8.8
        dns:
        - 8.8.8.8
        - 8.8.4.4
```

#### Devices

```Dockerfile
services:
    web:
        devices:
            - "/dev/ttyUSB0:/dev/ttyUSB0"
```


#### External links

```Dockerfile
services:
    web:
        external_links:
            - redis_1
            - project_db_1:mysql
```


#### Hosts

```Dockerfile

```

#### services


