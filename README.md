# golang-docker-compose-sample

This repository is a sample code to run an application written in golang with docker-compose.

# contents of directories

```
.
├── Dockerfile
├── README.md
├── cmd
│   └── main.go
├── docker-compose.yml
└── entity
    └── sample_entity.go
```

# Quick Start

```
docker-compose build
docker-compose up -d
curl "http://localhost:5000"
```
