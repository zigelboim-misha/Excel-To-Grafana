FROM golang:1.20 AS build
USER root
WORKDIR /app
COPY . .
RUN go get && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o excel-to-grafana main.go

FROM alpine:3.14 AS run
WORKDIR /app
COPY --from=build /app/excel-to-grafana /app/excel-to-grafana

ENV SPREADSHEET=Sheet1
ENV FILE_NAME=Book
EXPOSE 8081
ENTRYPOINT [ "/app/excel-to-grafana" ]