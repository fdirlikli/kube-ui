package kubernetes

import v1 "k8s.io/api/core/v1"

type Pod struct {
	Name      string      `json:"name"`
	NameSpace string      `json:"namespace"`
	Phase     v1.PodPhase `json:"phase"`
	Image     string      `json:"image"`
}

type NameSpace struct {
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Phase     v1.NamespacePhase `json:"phase"`
}
