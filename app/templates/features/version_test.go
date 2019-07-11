// +build feature

package features_test

import (
	"os/exec"

	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = FDescribe("Report version", func() {
	steps := NewSteps()

	Scenario("version command reports version", func() {
		steps.Given("the <%= appName %> command is built with a version")

		steps.When("version subcommand is run")

		steps.Then("the command exits without error")
		steps.And("the result is the version")
	})

	steps.Define(func(define Definitions) {
		var (
			<%= appName %>Path string
			commandSession *gexec.Session
		)

		define.Given(`^the <%= appName %> command is built with a version$`, func() {
			var err error
			<%= appName %>Path, err = gexec.Build(
				"<%= repoURL %>/cmd/<%= appName %>",
				"-ldflags",
				"-X <%= repoURL %>/version.Version=1.0.1",
				)
			Expect(err).NotTo(HaveOccurred())
		}, func() {
			gexec.CleanupBuildArtifacts()
		})

		define.When(`^version subcommand is run$`, func() {
			versionCommand := exec.Command(<%= appName %>Path, "version")
			var err error
			commandSession, err = gexec.Start(versionCommand, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.Then(`^the command exits without error$`, func() {
			Eventually(commandSession).Should(gexec.Exit(0))
		})

		define.Then(`the result is the version`, func() {
			Eventually(commandSession.Out).Should(Say("<%= appName %> version: 1.0.1"))
		})
	})
})
