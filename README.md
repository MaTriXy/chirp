Chirp!
======

A lightweight but modern Social Media client, written in Go & QML.

![chirp Screenshot](/assets/screenshot.png)

## Features

- [x] Live feed via Mastodon's Streaming API
- [x] Live feed via Twitter's Streaming API
- [x] Multi pane support
- [x] Linux/macOS/Windows (Android & iOS should be working, but aren't tested yet)
- [x] Media previews
- [x] Shortened URL resolving
- [ ] System notifications
- [ ] Multiple accounts (work-in-progress)
- [ ] Support for more networks

## Installation

Make sure you have a working Go environment (Go 1.8 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

### Dependencies

Before you can build Chirp you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

### Building Chirp!

    git clone https://github.com/muesli/chirp.git
    qtdeploy build desktop chirp/

### Run it

    ./chirp/deploy/linux/chirp

### Config

As of now you will need to create your own API apps & keys to use Chirp. You can do this on [https://dev.twitter.com/](https://dev.twitter.com/).
When you run Chirp for the first time, it will create an empty config file `chirp.conf` for you. Just edit it and enter
your consumer key and accesstoken.

![chirp logo](/assets/chirp.png)

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/chirp)
[![Build Status](https://travis-ci.org/muesli/chirp.svg?branch=master)](https://travis-ci.org/muesli/chirp)
[![Go ReportCard](http://goreportcard.com/badge/muesli/chirp)](http://goreportcard.com/report/muesli/chirp)
