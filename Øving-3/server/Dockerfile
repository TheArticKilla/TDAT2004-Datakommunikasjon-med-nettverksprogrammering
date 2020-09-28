FROM golang:alpine

# Install tools required for project
RUN apk add --no-cache git
RUN go get -d github.com/gorilla/mux
RUN apk add python3

# Run `docker build --no-cache .` to update dependencies
RUN go get github.com/golang/dep/cmd/dep

COPY main.go .

CMD ["go", "run", "main.go"]
