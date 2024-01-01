# Realworld Application Backend Implementation (Medium Clone)

This repository hosts the backend implementation of the Realworld Application, a clone of the Medium.com. It's designed with a modern tech stack using the Go programming language.
https://github.com/gothinkster/realworld. and use https://github.com/SyaibanAhmadRamadhan/gocatch
- [0.0.0.0:9001](http://0.0.0.0:9001) is endpoint minio console ui
- [0.0.0.0:9411](http://0.0.0.0:9411) is endpoint zipkin ui
- [0.0.0.0:16686](http://0.0.0.0:16686) is endpoint jaeger ui
- [0.0.0.0:9090](http://0.0.0.0:9090) is endpoint prometheus ui
- [0.0.0.0:8080](http://0.0.0.0:8080) is endpoint api

## Tech Stack

- **Go**: The backend service is developed using the Go programming language, with a strong focus on simplicity and scalability.

- **MongoDB**: MongoDB, a popular NoSQL database, is our choice for data persistence. It provides high performance, high availability, and easy scalability.

- **OpenTelemetry (Otel)**: OpenTelemetry is used for observability to collect and manage telemetry data (metrics, logging, and traces) to provide insight into the application's behavior.

- **Jaeger**: Jaeger is used for distributed tracing, helping to monitor and troubleshoot the flow of requests across various services in a distributed system.

- **Zipkin**: Zipkin is used for distributed tracing, similar to Jaeger, to monitor and trace requests as they travel through various components of a distributed system.

- **Datadog**: Datadog is used for monitoring and observability, providing a comprehensive platform for collecting, analyzing, and visualizing metrics, traces, and logs.

- **Prometheus**: Prometheus is used for monitoring and alerting, primarily focusing on collecting and storing time-series data.

- **Minio**: Minio is used for object storage, providing a scalable and high-performance solution for storing and retrieving data.

## how to init replica set mongodb
- running mongodb docker compose ```docker compose up```
- after running ```docker compose up```. you must mongodb init replica set in another terminal and execute 
   ```docker exec -it mongodb_realworld1 mongosh --eval "rs.initiate({_id:'dbrs', members:[{_id:0, host:'mongodb_realworld1'},{_id:1, host:'mongodb_realworld2'},{_id:2, host:'mongodb_realworld3'}]})"```
- after init replica set mongodb, you could ensure mongodb replica set is initiated by running the ```docker exec -it mongodb_realworld1 mongosh --eval "rs.status()"```