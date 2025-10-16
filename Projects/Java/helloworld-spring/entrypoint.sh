#!/bin/sh

# Add to Dockerfile
# COPY entrypoint.sh /entrypoint.sh
# RUN chmod +x /entrypoint.sh
# ENTRYPOINT ["/entrypoint.sh"]

set -e

echo "Running HelloWorld with:"
echo "Vector ID = $VECTOR"
echo "Tenant ID = $TENANT"

# Run the app with system properties
java -Da="$A" -Db="$B" -jar /app/app.jar
