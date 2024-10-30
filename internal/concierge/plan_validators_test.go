package concierge

import (
	"testing"

	"github.com/jnsgruk/concierge/internal/config"
	"github.com/jnsgruk/concierge/internal/runnertest"
)

func TestSingleK8sValidator(t *testing.T) {
	runner := runnertest.NewMockRunner()

	twoK8s := &config.Config{}
	twoK8s.Providers.K8s.Enable = true
	twoK8s.Providers.MicroK8s.Enable = true

	plan := NewPlan(twoK8s, runner)
	err := plan.validate()
	if err == nil {
		t.Fatalf("should not allow enabling two local kubernetes providers")
	}

	justK8s := &config.Config{}
	justK8s.Providers.K8s.Enable = true
	plan = NewPlan(justK8s, runner)
	err = plan.validate()
	if err != nil {
		t.Fatalf("single kubernetes provider should be permitted")
	}

	justMicroK8s := &config.Config{}
	justMicroK8s.Providers.MicroK8s.Enable = true
	plan = NewPlan(justMicroK8s, runner)
	err = plan.validate()
	if err != nil {
		t.Fatalf("single kubernetes provider should be permitted")
	}

}
