# PlantonCloud Kubernetes Pulumi Module

### API-Resource Highlights

- **Standardized Structure**: Follows the Kubernetes API resource model, including `apiVersion`, `kind`, `metadata`, `spec`, and `status`, ensuring consistency and ease of integration.
- **Unified API Approach**: Designed for the multi-cloud era, allowing seamless interaction with various cloud providers through a consistent API structure.
- **Environment Configuration**: Supports specifying environment information and Kubernetes cluster credentials to tailor deployments to different environments and clusters.
- **Extensibility**: Serves as a foundation for future enhancements, enabling the addition of more complex resource provisioning and management features.

### Module Features

- **Pulumi Integration**: Built in Go, the module integrates seamlessly with Pulumi, leveraging its powerful IaC capabilities for managing Kubernetes infrastructure.
- **Kubernetes Provider Setup**: Automates the setup of the Kubernetes provider using the provided cluster credentials, facilitating secure and efficient connections to target Kubernetes clusters.
- **Modular Design**: Designed to be easily extensible, allowing for the incorporation of additional functionalities and integrations as needed.
- **Consistent Deployment Pattern**: Utilizes a standardized pattern where API resources serve as inputs to Pulumi modules, promoting consistency and reducing complexity in infrastructure deployments.
- **Future-Proof Architecture**: Positioned to support a wide range of Kubernetes deployment scenarios across multiple cloud environments, aligning with the evolving needs of modern infrastructure management.

## Installation

To install the PlantonCloud Kubernetes Pulumi Module, follow the standard Pulumi module installation procedures. Ensure that you have Pulumi installed and configured in your development environment. The module is available on GitHub, and you can include it in your Pulumi projects by referencing the repository in your project’s dependencies.

```shell
pulumi plugin install resource plantoncloud-kubernetes <version>
```

Replace `<version>` with the desired version of the module.

## Configuration

The module requires a `PlantonCloudKubernetesStackInput` specification to define the desired state of the Kubernetes deployment. This specification includes configurations for Kubernetes cluster credentials and other essential parameters necessary for establishing a connection to the target Kubernetes cluster.

### Key Configuration Parameters

- **Kubernetes Cluster Credentials**: Provide the necessary credentials to authenticate and interact with the target Kubernetes cluster. This includes specifying the `kubernetes_cluster_credential_id` to set up the Kubernetes provider.
- **Pulumi Input**: Define the Pulumi-specific configurations required for the stack job, facilitating the integration between the API resource and Pulumi’s deployment processes.
- **Target API-Resource**: Specify the `PlantonCloudKubernetes` target, linking the API resource definition with the Pulumi module for deployment.

## Usage

Refer to the example section for usage instructions.

## Outputs

Upon successful execution, the module is expected to export outputs related to the Kubernetes namespace and other deployment details. However, as the module is not fully implemented, the current outputs are limited.

- **Namespace**: The name of the Kubernetes namespace where PlantonCloud Kubernetes resources are intended to be deployed.

Future iterations of the module will provide additional outputs to facilitate comprehensive management and integration of Kubernetes resources.

## Documentation

Comprehensive documentation for the API resources and module is available via [buf.build](https://buf.build). This documentation provides detailed insights into the API structure, configuration options, and operational guidelines, ensuring that developers can effectively utilize the module to deploy and manage Kubernetes services within the PlantonCloud ecosystem.

## Contributing

Contributions to the PlantonCloud Kubernetes Pulumi Module are welcome. Please refer to the `CONTRIBUTING.md` file in the repository for guidelines on how to contribute, report issues, and suggest enhancements. Your contributions help improve the module and expand its capabilities to better serve the community.

## License

This project is licensed under the [MIT License](LICENSE).