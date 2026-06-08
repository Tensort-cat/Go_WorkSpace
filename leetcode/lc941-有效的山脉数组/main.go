package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var i, j int

	asc := false
	for i = 0; i < len(arr)-1 && arr[i] < arr[i+1]; i++ {
		asc = true
	}

	desc := false
	for j = i + 1; j < len(arr) && arr[j] < arr[j-1]; j++ {
		desc = true
	}

	return j == len(arr) && asc && desc
}
