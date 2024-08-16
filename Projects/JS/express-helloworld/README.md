# Express Hello World Application

This project is a simple "Hello, World" web application built using Node.js and the Express framework. The application is containerized using Docker.

### Project Structure

```
express-helloworld/
│
├── app/
│   └── index.js          # Main application file
├── Dockerfile            # Dockerfile to build the Docker image
├── package.json          # Node.js dependencies and scripts
├── build_and_run.sh      # Script to build and run the Docker container
└── README.md             # Project documentation
```

### How to Build and Run

1. **Make the `build_and_run.sh` script executable:**

   ```bash
   chmod +x build_and_run.sh
   ```

2. **Build and Run the Docker Container:**

   Run the following script to build the Docker image and start the container:

   ```bash
   ./build_and_run.sh
   ```

   This will:

   - Build the Docker image with the tag `express-helloworld`.
   - Run the container, mapping port `3000` on your local machine to port `3000` in the container.

3. **Access the Application:**

   Open your web browser and go to `http://localhost:3000` to see the application running.


<hr>

<br><br>

### Troubleshooting

If you encounter issues when running the application, consider the following troubleshooting steps:

1. **Container Exits Immediately:**

   If the container exits immediately after starting, check the logs to see the error message:

   ```bash
   docker logs [container_id]
   ```

   Replace `[container_id]` with the actual container ID, which you can find by running:

   ```bash
   docker ps -a
   ```

2. **Port 3000 Already in Use:**

   If you get an error that port `3000` is already in use, either stop the service using that port or run the container on a different port:

   ```bash
   docker run -d -p 8080:3000 express-helloworld
   ```

   Then access the application at `http://localhost:8080`.
