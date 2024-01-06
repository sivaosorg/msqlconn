# msqlconn

![GitHub contributors](https://img.shields.io/github/contributors/sivaosorg/gocell)
![GitHub followers](https://img.shields.io/github/followers/sivaosorg)
![GitHub User's stars](https://img.shields.io/github/stars/pnguyen215)

A Golang MySQL connector library with built-in functionality to create a new database, execute batch operations, and manage transactions.

## Table of Contents

- [msqlconn](#msqlconn)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Modules](#modules)
    - [Running Tests](#running-tests)
    - [Tidying up Modules](#tidying-up-modules)
    - [Upgrading Dependencies](#upgrading-dependencies)
    - [Cleaning Dependency Cache](#cleaning-dependency-cache)

## Introduction

This Golang repository provides a MySQL connector package along with a set of services for handling common database operations. It simplifies the process of creating new databases, executing batch queries, and managing transactions in a MySQL database using Golang.

## Prerequisites

Golang version v1.20

## Installation

- Latest version

```bash
go get -u github.com/sivaosorg/msqlconn@latest
```

- Use a specific version (tag)

```bash
go get github.com/sivaosorg/msqlconn@v0.0.1
```

## Modules

Explain how users can interact with the various modules.

### Running Tests

To run tests for all modules, use the following command:

```bash
make test
```

### Tidying up Modules

To tidy up the project's Go modules, use the following command:

```bash
make tidy
```

### Upgrading Dependencies

To upgrade project dependencies, use the following command:

```bash
make deps-upgrade
```

### Cleaning Dependency Cache

To clean the Go module cache, use the following command:

```bash
make deps-clean-cache
```
