FROM golang:1.14 as builder

ADD src /go/src/workdir/src

ENV HOSTNAME postgres
ENV USER postgres
ENV DBNAME wantedly
ENV PASSWORD password
ENV CGO_ENABLED 0 
ENV GO111MODULE on
ENV PORT 8080
ENV DB_PORT 5432



WORKDIR /go/src/workdir/src

RUN go build -o APIServer

# runtime image
FROM alpine 
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/workdir/src/APIServer /APIServer
EXPOSE 8080
#ENV PORT 8080

ENTRYPOINT [ "./APIServer" ]
