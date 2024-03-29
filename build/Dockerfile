FROM registry.ci.openshift.org/stolostron/builder:go1.19-linux AS builder

ARG REMOTE_SOURCE
ARG REMOTE_SOURCE_DIR

COPY $REMOTE_SOURCE $REMOTE_SOURCE_DIR/app/
WORKDIR $REMOTE_SOURCE_DIR/app

# compile go tests in build image
# RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@vv1.16.5 && GOFLAGS="" ginkgo build -tags e2e pkg/tests
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/metrics
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/create_cluster
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/create_cluster_bm
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/import_cluster
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/detach_destroy
RUN GOFLAGS="" go install github.com/onsi/ginkgo/ginkgo@v1.16.5 && GOFLAGS="" ginkgo build pkg/tests/destroy_bm

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ARG VCS_REF
ARG VCS_URL
ARG IMAGE_NAME
ARG IMAGE_DESCRIPTION
ARG ARCH_TYPE
ARG REMOTE_SOURCE_DIR

RUN microdnf update && \
    microdnf clean all

# expose env vars for runtime
ENV KUBECONFIG "/opt/.kube/config"
ENV IMPORT_KUBECONFIG "/opt/.kube/import-kubeconfig"
ENV OPTIONS "/resources/options.yaml"
# ENV REPORT_FILE_IMPORT "/results/results-import.xml"
# ENV REPORT_FILE_CREATE "/results/results-provision"
# ENV REPORT_FILE_DESTROY "/results/results-destroy"
# ENV REPORT_FILE_METRICS "/results/results-metrics.xml"
# ENV REPORT_FILE_BM "/results/results-baremetal.xml"
ENV GINKGO_DEFAULT_FLAGS "-slowSpecThreshold=120 -timeout 7200s"
# ENV GINKGO_NODES "11"
# ENV GINKGO_FLAGS=""
# ENV GINKGO_FOCUS=""
# ENV GINKGO_SKIP=""

# run as root
USER root

# install ginkgo into built image
COPY --from=builder /go/bin/ /usr/local/bin

# Copy resources which contains the options.yaml
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/resources /test/resources

# install operator binary
# COPY --from=builder $REMOTE_SOURCE_DIR/app/e2e-test.test /test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/metrics/metrics.test /test/metrics/metrics.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/create_cluster/create_cluster.test /test/create_cluster/create_cluster.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/create_cluster_bm/create_cluster_bm.test /test/create_cluster_bm/create_cluster_bm.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/import_cluster/import_cluster.test /test/import_cluster/import_cluster.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/detach_destroy/detach_destroy.test /test/detach_destroy/detach_destroy.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/pkg/tests/destroy_bm/destroy_bm.test /test/destroy_bm/destroy_bm.test
COPY --from=builder $REMOTE_SOURCE_DIR/app/build/start-tests.sh /test/start-tests.sh
VOLUME /results
WORKDIR "/test"

#CMD ["/bin/bash", "-c", "ginkgo --v --focus=${GINKGO_FOCUS} --skip=${GINKGO_SKIP} --p --trace --debug e2e-test.test -- -v=3"]

RUN ["chmod", "+x", "start-tests.sh"]

ENTRYPOINT ["./start-tests.sh"]