# Remote Dockerfile build

When using `docker build <URL>`, you can specify a **remote location** as the build context. This allows Docker to fetch the necessary files from a repository rather than using local files. This is useful in **CI/CD pipelines** and **automated builds**.  


## How Does It Work?  
When you provide a URL, Docker performs the following steps:  
1. **Clones the repository** or downloads the archive from the provided URL.  
2. **Looks for a `Dockerfile`** in the repository root (or in a specified directory).  
3. **Uses the downloaded repository as the build context** to create the image.  

**Supported remote sources:**  
- **Git repositories** (`https://github.com/user/repo.git`)
- **Git archives** (`https://github.com/user/repo/archive/branch.zip`)
- **HTTP/HTTPS URLs** pointing to compressed build contexts (e.g., `.tar.gz`)

## Examples  
### **Build from a Public Git Repository**
```sh
docker build https://github.com/docker-library/python.git
```
**Docker will**:
- Clone `https://github.com/docker-library/python.git`
- Look for a `Dockerfile` in the repository root  
- Use that repository as the build context  

### **Build from a Specific Branch**
```sh
docker build https://github.com/docker-library/python.git#3.10
```
**This will**:  
- Clone the `3.10` branch of the repository  
- Use its files as the build context  

### **Build from a Subdirectory**
If your `Dockerfile` is in a subdirectory (e.g., `app`), specify it after a `:`  
```sh
docker build https://github.com/example/repo.git#:app
```
**Docker will**:
- Clone `repo.git`
- Use `repo/app/` as the build context  

### **Build from a Private Repository**
You **cannot** pass authentication credentials via `docker build <URL>`, but you can manually clone the repo and then build:
```sh
git clone https://github.com/private/repo.git
cd repo
docker build .
```
For private repos in **CI/CD**, you typically:
- **Use SSH keys** for authentication  
- **Store credentials securely** in a secret management system  


## When to Use This?  
- Useful for **CI/CD pipelines** where build context needs to be fetched dynamically  
- Keeps local systems clean without downloading files manually  
- Ensures builds always use the latest version from source control  


<br><br><br>

## **Building a Docker Image from a Remote Repository in Jenkins**  

### **Prerequisites**
- **Jenkins Installed** with Docker support  
- **Docker Installed** on Jenkins agents  
- **Jenkins Pipeline Plugin** enabled  
- **Git Installed** on Jenkins agents (if cloning manually for private repos)  

### **Example: Jenkinsfile for Public GitHub Repository**  

This pipeline will:  
1. Pull the Dockerfile from a **public GitHub repository**  
2. Build a Docker image using `docker build <URL>`  
3. Tag the image  
4. Push the image to **Docker Hub**  

```groovy
pipeline {
    agent any
    environment {
        DOCKER_IMAGE = "myrepo/myapp"
        DOCKER_TAG = "latest"
    }
    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $DOCKER_IMAGE:$DOCKER_TAG https://github.com/docker-library/python.git'
                }
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script {
                    sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin'
                    sh 'docker push $DOCKER_IMAGE:$DOCKER_TAG'
                }
            }
        }
    }
}
```



### **Example: Handling a Private Repository**  
Since `docker build <URL>` **cannot authenticate private repositories**, we must:  
1. **Manually clone the repo**  
2. **Use the cloned directory as the build context**  

### **Jenkinsfile for Private GitHub Repo**
```groovy
pipeline {
    agent any
    environment {
        GIT_REPO = "git@github.com:private/repo.git"
        BRANCH = "main"
        DOCKER_IMAGE = "myrepo/myapp"
        DOCKER_TAG = "latest"
    }
    stages {
        stage('Clone Repo') {
            steps {
                script {
                    sh 'git clone --branch $BRANCH $GIT_REPO app_repo'
                }
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $DOCKER_IMAGE:$DOCKER_TAG ./app_repo'
                }
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script {
                    sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin'
                    sh 'docker push $DOCKER_IMAGE:$DOCKER_TAG'
                }
            }
        }
    }
}
```


### **Storing Credentials Securely in Jenkins**
- **GitHub Credentials**: Store SSH private key in **Jenkins Credentials**  
- **Docker Hub Credentials**: Add **DOCKER_USERNAME & DOCKER_PASSWORD** as Jenkins environment variables  
