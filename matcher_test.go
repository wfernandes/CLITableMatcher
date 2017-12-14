package cli_table_matcher_test

import (
	. "github.com/wfernandes/CLITableMatcher"

	"bytes"

	"code.cloudfoundry.org/cli/cf/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLITableMatcher", func() {

	var table *terminal.Table

	BeforeEach(func() {
		table = terminal.NewTable([]string{"Name", "Email"})
		table.Add("foo", "foo@email.com")
	})

	It("returns error if the actual isn't stringable", func() {
		result, err := MatchCLITable(table).Match(nil)
		Expect(err).To(HaveOccurred())
		Expect(result).To(BeFalse())
	})

	It("matches cli table stringified", func() {
		actualTable := terminal.NewTable([]string{"Name", "Email"})
		actualTable.Add("foo", "foo@email.com")
		var b []byte
		actualBuf := bytes.NewBuffer(b)
		actualTable.PrintTo(actualBuf)

		result, err := MatchCLITable(table).Match(actualBuf.String())
		Expect(result).To(BeTrue())
		Expect(err).ToNot(HaveOccurred())
	})

	It("matches cli table", func() {
		actualTable := terminal.NewTable([]string{"Name", "Email"})
		actualTable.Add("foo", "foo@email.com")
		var b []byte
		actualBuf := bytes.NewBuffer(b)
		actualTable.PrintTo(actualBuf)

		Expect(actualBuf.String()).To(MatchCLITable(table))
	})

	It("does not match cli table", func() {
		actualTable := terminal.NewTable([]string{"Name", "Email"})
		actualTable.Add("foo", "bar@email.com")
		var b []byte
		actualBuf := bytes.NewBuffer(b)
		actualTable.PrintTo(actualBuf)

		Expect(actualBuf.String()).ToNot(MatchCLITable(table))
	})

})
