# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.13-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates

# Copy local code to the container image.
WORKDIR /go/src/github.com/hirac1220/go-clean-architecture/api
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# WORKDIR /github.com/hirac/graphql-grpc-testpro/questionnaire
RUN go build -o /go/bin/api ./main.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest

# Copy the binary to the production image from the builder stage.
WORKDIR /usr/bin
COPY --from=build /go/bin .

# Run the web service on container startup.
ENV PORT 8080
EXPOSE 8080
CMD ["api"]
