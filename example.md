---

### Example 1: Basic Kubernetes Cluster Connection

This example demonstrates how to configure the module to connect to a basic Kubernetes cluster using the provided credentials. The `kubeconfig` is specified in the `cluster_credentials` section.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: basic-kubernetes-cluster
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://k8s-api.mycompany.com
          name: my-kubernetes-cluster
      contexts:
        - context:
            cluster: my-kubernetes-cluster
            user: cluster-admin
          name: my-cluster-context
      current-context: my-cluster-context
      users:
        - name: cluster-admin
          user:
            token: my-secret-token
```

---

### Example 2: Kubernetes Cluster on GKE (Google Kubernetes Engine)

This example connects to a GKE cluster using a GCP service account for authentication. It assumes that the service account has appropriate permissions to access the Kubernetes API.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: gke-cluster-connection
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://gke-cluster-api.example.com
          name: gke-cluster
      contexts:
        - context:
            cluster: gke-cluster
            user: gke-user
          name: gke-cluster-context
      current-context: gke-cluster-context
      users:
        - name: gke-user
          user:
            auth-provider:
              name: gcp
              config:
                access-token: your-access-token
                cmd-path: gcloud
                cmd-args: config config-helper --format=json
                expiry-key: '{.credential.token_expiry}'
                token-key: '{.credential.access_token}'
```

---

### Example 3: EKS (Elastic Kubernetes Service) Cluster on AWS

This example configures a connection to an EKS cluster on AWS using the AWS command-line tool to retrieve an authentication token. It allows users to securely access their EKS cluster.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: eks-cluster-connection
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://eks-cluster-api.example.com
          name: eks-cluster
      contexts:
        - context:
            cluster: eks-cluster
            user: eks-user
          name: eks-cluster-context
      current-context: eks-cluster-context
      users:
        - name: eks-user
          user:
            exec:
              apiVersion: client.authentication.k8s.io/v1beta1
              command: aws
              args:
                - eks
                - get-token
                - --cluster-name
                - eks-cluster
```

---

### Example 4: Azure Kubernetes Service (AKS) Cluster Connection

This example connects to an Azure Kubernetes Service (AKS) cluster using Azure Active Directory authentication. The `kubeconfig` is configured with the appropriate authentication details from Azure.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: aks-cluster-connection
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://aks-cluster-api.example.com
          name: aks-cluster
      contexts:
        - context:
            cluster: aks-cluster
            user: aks-user
          name: aks-cluster-context
      current-context: aks-cluster-context
      users:
        - name: aks-user
          user:
            auth-provider:
              name: azure
              config:
                access-token: your-access-token
                client-id: your-client-id
                tenant-id: your-tenant-id
```

---

### Example 5: Private Kubernetes Cluster with Bearer Token Authentication

This example connects to a private Kubernetes cluster using a bearer token for authentication. This setup ensures that only authorized users can access the Kubernetes API.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: private-kubernetes-cluster
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://private-k8s-api.example.com
          name: private-cluster
      contexts:
        - context:
            cluster: private-cluster
            user: private-user
          name: private-cluster-context
      current-context: private-cluster-context
      users:
        - name: private-user
          user:
            token: private-cluster-bearer-token
```

---

### Example 6: Kubernetes Cluster Connection with TLS Certificates

This example configures a Kubernetes cluster connection using TLS client certificates for secure authentication. The `kubeconfig` file includes both the client certificate and key.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: tls-authenticated-cluster
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://tls-cluster-api.example.com
          name: tls-cluster
      contexts:
        - context:
            cluster: tls-cluster
            user: tls-user
          name: tls-cluster-context
      current-context: tls-cluster-context
      users:
        - name: tls-user
          user:
            client-certificate-data: your-base64-encoded-client-cert
            client-key-data: your-base64-encoded-client-key
```

---

### Future Example: Kubernetes Resource Deployment (To Be Supported)

Once the module is fully implemented, it will support provisioning Kubernetes resources, such as deployments, services, and ingresses. Below is a hypothetical example for a future iteration of the module:

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PlantonCloudKubernetes
metadata:
  name: kubernetes-resource-deployment
spec:
  cluster_credentials:
    kubeconfig: |
      apiVersion: v1
      kind: Config
      clusters:
        - cluster:
            server: https://k8s-api.example.com
          name: my-cluster
      contexts:
        - context:
            cluster: my-cluster
            user: admin-user
          name: my-cluster-context
      current-context: my-cluster-context
      users:
        - name: admin-user
          user:
            token: my-token
  resources:
    deployments:
      - name: my-deployment
        spec:
          replicas: 3
          containers:
            - name: nginx
              image: nginx:latest
              ports:
                - containerPort: 80
    services:
      - name: my-service
        spec:
          type: LoadBalancer
          ports:
            - port: 80
          selector:
            app: my-deployment
```

---

### Applying the Configuration

Once the desired YAML file is created with the configuration, you can apply it using the following command:

```shell
planton apply -f <yaml-path>
```
