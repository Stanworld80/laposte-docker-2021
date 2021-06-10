
# Storage in Docker

## Problematics

By default, all the files created inside a container are stored on a writable container layer. This means:

- The data in a container is not persistent
- You can’t easily move the data somewhere else   
  A container’s writable layer is tightly coupled to the host machine.
- Data is only available from its container
- read and write speeds via the storage driver are lower than native file system performance   
  Writing into a container’s writable layer requires a storage driver to manage the filesystem. The storage driver provides a union filesystem, using the Linux kernel. This extra abstraction reduces performance as compared to using data volumes, which write directly to the host filesystem.

## Storage types

  - [volumes](https://docs.docker.com/storage/volumes/) - prefered
  - [bind mounts](https://docs.docker.com/storage/bind-mounts/)
  - [tmpfs mount](https://docs.docker.com/storage/tmpfs/)

![Docker storage types](image/storage-types.png)

## Storage types: Volumes

[Docker volumes](https://docs.docker.com/storage/volumes/) are:

- are the best way to persist data in Docker
- are stored in a part of the host filesystem which is managed by Docker (`/var/lib/docker/volumes/` on Linux)
- Non-Docker processes should not modify this part of the filesystem

**Use cases:**

- Sharing data among multiple running containers (eg.: databases)
- When the Docker host is not guaranteed to have a given directory or file structure   
  Volumes help you decouple the configuration of the Docker host from the container runtime.
- To store your container’s data on a remote host or a cloud provider, rather than locally.
- To back up, restore, or migrate data from one Docker host to another   
  You can stop containers using the volume, then back up the volume’s directory (such as /var/lib/docker/volumes/<volume-name>).
- When an application requires high-performance I/O on Docker Desktop   
  Volumes are stored in the Linux VM rather than the host, which means that the reads and writes have much lower latency and higher throughput.
- When your application requires fully native file system behavior on Docker Desktop   
  For example, a database engine requires precise control over disk flushing to guarantee transaction durability. Volumes are stored in the Linux VM and can make these guarantees, whereas bind mounts are remoted to macOS or Windows, where the file systems behave slightly differently.

## Storage types: Bind mounts

 - may be stored anywhere on the host system
 - may even be important system files or directories
 - Non-Docker processes on the Docker host or a Docker container can modify them at any time

**Use cases:**

- In general, you should use volumes where possible
- Sharing configuration files from the host machine to containers   
  This is how Docker provides DNS resolution to containers by default, by mounting `/etc/resolv.conf` from the host machine into each container.
- Sharing source code or build artifacts between a development environment on the Docker host and a container   
  For instance, you may mount a Maven target/ directory into a container, and each time you build the Maven project on the Docker host, the container gets access to the rebuilt artifacts.
- When the file or directory structure of the Docker host is guaranteed to be consistent with the bind mounts the containers require

## Storage types: `tmpfs` mounts

- are stored in the host system’s memory only
- are never written to the host system’s filesystem
- can’t be shared between containers

**Use cases:**

- when you do not want the data to persist either on the host machine or within the container
- to store non-persistent state, cache or sensitive information
- to increase container performance (by avoiding writing into the container’s writable layer using storage driver)
- used only Linux ([named pipe](https://docs.microsoft.com/en-us/windows/win32/ipc/named-pipes) for Windows)

## CLI example

To bind mount a volume run:

```
docker run -v path/on/host:/app nginx:latest
```
