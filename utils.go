package main

func contains(sl []string, el string) bool {
    for _, x := range sl {
        if x == el {
            return true
        }
    }
    return false
}

func intersection(sliceA []string, sliceB []string) []string {
    longer := sliceA
    shorter := sliceB
    newSlice := []string{}
    if len(sliceB) > len(sliceA) {
        longer = sliceB
        shorter = sliceA
    }
    for _, w := range shorter {
        if contains(longer, w) {
            newSlice = append(newSlice, w)
        }
    }
    return newSlice
}

func filterOut(sliceA []string, sliceB []string) []string {
    newSlice := []string{}
    for _, w := range sliceA {
        if !contains(sliceB, w) {
            newSlice = append(newSlice, w)
        }
    }
    return newSlice
}
