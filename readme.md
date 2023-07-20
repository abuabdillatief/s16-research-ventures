# Explanation

This api is using grpcui for endpoint testing. Makesure you can run `make` command for simplicity. To run:

- Run `go get`
- Run `make pb` to generate the protobuf
- Run `make run` to spin up the api
- Run `make test` to see the testing result
- Run `make grpcui` to spin up the grpc interactive UI. To install `grpcui`, kindly follow this [repository](https://github.com/fullstorydev/grpcui)

The tests might not be too sophisticated due to I only focus on several important functions. Because it's a quite simple API, I designed the folder structure to be quite simple too, separating the service from the 'handlers'. Also I added some functionality like middleware (interceptors), config variables using .yaml file and simple `api_key` for security. Also I use `makefile` to make it easier to run on the terminal. Hope this satisfies your requirements. Thank you!
