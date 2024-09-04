# SFTPGo

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development purposes.

### Prerequisites

- Docker (for building and pushing the Docker image)

### Run application

Change sftpgo-dir permissions recursively

```sh
chmod -R 0777 sftpgo-dir/
```

Start application with docker-compose

```sh
docker-compose up --build
```
