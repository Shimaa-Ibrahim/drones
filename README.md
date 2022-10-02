# Drones
Golang

### Description
REST API service allows clients to communicate with the drones capable of loading medications.
## Built With
- [GOLANG](https://go.dev/) - REST APIs backend
- [Postgresql](https://www.postgresql.org/) - DataBase
- [GORM](https://gorm.io/) - ORM

#### Features:
- registering a drone
- loading a drone with medication items
- checking battery level and loaded medication items for a given drone
- checking available drones for loading
- registering a medication
- periodic task to check drones battery levels and create history/audit event log

#### Behaviors
- load drone with medication item one item per each api call changing the drone state from idle to loading state (default) or loaded state if loaded = true
- if loaded items reach the weight limit of the drone the state changed to loaded
- available drones should be in IDLE or LOADING state
- the drone cannot being loaded with more weight that it can carry;
- the drone cannot being in LOADING state if the battery level is **below 25%**

 ##### APIs Documentation:
###### swagger: https://app.swaggerhub.com/apis/shimaa/drones/1.0.0


 ### Getting started:
##### Installation:
###### Make sure Golang and Postgresql are installed
* Install all project dependencies with `go get ./...`
* Run migrations `gorm-goose -path=repository/db -pgschema=drones up`
* required environmental variable:
  - DB_ENGINE: db driver
  - MAIN_DB: main db connection
  - DEV_DRONE_DATABASE: drone db connection for development
  - TEST_DRONE_DATABASE : drone db connection for testing
