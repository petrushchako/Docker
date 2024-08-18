# Stage 1: Setup and Build Angular Application

# Use a Node.js image as the base
FROM node:18 as build

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Create a directory for the Angular app
WORKDIR /app

# Create a new Angular project (without interactive prompts)
RUN ng new hello-angular --defaults --skip-install

# Change to the new Angular app directory
WORKDIR /app/hello-angular

# Copy package.json and package-lock.json for dependency installation
COPY package.json package-lock.json ./

# Install project dependencies
RUN npm install

# Build the Angular app in production mode
RUN npm run build --prod





# Stage 2: Serve the Angular Application with Nginx

# Use an Nginx image as the base
FROM nginx:alpine

# Copy the built Angular app from the previous stage
COPY --from=build /app/hello-angular/dist/hello-angular /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
