# Simple go http server

### how to use :
1- clone this repository <br>
2- in folder config rename `configs.yaml.sample` to `configs.yaml` then complete fields base your own usage <br>
3- use command  ``Make all`` to download and vendor dependencies and build the project <br>
4- run ``./go-http-server serve`` to serve built project, and it's done! <br>
***Note:*** there is a simple example for get request ``localhost:port/ping`` witch return html file contained body `Pong`
****

### Project Structure

```bash
.
├── Makefile
├── config
│   ├── config.go
│   ├── configs.yaml
│   └── configs.yaml.sample
├── go.mod
├── main.go
├── server
│   └── panel
│       └── html
│           ├── 404.html
│           └── ping.html
└── src
    ├── main
    │   └── cli
    │       ├── root.go
    │       └── serve.go
    ├── middleware
    │   ├── middleware.go
    │   └── ports.go
    ├── modules
    │   └── panel
    │       ├── delete.go
    │       ├── get.go
    │       ├── main_handler.go
    │       ├── ports.go
    │       ├── post.go
    │       ├── put.go
    │       └── response.go
    └── service
        ├── providers
        │   ├── http_server.go
        │   └── ports.go
        └── wiring
            ├── service.go
            └── wiring.go
```