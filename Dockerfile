FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.20-openshift-4.16 AS builder
WORKDIR /go/src/github.com/openshift/azure-storage-azcopy
COPY . .
RUN go build -o ./bin/azcopy .

FROM registry.ci.openshift.org/ocp/4.16:base
COPY --from=builder /go/src/github.com/openshift/azure-storage-azcopy/bin/azcopy /usr/bin/