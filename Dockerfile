FROM debian:bookworm

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Update and install essential packages
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    build-essential \
    ca-certificates \
    gnupg \
    lsb-release \
    software-properties-common \
    && rm -rf /var/lib/apt/lists/*

# Install Fish shell
RUN apt-get update && apt-get install -y fish \
    && rm -rf /var/lib/apt/lists/*

# Install Neovim
RUN apt-get update && apt-get install -y neovim \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION=1.21.3
RUN curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Create Go workspace directory
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOPATH/pkg"

# Create a directory for mounting host code
RUN mkdir -p /workspace

# Set Fish as default shell
SHELL ["/usr/bin/fish", "-c"]

# Set the working directory to the mount point
WORKDIR /workspace

# Set the default command to fish shell
CMD ["fish"]
