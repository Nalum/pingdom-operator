package main

import (
	"log"
	"os"

	"github.com/nalum/pingdom-operator/pkg/pingdom"
	extensionsobj "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	k8sConfig, err := rest.InClusterConfig()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	k8sClient, err := kubernetes.NewForConfig(k8sConfig)

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	r := k8sClient.ExtensionsV1beta1()
	rc := r.RESTClient()
	rcg := rc.Get()
	rcg.Name("pingdom.luke.mallon.io")
	rcg.

	crd := extensionsobj.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pingdom.luke.mallon.io",
		},
	}

	client := pingdom.NewClient()
	httpCheck, err := pingdom.NewHTTPCheck("Testing", "https://google.ie")

	if err != nil {
		log.Println(err)
	}

	err = client.CreateCheck(httpCheck)

	if err != nil {
		log.Println(err)
	}
}
