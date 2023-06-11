<div align="center">
    <h1>file-server-api</h1>
</div>

## Description

File Server API is a file storage service for storing and accessing data.

## Requirements

* [Go](https://golang.org) 

## Setup

In order to set up this app, run the following commands:

1. `git clone git@github.com:qosimmax/file-server-api.git`
2. `make build`

## Run

In order to run this app, run the following commands:

1. `make run`


## Run scale instances

1. `docker-compose up --build  --scale file-server-api=7`