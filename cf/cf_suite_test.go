package cf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/cf-test-helpers/cf"
	"github.com/peterellisjones/cf-test-helpers/runner"
)

var originalCf = cf.Cf
var originalStarter = runner.SessionStarter

var _ = AfterEach(func() {
	cf.Cf = originalCf
	runner.SessionStarter = originalStarter
})

func TestCf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cf Suite")
}
