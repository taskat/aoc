# This is the Dockerfile for the devcontainer
FROM mcr.microsoft.com/devcontainers/go:1.25-bookworm

# Set working directory
WORKDIR /workspaces/aoc

RUN rm -f /etc/apt/sources.list.d/yarn.list

# Install additional packages
RUN apt-get update
RUN apt-get install -y git-crypt liblpsolve55-dev libsuitesparse-dev

# Set AOC_HOME env var
ENV AOC_HOME=/workspaces/aoc

# Set CGO flags for lpsolve
ENV CGO_CFLAGS="-I/usr/include/lpsolve"
ENV CGO_LDFLAGS="-lm -lcolamd"
