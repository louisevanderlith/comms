# comms
Mango API: Comms
Email, SMS and other Messages are all to handled by Comms.

## Run with Docker
* $ docker build -t avosa/comms:dev .
* $ docker rm commsDEV
* $ docker run -d -p 8085:8085 --network mango_net --name CommsDEV avosa/comms:dev
* $ docker logs commsDEV

--ENV to be moved.
smtpUsername = hello@gmail.com
smtpPassword = HXXXfP59z4G7t
smtpAddress = smtp.gmail.com
smtpPort = 587
defaultEmail = backup@yahoo.com