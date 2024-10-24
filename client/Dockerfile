
FROM golang:1.22.4-alpine as build


WORKDIR /app


COPY go.mod ./
RUN go mod download


COPY . .


RUN go build -o client


FROM alpine:3.18

# Set working directory
WORKDIR /app


COPY --from=build /app/client .

# Command to run the client
CMD ["./client"]
