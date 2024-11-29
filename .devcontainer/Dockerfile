# This is the Dockerfile for the devcontainer
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Install additional packages
RUN apt-get update
RUN apt-get install -y git-crypt

# Add a script to source the SSH agent environment automatically
RUN echo 'export SSH_AUTH_SOCK=/root/.ssh/agent.sock' >> /root/.bashrc

# Set AOC_HOME env var
ENV AOC_HOME=/workspaces/aoc
