package bosh_test

import (
	"github.com/pivotal-cf/p-mysql-manifest-validation/bosh"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manifest", func() {

	var (
		manifest *bosh.Manifest
	)

	Describe("JobNamed", func() {
		BeforeEach(func() {
			job := bosh.NewJob("existentJob-partition-random-guid")
			manifest = &bosh.Manifest{
				Jobs: []*bosh.Job{job},
			}
		})

		It("returns a Job matching the given name", func() {
			expectedJob := manifest.JobNamed("existentJob")
			Expect(expectedJob.Name()).To(HavePrefix("existentJob"))

		})

		It("panics when no match is found", func() {
			Expect(func() { manifest.JobNamed("nonExistentJob") }).To(Panic())
		})
	})

})