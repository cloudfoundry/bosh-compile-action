FROM --platform=${BUILDPLATFORM} ubuntu:18.04

RUN apt-get update && apt-get install -y make ca-certificates

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

COPY build/linux/bc /usr/bin/bc
COPY github-actions-entrypoint.sh /usr/bin/github-actions-entrypoint.sh

RUN mkdir /var/vcap \
	&& chmod a+w /var/vcap \
	&& mkdir /var/vcap/packages

ENTRYPOINT [ "/usr/bin/github-actions-entrypoint.sh" ]
