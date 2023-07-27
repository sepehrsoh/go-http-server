# Simple Go HTTP Server

Welcome to the Simple Go HTTP Server project! This repository contains a basic Go HTTP server that you can use as a starting point for your own projects. Follow the steps below to set up and run the server.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

1. Go 
2. Git 

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/go-http-server.git
   ```

2. In the `config` folder, rename `configs.yaml.sample` to `configs.yaml` and fill in the necessary configuration fields based on your usage.

## Building and Running the Server

1. Use the provided Makefile to download and vendor dependencies, and build the project:

   ```bash
   make all
   ```

2. Run the following command to start the HTTP server:

   ```bash
   ./go-http-server serve
   ```

3. The server is now running! You can access it at `localhost:port` where `port` is the port you specified in the `configs.yaml` file.

## Example Endpoint

The server comes with a simple example endpoint:

- GET request to `localhost:port/ping` will return an HTML file with the body "Pong."

## Project Structure

The project has the following structure:

```
.
├── Makefile
├── config
│   ├── config.go
│   ├── configs.yaml
│   └── configs.yaml.sample
├── go.mod
├── main.go
├── server
│   └── panel
│       └── html
│           ├── 404.html
│           └── ping.html
└── src
    ├── main
    │   └── cli
    │       ├── root.go
    │       └── serve.go
    ├── middleware
    │   ├── middleware.go
    │   └── ports.go
    ├── modules
    │   └── panel
    │       ├── delete.go
    │       ├── get.go
    │       ├── main_handler.go
    │       ├── ports.go
    │       ├── post.go
    │       ├── put.go
    │       └── response.go
    └── service
        ├── providers
        │   ├── http_server.go
        │   └── ports.go
        └── wiring
            ├── service.go
            └── wiring.go
```

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues. Your contributions are greatly appreciated!

## License

This project is licensed under the [MIT License](LICENSE).

Happy coding!
