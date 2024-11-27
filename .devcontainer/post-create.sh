#!/bin/bash

# Define socket and environment variable file
SOCKET="/root/.ssh/agent.sock"
ENV_FILE="/root/.ssh/agent.env"

# Function to kill the existing ssh-agent process if it exists
cleanup_stale_agent() {
    if [[ -e $SOCKET ]]; then
        # Check if the socket is active and belongs to an ssh-agent process
        if ! ssh-add -l > /dev/null 2>&1; then
            echo "Removing stale socket file and cleaning up existing ssh-agent..."
            # Find the PID of the ssh-agent process
            AGENT_PID=$(ps aux | grep '[s]sh-agent -a' | awk '{print $2}')
            if [[ -n $AGENT_PID ]]; then
                # Kill the ssh-agent process if PID exists
                kill -9 $AGENT_PID
            else
                echo "No ssh-agent process found to clean up."
            fi
            rm -f $SOCKET
        fi
    fi
}

# Cleanup stale agents and socket files
cleanup_stale_agent

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
# Add the SSH key to the SSH agent
ssh-add /root/.ssh/github-dev-container-copy