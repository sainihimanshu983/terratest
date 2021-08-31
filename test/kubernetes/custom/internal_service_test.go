package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesCustomInternalService(t *testing.T) {
	t.Parallel()
	options := k8s.NewKubectlOptions("", "", "hello-world")
	service := k8s.GetService(t, options, "hello-world-service")
	pod := k8s.GetPod(t, options, "busybox")
	m, e := k8s.RunKubectlAndGetOutputE(t, options, "exec", pod.Name, "--", "wget", "-qO", "-", service.Name+":80")
	if e == nil {
		t.Log(m)
	}
	assert.Equal(t, "Hello world!", m)
}
