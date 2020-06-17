package main

// function to remove duplicate value in slices
func fungsiA(slice []string) []string {
	// new slices of map (collection of key-value pairs)
	fungsiMap := make(map[string]struct{})
	// loop through input slice
	for _, v := range slice {
		// set value of slice as fungsiMap map key
		fungsiMap[v] = struct{}{}
	}

	// new slices with length 0 and max length = length of fungsiMap
	fungsiSlice := make([]string, 0, len(fungsiMap))
	// loop through fungsiMap slices
	for v := range fungsiMap {
		// adding element into fungsiSlice with value of fungsiMap map key
		fungsiSlice = append(fungsiSlice, v)
	}

	// return slice fungsiSlice
	return fungsiSlice
}
