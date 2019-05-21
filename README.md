# comms
Mango API: Comms
Email, SMS and other Messages are all to handled by Comms.

## Run with Docker
* $ docker build -t avosa/comms:dev .
* $ docker rm commsDEV
* $ docker run -d -e RUNMODE=DEV -p 8085:8085 --network mango_net --name CommsDEV avosa/comms:dev
* $ docker logs commsDEV