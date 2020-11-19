ARG GOLANG_VERSION=1.15.3
ARG ALPINE_VERSION=3.12

# build targets.
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as build

RUN mkdir /project

WORKDIR /project

# Copy source code to the container.
COPY . .

RUN go get all

# expose debug api ports
#EXPOSE 3002

ENTRYPOINT [ "go", "build", "cmd/server.go" ]

# build a binary & move to /var
#RUN CGO_ENABLED=0 go build -o /var/api cmd/server.go

# expose initial ports
#EXPOSE 3002

# copy the built binary from the build process.
#COPY --from=build /var/api /var/api

# set the work directory.
#WORKDIR /var

# run the binary.
#CMD ./api

# --- development stage ---
# https://hub.docker.com/_/golang/
#FROM build AS development

# expose debug api port.
#EXPOSE 3002

#CMD [ "go", "run", "cmd/server.go" ]
