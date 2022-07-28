# godice
This is my first Go application.
It rolls two dice.

/godice/cmd/dice > go run dice.go
/godice/cmd/dice > go build dice.go
/godice/cmd/dice > ./dice

# dice-app
# http://localhost:8080/
# http://localhost:8080/roll
/godice > docker build -t dice-app -f cmd/dice-app/Dockerfile .
/godice > docker run -p 8080:8081 -it dice-app
