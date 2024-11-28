# Use the Go image as the base image
FROM mcr.microsoft.com/devcontainers/go:1.2

# Install Netcat (nc)
RUN apt-get update && apt-get install -y netcat

# Set the working directory
WORKDIR /workspace

# Default command (just to keep the container running)
CMD [ "bash" ]

