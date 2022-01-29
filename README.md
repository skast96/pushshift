# pushshift

[![](https://godoc.org/github.com/skast96/pushshift?status.svg)](http://godoc.org/github.com/skast96/pushshift)

[![Go Report Card](https://goreportcard.com/badge/github.com/skast96/pushshift)](https://goreportcard.com/report/github.com/skast96/pushshift)

A Go client for the [pushshift.io reddit API](https://pushshift.io/api-parameters/).

## Installation

Install the package with

`go get -u github.com/skast96/pushshift`

## Authentication

The pushshift API does not require authentication. However, the API is rate-limited server-side.

## Examples

For detailed examples, check out the examples folder.

```Go
// Returns a new unauthenticated client for invoking the API
client := pushshift.NewClient("myApp/0.1.0")

```
