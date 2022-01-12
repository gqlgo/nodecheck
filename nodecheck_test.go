package nodecheck_test

import (
	"testing"

	"github.com/bannzai/nodecheck"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nodecheck.Analyzer(""), "a")
}

func TestWithSingleExclusion(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nodecheck.Analyzer("Community"), "b")
}

func TestWithMultipleExclusion(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nodecheck.Analyzer("Community,Item"), "c")
}
