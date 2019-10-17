package lib

type DisjointSet struct {
	parents []int
}

func NewDisjointSet(size int) DisjointSet {
	parents := make([]int, size)
	for i := 0; i < size; i++ {
		parents[i] = i
	}
	return DisjointSet{parents: parents}
}

func (ds *DisjointSet) Find(i int) int {
	if ds.parents[i] == i {
		return i
	}

	finalParent := i
	for ds.parents[i] != i {
		finalParent = ds.parents[i]
		i = ds.parents[i]
	}
	current := i
	for ds.parents[current] != current {
		ds.parents[current] = finalParent
		current = ds.parents[current]
	}

	return ds.parents[i]
}

func (ds *DisjointSet) Union(i int, j int) {
	ds.parents[ds.Find(i)] = ds.parents[j]
}

func (ds *DisjointSet) Count() int {
	count := 0
	for i := 0; i < len(ds.parents); i++ {
		if ds.parents[i] == i {
			count++
		}
	}
	return count
}
