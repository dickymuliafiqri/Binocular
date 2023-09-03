FROM node:lts

ENV WORKDIR /usr/src/binocular
ENV WEBROOT /var/www/html

WORKDIR ${WORKDIR}

RUN curl -Lo go.tar.gz "https://go.dev/dl/go1.21.0.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf go.tar.gz 
RUN rm go.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc

RUN git clone https://github.com/dickymuliafiqri/Binocular .
RUN cd ${WORKDIR}/web && npm install && npm run generate && cp -r .output/public/* ${WEBROOT}

RUN cd ${WORKDIR} && go mod download && go mod tidy && go mod verify
RUN go build -o /usr/local/bin/binocular ./cmd/binocular/main.go
RUN rm -rf *

EXPOSE 8080

CMD ["binocular"]