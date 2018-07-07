# Dockerize farmers-market app.

# This image has the value of $GOPATH set to /go.
# All packages installed in /go/src will be accessible to the go command.
FROM golang:1.10

WORKDIR /go/src/farmers-market
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Expose the application on port 8080
EXPOSE 8080

CMD ["farmers-market"]