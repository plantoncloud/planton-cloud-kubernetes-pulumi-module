package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/plantoncloudkubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *plantoncloudkubernetes.PlantonCloudKubernetesStackInput) *plantoncloudkubernetes.PlantonCloudKubernetesStackOutputs {
	return &plantoncloudkubernetes.PlantonCloudKubernetesStackOutputs{}
}
