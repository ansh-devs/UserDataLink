## **User Service with Search Functionality using gRPC as transportation layer**

#### Here is the documentation of the TotalityCorp's Assignment

### **Aspects Taken in considerations**

- &#9745; Get-Users with a list of IDs provided
- &#9745; Get-User-Details with a specific ID provided

<!-- -  &#9746; -->

### **Software Architecture Patterns and paradiagms used -**

- **_Repository Pattern_** &#8594; used in Repository-Layer/Database layer inside the _`internal/database`_ directory to make business logic and database layer logic decoupled. In future if needed just make the struct and implement the [_`UserRepository`_](/internal/database/repository.go) interface to the desired database such that there will be no need to make any changes in the [_`UserService`_](/internal/gapi/server.go) Service layer.
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

For the request and response of RPC calls checkout the [`user_service.proto`](/protos/user_service.proto)

### **Environment Variable**

To run this project, The environment variable `PORT` must be in the accessible scope which is used to set the port where the _gRPC_ server will be listening for the incoming connections.
