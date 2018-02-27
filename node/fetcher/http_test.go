package fetcher

import "testing"

func TestImplementsFetcherInterface(t *testing.T) {
	var fetcher Fetcher
	fetcher = &HTTPFetcher{}

	if fetcher == nil {
		t.Error("Pointer to Fetcher is nil")
		t.Fail()
	}
}
