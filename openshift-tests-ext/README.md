# ACM Lifecycle OpenShift Tests Extension (OTE)

This directory contains the OpenShift Tests Extension (OTE) binary for ACM Lifecycle testing.

## Overview

The `acm-tests-ext` binary implements the OTE v1.1 protocol so ACM hub/lifecycle checks can run under `openshift-tests` with `--extension-local-binaries`.

## Building

```bash
make build
```

## Running Tests

```bash
./acm-tests-ext run-suite acm/lifecycle/all
./acm-tests-ext run-suite acm/lifecycle/smoke
./acm-tests-ext list tests
./acm-tests-ext info
```

## Test Suites

- `acm/lifecycle/all`
- `acm/lifecycle/smoke`

Scripts are embedded in the binary. The adapter extracts them to a temp directory at runtime and sets `INTEROP_SCRIPTS_DIR`. Shell scripts use exit `2` for skip conditions; the adapter maps that to OTE `skipped`.

## Related

- JIRA: https://redhat.atlassian.net/browse/INTEROP-9589
- OpenShift Tests Extension framework: https://github.com/openshift-eng/openshift-tests-extension
