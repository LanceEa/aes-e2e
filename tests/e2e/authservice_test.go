package e2e_test

import (
	"context"
	"fmt"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// TestHello shows an example of a test environment
// that uses a simple setup to assess a feature (test)
// in a test function directly (outside of test suite TestMain)
func TestHello(t *testing.T) {

	feat := features.New("Hello Feature").
		WithLabel("type", "simple").
		Assess("test message", func(ctx context.Context, t *testing.T, _ *envconf.Config) context.Context {
			result := Hello("foo")
			if result != "Hello foo" {
				t.Error("unexpected message")
			}
			return ctx
		})

	testenv.Test(t, feat.Feature())
}
