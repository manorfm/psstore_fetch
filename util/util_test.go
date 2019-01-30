package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
 
func TestUpdatePathPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20"
	UpdatePathPagination(&path, 21, 50)
	
	assert.Equal(t, path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=21&size=50", "Error while trying to replae path")
}

func TestAddPathPaginationWithOutPaginationOnPath(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES"
	UpdatePathPagination(&path, 0, 20)

	assert.Equal(t, path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20", "Error while trying to replae path")
}
func TestAddPathSizeOnPaginationData(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0"
	UpdatePathPagination(&path, 0, 20)

	assert.Equal(t, path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20", "Error while trying to replae path")
}
func TestAddPathStartOnPaginationData(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?size=0"
	UpdatePathPagination(&path, 0, 20)

	assert.Equal(t, path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?size=20&start=0", "Error while trying to replae path")
}

func TestInitialPathWithStartPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=10"
	start := FindStartPagination(path)

	assert.Equal(t, 10, start, "Pagination parameter not found")
}

func TestInitialPathWithoutStartPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES"
	start := FindStartPagination(path)

	assert.Equal(t, 0, start, "Pagination parameter not found")
}