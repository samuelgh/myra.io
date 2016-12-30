## MYRA.IO
Some notes so you don't forget these things

### Usage:

#### Create new device
POST `http://localhost:8080/api/device/<name>` 

Body should contain optional information about field names.
```
{
"id" : "boaty",
"field1" : "volt",
"field2" : "ampere",
"field3" : "temperature"
...
"field10": "last-name-for-sensor"
}
```

#### Send sensor data
GET `http://localhost:8080/api/device/<name>?Field1=12&Field2=12.56`
#### Get all data
GET `http://localhost:8080/api/device/<name>/all`
