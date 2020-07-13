package integration

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/testhelpers/e2e"
)

type ConditionIntegrationTestSuite struct {
	suite.Suite
}

func (suite *ConditionIntegrationTestSuite) TestCondition() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	suite.PrepareActiveStateYAML(ts)

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("run", "test"),
		e2e.AppendEnv("VERBOSE=true"),
	)
	cp.Expect(`projectNameValue`)
	cp.Expect(`projectOwnerValue`)
	cp.Expect(`projectNamespaceValue`)
	cp.Expect(`osNameValue`)
	cp.Expect(`osVersionValue`)
	cp.Expect(`osArchValue`)
	cp.Expect(`shellValue`)
	cp.ExpectExitCode(0)

	cp = ts.SpawnWithOpts(
		e2e.WithArgs("activate"),
	)
	cp.Expect(`Activation Event Ran`)
	cp.WaitForInput()
	cp.SendLine("exit")
	cp.ExpectExitCode(0)

	cp = ts.SpawnWithOpts(
		e2e.WithArgs("run", "complex-true"),
	)
	cp.Expect(`I Exist`)
	cp.ExpectExitCode(0)

	cp = ts.SpawnWithOpts(
		e2e.WithArgs("run", "complex-false"),
	)
	cp.ExpectExitCode(1)
}

func (suite *ConditionIntegrationTestSuite) TestConditionOSName() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	suite.PrepareActiveStateYAML(ts)

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("run", "OSName"),
	)
	if runtime.GOOS == "windows" {
		cp.Expect(`Windows`)
	} else if runtime.GOOS == "darwin" {
		cp.Expect(`MacOS`)
	} else {
		cp.Expect(`Linux`)
	}
	cp.ExpectExitCode(1)
}

func (suite *ConditionIntegrationTestSuite) TestConditionSyntaxError() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	suite.PrepareActiveStateYAMLWithSyntaxError(ts)

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("run", "test"),
	)
	cp.Expect(`not defined`) // for now we aren't passing the error up the chain, so invalid syntax will lead to empty result
	cp.ExpectExitCode(1)
}

func (suite *ConditionIntegrationTestSuite) PrepareActiveStateYAML(ts *e2e.Session) {
	asyData := strings.TrimSpace(`
project: https://platform.activestate.com/ActiveState-CLI/test?commitID=9090c128-e948-4388-8f7f-96e2c1e00d98
constants:
  - name: projectName
    value: invalidProjectName
    if: false
  - name: projectName
    value: projectNameValue
    if: ne .Project.Name ""
  - name: projectName
    value: invalidProjectName
    if: false
  - name: projectOwner
    value: projectOwnerValue
    if: ne .Project.Owner ""
  - name: projectNamespace
    value: projectNamespaceValue
    if: ne .Project.Namespace ""
  - name: osName
    value: osNameValue
    if: ne .OS.Name ""
  - name: osVersion
    value: osVersionValue
    if: ne .OS.Version.Name ""
  - name: osArch
    value: osArchValue
    if: ne .OS.Architecture ""
  - name: shell
    value: shellValue
    if: ne .Shell ""
scripts:
  - name: complex-true
    language: bash
    standalone: true
    value: echo "I exist"
    if: or (eq .OS.Architecture "") (contains .OS.Architecture "64")
  - name: complex-false
    language: bash
    standalone: true
    value: echo "I exist"
    if: and (eq .OS.Architecture "") (contains .OS.Architecture "64")
  - name: test
    language: bash
    standalone: true
    value: echo wrong script
    if: false
  - name: test
    standalone: true
    language: bash
    value: |
      echo ${constants.projectName}
      echo ${constants.projectOwner}
      echo ${constants.projectNamespace}
      echo ${constants.osName}
      echo ${constants.osVersion}
      echo ${constants.osArch}
      echo ${constants.shell}
    if: ne .Shell ""
  - name: test
    language: bash
    standalone: true
    value: echo wrong script
    if: false
  - name: OSName
    language: bash
    standalone: true
    value: echo using-windows
    if: eq .OS.Name "Windows"
  - name: OSName
    language: bash
    standalone: true
    value: echo using-macos
    if: eq .OS.Name "MacOS"
  - name: OSName
    language: bash
    standalone: true
    value: echo using-linux
    if: eq .OS.Name "Linux"
events:
  - name: ACTIVATE
    value: echo "Wrong event"
    if: false
  - name: ACTIVATE
    value: echo "Activation Event Ran"
    if: ne .Shell ""
  - name: ACTIVATE
    value: echo "Wrong event"
    if: false
`)

	ts.PrepareActiveStateYAML(asyData)
}
func (suite *ConditionIntegrationTestSuite) PrepareActiveStateYAMLWithSyntaxError(ts *e2e.Session) {
	asyData := strings.TrimSpace(`
project: https://platform.activestate.com/ActiveState-CLI/test?commitID=9090c128-e948-4388-8f7f-96e2c1e00d98
scripts:
  - name: test
    language: bash
    standalone: true
    value: echo invalid value
    if: not a valid conditional
  - name: test
    language: bash
    standalone: true
    value: echo valid value
    if: true
`)

	ts.PrepareActiveStateYAML(asyData)
}

func TestConditionIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ConditionIntegrationTestSuite))
}
