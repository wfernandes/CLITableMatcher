package main

import (
	"fmt"

	"bytes"

	"strings"

	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/util/testhelpers/io"
	"github.com/wfernandes/CLITableMatcher"
)

func main() {

	actualTable := terminal.NewTable([]string{"Name", "Type"})
	actualTable.Add("name1", "type1")
	actualTable.Add("name2", "type2")

	// copy the object
	expectedTable := &terminal.Table{}
	*expectedTable = *actualTable

	var b []byte
	actualOutput := bytes.NewBuffer(b)
	actualTable.PrintTo(actualOutput)
	actual := actualOutput.String()
	capturedOutput := io.CaptureOutput(func() {
		fmt.Println(actualOutput)
	})
	println("len of captured output:", len(capturedOutput))
	println("len of actual output", len(strings.Split(actualOutput.String(), "\n")))

	formattedCapturedOutput := strings.Join(capturedOutput[:len(capturedOutput)-1], "\n")
	fmt.Printf("Formatted Captured:\n%s", formattedCapturedOutput)
	fmt.Printf("Actual:\n%s", actual)
	success, err := cli_table_matcher.ContainCLITable(expectedTable).Match(formattedCapturedOutput)
	if err != nil {
		panic(err)
	}

	if success {
		fmt.Println("Match was successful")
	} else {
		fmt.Println("Match was unsuccessful")
	}
}

func demonstratePrintToWeirdness(t *terminal.Table) {
	var b1 []byte
	output1 := bytes.NewBuffer(b1)
	var b2 []byte
	output2 := bytes.NewBuffer(b2)
	t.PrintTo(output1)
	// For some reason once t.PrintTo is called once for against the object
	// it is pointing to, it is unable to print again. That is, output2 is
	// empty.
	t.PrintTo(output2)
	actual1 := output1.String()
	actual2 := output2.String()
	fmt.Printf("Actual1:\n%s", actual1)
	fmt.Printf("Actual2:\n%s", actual2)
}
