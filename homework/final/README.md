## go advanced bootcamp graduation project

## Graduation Project Requirements

A microservice transformation of the business in your own project at the moment, you need to consider the following technical points:

1) Microservice architecture (BFF, Service, Admin, Job, Task in modules)

(2) API design (including API definition, error code specification, the use of Error)

(3) the use of gRPC

(4) Go project engineering (project structure, DI, code layering, ORM framework)

(5) the use of concurrency (errgroup parallel link request)

(6) the use of microservice middleware (ELK, Opentracing, Prometheus, Kafka)

7) Optimisation of the use of caching (consistency handling, Pipeline optimisation)

## Ideas

### Business Scenario - Airline Ticket Inquiry Service (ShoppingService)

- The Airline Ticket Inquiry interface returns the following three pieces of information, which are provided by three different interfaces of the two services (Shopping)

    - Fare Inquiry returns fares for a specified route and date (FareService.Pring)
    - return to the specified route before and after the week of the lowest calendar (FareService. PriceCalendar)
    - return destination travel recommendations (Travel.Query)

- The air ticket inquiry interface belongs to the BFF layer, using errgroup to query the three interfaces, all three interface queries are successful, and then aggregate the data
- Query fare, first query the cache, not hit, and then query the DB, and then send a task to MQ to build the cache
- The fare cache updates the JOB, consumes the build task in the MQ, and fishes out the information from the DB to hit Redis

### Architecture design

! [architecture](doc/img/architecture.png)

## TODO

1. [X] Fare, calendar and travel recommendation interface design

2. [X] Quote Service API Design

3. [X] Table Structure Design

4. [X] Fare Query Implementation

5. [X] Calendar and Travel Recommendation Implementation

6. [X] Cache Update Job

7. [X] Testing

## Reference

- [Go-gRPC Practice Guide](https://www.bookstack.cn/read/go-grpc/summary.md)
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
- [kratos blog example](https://github.com/go-kratos/kratos/blob/main/examples/blog)
- [ent documentation](https://entgo.io/zh/docs/getting-started)

## Compile

### kratos generates interfaces

``sh
cd api\fare\v1
kratos proto client fare.proto --proto_path=... /... /... /third_party --proto_path=... /... --proto_path=. /. /. /third_party --proto_path=. /.
```

/.../...-I=.../...-I=.

1. generate structure Fare in `<project>/ent/schema/` directory `ent init Fare`

2. generate various operations on the DB `ent generate . /ent/schema`

### conf

``sh
protoc --go_out=. --proto_path=... /... /third_party -I=. conf.proto
```

### Project compilation

```sh
cd cmd\fare
## Generate injection code
wire
go build github.com/webmin7761/go-school/homework/final/cmd/fare

cd cmd\shop
# Generate injection code
cd cmd\shop
go build github.com/webmin7761/go-school/homework/final/cmd/shop

cd cmd\travel
# Generate injection code
cd cmd/travel
go build github.com/webmin7761/go-school/homework/final/cmd/travel

cd cmd\job
cd cmd\job
go build github.com/webmin7761/go-school/homework/final/cmd/job
``

### Run the project

``##sh
shop -conf ... /... /configs/shop/config.yaml 
travel -conf ... /... /.../configs/travel/config.yaml
fare -conf ... /... /.../configs/fare/config.yaml
job -conf ... /... /configs/job/config.yaml
```

## Test

- [Postman test case](test/data/go-school-final.postman_collection.json)
