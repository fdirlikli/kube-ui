An application is a logical component of the with below attributes:
* Name
* Label
* Owner

Application groups together specific Kubernetes constructs to ease the monitoring of these constructs all at once. Below are the eligible Kubernetes constructs:
* Pod
* Deployment
* Statefulset
* Daemonset

When a Kubernetes construct is added to an application we call it an application component. Application components can span across Kubernetes namespaces.
