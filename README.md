# Go Advanced Training Camp Graduation Project

## Graduation Project Requirements

For the current project, perform a microservice transformation, taking into account the following technical points:

1) Microservice architecture (BFF, Service, Admin, Job, Task modules)

2) API design (including API definitions, error code specifications, and usage of errors)

3) Use of gRPC

4) Go project engineering (project structure, DI, code layering, ORM framework)

5) Use of concurrency (parallel link requests with errgroup)

6) Use of microservice middleware (ELK, Opentracing, Prometheus, Kafka)

7) Cache optimization (consistency handling, pipeline optimization)

## Concept

### Business Scenario - Airline Ticket Inquiry Service (ShoppingService)

- The ticket inquiry interface returns the following three items, which are provided by three different interfaces from two services (Shopping):

    - Query the ticket price for a specific route and date (FareService.Pring)
    - Return the lowest price calendar for a week before and after a specific route (FareService.PriceCalendar)
    - Return tourist recommendations for the destination (Travel.Query)

- The ticket inquiry interface belongs to the BFF layer, using errgroup to query the three interfaces. All three queries must succeed, and then the data is aggregated.
- For ticket price inquiries, first query the cache. If not found, query the database, and then send a message to MQ to build a cache task.
- Ticket price cache update job: consume the build task from MQ and fetch information from the DB to put into Redis.


## TODO

1. [X] Ticket price, calendar, and travel recommendation interface design
2. [X] Inquiry service API design
3. [X] Database schema design
4. [X] Ticket price query implementation
5. [X] Calendar and travel recommendation implementation
6. [X] Cache update job
7. [X] Testing
