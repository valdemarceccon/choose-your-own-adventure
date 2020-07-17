FROM golang:1.13-alpine as BUILDER

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/cyoaweb/main.go  

FROM scratch
RUN ABOBRINHA
COPY --from=BUILDER /server /server
COPY --from=BUILDER /app/gopher.json /gopher.json

ENTRYPOINT [ "/server" ]
