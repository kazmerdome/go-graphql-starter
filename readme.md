
# go-graphql-starter

Lightweight, easy-to-develop starter kit for Golang projects with graphql, mongodb, dataloaden and standard project-layout 🔥

- 🔮 **Gqlgen** — Generated, type safe Graphql for Go
- 👽 **Mongo Driver** - The official Mongodb driver for Go
- 🐶 **Dataloaden** — Generated type safe data loaders for Go
- 📄 **Echo** - High performance, extensible, minimalist Go web framework
- 🤟🏻 **Project-layout** — Based on https://github.com/golang-standards/project-layout


## Todo
- authZ/authN implementation
- docker/docker-compose implementation


## Features:

Server:
- Using Labstack Echo
- Restapi example
- prometheus implementation
- healthz - basic (demo only) health probe implementation

Graphql:
- using the latest (0.13.0) Gqlgen version
- playground security with http header password (Disable Introspection)
- custom scalar example
- dataloader examples (https://github.com/vektah/dataloaden) for n+1 problems

MongoDB
- model examples
- multiple db connections implementation

Other
- well separated module structure
- Using Makefile


## 🚀 Getting started


### Development Mode
start the server
```console
  make gateway
```

regenerate gqlgen files
```console
  make generate
```

default port: `http://localhost:9090`.


### Enable Graphql Documentation:
 ```add "Playground-Password": <GRAPHQL_PLAYGROUND_PASS> to request header ```
 change Header key -> pkg/gateway/handler.go.28 

### Working with Auth
```add "Authorization": Bearer <JWT TOKEN> to request header ```
change Header key -> pkg/gateway/handler.go.27


## Production Mode
step0 - make sure to provide envs (copy .env file to build folder or provide global envs)

step1 - binary build
```console
  make generate
  make build
```

step2 - run
```console
  ./build/go-graphql-starter or make run
```
