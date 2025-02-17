FROM golang:1.20 as builder

WORKDIR /go/src/github.com/liuds832/micromdm/

ARG TARGETARCH
ARG TARGETOS

ENV CGO_ENABLED=0 \
	GOARCH=$TARGETARCH \
	GOOS=$TARGETOS

COPY . .

RUN make deps
RUN make


FROM alpine:latest

RUN apk --update add ca-certificates

COPY --from=builder /go/src/github.com/liuds832/micromdm/build/linux/micromdm /usr/bin/
COPY --from=builder /go/src/github.com/liuds832/micromdm/build/linux/mdmctl /usr/bin/

EXPOSE 80 443
VOLUME ["/var/db/micromdm"]
CMD ["micromdm", "serve"]
