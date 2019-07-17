package cuckoofilter

import (
	"hash"
	"math/rand"
)

type CFilter struct {
	hashfn  hash.Hash // Hash function used for fingerprinting
	buckets []bucket  // Buckets where fingerprints are stored
	count   uint      // Total number of elements currently in the Filter

	bSize  uint8 // Bucket size
	fpSize uint8 // Fingerprint size
	size   uint  // Number of buckets in the filter
	kicks  uint  // Maximum number of times we kick down items from buckets
}

// New returns a new CFilter object. It's Insert, Lookup, Delete and
// Size behave as their names suggest.
// Takes zero or more of the following option functions and applies them in
// order to the Filter:
//      - cuckoofilter.Size(uint) sets the number of buckets in the filter
//      - cuckoofilter.BucketSize(uint8) sets the size of each bucket
//      - cuckoofilter.FingerprintSize(uint8) sets the size of the fingerprint
//      - cuckoofilter.MaximumKicks(uint) sets the maximum number of bucket kicks
//      - cuckoofilter.HashFn(hash.Hash) sets the fingerprinting hashing function
func NewCuckooFilter(opts ...option) *CFilter {
	cf := new(CFilter)
	for _, opt := range opts {
		opt(cf)
	}
	configure(cf)

	cf.buckets = make([]bucket, cf.size, cf.size)
	for i := range cf.buckets {
		cf.buckets[i] = make([]fingerprint, cf.bSize, cf.bSize)
	}

	return cf
}

// Insert adds an element (in byte-array form) to the Cuckoo filter,
// returns true if successful and false otherwise.
func (f *CFilter) Insert(item []byte) bool {
	fp := fprint(item, f.fpSize, f.hashfn)
	j := hashfp(item) % f.size
	k := (j ^ hashfp(fp)) % f.size

	if f.buckets[j].insert(fp) || f.buckets[k].insert(fp) {
		f.count++
		return true
	}

	i := [2]uint{j, k}[rand.Intn(2)]
	for n := uint(0); n < f.kicks; n++ {
		fp = f.buckets[i].swap(fp)
		i = (i ^ hashfp(fp)) % f.size

		if f.buckets[i].insert(fp) {
			f.count++
			return true
		}
	}

	return false
}

// Lookup checks if an element (in byte-array form) exists in the Cuckoo
// Filter, returns true if found and false otherwise.
func (f *CFilter) Lookup(item []byte) bool {
	fp := fprint(item, f.fpSize, f.hashfn)
	j := hashfp(item) % f.size
	k := (j ^ hashfp(fp)) % f.size

	return f.buckets[j].lookup(fp) || f.buckets[k].lookup(fp)
}

// Delete removes an element (in byte-array form) from the Cuckoo Filter,
// returns true if element existed prior and false otherwise.
func (f *CFilter) Delete(item []byte) bool {
	fp := fprint(item, f.fpSize, f.hashfn)
	j := hashfp(item) % f.size
	k := (j ^ hashfp(fp)) % f.size

	if f.buckets[j].remove(fp) || f.buckets[k].remove(fp) {
		f.count--
		return true
	}

	return false
}

// Count returns the total number of elements currently in the Cuckoo Filter.
func (f *CFilter) Count() uint {
	return f.count
}
