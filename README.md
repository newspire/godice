# godice
This is my first Go application.
It rolls two dice.

/godice/dice > go run dice.go
/godice/dice > go build dice.go
/godice/dice > ./dice

# dice-app
# http://localhost:8080/
# http://localhost:8080/roll
/godice > docker build -t dice-app -f dice-app/Dockerfile .
/godice > docker run -p 8080:8081 -it dice-app
