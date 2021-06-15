FROM golang:alpine

WORKDIR /build



COPY . .

RUN go build -o server .

WORKDIR /dist

RUN cp /build/server .

EXPOSE 8081

CMD [ "/dist/server" ]