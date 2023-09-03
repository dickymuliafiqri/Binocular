FROM node:lts

ENV WORKDIR /usr/src/binocular
ENV WEBROOT /var/www/html

WORKDIR ${WORKDIR}

RUN apt-get update
RUN apt -y install golang

RUN git clone https://github.com/dickymuliafiqri/Binocular .
RUN cd ${WORKDIR}/web && npm install && npm run generate && cp -r .output/public/* ${WEBROOT}

RUN cd ${WORKDIR} && go mod download && go mod tidy && go mod verify
RUN go build -o /usr/local/bin/binocular ./cmd/binocular/main.go
RUN rm -rf *

EXPOSE 8080

CMD ["binocular"]