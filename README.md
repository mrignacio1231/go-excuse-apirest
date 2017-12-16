# Random Excuse REST API Example 

A RESTful API example for random excuse application with Go using libraries like fasthttp

## Installation and run

```bash
# Download this project
git clone https://github.com/mrignacio1231/go-excuse-apirest.git
# Build an run
$ cd go-excuse-apirest
$ go build
$ ./go-excuse-apirest

# API Endpoint : http://127.0.0.1:8080
```

## Excuses list

The list is taken from [here](http://pages.cs.wisc.edu/~ballard/bofh/).

## API

#### /excuse
* `GET` : Returns a random excuse
