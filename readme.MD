# Realworld Application Backend Implementation (Medium Clone)

This repository hosts the backend implementation of the Realworld Application, a clone of the Medium.com. It's designed with a modern tech stack using the Go programming language.
https://github.com/gothinkster/realworld. and use https://github.com/SyaibanAhmadRamadhan/gocatch

## Tech Stack

- **Go**: The backend service is developed using the Go programming language, with a strong focus on simplicity and scalability.

- **MongoDB**: MongoDB, a popular NoSQL database, is our choice for data persistence. It provides high performance, high availability, and easy scalability.

[//]: # (- **Elasticsearch**: Elasticsearch, a powerful open source search and analytics engine, is used for implementing the search functionality in this project.)

[//]: # ()
[//]: # (- **Debezium**: Change Data Capture &#40;CDC&#41; is implemented using Debezium, which is an open source distributed platform for change data capture.)

- **OpenTelemetry (Otel)**: OpenTelemetry is used for observability to collect and manage telemetry data (metrics, logging, and traces) to provide insight into the application's behavior.

[//]: # (## Architecture)

[//]: # ()
[//]: # (The backend process is as follows:)

[//]: # ()
[//]: # (1. User-related actions are persisted in the MongoDB database.)

[//]: # (2. Debezium captures the change events from MongoDB Transaction Logs.)

[//]: # (3. These events are streamed to Elasticsearch, making the data searchable in near real-time.)

[//]: # ()
[//]: # (In addition to the primary functionality, we utilized Elasticsearch for logging to provide efficient debugging and traceability.)

[//]: # ()
[//]: # (## Getting Started)

[//]: # ()
[//]: # (Follow these steps to get the project up and running on your local machine for development and testing purposes:)

[//]: # ()
[//]: # (&#40;Assuming Go, MongoDB, Elasticsearch, and Debezium are installed on your machine&#41;)

[//]: # ()
[//]: # (**1. Clone the repository**)

[//]: # ()
[//]: # (More detailed sections like 'Installation', 'Running the tests', 'Deployment', 'Built With', 'Versioning', 'Authors' and 'License' can be added depending on the scope and requirements of your project.)

[//]: # ()
[//]: # (## Contributing)

[//]: # ()
[//]: # (Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.)