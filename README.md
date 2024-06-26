## **User Service with Search Functionality using gRPC as transportation layer**

[](https://github.com/ansh-devs/userdatalink/actions/workflows/go-test.yml/badge.svg)


<!-- ### **Aspects Taken in considerations**

- &#9745; Get-Users with a list of IDs provided
- &#9745; Get-User-Details with a specific ID provided -->

<!-- -  &#9746; -->

### **Software Architecture Patterns and paradiagms used -**

- **_Repository Pattern_** &#8594; used in Repository-Layer/Database layer inside the _`internal/database`_ directory to make business logic and database layer logic decoupled. In future if needed just make the struct and implement the [_`UserRepository`_](/internal/database/repository.go) interface to the desired database (Postgres/Cassandra/MongoDb) such that there will be no need to make any changes in the [_`UserService`_](/internal/gapi/server.go) Service layer.
- **_Separation Of Concern (from SOLID)_** &#8594; there is a separation between the transportation layer **_gRPC_** and the Service layer and the repository layer
- **_Dependency Inversion Principle (from SOLID)_** &#8594; Dependency of `UserRepository` is injected in the `UserService` to perform **Db-calls**.
- **_Builder Pattern from Design Principles_** &#8594; used in the constructor of the `UserService` which itself configure the **gRPC's Reflection** and bind the **PORT** to the service.
- **_Idiomatic practices encouraged by the Golang Community_** &#8594; efficient use of `maps`, `slices` etc.

### **Protocol Buffers v3 Definition**

- UserService Proto Definition-

```
+-------------+--------------------+---------------------------+----------------------------+
|   SERVICE   |        RPC         |       REQUEST TYPE        |       RESPONSE TYPE        |
+-------------+--------------------+---------------------------+----------------------------+
| UserService | GetUserById        | GetUserByIdRequest        | GetUserByIdResponse        |
| UserService | GetUsersListByIds  | GetUsersListByIdsRequest  | GetUsersListByIdsResponse  |
| UserService | GetUsersByCriteria | GetUsersByCriteriaRequest | GetUsersByCriteriaResponse |
+-------------+--------------------+---------------------------+----------------------------+
```

- Available Proto Message Types-
```
+----------------------------+
|          MESSAGE           |
+----------------------------+
| GetUserByIdRequest         |
| GetUserByIdResponse        |
| GetUsersByCriteriaRequest  |
| GetUsersByCriteriaResponse |
| GetUsersListByIdsRequest   |
| GetUsersListByIdsResponse  |
+----------------------------+
```

* **GetUserById RPC types-**
    * _Request_

    ```
    +-------+------------+----------+
    | FIELD |    TYPE    | REPEATED |
    +-------+------------+----------+
    | id    | TYPE_INT64 | false    |
    +-------+------------+----------+

    ```

    * _Response_
    ```
    +-------+---------------------+----------+
    | FIELD |        TYPE         | REPEATED |
    +-------+---------------------+----------+
    | user  | TYPE_MESSAGE (User) | false    |
    +-------+---------------------+----------+
    ```

* **GetUsersListByIds RPC types-**
    * _Request_

    ```
    +-------+------------+----------+
    | FIELD |    TYPE    | REPEATED |
    +-------+------------+----------+
    | ids   | TYPE_INT64 | true     |
    +-------+------------+----------+

    ```

    * _Response_
    ```
    +-------+---------------------+----------+
    | FIELD |        TYPE         | REPEATED |
    +-------+---------------------+----------+
    | users | TYPE_MESSAGE (User) | true     |
    +-------+---------------------+----------+
    ```

* **GetUsersByCriteria RPC types-**
    * _Request_

    ```
    +-------+---------------------------+----------+
    | FIELD |           TYPE            | REPEATED |
    +-------+---------------------------+----------+
    | type  | TYPE_ENUM (UserCriterias) | false    |
    | value | TYPE_STRING               | false    |
    +-------+---------------------------+----------+

    ```

    * _Response_
    ```
    +-------+---------------------+----------+
    | FIELD |        TYPE         | REPEATED |
    +-------+---------------------+----------+
    | users | TYPE_MESSAGE (User) | true     |
    +-------+---------------------+----------+
    ```


For the request and response of RPC calls checkout the [`user_service.proto`](/protos/user_service.proto)

* **Grpc Interceptor Logging for Incoming Requests**
![alt-logging-image](https://private-user-images.githubusercontent.com/73169410/335441726-dbfb8e6b-90ba-4851-84be-b589080b09a0.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTcxMjg2MzcsIm5iZiI6MTcxNzEyODMzNywicGF0aCI6Ii83MzE2OTQxMC8zMzU0NDE3MjYtZGJmYjhlNmItOTBiYS00ODUxLTg0YmUtYjU4OTA4MGIwOWEwLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA1MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwNTMxVDA0MDUzN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWFmZjg0OGE2MTExZTk0ZTgzNmI3YzRiOTc0NzkwOTEyM2JlYWM5MzNkODQwNGI4MzY5MmZmODk4NzVmN2QzNWQmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.2thX25r1tv5-xawpcVycrkPIdv4QfKUmBT9JrG4WY38)

### **Environment Variable**

To run this project, The environment variable `PORT` must be in the accessible scope which is used to set the port where the _gRPC_ server will be listening for the incoming connections.

### **Steps To Run**
* Without Docker in Linux/WSL Environment:
    * `make run` - will execute the source code. 
    * `make build` - will build the executable binary from the source code. 
* With Docker:
    * `docker build . -t user-service` - will build image from Dockerfile.
    * `docker run -d --restart=always -p 8080:8080 user-service` - will run the image inside container. 
