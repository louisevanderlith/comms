# comms
Mango API: Communications

## Run with Docker
*$ go build
*$ docker build -t avosa/comms:dev .
*$ docker rm commsDEV
*$ docker run -d -p 8085:8085 --network mango_net --name commsDEV avosa/comms:dev
*$ docker logs commsDEV