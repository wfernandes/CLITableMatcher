package cli_table_matcher

import (
	"bytes"
	"strings"

	"fmt"

	"code.cloudfoundry.org/cli/cf/terminal"
	"github.com/onsi/gomega/types"
)

type tableMatcher struct {
	expected    *terminal.Table
	expectedStr string
}

func ContainCLITable(expected *terminal.Table) types.GomegaMatcher {
	return &tableMatcher{
		expected: expected,
	}
}

func (m *tableMatcher) Match(actual interface{}) (bool, error) {
	a, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("CLITable matcher expects a string")
	}

	var b []byte
	expectedBytes := bytes.NewBuffer(b)
	err := m.expected.PrintTo(expectedBytes)
	if err != nil {
		return false, fmt.Errorf("unable to process expected")
	}
	m.expectedStr = strings.TrimSpace(expectedBytes.String())

	if strings.Contains(a, m.expectedStr) {
		return true, nil
	}

	return false, nil
}

func (m *tableMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected \n%s  to match \n%s", actual, m.expectedStr)
}

func (m *tableMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected \n%s  to not match \n%s", actual, m.expectedStr)
}
