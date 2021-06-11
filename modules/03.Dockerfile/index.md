# Dockerfile reference

## Features

- It is a plain text file
- Provide instructions to build the image and run the container
- Automate the commands a user could call on the command line
- Reproducible build
- Compatible with GitOps

## Good to know

* Since version 18.09, Docker use a new backend called [Buildkit](https://github.com/moby/buildkit)

  * Detect and skip executing unused build stages
  * Parallelize building independent build stages
  * Incrementally transfer only the changed files in your build context between builds
  * Detect and skip transferring unused files in your build context
  * Use external Dockerfile implementations with many new features

* Set an environment variable `DOCKER_BUILDKIT=1` on the CLI before invoking `docker build`.

* USE bash default value

  `USER ${user:-some_user}`

* `shells` form: `RUN <command>`, use `/bin/sh -c`

* `exec` form: `RUN ["executable", "param1", "param2"]`

* Proxy variables defined as [predefined ARGs](https://docs.docker.com/engine/reference/builder/#predefined-args)
  `docker build --build-arg HTTPS_PROXY=https://my-proxy.example.com .`

## Most common instructions

* `FROM`

  The base image, if any

* [`RUN`](https://docs.docker.com/engine/reference/builder/#run)

  * `shell` form: `RUN <command>`

    `exec` form: `RUN ["executable", "param1", "param2"]`

  * Execute any commands in a new layer on top of the current image and commit the result

* [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd)

  * `exec` form and `shell` form
  * Provide defaults for an executing container, as an executable or as argument from `ENTRYPOINT`.

  * Only one `CMD` instruction, last one takes precedence.
  * Provide default arguments to `ENTRYPOINT` when defined.
  * Overwritable with the command line `docker run [OPTIONS] IMAGE [COMMAND] [ARG...]`

* `LABEL`

  Adds metadata to an image.

  `LABEL <key>=<value> <key>=<value> <key>=<value> ...`

* `EXPOSE`

  * `EXPOSE <port> [<port>/<protocol>...]`
  * Listens on the specified network ports at runtime.

  * Persist when a container is run from the resulting image.

* `ENV`

  Sets the environment variable `<key>` to the value `<value>`.

  `ENV <key>=<value> ...`
  
* `ADD`

  * `ADD [--chown=<user>:<group>] <src>... <dest>`
    `ADD [--chown=<user>:<group>] ["<src>",... "<dest>"]`
  * Copies new files, directories or remote file URLs from `<src>` and adds them to the filesystem of the image at the path `<dest>`.
  
* `COPY`

  * `COPY [--chown=<user>:<group>] <src>... <dest>`

    `COPY [--chown=<user>:<group>] ["<src>",... "<dest>"]`

  * copies new files or directories from `<src>` and adds them to the filesystem of the container at the path `<dest>`.

* `ENTRYPOINT`

  * `exec` form : `ENTRYPOINT ["executable", "param1", "param2"]`

    `shell` form: `ENTRYPOINT command param1 param2`

  * `shell` form prevents any `CMD` or `run` command line arguments from being used.

  * with `shell` form, executable will not be the container’s PID 1 and will not receive Unix signals (with `docker stop`).

* `VOLUME`

  * `VOLUME ["/data"]`

    `VOLUME /data`

* `USER`

  * `USER <user>[:<group>]`

    `USER <UID>[:<GID>]`

  * Sets the user name (or UID) and optionally the user group (or GID) to use

* `WORKDIR`

  * `WORKDIR /path/to/workdir`
  * Sets the working directory for any `RUN`, `CMD`, `ENTRYPOINT`, `COPY` and `ADD` instruction.

## Starter scripts

* Use `ENTRYPOINT` to execute a script instead of a command.
* Detect arguments with bash control, eg `if [ "$1" = 'postgres' ]; then .. ; fi`.
* Exec all arguments with `exec "$@"`.
* Receives the Unix signals with `exec gosu <cmd>`.

## Build vs Runtime

* Start with build instructions.
* End with runtime instruction.
* Some runtime instructions can be overwritten by the `docker run` command.

## `docker build`

- Build an image from a Dockerfile and a context.
- Context is the set of files at a specified location `PATH` or `URL`.
- Use `docker build -f /path/to/a/Dockerfile .` to build an image.
- Where `-f` is the optional path to the `Dockerfile` and `.` is the context.
- Do not use `/`, it transfers the entire contents of your hard drive to the Docker daemon.
- Use `-t` to name and tag the image, eg `docker build -t shykes/myapp:1.0.2 -t shykes/myapp:latest .`.

## Buildpack

* An alternative to `Dockerfile`.
* Turns source code into a runnable container image.
* Usually encapsulate a single language ecosystem toolchain (Ruby, Go, NodeJs, Java, Python, ...).
* A collection of multiple build packs is called a **builder.**
* Inspect the application source code and **detect** if it should participate in the build process, eg: a go build pack search for `*.go`.
*  Once a buildpack (or set of buildpacks) has matched, it moves on to the **build** step, eg:
  * add a layer with a dependency, eg the go distribution, the Java JDK, ...
  * run a command, eg `go build`, `npm run build`.
  * build an OCI image.
* Kubernetes integration with [kpack](https://github.com/pivotal/kpack), declarative resource definitions for mapping source code to buildpacks.
* Increase developer productivity with reduced developer actions.
* Avoid sharing a `Dockerfile` requiring Docker familiarity and knowledge to quickly build (and rebuild) small and secure images.
* Developer don't need to write anything, to take care of size and security.
* Rely upon an open-source project and its community contributions.
* Less control but with great power comes great responsibilities.

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

## [Best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

* Use [`docker scan`](https://docs.docker.com/engine/scan/) to check the [vulnerability](https://docs.docker.com/docker-hub/vulnerability-scanning/) of your image.
* Clean any unnecessary files (cache, temporary files, downloads, dev/test/audit tools, ...).
* Create ephemeral containers: container can be stopped and destroyed, then rebuilt and replaced with an absolute minimum set up and configuration.
* Pipe Dockerfile through stdin.
  * Perform one-off builds without writing a Dockerfile to disk.
  * Generated `Dockerfile`.
  * Hide sensitive information from commands and files by interpolating environmental variables.
* Exclude with .dockerignore, avoid sending large or sensitive information to the context
* Use [multi-stage builds](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#use-multi-stage-builds), reduce the size of the final image, without struggling to reduce the number of intermediate layers and files.
* Don’t install unnecessary packages
* Decouple applications
* Minimize the number of layers
* Facilitate comprehension, sort multi-line arguments, use `\` (backslash)
* Leverage and understand build cache
* Use/extend [trusted images](https://docs.docker.com/engine/security/trust/)
* Configure container's application to [run as non-root user](https://docs.docker.com/engine/security/userns-remap/) to prevent privilege-escalation

