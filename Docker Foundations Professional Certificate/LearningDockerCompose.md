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



<br><br><br>



## Named Volumes in Docker Compose
Named volumes in Docker Compose allow for better management of persistent storage by linking volumes to specific containers. Unlike anonymous volumes, named volumes persist even when containers are removed, ensuring data is retained across container restarts.

### Defining Named Volumes
To define a named volume in a `docker-compose.yml` file, include a `volumes` section at the same level as `services` and provide a custom name for the volume.

Example:
```yaml
version: "3.8"
services:
  db:
    image: mysql:latest
    volumes:
      - db_data:/var/lib/mysql
volumes:
  db_data:
```
In this example:
- `db_data` is the named volume.
- It is mounted to `/var/lib/mysql` inside the `db` container.
- The `volumes` section at the bottom ensures `db_data` is managed by Docker Compose.

### Creating Named Volumes
Named volumes can be created explicitly using the `docker volume create` command:
```sh
docker volume create db_data
```
Alternatively, Docker Compose will automatically create the volume when `docker-compose up` is run if it is defined in the `docker-compose.yml` file.

### Mounting Named Volumes
Once the volume is created, you can mount it to a container manually:
```sh
docker run -d --name my_container -v db_data:/var/lib/mysql mysql:latest
```
This ensures that MySQL data is stored persistently in the `db_data` volume.

### Benefits of Named Volumes
1. **Persistent Storage**: Data is retained across container stops and removals.
2. **Automatic Data Management**: Docker handles volume storage and cleanup.
3. **Avoids Orphaned Volumes**: Prevents excessive disk usage from accumulating unnamed volumes.
4. **Data Migration**: Allows seamless data reuse between container instances.

### Managing Named Volumes
- **List all volumes**:
  ```sh
  docker volume ls
  ```
- **Inspect volume details**:
  ```sh
  docker volume inspect db_data
  ```
- **Remove unused volumes**:
  ```sh
  docker volume prune
  ```
- **Delete a specific named volume**:
  ```sh
  docker volume rm db_data
  ```

### Cleanup Considerations
- Running `docker-compose down` **without** `--volumes` retains named volumes.
- Adding `--volumes` removes named volumes, freeing up disk space.
  ```sh
  docker-compose down --volumes
  ```

Using named volumes ensures better control over persistent data storage while preventing unintentional data loss or excessive disk usage.



<br><br><br>



## Exposing Ports in Docker Compose
By default, Docker containers are isolated and do not expose any ports to the host machine or other containers. To allow communication, developers need to explicitly expose ports, mapping them between the host and the container.

### Port Mapping Syntax
Ports in Docker Compose are mapped using the following syntax:

```yaml
ports:
  - "<host_port>:<container_port>"
```

This means traffic sent to `<host_port>` on the host machine will be forwarded to `<container_port>` inside the container.

### Example Configuration
Example `docker-compose.yml` file:

```yaml
version: "3.8"
services:
  storefront:
    image: storefront:latest
    ports:
      - "80:80"
  scheduler:
    image: scheduler:latest
    ports:
      - "81:80"
```

#### Explanation:
- The `storefront` service exposes port **80** on the host and maps it to port **80** inside the container.
- The `scheduler` service is also configured to run on port **80** internally, but to prevent conflicts, it is mapped to port **81** on the host.

### Handling Port Collisions
If two services attempt to bind to the same host port, a port collision occurs. To avoid this, always ensure that the ports on the host machine are unique. For example:

```yaml
ports:
  - "8080:80"  # Maps container's port 80 to host port 8080
```

### Multiple Port Exposures
A service may require multiple ports to be exposed. For instance, a monitoring tool might collect metrics on a different port:

```yaml
version: "3.8"
services:
  storefront:
    image: storefront:latest
    ports:
      - "80:80"
      - "443:443"
```

Here, the `storefront` container exposes both port **80** (for HTTP traffic) and port **443** (for HTTPS traffic).

### YAML Parsing Considerations
When defining ports, it is recommended to use quotes around port mappings, especially for numbers below 60, due to YAML’s handling of base-60 integers. For example:

```yaml
ports:
  - "50:50"  # Correct
  - 50:50    # May cause parsing issues
```

### Key Takeaways
1. Ports are exposed using `<host_port>:<container_port>` syntax.
2. Avoid port collisions by assigning unique host ports.
3. Use quotes around port numbers, especially for values below 60.
4. Multiple ports can be mapped for a single service.
5. Port mapping allows external access to containerized applications.

Exposing ports is crucial for enabling communication between Docker services and external systems, making it one of the most common configurations in Docker Compose setups.



<br><br><br>



## Enforcing Start-up Order in Docker Compose
Many modern applications have **service dependencies**, meaning certain services cannot function properly unless others are already running. For example, a web application may rely on a database being up before it can fully initialize. Docker Compose offers tools to help manage **start-up order** for these types of architectures.

### Enforcing Start-up Order with `depends_on`
Docker Compose provides the `depends_on` key, which allows you to **define dependencies between services** in the `docker-compose.yml` file. This ensures:

- When running `docker-compose up`, the dependent service (like a database) starts **before** the dependent application (like a web app).
- When running `docker-compose down`, the dependent application stops **before** the database.
- Example:
    ```yaml
    version: "3.8"
    services:
      storefront:
        image: storefront-app:latest
        depends_on:
          - db
      db:
        image: mysql:latest
    ```

In this case:
- `storefront` depends on `db`.
- When you run `docker-compose up`, Compose will start `db` **first**, then `storefront`.
- When stopping the services, `storefront` will **stop first**, then `db`.


### Important Limitation: Start ≠ Healthy
**Compose does NOT guarantee that a dependent service is fully initialized or healthy** — it only guarantees the service **has been started**.

This is important in cases where:
- The `db` service might need a few seconds to initialize.
- Compose will start `storefront` immediately after `db` starts, even if `db` is not ready to accept connections yet.


### Starting Individual Services
If you want to start only one service (like `storefront`), Compose will automatically **start any services listed in `depends_on` first**.

```sh
docker-compose up storefront
```

This command will:
- Start `db` first.
- Start `storefront` next.


### Handling Health Checks (External Solutions)
Because Docker Compose does not natively **wait for services to become healthy**, if your application absolutely requires a dependency to be fully available, you will need to:

- Use **health check wrappers**.
- Use **init scripts** to check service availability (for example, wait-for-it scripts).
- Add retry logic directly inside your application.


### Recommendations
- **Avoid tightly coupling service start-up to the availability of dependencies.** 
- Design services to **handle temporary unavailability** of dependencies.
- Use **retry mechanisms** within your application code to reconnect if a dependency isn’t available at start-up.


### Key Commands
- Start all services in dependency order:
    ```sh
    docker-compose up
    ```
- Start one service and all its dependencies:
    ```sh
    docker-compose up storefront
    ```
- Stop all services in reverse dependency order:
    ```sh
    docker-compose down
    ```



<br><br><br>



## Dynamic Configurations in Docker Compose
## Named Subsets of Services (Service Profiles) in Docker Compose
Docker Compose allows you to create **named subsets of services** within a single `docker-compose.yml` file using **service profiles**. This is especially useful in larger projects where different teams or environments may only need specific sets of services running at any given time.

### Use Case Example
Imagine a company, Kinetico, with two engineering teams:
- One team works on **storefront** and services related to purchasing.
- Another team works on **scheduler** and services related to installation.

Both teams need to share a common **database**, but they do not need to run each other’s services during development. With **profiles**, you can group services and choose which subset to run.

### Defining Service Profiles
In the `docker-compose.yml`, you can add a `profiles` key under each service. This defines **which group (profile)** the service belongs to.

#### Example

```yaml
version: "3.8"

services:
  db:
    image: mysql:latest
    # No profile assigned, so db runs in the default profile (always available)

  storefront:
    image: storefront-app:latest
    profiles:
      - storefront_services

  scheduler:
    image: scheduler-app:latest
    profiles:
      - scheduling_services
```

In this example:
- `storefront` belongs to the **storefront_services** profile.
- `scheduler` belongs to the **scheduling_services** profile.
- `db` does not have a profile, so it automatically belongs to the **default** profile, meaning it runs in **all cases**, regardless of which profile is selected.

### Running Services with Profiles
#### Default Behavior
If you just run:
```sh
docker-compose up
```
Only services in the **default profile** (like `db`) will run.

#### Running Specific Profiles
To start only services in a specific profile (plus any default services like `db`), use:
```sh
docker-compose --profile storefront_services up
```

This starts:
- `db` (because it's in the default profile)
- `storefront` (because it’s part of `storefront_services`)

You can also stop, restart, or manage services using profiles:
```sh
docker-compose --profile scheduling_services down
```

### Multiple Profiles
You can assign a service to **more than one profile** if needed:

```yaml
storefront:
  image: storefront-app:latest
  profiles:
    - storefront_services
    - shared_services
```

This makes `storefront` part of **both `storefront_services` and `shared_services` profiles**, allowing flexibility when selecting service groups.

### Why Use Profiles?
- **Selective Start-up**: Avoid running unnecessary services, saving resources.
- **Team Separation**: Each team can run only the services they care about.
- **Shared Config**: You can still have **one Compose file** for all services, simplifying configuration management.
- **Flexible Testing**: Easily switch between isolated services and full integration environments.

### Key Benefits
- Simplifies **running only what you need**  
- Avoids **duplicating Compose files** for subsets of services  
- Works with all Compose commands like `up`, `down`, `restart`

### Key Commands Summary
- Run **default profile only**:
    ```sh
    docker-compose up
    ```
- Run **storefront services and default services**:
    ```sh
    docker-compose --profile storefront_services up
    ```
- Stop all **scheduler services and default services**:
    ```sh
    docker-compose --profile scheduling_services down
    ```



<br><br><br>



## Multiple Compose Files in Docker Compose
### When to Use Multiple Compose Files
- **Best for**: Different environments (local, staging, CI, production) where the configurations:
    - **Are completely distinct**.
    - **Will never run together on the same machine**.
- **Not recommended for**: Different parts of the same system (use **service profiles** for that instead).

### Default Files
By default, Docker Compose automatically looks for:
- `docker-compose.yaml` — The **main configuration file**.
- `docker-compose.override.yaml` — Optional override file that **extends** the main file.

### How Overrides Work
- Docker Compose **merges** the two files.
- Fields like `depends_on` (arrays) will **combine values** from both files.
- Fields like `build` (single values) will **prefer the value from the override file**.

### Example
```yaml
# docker-compose.yaml
services:
  web:
    image: myapp:latest
    environment:
      - ENV=prod

# docker-compose.override.yaml
services:
  web:
    environment:
      - ENV=dev
```
In this case, the `ENV` variable would be set to **dev** when using both files.


### File Paths in Overrides
- **Paths (like volume mounts or build contexts) in override files must be relative to the primary file.**
- Override files don’t need to be complete; they can **contain only the parts you want to change**.

### Custom Override Files
- You can create **custom override files** (not just `override`).
- Examples:
    - `docker-compose.local.yaml`
    - `docker-compose.staging.yaml`

- These are **not used automatically** — you have to specify them explicitly when running Docker Compose.


### Running with Custom Override Files
#### Command Syntax
```bash
docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up
```
Explanation:
- `-f` specifies the **files to use**.
- **Order matters** — the **last file** overrides any settings from earlier files.
- After the files, you provide the actual command (e.g., `up`).

### Benefits
- Reuse the same **base compose file** across different environments.  
- Customize only the necessary fields in **small override files**.  
- Avoid duplication across environments.  
- Supports multiple overrides (you can chain `-f` options).

### Example for Local and Staging
```bash
docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up
```
Runs with **local overrides**.

```bash
docker-compose -f docker-compose.yaml -f docker-compose.staging.yaml up
```
Runs with **staging overrides**.

### Key Takeaways
| Use Case | Recommended Approach |
|---|---|
| Different environments (local, staging, etc.) |  **Multiple compose files with overrides** |
| Different service groups within one system | Use **Service Profiles** instead |
