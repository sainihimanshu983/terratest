package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesCustomIngressStatus(t *testing.T) {
	tlsConfig := tls.Config{InsecureSkipVerify: true}
	options := k8s.NewKubectlOptions("", "", "hello-world")
	t.Parallel()
	ingress := k8s.GetService(t, k8s.NewKubectlOptions("", "", "kube-system"), "addon-http-application-routing-nginx-ingress")
	publicUrl := fmt.Sprintf("https://%s", k8s.GetServiceEndpoint(t, options, ingress, 443))
	http_helper.HttpGetWithRetry(t, publicUrl+"/hello-world", &tlsConfig, 200, "Hello world!", 3, 3*time.Second)
}
