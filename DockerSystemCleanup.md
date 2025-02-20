Here's a Bash script to automate daily Docker cleanup. It removes unused containers, images, networks, and optionally volumes.  
# **Docker Cleanup Script**

```bash
#!/bin/bash

echo "Starting Docker cleanup..."

# Remove stopped containers
docker container prune -f

# Remove dangling images
docker image prune -f

# Remove all unused images
docker image prune -a -f

# Remove unused networks
docker network prune -f

# Remove unused volumes (optional: uncomment to enable)
# docker volume prune -f

# Full system cleanup (excluding volumes)
docker system prune -f

echo "Docker cleanup completed!"
```

### **Usage Instructions**
1. **Save the script:**  
   ```sh
   nano docker-cleanup.sh
   ```
   Paste the script and save (`CTRL+X`, then `Y`, then `ENTER`).

2. **Make it executable:**  
   ```sh
   chmod +x docker-cleanup.sh
   ```

3. **Run the script:**  
   ```sh
   ./docker-cleanup.sh
   ```
