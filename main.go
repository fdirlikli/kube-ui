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
	"github.com/fdirlikli/kube-ui/data"
	route "github.com/fdirlikli/kube-ui/routing"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	//
	//http.HandleFunc("/", getAllPods)
	//http.HandleFunc("/create", createPod)
	//http.HandleFunc("/delete", deletePod)
	handler := cors.Default().Handler(route.GetRouter())
	data.Test()
	err := http.ListenAndServe(":8080", handler)
	panic(err.Error())
}
