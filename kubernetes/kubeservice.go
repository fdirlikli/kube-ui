package kubernetes

import (
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var Service *kubeService

type kubeService struct {
	kubeconfig *string
	clientset  *kubernetes.Clientset
}

func init() {
	Service = new(kubeService)
	Service.init()
}

//func test() (*rest.Config, error) {
//	kubeconfigPath := ""
//	masterUrl := ""
//	if kubeconfigPath == "" && masterUrl == "" {
//		klog.Warningf("Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.")
//		//kubeconfig, err := restclient.InClusterConfig()
//		//
//		klog.Warning("error creating inClusterConfig, falling back to default config: ", err)
//	}
//	  c,b := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
//		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
//		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterUrl}}).RawConfig()
//	  for index,a := range c.Contexts{
//	  	a.Namespace
//	  }
//}

func (s *kubeService) init() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)
	s.clientset = clientset

}

func (s *kubeService) DeletePod(podname *string) {
	err := s.clientset.CoreV1().Pods("default").Delete(*podname, nil)
	if err != nil {
		panic(err.Error())
	}
}

func (s *kubeService) CreatePod(podname *string, namespace *string) *Pod {

	kpod, err := s.clientset.CoreV1().Pods(*namespace).Create(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: *podname,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}
	return &Pod{kpod.Name, kpod.Namespace, kpod.Status.Phase, kpod.Spec.Containers[0].Image}

}

func (s *kubeService) GetAllPods(namespace *string) []Pod {
	kpods, err := s.clientset.CoreV1().Pods(*namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var pods []Pod

	for _, pod := range kpods.Items {

		pods = append(pods, Pod{pod.Name, pod.Namespace, pod.Status.Phase, pod.Spec.Containers[0].Image})
	}

	return pods
}

func (s *kubeService) GetNamespaces() []NameSpace {
	knamespaces, err := s.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var namespaces []NameSpace
	for _, namespace := range knamespaces.Items {
		namespaces = append(namespaces, NameSpace{
			Name:      namespace.Name,
			NameSpace: namespace.Namespace,
			Phase:     namespace.Status.Phase,
		})
	}

	return namespaces

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
