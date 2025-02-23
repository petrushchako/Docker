# Learning Docker Compose

### Table of contents
- **Understanding Docker Compose**
    -  Compose in the Docker tool ecosystem
    -  Docker Compose basics
    -  Where to use Docker Compose
- **Getting Started with Docker Compose**
    -  Writing a Docker Compose configuration
    - Core Docker Compose commands
- **Docker Compose Core Features**
    - Build arguments
    - Mounting volumes
    - Named volumes
    - Exposing ports
    - Enforcing start-up order
- **Dynamic Configurations in Docker Compose**
    - Named subsets of services
    - Multiple compose files
    - Environment variables


<br>

## Understanding Docker Compose

## Compose in the Docker tool ecosystem
### **Introduction to Docker Compose**
- Docker tutorials often focus on starting one or two containers.
- Real-world software systems are **more complex**, often involving multiple services with distinct configurations.
- Many applications require **third-party dependencies** and **service orchestration** (specific startup order).
- **Microservices architectures** can have hundreds of containers, making manual startup impractical.

### **The Role of Docker Compose**
- Docker Compose **simplifies** managing multiple containers.
- It is an **independent tool** included with most Docker distributions.
- Allows developers to define configurations in a single file rather than manually setting up containers.

### **Key Features of Docker Compose**
- **Configuration as Code**: Uses YAML files to define container settings.
- **Scaffolding for Services**: Specifies how services interact, their dependencies, and environment settings.
- **Order Control**: Ensures services start in the required order.
- **Ease of Use**: Reduces complexity by replacing multiple manual commands with a single configuration file.
- **No Additional Functionality**: It does not add new Docker features but makes existing ones easier to use.

### **Purpose of This Course**
- Clarifies the role and capabilities of Docker Compose.
- Provides detailed, practical examples of how to use Docker Compose effectively.
- Helps integrate Compose into Docker-based software systems efficiently.



<br><br><br>



## Docker Compose Basics
### Understanding Configuration as Code
Configuration as code is a key concept in Docker Compose. Configuration defines the settings required for a system to run, including:
- Persistent data locations
- Internal and external service communication
- Environment-specific values

Most configuration-as-code tools, including Docker Compose, are **declarative** rather than **procedural**.

<br>

### Declarative vs. Procedural Approaches
#### Procedural Configuration
- Involves executing a sequence of steps in a predefined order.
- Each step depends on the previous step and assumes a specific system state.
- Can lead to errors if the environment does not match the expected state.
- Example issue: Running a container twice without stopping the first instance may cause failures.
#### Declarative Configuration
- Specifies the desired end state rather than the steps to achieve it.
- Docker Compose ensures the system reaches the defined state automatically.
- Running the same configuration multiple times should always yield the same result.

<br>

### Advantages of Configuration as Code
- **Version Control**: Configuration files can be stored in a repository, making it easy to track and revert changes.
- **Self-Documenting**: Eliminates the need to recall commands manually.
- **Environment Management**: Different environments (e.g., development and testing) can have their own independent configuration files.

Docker Compose simplifies and standardizes Docker container management by ensuring consistent results and reducing manual intervention.



<br><br><br>

## Where to Use Docker Compose
Docker Compose was designed as a tool for managing containers on a single hosted server. It is particularly useful for local development, staging environments, and continuous integration (CI) testing. However, it is not designed for distributed systems and lacks functionality for running containers across multiple hosts.

### Ideal Use Cases
- **Local Development**: Developers can define and configure multi-container applications easily using `docker-compose.yml` files.
- **Staging Environments**: Useful for testing deployments before moving to production.
- **Continuous Integration (CI) Testing**: Automates testing processes by spinning up required services in a controlled environment.

### Limitations in Production
For production environments with high user traffic, Docker Compose is not ideal because:
- **No Independent Scaling**: All services in a `docker-compose.yml` file scale together, which can waste resources.
- **Lack of Auto Scaling**: It does not provide mechanisms for automatically adjusting the number of running containers based on demand.

### Example Scenario
Consider **KinetEco**, a clean energy business with two primary applications:
1. **Online Storefront**: Sells solar panels and inverters.
2. **Scheduler**: Manages professional equipment installation appointments.

If KinetEco runs a promotional sale, the storefront may require additional instances to handle increased web traffic. However, using Docker Compose, scaling up the storefront would also scale the scheduler, even though it does not need additional resources. This inefficiency makes Docker Compose unsuitable for production environments requiring dynamic scaling.

### Alternatives for Production
For scalable and efficient production deployments, dedicated container orchestration tools should be used, such as:
- **Docker Swarm**: A native clustering and orchestration solution for Docker.
- **Kubernetes**: A robust and widely used container orchestration platform that supports auto-scaling, load balancing, and distributed deployments.


<br><br><br>


## Getting Started with Docker Compose
## Writing a Docker compose configuration
To use Docker Compose, the first step is to create a configuration file within the application directory. This file must be in YAML format and named `docker-compose.yml`.

### Structure of a Docker Compose File
A Docker Compose configuration file typically includes the following:
1. **Version Specification**: Defines which version of Docker Compose is being used.
2. **Services Definition**: Lists all the containers needed for the application.
3. **Build or Image Instructions**: Specifies how to create or retrieve container images.

### Example Configuration
Let's consider an example for **KinetEco**, a clean energy business that operates an online storefront backed by a MySQL database.

```yaml
version: '3'

services:
  storefront:
    build: .  # Uses the Dockerfile in the current directory

  database:
    image: mysql  # Pulls the MySQL image from Docker Hub
```

### Key Components
- **`services`**: Defines individual containers.
- **`build`**: Points to the directory containing the Dockerfile for the custom application.
- **`image`**: Specifies a prebuilt image from a registry like Docker Hub.

### Naming Services
Service names in Docker Compose are customizable and should be meaningful. While you could name services arbitrarily (e.g., `storefront` as `Alice` and `database` as `Bob`), using clear, descriptive names ensures better maintainability.

### Conclusion
With this basic configuration, you now have a working `docker-compose.yml` file that defines multiple services, builds images locally, and fetches prebuilt images from Docker Hub. Future enhancements may include adding networking, environment variables, and volume definitions to optimize the deployment further.


<br><br><br>


## Core Docker Compose Commands
Docker Compose provides essential commands for managing the lifecycle of Docker services. The most commonly used commands include `up`, `down`, `stop`, and `restart`.

### Running Docker Compose Commands
Before running any Docker Compose command, navigate to the directory containing the `docker-compose.yml` file. All commands start with:
```sh
docker-compose
```

### Starting Services
- **`docker-compose up`**: Builds images (if necessary), creates containers, and starts all services defined in the `docker-compose.yml` file.
- To start a specific service:
  ```sh
  docker-compose up <service_name>
  ```
  Example:
  ```sh
  docker-compose up storefront
  ```
  This will start only the `storefront` service and its dependencies.
- Alternative individual commands:
  - `docker-compose build`: Builds the Docker images.
  - `docker-compose create`: Creates the containers but does not start them.
  - `docker-compose start`: Starts existing containers.

### Stopping Services
- **`docker-compose stop`**: Stops all running containers without removing them. Useful for saving resources.
- **`docker-compose down`**: Stops and removes all containers, networks, and artifacts created by `docker-compose up`.
- Alternative individual command:
  - `docker-compose rm`: Removes stopped containers.

### Restarting Services
- **`docker-compose restart`**: Stops and starts all running containers in one step. Equivalent to running `stop` followed by `start`.

### Getting Help
- **`docker-compose --help`**: Displays a list of available commands along with their descriptions.

### Summary
| **Command** | **Description** |
|---|---|
| `docker-compose up` | Builds, creates, and starts services |
| `docker-compose up <service_name>` | Starts a specific service |
| `docker-compose build` | Builds images without starting containers |
| `docker-compose stop` | Stops running containers but retains them |
| `docker-compose down` | Stops and removes all containers and artifacts |
| `docker-compose restart` | Restarts all running containers |
| `docker-compose --help` | Displays available commands and descriptions |



<br><br><br>



## Docker Compose Core Features
## Build Arguments and Environment Variables
Build arguments and environment variables enhance Docker builds by making them more flexible. While environment variables are accessible inside a running container, build arguments are only available at build time.

### Build Arguments
Build arguments are useful for specifying build-specific configurations, such as selecting a tool version or defining a cloud provider region.

#### Defining Build Arguments in Docker Compose
To specify a build argument in a `docker-compose.yml` file, the build path must use an explicit syntax:
```yaml
services:
  app:
    build:
      context: .
      args:
        REGION: us-east-1
```
- `context`: Specifies the build directory (location of the Dockerfile)
- `args`: Defines a list of build arguments.
- Arguments follow the syntax `ARG_NAME=VALUE` with no spaces.

These arguments are available only at build time and do not persist inside the running container.

### Environment Variables
Environment variables provide a way to pass runtime configurations into a container. They can be useful for setting application modes (`dev`, `test`, `prod`), enabling feature flags, or passing credentials.

#### Defining Environment Variables in Docker Compose
```yaml
services:
  app:
    environment:
      - RUNTIME_ENV=dev
```
If an environment variable is set on the host machine, it can be passed to the container without defining a value:
```yaml
services:
  app:
    environment:
      - RUNTIME_ENV
```
Before running `docker-compose up`, ensure the variable is set on the host:
```sh
export RUNTIME_ENV=dev
```

### Using an Environment File
Instead of listing each variable individually, Docker Compose allows using an environment file:
```yaml
services:
  app:
    env_file:
      - .env
```
The `.env` file should contain key-value pairs:
```
RUNTIME_ENV=dev
DB_USER=root
DB_PASSWORD=securepassword
```

### MySQL Example
Many official Docker images, such as MySQL, rely on environment variables:
```yaml
services:
  mysql:
    image: mysql
    env_file:
      - mysql_env_vars
```
The `mysql_env_vars` file might contain:
```
MYSQL_ROOT_PASSWORD=rootpass
MYSQL_USER=myuser
MYSQL_PASSWORD=mypassword
MYSQL_DATABASE=mydb
```

#### .env file extra
In an `.env` file, wrapping secrets (or any values) in quotation marks (`" "` or `' '`) is optional, but it can affect how the values are interpreted.

##### **Key Rules for Secrets in an `.env` File:**
1. **No Quotes Needed (Default Behavior)**
   - The value is taken as-is.
   ```ini
   DB_PASSWORD=mysecretpassword
   ```
   - The value of `DB_PASSWORD` will be `mysecretpassword`.

2. **Using Quotes (`"` or `'`)**
   - If wrapped in double or single quotes, the quotes become part of the value unless the parser trims them.
   ```ini
   DB_PASSWORD="mysecretpassword"
   ```
   - Some applications **strip the quotes**, so `DB_PASSWORD` is still `mysecretpassword`.
   - Others **retain the quotes**, making it `"mysecretpassword"` (including the quotes).

3. **Special Cases**
   - If the secret contains spaces or special characters, **use quotes**.
   ```ini
   DB_PASSWORD="my secret password"
   ```
   - Without quotes, only `DB_PASSWORD=my` would be read, and the rest would be ignored.

4. **Escape Characters in Quotes**
   - If using quotes, you might need to escape `$` or `\`:
   ```ini
   API_KEY="my\$ecretKey"
   ```
   - The `$` won't be interpreted as a variable.

### **Best Practice**
- If the value contains **spaces, special characters, or leading/trailing whitespace**, use **double quotes** (`"`).
- Otherwise, you can leave them unquoted.



### Summary
| **Feature** | **Availability** | **Use Case** |
|---|---|---|
| **Build Arguments** | Build Time Only | Configuring build-time settings |
| **Environment Variables** | Runtime | Setting application configurations |
| **Environment Files** | Runtime | Managing multiple variables easily |



<br><br><br>



## Mounting Volumes
Volumes are a Docker construct used for persisted storage, ensuring that important data is not lost when a container stops running.

### Defining a Volume
When mounting a volume to a Docker service, at a minimum, you must specify a **target** directory. Typically, you also include a **source** and an **access mode**:
- **Target**: The directory inside the running container where the volume data is mounted.
- **Source**: The location on the host machine where the volume data resides.
- **Mode**: Defines access permissions, such as read-write (default) or read-only.

### Syntax in Docker Compose
In `docker-compose.yml`, volumes are defined within a service using the `volumes` object. The basic syntax is:
```yaml
services:
  mysql:
    image: mysql:latest
    volumes:
      - source_path:target_path:mode
```
#### Example:
Mounting a MySQL data directory:
```yaml
services:
  mysql:
    image: mysql:latest
    volumes:
      - ./mysql:/var/lib/mysql
```
This mounts the `./mysql` directory from the host to `/var/lib/mysql` inside the container.

### Specifying Paths
Docker Compose follows Bash standards for specifying file paths:
- `./directory_name` → Relative to the `docker-compose.yml` file.
- `../directory_name` → One level above the `docker-compose.yml` file.
- `/absolute/path` → Specifies an absolute path on the host machine.

### Setting Access Modes
The default mode is **read-write (rw)**, but for read-only access, set it to **ro**:
```yaml
volumes:
  - ./mysql_dump:/docker-entrypoint-initdb.d:ro
```
This ensures the container can read the data but cannot modify it.

### Summary
| **Parameter**  | **Description** |
|---|---|
| **Target** | Directory inside the container where data is stored. |
| **Source** | Directory on the host machine providing the data. |
| **Mode** | Defines access level (`rw` for read-write, `ro` for read-only). |
