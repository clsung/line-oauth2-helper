# line-oauth2-helper

[![Build Status](https://travis-ci.org/clsung/line-oauth2-helper.svg?branch=master)](https://travis-ci.org/clsung/line-oauth2-helper)
[![codecov](https://codecov.io/gh/clsung/line-oauth2-helper/branch/master/graph/badge.svg)](https://codecov.io/gh/clsung/line-oauth2-helper)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/clsung/line-oauth2-helper)
[![Go Report Card](https://goreportcard.com/badge/github.com/clsung/line-oauth2-helper)](https://goreportcard.com/report/github.com/clsung/line-oauth2-helper)


## Introduction
Helper to [generate LINE JSON Web Token (JWT)](https://developers.line.biz/en/docs/messaging-api/generate-json-web-token/) for LINE OAuth2 v2.1.

## Install

`% go get github.com/clsung/line-oauth2-helper/cmd/line_jwt`

## Usage

### Command line

`% line_jwt -file ${LINE_PRIVATEKEY_FILE} -channel_id ${CHANNEL_ID}`

or

`% line_jwt -channel_id ${CHANNEL_ID} < {LINE_PRIVATEKEY_FILE}`

### Docker

Pull the image:

`% docker pull clsung/line-oauth2-helper`

then run the following

`% docker run -ti -p 8080:8080 clsung/line-oauth2-helper`

and connect to http://localhost:8080/
