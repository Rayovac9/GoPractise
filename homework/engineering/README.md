## Go Advanced Bootcamp Week 4 Assignment

## Question

 Write a project that meets the basic directory structure and project as per your own conception. The code should include registration of data layer, business layer, API registration, registration and startup of services by the main function, signal processing, and dependency building using Wire. You can use your own familiar framework.

## Ideas

1. Set the business scenario as - Airline ticket price service (FareService), provide the following five interfaces

    - CreateFare Airline Ticket Pricing Rules Add
        - StartAirport, DestinationAirport, StartTravelDate, EndTravelDate, PassengerCategory (Adult, Child, Infant), Fare

    - UpdateFare Airfare rule update

    - DeleteFare airfare rule deletion

    - GetFare Airfare Rule by ID

    - Pricing Airfare calculation

        - Calculate the applicable fare based on the origin airport, destination airport, date of travel, passenger category (adult, child, and infant), and conditions.

2. define the interface protocol using Protobuf

3. generate HTTP and grpc using framework using Kratos

4. using ent to generate entity framework to handle DB operations

5. using wrie to generate injection code

## TODO

1. [X] Engineering Catalogue

2. [X] Service API Design

3. [X] Table Structure Design

4. [X] kratos usage

5. [X] ent usage

6. [x] wrie usage

7. [x] Testing

## Reference

- [Go-gRPC Practice Guide](https://www.bookstack.cn/read/go-grpc/summary.md)
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
- [kratos blog example](https://github.com/go-kratos/kratos/blob/main/examples/blog)
- [ent documentation](https://entgo.io/zh/docs/getting-started)

## Compile

### kratos generates interfaces

``sh
cd api\fare\v1
kratos proto client fare.proto --proto_path=... /... /... /third_party -I=.
```

### ent

1. generate structure Fare in `<project>/ent/schema/` directory `ent init Fare`

2. generate various operations on the DB `ent generate . /ent/schema`.

### Project compilation

``sh
cd cmd\fare
## Generate injection code
wrie
go build github.com/webmin7761/go-school/homework/engineering/cmd/fare
```

## Test

- [Postman test case](test/data/_postman_engineering.json)

## protobuf memo

``sh
go get -u github.com/golang/protobuf

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/lazada/protoc-gen-go-http

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-http_out=. --proto_path=d:/protobuf -I=. price.proto
``

Translated with DeepL.com (free version)
