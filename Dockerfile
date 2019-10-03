############################
# STEP 1 build executable binary
############################
FROM golang:alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/khawerrind/go-aws-face-rekognition
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get ./...
# Build the binary.
RUN CGO_ENABLED=0 go build -o main .

# Run the binary.
CMD ["/go/src/github.com/khawerrind/go-aws-face-rekognition/main"]