# syntax=docker/dockerfile:1
FROM gcr.io/distroless/base

LABEL maintainer="hello@lukasbahr.de"

COPY deconz-ctl /usr/bin/deconz-ctl
ENTRYPOINT ["/usr/bin/deconz-ctl"]