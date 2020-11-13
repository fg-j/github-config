package internal_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestInternal(t *testing.T) {
	suite := spec.New("scripts/first-contact/internal", spec.Report(report.Terminal{}))
	suite("TestIssue", testIssue)
	suite("TestAPIClient", testAPIClient)
	suite.Run(t)
}
