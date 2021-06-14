
# Docker Images

An image is a read-only template with instructions for creating a Docker container. Often, an image is based on another image, with some additional customization. For example, you may build an image which is based on the ubuntu image, but installs the Apache web server and your application, as well as the configuration details needed to make your application run.

## Docker Images

- is a read-only template
- contains instructions for creating a Docker container
- (often) is based on another image, with some additional customization
- you might create your own images or you might use them published in a registry

## Image layers

- Images is made of layers
- Layers represent the filesystem differences
- Layers are stacked on top of each other to form a base for the container's filesystem
- Layers are `read-only`

## Storage driver

- The [Docker storage driver](https://docs.docker.com/storage/storagedriver/) stacks it and provides a single unified view
- Handles the details about the way these layers interact with each other
- Different storage drivers, with advantages and disadvantages in different situations

![Image layers](image/docker-images.png)

![Image layers](image/container-layers.jpg)

## Containers and layers

- Containers have a top **writable layer** - the major difference from images
- All writes to the container are stored in this writable layer
- When the container is deleted the writable layer is also deleted
- Multiple containers share the same underlying image

[Read more](https://medium.com/@BeNitinAgarwal/docker-containers-filesystem-demystified-b6ed8112a04a)

![Containers and layers](image/sharing-layers.jpg)

## Cache

- When possible, Docker uses a build-cache to accelerate the build
- Order your layers and group commands accordingly
- build-cache can be shared and distributed through an image registry with `--cache-from`

## Copy on Write strategy (CoW)

- Used by all storage drivers
- Copy-on-write is a mechanism allowing to share data.
- The data appears to be a copy, but is only a link (or reference) to the original data.
- The actual copy happens only when someone tries to change the shared data.
- Whoever changes the shared data ends up using their own copy instead of the shared data.

## Copy on Write strategy (CoW) advantages

* First time a file is modified in the lower layer, it is created in the upper layer
* Optimizes image disk space usage
* Optimize container disk space usage, see `du -hs ls /var/lib/docker/containers/*`
* Improve container start times by not having to copy the entire image
* Creating a new container (from an existing image) is "free"   
  Otherwise, we would have to copy the image first.
* Customizing a container (by tweaking a few files) is cheap   
  Adding a 1 KB configuration file to a 1 GB container takes 1 KB, not 1 GB.
* We can take snapshots, i.e. have "checkpoints" or "save points" when building images

## Container size on disk

* use the `docker ps -s` command
  * `size`: the amount of data (on disk) that is used for the writable layer of each container.
  * `virtual size`: the amount of data used for the read-only image data used by the container plus the container’s writable layer `size`.  Two containers started from the same image share 100% of the read-only data, while two containers with different images which have layers in common share those common layers.
* Total disk space is not the sum of `virtual size`
* beside the container and image size, we shall take into account:
  * Disk space used for log files with the `json-file` logging driver
  * Volumes and bind mounts used by the container
  * Disk space used for the container’s configuration files, small
  * Memory was written to disk (if swapping is enabled)
  * Checkpoints, if you’re using the experimental checkpoint/restore feature

## Building Docker images

To build an image, you create a `Dockerfile`:

- is a configuration file for building images.
- it has a simple syntax for defining the steps needed to create the image and run it
- each instruction in a `Dockerfile` creates a layer in the image
- when you change the `Dockerfile` and rebuild the image, only those layers which have changed are rebuilt. It makes it lightweight, small, and fast.

[Learn Dockerfile instructions](https://docs.docker.com/engine/reference/builder/#format)

## CLI to build an image

```
docker build .
# or
docker build -t my_image:v1 .
```

Cleanup

* Remove dangling images

  `docker image prune`

  A new build of the image is created without a new name, the old images become the "dangling image".

* Remove both unused and dangling images

  `docker image prune -a`
