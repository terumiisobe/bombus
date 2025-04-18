package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFetchColmeias_EmptyDataset_ShouldExpectEmptyList(t *testing.T) {
	result := FetchColmeias()
	expected := []
	assert.Equal()	
}
