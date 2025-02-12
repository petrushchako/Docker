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


