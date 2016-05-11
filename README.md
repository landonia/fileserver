# fileserver

The simplest http file server you can think of.

## Overview

If you start the program it will start an HTTP file server using localhost:8090
serving from the working directory.

## Maturity

It uses the underlying go http.FileServer so it is ready to go. I quickly created
this as it is a very useful utility to have when you must provide a valid URL to
some resource.

## Installation

simply run `go get github.com/landonia/goat`

go run using `go run fileserver.go -addr=:8090 -dir=.`

install by `go install fileserver.go`

once installed `fileserver -addr=:8090 -dir=.`

## About

fileserver was written by [Landon Wainwright](http://www.landotube.com) | [GitHub](https://github.com/landonia).

Follow me on [Twitter @landoman](http://www.twitter.com/landoman)!
