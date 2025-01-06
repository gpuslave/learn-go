
ARG GO_VERSION=1.23.4

FROM golang:${GO_VERSION}-alpine AS build

RUN apk add --no-cache git

WORKDIR /src
COPY ./go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app


FROM gcr.io/distroless/static-debian12 AS final
LABEL maintainer="gpuslave"
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /app /app

ENTRYPOINT ["/app"]
