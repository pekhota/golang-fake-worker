# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.# Using go get.
RUN go get -d -v

# Build the binary.

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/async-worker ./

#ENTRYPOINT ["/go/bin/async-worker"]
############################
# STEP 2 build a small image
############################
FROM alpine

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/async-worker /go/bin/async-worker

# Use an unprivileged user.
USER appuser:appuser

ENV PORT=9292

# Port on which the service will be exposed.
EXPOSE ${PORT}

# Run the hello binary.
ENTRYPOINT ["/go/bin/async-worker"]