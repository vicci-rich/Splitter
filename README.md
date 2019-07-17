# BDS 
![logo](./docs/bds-logo.png)

## Introduction
Blockchain Data Service (BDS) is that JD Cloud uses technical ways to store the chain-like and unstructured data of blockchain and synchronizes them to  database service in real time, meanwhile provides a tool of data visualization.

Splitter is a separate module of the open source project which is blockchain data service (BDS) and provides data analysis services

Splitter is responsible for consuming blockchain data from message queue (kafka) and inserting data into persistent data storage services (relational database, data warehouse, etc.).

## Architecture 
![Architecture](./docs/bds-architecture.jpg)

## Environment Deployment
### Install BDS 

#### Environment initialization
Before compiling and running BDS, you must install go's compilation environment locally: [go install](https://golang.org/doc/install)

#### Install Splitter steps

1. Set the path of project : `$GOPATH/src/github.com/jdcloud-bds/bds/`
2. Input`go build -v github.com/jdcloud-bds/bds/cmd/bds-splitter`，compile to get executable file *bds-splitter*
3. Build new configuration file *splitter.conf*,  see `/config/splitter_example.conf` configuration file template
4. Run programe `./bds-splitter -c splitter.conf`

### Install confluent and kafka
#### Install kafka
See [kafka](https://kafka.apache.org/quickstart)

##### Modify config/server.properties 

* message.max.bytes=1048576000

#### Install confluent 
see [confluent](https://docs.confluent.io/current/installation/installing_cp/zip-tar.html#prod-kafka-cli-install)

Unzip the confluent package and run Confluent REST Proxy

##### Modify  <path-to-confluent>/etc/kafka-rest/kafka-rest.properties 

* max.request.size = 1048576000
* buffer.memory = 1048576000
* send.buffer.bytes = 1048576000

### Database
Database we now support SQL Server, PostgreSQL, you can choose one as a data storage method.

#### SQL Server
Buy [JCS For SQL Server](https://www.jdcloud.com/cn/products/jcs-for-sql-server)

#### PostgreSQL 
Buy [JCS For PostgreSQL](https://www.jdcloud.com/cn/products/jcs-for-postgresql)

After you run the database, you need to manually create new database and use the database name initialization splitter.conf.

### Install Grafana
See [Grafana Official](https://grafana.com/)

## Source code
[Splitter Modules](./splitter/README.md)

### Development Steps
1. Define the data structure of Kafka messages.
2. Define table structure.
3. Analyze Kafka message and store data in database.

## Contributing
[Contributing guide](./CONTRIBUTING.md)

## License
[Apache License 2.0](./LICENSE)

## Project Demonstration
[Blockchain Data Service](https://bds.jdcloud.com/)
