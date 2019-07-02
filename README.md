# comms
Mango API: Comms
Email, SMS and other Messages are all to handled by Comms.

## Run with Docker
* $ docker build -t avosa/comms:dev .
* $ docker rm CommsDEV
* $ docker run -d -p 8085:8085 -e SMTPUsername=frik@avosa.co.za -e SMTPPassword=not_real -e SMTPAddress=smtp.gmail.com -e SMTPPort=587 --network mango_net --name CommsDEV avosa/comms:dev
* $ docker logs CommsDEV