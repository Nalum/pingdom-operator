FROM alpine:3.7

RUN apk add --no-cache ca-certificates

COPY ./pingdom-operator /pingdom-operator

ENTRYPOINT ["/pingdom-operator"]
