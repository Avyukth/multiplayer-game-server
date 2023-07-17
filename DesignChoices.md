# Design Choices & Scalability of game-APP Service

When designing the game-app service, a few key considerations were made in order to ensure it is efficient, resilient, and scalable, capable of handling millions of concurrent gamers. Below, we dive into these design decisions and how they ensure scalability.

## gRPC and Protobuf

The choice to use gRPC and Protobuf was largely driven by performance considerations. gRPC uses HTTP/2, which is significantly more efficient than HTTP/1.1 (used by REST). It supports simultaneous requests over a single TCP connection, reducing network latency. Moreover, Protobuf ensures small payload size and fast serialization/deserialization, both of which are crucial for high throughput.

## Caching with Redis

We use Redis as a cache to store frequently accessed data. This design choice significantly reduces the amount of time needed to retrieve game stats by preventing unnecessary repeated reads from the database. Redis stores key-value pairs in memory, which allows for faster access times compared to disk-based databases.

## MongoDB for Storage

MongoDB, a NoSQL database, was chosen because of its ability to store and process large amounts of data with varying structures. It supports horizontal scaling through sharding, allowing the service to distribute data across multiple servers.

## Stateless Design

The service is designed to be stateless, meaning each request can be processed independently without requiring any knowledge of previous requests. This makes the service highly scalable as new instances can be added or removed as needed without affecting the service's functionality.

## Containerization with Docker

I've containerized the service using Docker which encapsulates the service along with its dependencies into a single standalone unit. This makes the service easy to deploy on any platform. Additionally, Docker enables easy horizontal scaling as we can quickly spin up multiple containers across multiple servers.

# Scaling to More than a Million Concurrent Gamers

To handle millions of concurrent users, several strategies could be employed:

## 1 Microservices

Breaking down the application into microservices would enable each component to scale individually based on the load it is experiencing. This can be especially beneficial for larger applications where different components may have different resource requirements.

#### ( Implemented, This project has been developed using the architecture of a microservices point of view.)

##### separate services for each microservice e.g game-app, game-app-redis, game-app-mongodb.

## 2 Auto-Scaling

With an auto-scaling setup, new instances of the service can be automatically spun up as demand increases, and spun down as demand decreases. This ensures the service has just enough resources to handle the current load, and can scale up to handle sudden surges in traffic.

#### ( Not Implemented , Can be done using kubernetes in this Project)

### 3 Load Balancing

By distributing incoming network traffic across multiple servers, a load balancer ensures no single server bears too much demand. This can significantly increase the number of concurrent gamers our service can handle.

#### (Not Implemented in this Project)

## 4 Database Sharding

Database sharding can also be implemented to increase the ability of the service to handle more concurrent users. By breaking our database into smaller chunks, or "shards", and distributing those across multiple servers, we can ensure that the database remains performant even under heavy loads.

#### (Not Implemented in this Project)

By employing these strategies and principles, the game-app service is designed to efficiently scale and handle the load of more than a million concurrent gamers.
