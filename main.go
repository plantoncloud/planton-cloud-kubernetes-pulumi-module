package main

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton-cloud-kubernetes-pulumi-module/pkg"
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/kubernetes/plantoncloudkubernetes"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &plantoncloudkubernetesv1.PlantonCloudKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
