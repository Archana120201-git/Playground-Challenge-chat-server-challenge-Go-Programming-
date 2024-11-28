# Use an official Go image as the base image
FROM golang:1.23

# Install Netcat (nc)
RUN apt-get update && apt-get install -y netcat

# Set the working directory
WORKDIR /workspace

# Default command
CMD [ "bash" ]
