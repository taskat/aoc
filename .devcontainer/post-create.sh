#!/bin/bash

# Define socket and environment variable file
SOCKET="/root/.ssh/agent.sock"
ENV_FILE="/root/.ssh/agent.env"

# Remove stale socket file if it exists and is not in use
if [[ -e $SOCKET ]]; then
    if ! ssh-add -l > /dev/null 2>&1; then
        echo "Removing stale socket file..."
        rm -f $SOCKET
    fi
fi

# Start the SSH agent if it's not already running
if [[ ! -S $SOCKET ]]; then
    echo "Starting SSH agent..."
    ssh-agent -a $SOCKET > $ENV_FILE
else
    echo "SSH agent is already running."
fi

# Source the environment variables
if [[ -f $ENV_FILE ]]; then
    source $ENV_FILE
else
    echo "Environment file not found. Exiting..."
    exit 1
fi

# Copy the SSH key to be able to change the permissions
cp /root/.ssh/github-dev-container /root/.ssh/github-dev-container-copy 
# Change the permissions of the SSH key
chmod 600 /root/.ssh/github-dev-container-copy 
# Start the SSH agent
ssh-agent -a /root/.ssh/agent.sock > /root/.ssh/agent.env 
# Add the SSH key to the SSH agent
ssh-add /root/.ssh/github-dev-container-copy