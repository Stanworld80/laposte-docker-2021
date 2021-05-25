
# Docker Images

An image is a read-only template with instructions for creating a Docker container. Often, an image is based on another image, with some additional customization. For example, you may build an image which is based on the ubuntu image, but installs the Apache web server and your application, as well as the configuration details needed to make your application run.

## Docker Images

- is a read-only template
- contains instructions for creating a Docker container
- (often) is based on another image, with some additional customization
- you might create your own images or you might use published in a registry

## Image layers

- Images is made of layers
- Layers represent the filesystem differences
- Layers are stacked on top of each other to form a base for containers filesystem
- The Docker storage driver stacks it and provides a single unified view

![Image layers](image/docker-images.png)

![Image layers](image/container-layers.jpg)

## Containers and layers

- Containers have a top **writable layer** - the major difference from images
- All writes to the container are stored in this writable layer
- When the container is deleted the writable layer is also deleted
- Multiple containers share the same underlying image

[Read more](https://medium.com/@BeNitinAgarwal/docker-containers-filesystem-demystified-b6ed8112a04a)

![Containers and layers](image/sharing-layers.jpg)

## Copy on Write strategy (CoW)

- Copy-on-write is a mechanism allowing to share data.
- The data appears to be a copy, but is only a link (or reference) to the original data.
- The actual copy happens only when someone tries to change the shared data.
- Whoever changes the shared data ends up using their own copy instead of the shared data.

## Copy on Write strategy (CoW) advantages

Copy-on-write is essential to give us "convenient" containers:

- Optimizes both image disk space usage and the performance of container start times.
- Creating a new container (from an existing image) is "free"   
  Otherwise, we would have to copy the image first.
- Customizing a container (by tweaking a few files) is cheap   
  Adding a 1 KB configuration file to a 1 GB container takes 1 KB, not 1 GB.
- We can take snapshots, i.e. have "checkpoints" or "save points" when building images

## Building Docker images

To build an image, you create a `Dockerfile`:

- is a configuration file for building images.
- it has simple syntax for defining the steps needed to create the image and run it
- each instruction in a `Dockerfile` creates a layer in the image
- when you change the `Dockerfile` and rebuild the image, only those layers which have changed are rebuilt. It makes it lightweight, small, and fast.

[Learn Dockerfile instructions](https://docs.docker.com/engine/reference/builder/#format)

## Dockerfile examples

**Python app:**

```
FROM ubuntu:15.04
COPY . /app
RUN make /app
CMD python /app/app.py
```

**Node.js app:**

```
FROM node:12
WORKDIR /usr/src/app
COPY package.json .
RUN npm install
COPY . .
CMD [ "npm", "start" ]
```



## CLI to build an image

```
docker build .
# or
docker build -t my_image:v1 .
```
