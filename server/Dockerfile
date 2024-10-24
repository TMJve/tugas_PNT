
FROM golang:1.22.4-alpine as build


WORKDIR /app


COPY go.mod ./
RUN go mod download

COPY . .


RUN go build -o server

FROM alpine:3.18


WORKDIR /app


COPY --from=build /app/server .


EXPOSE 8080

# Command to run the server
CMD ["./server"]
