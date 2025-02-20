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