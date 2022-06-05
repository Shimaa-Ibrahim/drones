# drones
Golang

### Description
REST API service allows clients to communicate with the drones capable of loading medications.
## Built With
- [GOLANG](https://go.dev/) - REST APIs backend
- [Postgresql](https://www.postgresql.org/) - DataBase
- [GORM](https://gorm.io/) - ORM

Features:
- registering a drone;
- loading a drone with medication items;
- checking loaded medication items for a given drone; 
- checking available drones for loading;
- check drone battery level for a given drone;
- registering a medication;
- periodic task to check drones battery levels and create history/audit event log

Behaviors
- Prevent the drone from being loaded with more weight that it can carry;
- Prevent the drone from being in LOADING state if the battery level is **below 25%**;

 ##### APIs:
    | Method | Api | functionality |
    |------| ------ | ------ |
    | POST | /drone/register/ | register drone|
    | GET |  /drone/checkdroneloaded/{id}/  | Check Drone's Loaded Item |
    | GET | /drone/availabledrones/ | Get Available Drones For Loading |
    | GET | /drone/checkbattery/{id}/  | Check drone's battery level |
    | POST | /medication/register/ | register Medication |
    | POST | /medications/load/  | Load Drone With Medication Items |
 
 ### Getting started:
##### Installation:
###### Make sure Golang is installed
* Install all project dependencies with `go get ./...`
* Run migrations gorm-goose -path=repository/db -pgschema=drones up
* Load db seeds run `./repository/db/load-data.sh`

 sample of medication register
```
curl --request POST \
  --url http://localhost:3000/drone/register/ \
  --header 'Content-Type: application/json' \
  --data '{
	"serial_number": "serial-1320-number",
		"drone_model_id": 1,
		"weight_limit": 500,
		"battery_capacity": 75,
		"drone_state_id": 1
}'
```   

 sample of load drone with medication items
```
curl --request POST \
  --url http://localhost:3000/medications/load/ \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 1, 
	"medications_ids": [1,2,3]
}'
```   
    
sample of medication register
```
curl --request POST \
  --url http://localhost:3000/medication/register/ \
  --header 'Content-Type: multipart/form-data' \
  --header 'content-type: multipart/form-data; boundary=---011000010111000001101001' \
  --form image=@/home/shimaa/Downloads/DcKMUlbVAAAyP1a.jpg \
  --form 'medication={
	"name": "medecationOne",
	"code": "code",
	"weight": 300
}'
```
