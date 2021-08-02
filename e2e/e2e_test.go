package e2e_test

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("E2E", func() {
	type CompileRelease struct {
		Release     string
		ExpectError bool
		Output      []string
		Timeout     int
	}
	DescribeTable("can compile release",
		func(c CompileRelease) {
			wd, err := os.Getwd()
			Expect(err).ShouldNot(HaveOccurred())
			// #nosec G204
			command := exec.Command("docker",
				"run",
				"-t",
				"-e", "INPUT_FILE=/tmp/"+c.Release,
				"-e", "INPUT_ARGS=--guess",
				"-v", path.Join(wd, "..", "test_releases")+":/tmp",
				"cloudfoundry/bc:latest")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())

			contents := string(session.Wait(c.Timeout).Out.Contents())
			fmt.Println(contents)

			for _, out := range c.Output {
				Expect(contents).Should(ContainSubstring(out))
			}

			if c.ExpectError {
				Expect(session.ExitCode()).ShouldNot(Equal(0))
			} else {
				Expect(session.ExitCode()).Should(Equal(0))
			}
		},
		Entry("with simple bosh release", CompileRelease{
			Release: "simple.tgz",
			Timeout: 5000,
			Output: []string{
				"passing -> [passing]",
				"Building Simple Package...",
			},
		}),
		Entry("with dependant bosh release", CompileRelease{
			Release: "dependant.tgz",
			Timeout: 5000,
			Output: []string{
				"dependant-package -> [dependant-package]",
				"top-level-package -> [dependant-package top-level-package]",
				"Building Dependant Package...",
				"Building Top Level Package...",
				".compiled",
			},
		}),
	)
})
