package util

import "testing" 
 
func TestReplacePathPagination(t *testing.T) {
	oldPath := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20"
	path := ReplacePathPagination(oldPath, 21, 50)
	
	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=21&size=50", "Error while trying to replae path", t)
}

// func TestAddPathPagination(t *testing.T) {
// 	oldPath := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES"
// 	path := ReplacePathPagination(oldPath, 0, 20)
	
// 	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20", "Error while trying to replae path", t)
// }

func assert(value, compare string, message string, t *testing.T) {
	if value != compare {
		t.Errorf("%s, got: %s, want: %s.", message, value, compare)
	}
}