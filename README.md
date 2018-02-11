# KVSS Simple Charity Server

The simple charity server is a small Go API that is designed to be a quick and simple server to power a non-profit. Features include user management, program management, donation tracking, and contact management. The goal is not to be an all-encompassing application that is difficult to learn. This is supposed to be small and simple.

**Requires Go 1.9 and the dep tool**

## Features

Coming Soon

## Environment

The server is configured largely through environment variables. The following variables are required, although we try to use sane defaults when possible.

* PORT The API port, defaults to `8080`

* DB_USER The MySQL DB user, defaults to `root`

* DB_PASSWORD The MySQL DB password, defaults to `password`

* DB_HOST The MySQL DB host, defaults to `localhost`

* DB_PORT The MySQL DB port, defaults to `3306`

* DB_NAME The MySQL DB name, defaults to `SimpleCharity`

Please note that using the defaults above is a **Bad Idea** in production. Please don't use those like that for anything other than local testing.

## Running

Two tools are required to be globally installed.

* [dep](https://github.com/golang/dep) is used for dependency management: `go get -u github.com/golang/dep/cmd/dep`

* [Task](https://github.com/go-task/task) is used for consistent tasks: `go get -u -v github.com/go-task/task/cmd/task`

Check the `Taskfile.yml` document for all available tasks or run `task -l`

## Testing

The `Taskfile.yml` document contains two commands for testing. `task test` will run the tests without coverage whereas `task cover` will run the tasks and open a coverage report.

## Docker

To just build the Docker image, run:

`task docker-build`

To build AND run the image, run:

`task docker-run`