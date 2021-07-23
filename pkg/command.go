package pkg

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func Exec(workDir string, installTarget string, command string, args ...string) (string, error) {
	logrus.Infof("cd %s", workDir)
	logrus.Infof("BOSH_INSTALL_TARGET=%s BOSH_COMPILE_TARGET=%s %s %s",
		installTarget, workDir, command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Dir = workDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	environment := os.Environ()
	environment = append(environment, "BOSH_INSTALL_TARGET="+installTarget)
	environment = append(environment, "BOSH_COMPILE_TARGET="+workDir)
	//environment = append(environment, "PATH="+path)
	logrus.Debugf("Configuring environment as %s", environment)
	cmd.Env = environment

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Print the output
	return "", nil
}
