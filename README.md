# API Http server GoLang
Small  Api HTTP go server .
### Prerequisites

What things you need to install the software and how to install them.

* Install [Golang](https://golang.org/) for Windows.
* Update dependencies
### Installing

A step by step series of examples that tell you have to get a Development env. running.

* Clone or copy the project to your src Golang project folder. (Ex Windows. C:\\Users\YOUR USER\go\src)
* Inside the project folder, open a command prompt and type:
```
dep ensure -v
```
* Execute the http server:
```
go run main.go
```

* Open an Internet Browser and go to http://localhost:9000/documents

### Output

* The program will return all MD5 checksum, names and sizes of the files located in the "files" folder in JSON format.

Example:
Making a GET request to [http://localhost:9000/documents](http://localhost:9000/documents),
The server will respond with a dictionary of all files located in "files" folder in JSON format:
```
[{"ID":"b2a1f484840f512e93838cbd543acea5","Name":"fileA","Size":17},{"ID":"765e6aec6d47d10b08dd325aa6b5128d","Name":"fileB","Size":7},{"ID":"8d74c534c15a4ba83c71100a10374075","Name":"fileC","Size":13},{"ID":"6e2e7dab4a04889e27c27f2aa2048b67","Name":"fileD","Size":20}]
```

*Also you can get a document

Example:
Making a GET request to  http://localhost:9000/documents/{ID} : http://localhost:9000/documents/b2a1f484840f512e93838cbd543acea5,
### Output
The server will respond with {"ID":"b2a1f484840f512e93838cbd543acea5","Name":"fileA","Size":17}

## Authors

* **Josue Ferrufino** - *Developer*

## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE) file for details
