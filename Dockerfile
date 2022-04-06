FROM golang:1.14-alpine

ARG commit_id=master
LABEL maintainer="tom"
LABEL org.label-schema.description="Tempus"
LABEL org.label-schema.name="tempus"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.vcs-url="https://github.com/roundpartner/tempus"
LABEL org.label-schema.vcs-ref="${commit_id}"
LABEL org.label-schema.vendor="RoundPartner"

ARG build_number=unknown
ENV VERSION=${build_number}
ENV PATH=/

WORKDIR /
COPY tempus tempus

ENTRYPOINT ["tempus"]
