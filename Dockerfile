FROM golang:latest AS build

WORKDIR /build

COPY go.sum go.mod ./

RUN go mod download

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

COPY main.go ./

RUN CGO_ENABLED=0 GOOS="${TARGETOS:-linux}" GOARCH="${TARGETARCH:-amd64}" GOARM="$(echo ${TARGETVARIANT:-v7} | tr -d -c 0-9)" go build -o tftp

FROM scratch

COPY --from=build /build/tftp /

ENTRYPOINT ["/tftp"]
