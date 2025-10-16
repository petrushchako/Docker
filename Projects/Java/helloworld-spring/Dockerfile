# Start with a Maven image to build the application
FROM maven:3.8.5-openjdk-17-slim AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the pom.xml and download dependencies
COPY pom.xml .
RUN mvn dependency:go-offline -B

# Copy settings.xml
COPY settings.xml /root/.m2/settings.xml

# Copy the source code
COPY src ./src

# Package the application
RUN mvn -s /root/.m2/settings.xml clean package -DskipTests

# Start a new stage with an OpenJDK image
FROM openjdk:17-jdk-slim

# Set the working directory
WORKDIR /app

# Copy the jar from the previous stage
COPY --from=build /app/target/helloworld-spring-1.0-SNAPSHOT.jar /app/helloworld-spring.jar

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["java", "-jar", "helloworld-spring.jar"]
