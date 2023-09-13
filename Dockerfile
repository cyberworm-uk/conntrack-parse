FROM --platform=$BUILDPLATFORM docker.io/library/golang:alpine AS build
ARG TARGETOS TARGETARCH
ENV GOOS="$TARGETOS" GOARCH="$TARGETARCH" GOFLAGS="-buildvcs=false -trimpath" CGO_ENABLED=0 
WORKDIR /go/src/conntrack
COPY . .
RUN --mount=type=cache,target=/go/pkg go mod download
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build go build -ldflags '-w -s -buildid=' ./cmd/cli

FROM scratch
COPY --from=build /go/src/conntrack/cli /conntrack
ENTRYPOINT [ "/conntrack" ]