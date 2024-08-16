# Use an official Node.js runtime as a parent image
FROM node:18-slim

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install the dependencies
RUN npm install --only=production

# Copy the current directory contents into the container at /app
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the app
CMD ["npm", "start"]
