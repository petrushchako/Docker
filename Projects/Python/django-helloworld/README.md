
# Django "Hello, World!" Docker Lab

This is a simple Dockerized Django project that demonstrates a basic "Hello, World!" application. The project is intended for learning and testing purposes, showcasing how to set up a Django application within a Docker container.

## Project Structure

```
django_docker_lab/
│
├── Dockerfile
├── requirements.txt
├── build_and_run.sh
├── cleanup.sh
└── helloworld/
    ├── manage.py
    └── helloworld/
        ├── __init__.py
        ├── asgi.py
        ├── settings.py
        ├── urls.py
        └── wsgi.py
```

### Files Overview

- **Dockerfile**: 
  - This file contains the instructions to build a Docker image for the Django application. It specifies the base Python image, installs the required packages, copies the application files, and runs the Django development server.
  
- **requirements.txt**: 
  - This file lists the Python dependencies required for the Django application. In this case, it only includes Django. The `pip` command inside the Dockerfile uses this file to install dependencies.
  
- **build_and_run.sh**: 
  - A shell script that builds the Docker image and runs the Docker container. It simplifies the process of setting up the application by combining the `docker build` and `docker run` commands into a single script.

- **cleanup.sh**: 
  - A shell script that finds, stops and removed instance of **django-helloworld** container. It simplifies the process of cleaning up.

- **helloworld/**: 
  - The root directory of the Django project.

  - **manage.py**: 
    - A command-line utility that lets you interact with this Django project in various ways. It allows you to run the development server, initiate database migrations, and more.

  - **helloworld/**: 
    - The inner directory named `helloworld` contains the core settings and configuration files for the Django project.

    - **\_\_init\_\_.py**: 
      - An empty file that marks this directory as a Python package.

    - **asgi.py**: 
      - ASGI (Asynchronous Server Gateway Interface) configuration for the Django project. It's used for asynchronous web servers and applications.

    - **settings.py**: 
      - The primary configuration file for the Django project. It contains settings for installed apps, middleware, URL configuration, and more. This version is simplified to only include necessary settings for this example.

    - **urls.py**: 
      - The URL configuration file. It maps URL patterns to views. In this example, it maps the root URL (`/`) to a simple view that returns a "Hello, World!" message.

    - **wsgi.py**: 
      - WSGI (Web Server Gateway Interface) configuration for the Django project. It's used for deploying the Django application on production servers.

<br><br>

## How to Use

### Steps to Run

1. **Build and Run the Docker Container**:
   - Run the following script to build the Docker image and start the container:
     ```bash
     ./build_and_run.sh
     ```

2. **Access the Application**:
   - Open your browser and navigate to `http://localhost:8000/`.
   - You should see the message "Hello, World!".

### Troubleshooting

- If you don't see the expected output, try the following:
  - Check the Docker container logs for any errors:
    ```bash
    docker logs <container_id>
    ```
  - Ensure that your browser cache is cleared or try accessing the page in an incognito window.
