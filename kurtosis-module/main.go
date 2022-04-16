package main

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/lib/execution"
	"github.com/tedim52/practice-kurtosis-module/kurtosis-module/impl"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

func main() {

	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<
	configurator := impl.NewExampleExecutableKurtosisModuleConfigurator()
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<


	executor := execution.NewKurtosisModuleExecutor(configurator)
	if err := executor.Run(); err != nil {
		logrus.Errorf("An error occurred running the Kurtosis module executor:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}
