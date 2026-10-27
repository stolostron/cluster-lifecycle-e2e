#!/bin/bash
set -euxo pipefail
shopt -s inherit_errexit 2>/dev/null || true

if [[ -z "${KUBECONFIG:-}" ]]; then
    echo "KUBECONFIG not set, skipping" >&2
    exit 2
fi

if ! oc get crd managedclusters.cluster.open-cluster-management.io &>/dev/null; then
    echo "WARN: CRD managedclusters.cluster.open-cluster-management.io is not installed — ACM Lifecycle may not be deployed" >&2
    exit 2
fi

echo "PASS: ACM ManagedCluster CRD is installed"
exit 0
