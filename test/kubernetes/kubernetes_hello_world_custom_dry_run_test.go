package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
)

var tlsConfig = tls.Config{InsecureSkipVerify: true}
var options = k8s.NewKubectlOptions("", "", "hello-world")

func TestKubernetesHelloWorldCustomIngressStatus(t *testing.T) {
	t.Parallel()
	ingress := k8s.GetService(t, k8s.NewKubectlOptions("", "", "kube-system"), "addon-http-application-routing-nginx-ingress")
	publicUrl := fmt.Sprintf("https://%s", k8s.GetServiceEndpoint(t, options, ingress, 443))
	http_helper.HttpGetWithRetry(t, publicUrl+"/hello-world", &tlsConfig, 200, "Hello world!", 3, 3*time.Second)
}

func TestKubernetesHelloWorldCustomIngressDefaultBackend(t *testing.T) {
	t.Parallel()
	ingress := k8s.GetService(t, k8s.NewKubectlOptions("", "", "kube-system"), "addon-http-application-routing-nginx-ingress")
	publicUrl := fmt.Sprintf("https://%s", k8s.GetServiceEndpoint(t, options, ingress, 443))
	http_helper.HttpGetWithRetry(t, publicUrl, &tlsConfig, 404, "default backend - 404", 3, 3*time.Second)
}

func TestKubernetesHelloWorldCustomInternalService(t *testing.T) {
	t.Parallel()
	service := k8s.GetService(t, options, "hello-world-service")
	pod := k8s.GetPod(t, options, "busybox")
	m, e := k8s.RunKubectlAndGetOutputE(t, options, "exec", pod.Name, "--", "wget", "-qO", "-", service.Name+":80")
	if e == nil {
		t.Log(m)
	}
	assert.Equal(t, "Hello world!", m)
}
