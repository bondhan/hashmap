# Hashmap

## What?
My own interpretation of hashmap algorithm written in golang

## What 2? 
A small web server application to call the hashmap function

## How?

### Test
```
 go test ./... -cover -race
```

### Build
#### Linux
```
docker build --platform linux/amd64 -f Dockerfile . -t hashmap-web:linux-v1
```

### Run
#### Linux
```
docker run -p 8080:3030 hashmap-web:linux-v1
```

### Test (Todo)
```
curl http://localhost:8080/init
curl http://localhost:8080/put
curl http://localhost:8080/get
curl http://localhost:8080/remove
```