# gocd-plumber
[![Build Status](https://travis-ci.org/dennisgranath/gocd-plumber.svg?branch=master)](https://travis-ci.org/dennisgranath/gocd-plumber)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/dennisgranath/gocd-plumber/blob/master/LICENSE.md)

Create gocd-pipelines from YAML config.

## Build
`go build`

## Dev
### GoCD
`git clone https://github.com/dennisgranath/gocd-docker.git` and run `docker-compose up`.

### Create example pipeline
`cd examples && ../gocd-plumber -config ../root/etc/gocd-plumber/config.ini`


