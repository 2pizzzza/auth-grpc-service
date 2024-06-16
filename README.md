# gRPC Authentication Service

This project is an example implementation of an authentication service using gRPC and Golang, following the onion architecture.

## Getting Started

### Prerequisites

- [Golang](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/)
- [Protocol Buffers Compiler (protoc)](https://github.com/protocolbuffers/protobuf)

### Installation

1. Clone the repository:

    ```sh
    git https://github.com/2pizzzza/auth-grpc-service.git
    cd auth-grpc-service
    ```

2. Download all dependencies:

    ```sh
    go mod tidy
    ```

3. Copy the example environment file and modify it as needed:

    ```sh
    cp .env.example .env
    ```

### Running the Project

To start the project, use the following command:

```sh
make run
