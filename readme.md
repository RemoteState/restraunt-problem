To run the code on local  

run the following command 
```shell
go run main.go
```

To run the test code on local

run the following command
```shell
go test
```

To build docker image

run the following command
```shell
 docker run build -t restraunt-app .
 ```

To run the docker image

run the following command
where ./logs is the directory that contains the log files on your host machine:
```shell
docker run run --rm -v "$(pwd)/logs:/app/logs" restraunt-app
 ```

