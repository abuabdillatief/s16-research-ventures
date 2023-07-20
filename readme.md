# Explanation

This api is using grpcui for endpoint testing. Makesure you can run `make` command for simplicity. To run:

- Run `go get`
- Run `make pb` to generate the protobuf
- Run `make run` to spin up the api
- Run `make test` to see the testing result
- Run `make grpcui` to spin up the grpc interactive UI. To install `grpcui`, kindly follow this [repository](https://github.com/fullstorydev/grpcui)

The tests might not be too sophisticated due to I only focus on several important functions.