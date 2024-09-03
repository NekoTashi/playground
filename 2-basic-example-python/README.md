# Basic Example Golang

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development purposes.

### Prerequisites

- Go version 1.21.9 or later
- Docker (for building and pushing the Docker image)
- Skaffold
- Colima (Kubernetes cluster)

### Run application

Create a local kubernetes cluster

```sh
colima start --with-kubernetes
```

Start development environment

```sh
skaffold dev --default-repo localhost:5000 --platform=linux/arm64
```
