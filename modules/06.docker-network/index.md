
# Docker Network

## Features

- you can connect containers together, or connect them to non-Docker workloads
- no need to be aware of where the containers are deployed

## Network drivers

Docker’s networking subsystem is pluggable, using drivers:

- `bridge`
  - the default network driver
  - usually used when your applications run in standalone containers that need to communicate
  - best when you need multiple containers to communicate on the same Docker host

- `host`
  - for standalone containers
  - remove network isolation between the container and the Docker host, and use the host’s networking directly
  - best when the network stack should not be isolated from the Docker host, but you want other aspects of the container to be isolated

- `overlay`
  - connect multiple Docker daemons together and enable swarm services to communicate with each other
  - you can also use overlay networks to facilitate communication between a swarm service and a standalone container, or between two standalone containers on different Docker daemons. This strategy removes the need to do OS-level routing between these containers.
  - best when you need containers running on different Docker hosts to communicate, or when multiple applications work together

- `macvlan`
  - allows to assign a MAC address to a container, making it appear as a physical device on your network   
    The Docker daemon routes traffic to containers by their MAC addresses.
  - using the `macvlan` driver is sometimes the best choice when dealing with legacy applications that expect to be directly connected to the physical network, rather than routed through the Docker host’s network stack
  - best when you are migrating from a VM setup or need your containers to look like physical hosts on your network, each with a unique MAC address.

- `none`
  - disable all networking
  - usually used in conjunction with a custom network driver

- Network plugins   
  - allows you to integrate Docker with specialized network stacks

[Read more](https://towardsdatascience.com/docker-networking-919461b7f498)

## Network drivers: `bridge`

TODO

## CLI examples: `bridge`

- `docker network create my-net` - create a network
- `docker network rm my-net` - remove a network
- connect container to a user-defined bridge   
  ```
  docker create --name my-nginx \
    --network my-net \
    --publish 8080:80 \
    nginx:latest
  ```
- `docker network connect my-net my-nginx` - connect a running container to an existing user-defined bridge
- `docker network disconnect my-net my-nginx` - disconnect a container from a user-defined bridge
