package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesHelloWorldCustomTest(t *testing.T) {
	t.Parallel()
	tlsConfig := tls.Config{InsecureSkipVerify: true}
	kubeResourcePath := "../../codebase/kubernetes/kubernetes-hello-world-custom/hello-world-deployment.yml"
	options := k8s.NewKubectlOptions("", "", "hello-world")
	defer k8s.KubectlDelete(t, options, kubeResourcePath)
	k8s.KubectlApply(t, options, kubeResourcePath)
	k8s.WaitUntilServiceAvailable(t, options, "hello-world-service", 10, 1*time.Second)
	service := k8s.GetService(t, options, "hello-world-service")
	ingress := k8s.GetService(t, k8s.NewKubectlOptions("", "", "kube-system"), "addon-http-application-routing-nginx-ingress")
	privateUrl := fmt.Sprintf("http://%s", k8s.GetServiceEndpoint(t, options, service, 5000))
	publicUrl := fmt.Sprintf("https://%s", k8s.GetServiceEndpoint(t, options, ingress, 443))

	http_helper.HttpGetWithRetry(t, publicUrl+"/hello-world", &tlsConfig, 200, "Hello world!", 10, 3*time.Second)
	http_helper.HttpGetWithRetry(t, publicUrl, &tlsConfig, 404, "default backend - 404", 10, 3*time.Second)
	http_helper.HttpGetWithRetry(t, privateUrl, nil, 200, "Hello world!", 10, 3*time.Second)
}
