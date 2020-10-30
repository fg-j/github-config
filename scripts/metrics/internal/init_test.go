package internal_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestInternal(t *testing.T) {
	suite := spec.New("scripts/metrics", spec.Report(report.Terminal{}))
	suite("PullRequest", testPullRequest)
	// suite("Repository", testRepository)
	suite.Run(t)
}
