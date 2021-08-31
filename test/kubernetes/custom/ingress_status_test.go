package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesCustomIngressDefaultBackend(t *testing.T) {
	t.Parallel()
	tlsConfig := tls.Config{InsecureSkipVerify: true}
	options := k8s.NewKubectlOptions("", "", "hello-world")
	ingress := k8s.GetService(t, k8s.NewKubectlOptions("", "", "kube-system"), "addon-http-application-routing-nginx-ingress")
	publicUrl := fmt.Sprintf("https://%s", k8s.GetServiceEndpoint(t, options, ingress, 443))
	http_helper.HttpGetWithRetry(t, publicUrl, &tlsConfig, 404, "default backend - 404", 3, 3*time.Second)
}
