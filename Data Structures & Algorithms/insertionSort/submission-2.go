// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
    res := [][]Pair{}
    if len(pairs) == 0 {
        return res
    }

    n := len(pairs)
    snapshot := make([]Pair, n)
    copy(snapshot, pairs)
    res = append(res, snapshot)

    for i := 1; i < n; i++ {
        key := pairs[i]
        j := i - 1

        for j >= 0 && pairs[j].Key > key.Key {
            pairs[j+1] = pairs[j]
            j--
        }
        pairs[j+1] = key

        snapshot := make([]Pair, n)
        copy(snapshot, pairs)
        res = append(res, snapshot)
    }

    return res
}
