# datagen
## Table of content
- [datagen](#datagen)
  - [Table of content](#table-of-content)
  - [Quick start](#quick-start)
  - [Config file](#config-file)
    - [Basic config](#basic-config)
    - [Column object](#column-object)
    - [Row object](#row-object)
## Quick start
+ Install Go >= 1.17:
+ Execute command below to install external library
```
~$ go mod tidy
```
+ Place config json file to config/settings
+ Execute command below to run application
```
~$ go run datagen.go
# or
~$ go build datagen.go && ./datagen
```
## Config file
### Basic config
| Column name | Type | Description |
| --- | --- | ---|
| format | string | Output file format(csv/json/xml) **(required)**<br> ```"format": "csv"``` |
| size | int | Number of output records or file **(required)**<br> ```"size": 100```
| columns | array | Array of [Column object](#column-object) <br> ```"columns": [```<br>&ensp;&ensp;```{"name": "column1", "format": "[A-Z]{1,3}"},```<br>&ensp;&ensp;```{"name": "column2", "format": "[0-9]"}```<br>```]``` |
| rows | object | [Row object](#row-object) <br> ```"rows": {```<br>&ensp;&ensp;```{"size": 10, "format": [0-9]{1,2}}```<br>```}```|
### Column object
| Column name | Type | Description |
| --- | --- | --- |
| name | string | Column name<br> ```"name": "column-name-1"``` |
| format | string | Field value format (regex) **(required)**<br>```"format": "[0-9]{1,2}"``` |
| columns | array | Array of [Column object](#column-object) <br> ```"columns": [```<br>&ensp;&ensp;```{"name": "column1", "format": "[A-Z]{1,3}"},```<br>&ensp;&ensp;```{"name": "column2", "format": "[0-9]"}```<br>```]``` |
| rows | object | [Row object](#row-bject) <br> ```"rows": {```<br>&ensp;&ensp;```{"size": 10, "format": [0-9]{1,2}}```<br>```}```
### Row object
| Column name | Type | Description |
| --- | --- | --- |
| size | int | Size of records **(required)**<br>```"size": 10``` |
| format | string | Field value format (regex)<br>```"format": "[0-9]{1,2}"``` |
| columns | array | Array of [Column object](#column-object) <br> ```"columns": [```<br>&ensp;&ensp;```{"name": "column1", "format": "[A-Z]{1,3}"},```<br>&ensp;&ensp;```{"name": "column2", "format": "[0-9]"}```<br>```]``` |