<p align="center">
  <a>
    <img src="product-fee.svg" height="250" width="500" alt="product-fee" />
  </a>
</p>
<p align='center'>
 <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  </a>
  <a href="https://kubernetes.io/pt-br/docs/tasks/tools/install-kubectl-linux/#instale-usando-o-gerenciador-de-pacotes-nativo">
  <img src="https://img.shields.io/badge/kubernetes-326ce5.svg?&style=for-the-badge&logo=kubernetes&logoColor=white">
  </a>
  <a href="https://www.mysql.com/downloads/">
  <img src="https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white">
  </a>
  <a href="https://www.rabbitmq.com/docs/download">
  <img src="https://img.shields.io/badge/rabbitmq-%23FF6600.svg?&style=for-the-badge&logo=rabbitmq&logoColor=white">
   </a>
  <a href="https://www.docker.com/">
  <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white">
   </a>
  <a href="https://grafana.com/grafana/download">
  <img src="https://img.shields.io/badge/Grafana-F2F4F9?style=for-the-badge&logo=grafana&logoColor=orange&labelColor=F2F4F9">
   </a>
  <a href="https://prometheus.io/download/">
  <img src="https://img.shields.io/badge/Prometheus-000000?style=for-the-badge&logo=prometheus&labelColor=000000">
</a>
</a>
</p>
<br />

## 💻 Project Description

Product Fee is an API centered around calculating the final price of a product, taking into account its base price and associated tax. It employs messaging for the generation, delivery, processing of messages, and data persistence for subsequent operations in conjunction with the development of a webserver.

## 📍 Motivation

The API and its services are Dockerized to streamline infrastructure, speed up development, and optimize resource usage. Kubernetes automates the deployment process. A relational database was selected for its suitability to our structured domain. A messaging service supports the microservices architecture and offers easy implementation. Prometheus and Grafana provide comprehensive monitoring, collecting and storing metrics to inform decision-making. These tools are industry standards and integrate seamlessly with our development workflow.

## 🚀 Installation

```bash

# Create a directory in the desired location
$ cd capital-gains

# Navigate to the directory and open it in a code editor
$ code .

# Clone this repository
$ git clone https://github.com/EvyOliveira/productFee.git (HTTPS)
$ git clone git@github.com:EvyOliveira/productFee.git (SSH)
$ gh repo clone EvyOliveira/productFee (Github CLI)

# Initialize module in Go
$ go init 

# To synchronize code dependencies
$ go mod tidy

# Start and run containers in background
$ docker compose up

# Stop and remove containers and their associated networks, defined in a docker-compose.yaml file
$ docker compose down

# Build Docker images defined in docker-compose
$ docker compose build

# Execute a shell command (bash) within the Docker container to start the 'goapp' service as defined in the docker-compose.yml file
$ docker compose exec goapp bash

# Execute a shell command (bash) within the Docker container to start the 'mysql' service as defined in the docker-compose.yml file
$ docker compose exec mysql bash

# Connect to the database
$ mysql -u<user> -p <database-name>

# Running a Go program without needing to compile
$ go run <file-path>

# Build a Go package to produce an executable file
$ go build <file-path>

# Forcefully delete the file named "main"
$ rm -f main

# Deploy a new version of the application using a production-ready Docker image
$ docker build -t EvyOliveira/product-fee:latest -f Dockerfile.prod .

# Expose the application inside the container on local port 8080 and remove the container when it exits
$ docker run --rm -p 8080:8080 EvyOliveira/product-fee:latest

# Lists all Docker images available locally on your system
$ docker images

# Create a Kubernetes cluster using the Kind tool
$ kind create cluster --name-product-fee

# List all nodes in your Kubernetes cluster
$ kubectl get nodes

# Apply Kubernetes resources defined in a YAML manifest
$ kubectl apply -f <resource-file-path>

# List all active pods in your Kubernetes cluster
$ kubectl get pods

# List all active services in your Kubernetes cluster
$ kubectl get svc

# Forward a local port to a specific service in your Kubernetes cluster
$ kubectl port-forward svc/goservice 8080:8080

```

---

## 🛠 Improvements
- Applying retry and resilience strategies to the message queue;
- Perform unit and integrated tests.
---

## 🦸 Author

[![Linkedin Badge](https://img.shields.io/badge/-evelyncristinioliveira-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/evelyncristinioliveira/)](https://www.linkedin.com/in/evelyncristinioliveira/)

---
