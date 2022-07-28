# godice
This is my first Go project. I use this as a knowledge store. When I learn something new about go or go projects I incorporate it into this project.

## dice
This rolls two dice a number of times and shows the results.

```console
/godice/cmd/dice > go run dice.go
/godice/cmd/dice > go build dice.go
/godice/cmd/dice > ./dice
```

## dice-app
This is a containerized web app that rolls a die

http://localhost:8080/

http://localhost:8080/roll

```console
/godice > docker build -t dice-app -f cmd/dice-app/Dockerfile .
/godice > docker run -p 8080:8081 -it dice-app
```
or
```console
/godice > docker compose up
```
