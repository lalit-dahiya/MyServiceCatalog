# MyServiceCatalog

* This is a Golang application named MyServiceCatalog that provides CRUDL (Create, Read, Update, Delete, List) 
  functionalities for managing Service and ServiceVersion data through a RESTful API based on Echo framework backed 
  by Mongo database.

## Why MongoDB (NoSQL) for Service Catalog?
This application utilizes MongoDB, a NoSQL document-oriented database, as the backend storage for several reasons:
* **Scalability**: MongoDB scales horizontally by adding more nodes to the cluster, making it suitable for storing and 
  managing potentially large volumes of service and version data.
* **Flexibility**: Service and ServiceVersion structs can evolve over time without requiring schema changes in MongoDB, 
  as it stores data as flexible JSON documents.
* **Performance**: MongoDB offers efficient performance for various data access patterns, including inserts, updates, 
  and queries, crucial for a service catalog application.

## Prerequisites
* **MongoDB Server**: Ensure you have a running MongoDB server instance. Refer to the official MongoDB documentation for installation instructions: https://www.mongodb.com/docs/
* **Golang**: Download and install Golang from the official website: https://go.dev/doc/install

## Installation
1. **Clone the repository**
   * `git clone https://github.com/lalit-dahiya/MyServiceCatalog.git`

2. **Install dependencies**
   * `cd MyServiceCatalog`
   * `go mod download`

3. **Run the MongoDB Server** <br>
   * `mongod --dbpath ~/mongo` 
   * default port that mongo runs on is 27017

4. **Run the application**
   * Run `go run server.go` from project root.

## Usage

* Use POST API `/api/v1/services` to create a Service
* Use GET API `/api/v1/services` to list all Services
* Use GET API `/api/v1/services/search/:search` to list service names that match the search string
* Use POST API `/api/v1/versions` to create a ServiceVersion
* Use GET API `/api/v1/versions?serviceId=:id` to list all ServiceVersions that are versions of Service serviceId.
