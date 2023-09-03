FROM node:lts as ui

WORKDIR /usr/src/ui

RUN git clone https://github.com/dickymuliafiqri/Binocular .
RUN cd ./web && npm install && npm run generate


FROM golang:latest

WORKDIR /usr/src/api

RUN mkdir -p /var/www/html

COPY --from=ui /usr/src/ui/web/.output/public/* /var/www/html

RUN git clone https://github.com/dickymuliafiqri/Binocular .
RUN go mod download && go mod tidy && go mod download
RUN go build -o /usr/local/bin/binocular ./cmd/binocular/main.go
RUN rm -rf *

EXPOSE 8080

CMD ["/usr/local/bin/binocular"]