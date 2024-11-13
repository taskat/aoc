#!/bin/bash

# Copy the SSH key to be able to change the permissions
cp /root/.ssh/github-dev-container /root/.ssh/github-dev-container-copy 
# Change the permissions of the SSH key
chmod 600 /root/.ssh/github-dev-container-copy 
# Start the SSH agent
ssh-agent -a /root/.ssh/agent.sock > /root/.ssh/agent.env 
# Add the SSH key to the SSH agent
ssh-add /root/.ssh/github-dev-container-copy