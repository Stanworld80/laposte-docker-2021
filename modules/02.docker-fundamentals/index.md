
# Docker fundamentals



## Docker avantages

  - **Portable anywhere** (Linux, Windows, Datacenter, Cloud, Serverless, etc.)   
    Docker created the industry standard for containers.
  - **Lightweight**   
    Containers share the machine’s OS system kernel and therefore do not require an OS per application, driving higher server efficiencies and reducing server and licensing costs.
  - **Almost Secure**   
    Applications are safer in containers than sharing a same host.

## What is Docker

- A company, Docker Inc
- A container runtime, Docker containers
- An image format, Docker images
- Some developers tools

## Docker history

- Docker Inc. was founded by Kamel Founadi, Solomon Hykes, and Sebastien Pahl during the Y Combinator Summer 2010 startup incubator group and launched in 2011
- released as open-source in March 2013
- Started with the LXC runtime until Docker 0.9
- Now use containerd

## [Docker and DevOps](https://www.docker.com/resources/white-papers/docker-and-three-ways-devops)

Docker provides global optimization around software:

- Increase "velocity": developer, integration and deployment flows
- Decrease "variation": infrastructure and application are included in the Docker image
- Provide "visualization": microservices model real world domains

![DevOps life cycle](./assets/devops.png)

## Example of a Docker workflow

1. Developers write code locally and replicate the targetted production enviornment dependencies with containers
2. They package and share their work with their colleagues using Git and Docker containers.
3. They use Docker to push their applications into a test environment and execute automated and manual tests.
4. When developers find bugs, they can fix them in the development environment and redeploy them to the test environment for testing and validation.
5. When testing is complete, getting the fix to the customer is as simple as pushing the updated image to the production environment.

## Docker architecture

- client (CLI via REST API)
- server (Docker host) - daemon, runtime, containers, images, volumes
- registry - Docker Hub

![Docker architecture](./assets/docker-architecture.png)

## Docker components

![Docker objects and Docker engine](./assets/docker-engine-components.png)

**Docker engine:**

- Server (daemon process)
- REST API
- Client (CLI)

**Docker objects**

- Images
- Containers
- Networks
- Volumes
- ... some more objects

Docker objects are available to be observed and controlled using the command `docker <object-name>`.

## Docker Images

- is a read-only template
- contains instructions for creating a Docker container
- (often) is based on another image, with some additional customization

## Docker Containter

- a runnable instance of an image
- you can create/start/stop/move/delete a container
- you can connect a container to one or more networks, attach storage to it, or even create a new image based on its current state.
- (by default) is relatively well isolated from other containers and its host machine
- is defined by its image with any configuration options you provide to it on create or start
- when it is removed, any changes to its state that are not stored in persistent storage disappear

## CLI commands

- `docker help` - explore commands
- `docker ps` - list running containers
- `docker run <CONTAINER_NAME>` - list running containers
- `docker container` - manage containers
- `docker image` - manage images
- `docker volume` - manage volumes
- `docker network` - manage networks
- `docker exec` - run a command in a running container
- ...

Example: `docker run -i -t --rm ubuntu /bin/bash`
