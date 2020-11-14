package internal_test

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/github-config/scripts/first-contact/internal"
	"github.com/paketo-buildpacks/github-config/scripts/first-contact/internal/fakes"
	"github.com/sclevine/spec"
)

func testRepository(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect
	var repo Repository
	var apiClient = &fakes.Client{}
	var clock = &fakes.Clock{}

	repo = Repository{
		Name: "example-org/example-repo",
	}
	context("GetRecentIssues", func() {
		it.Before(func() {
			clock.NowCall.Returns.Time = time.Date(2001, time.January, 1, 20, 20, 20, 0, time.UTC).Add(30 * 24 * time.Hour)
		})
		it("returns the issues from the repo", func() {
			_, err := repo.GetRecentIssues(apiClient, clock)
			Expect(err).NotTo(HaveOccurred())
			Expect(apiClient.GetCall.Receives.Path).To(Equal("/repos/example-org/example-repo/issues"))
			Expect(apiClient.GetCall.Receives.Params).To(ContainElement("per_page=100"))
			Expect(apiClient.GetCall.Receives.Params).To(ContainElement("since=2001-01-01T20:20:20Z"))
		})
	})
}
