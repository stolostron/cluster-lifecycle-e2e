package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/openshift-eng/openshift-tests-extension/pkg/cmd"
	e "github.com/openshift-eng/openshift-tests-extension/pkg/extension"
	"github.com/stolostron/cluster-lifecycle-e2e/openshift-tests-ext/pkg/adapter"
	"github.com/stolostron/cluster-lifecycle-e2e/openshift-tests-ext/scripts"
	"github.com/stolostron/cluster-lifecycle-e2e/openshift-tests-ext/test/acm"
)

func main() {
	defaultTimeout := 5 * time.Minute

	registry := e.NewRegistry()
	ext := e.NewExtension("acm", "external", "acm-tests-ext")

	ext.AddSuite(e.Suite{
		Name:             "acm/lifecycle/all",
		Description:      "All ACM Lifecycle tests",
		Qualifiers:       []string{`tags["component"] == "acm-lifecycle"`},
		ClusterStability: e.ClusterStabilityStable,
		Parallelism:      1,
		TestTimeout:      &defaultTimeout,
	})

	ext.AddSuite(e.Suite{
		Name:             "acm/lifecycle/smoke",
		Description:      "ACM Lifecycle smoke tests",
		Qualifiers:       []string{`labels.exists(l, l=="SMOKE")`},
		ClusterStability: e.ClusterStabilityStable,
		Parallelism:      1,
		TestTimeout:      &defaultTimeout,
	})

	specs, err := adapter.BuildExtensionTestSpecsFromScenarios(acm.ACMLifecycleScenarios, scripts.FS)
	if err != nil {
		panic(fmt.Sprintf("couldn't build extension test specs from ACM Lifecycle scenarios: %+v", err.Error()))
	}

	ext.AddSpecs(specs)
	registry.Register(ext)

	root := &cobra.Command{
		Use:   "acm-tests-ext",
		Short: "ACM Lifecycle Tests Extension for OpenShift",
		Long:  "OpenShift Tests Extension (OTE) binary for ACM Lifecycle testing",
	}

	root.AddCommand(cmd.DefaultExtensionCommands(registry)...)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
