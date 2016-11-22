# gocd-plumber
[![Build Status](https://travis-ci.org/dennisgranath/gocd-plumber.svg?branch=master)](https://travis-ci.org/dennisgranath/gocd-plumber)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/dennisgranath/gocd-plumber/blob/master/LICENSE.md)
[![Download](https://img.shields.io/github/release/dennisgranath/gocd-plumber.svg)](https://github.com/dennisgranath/gocd-plumber/releases/latest)


Create gocd-pipelines from YAML config.

## Install
Download static binary here:
https://github.com/dennisgranath/gocd-plumber/releases/latest

Edit config file (a sample is available in root/etc/gocd-plumber dir).
## Run

Run gocd-plumber in a directory containing pipeline configuration files:

`./gocd-plumber -config /path/to/config`

## Dev

### Build
`go build`

### GoCD
`git clone https://github.com/dennisgranath/gocd-docker.git` and run `docker-compose up`.

### Create example pipeline
`cd examples && ../gocd-plumber -config ../root/etc/gocd-plumber/config.ini`


