/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

//import "awesomeProject/testcli/cmd"
import (
	"encoding/json"
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"path/filepath"
)

var kubeconfig *string

func main() {
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	http.HandleFunc("/", getAllPods)
	http.HandleFunc("/create", createPod)
	http.HandleFunc("/delete", deletePod)
	http.ListenAndServe(":8080", nil)
}

func deletePod(w http.ResponseWriter, r *http.Request) {
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)
	clientset.CoreV1().Pods("default").Delete("test-pod", nil)

}

func createPod(w http.ResponseWriter, r *http.Request) {
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)
	pods, _ := clientset.CoreV1().Pods("default").Create(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
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
	if pods != nil {
		pod, _ := json.Marshal(pods);
		w.Write(pod)
	}
}

func getAllPods(w http.ResponseWriter, r *http.Request) {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	w.Header().Set("Content-Type", "application/json")
	var a []Pod
	for _, pod := range pods.Items {
		phase := pod.Status.Phase
		a = append(a, Pod{pod.ObjectMeta.Name, pod.ObjectMeta.Namespace, phase})

	}
	js, _ := json.Marshal(a)
	w.Write(js)

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

type Pod struct {
	Name      string      `json:"name"`
	NameSpace string      `json:"namespace"`
	Phase     v1.PodPhase `json:"phase"`
}
