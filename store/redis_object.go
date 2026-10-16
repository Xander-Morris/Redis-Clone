package store

import (
	"cmp"
	"container/list"
	"maps"
	"slices"
)

type objType int

const (
	typeString objType = iota + 1
	typeList
	typeHash
	typeSet
	typeZset
)

type redisObject struct {
	typ   objType
	value any
}

type sortedSet struct {
	scores map[string]float64 
}

type zEntry struct {
	member string 
	score float64 
}

func (z *sortedSet) add(member string, score float64) bool {
	_, exists := z.scores[member]
	z.scores[member] = score 
	return !exists 
}

func (z *sortedSet) remove(member string) bool {
	_, exists := z.scores[member]
	delete(z.scores, member)
	return exists
}

func (z *sortedSet) score(member string) (float64, bool) {
	s, ok := z.scores[member]
	return s, ok
}

func (z *sortedSet) len() int {
	return len(z.scores)
}

func (z *sortedSet) sorted() []zEntry {
	entries := make([]zEntry, 0, len(z.scores))
	for m, s := range z.scores {
		entries = append(entries, zEntry{member: m, score: s})
	}
	slices.SortFunc(entries, func(a, b zEntry) int {
		if c := cmp.Compare(a.score, b.score); c != 0 {
			return c
		}
		return cmp.Compare(a.member, b.member)
	})
	return entries
}

func (z *sortedSet) rank(member string) (int, bool) {
	if _, ok := z.scores[member]; !ok {
		return 0, false
	}
	for i, e := range z.sorted() {
		if e.member == member {
			return i, true
		}
	}
	return 0, false
}

func (z *sortedSet) rangeByRank(start, stop int) []zEntry {
	entries := z.sorted()
	n := len(entries)

	if start < 0 {
		start += n
	}
	if stop < 0 {
		stop += n
	}
	if start < 0 {
		start = 0
	}
	if stop >= n {
		stop = n - 1
	}
	if start > stop || start >= n {
		return nil
	}

	return entries[start : stop+1]
}

func newStringObject(str string) *redisObject {
	return &redisObject{typ: typeString, value: str}
}

func newListObject(strs ...string) *redisObject {
	lst := list.New()

	for _, s := range strs {
		lst.PushBack(s)
	}

	return &redisObject{typ: typeList, value: lst}
}

func newTypeHashObject(hash map[string]string) *redisObject {
	return &redisObject{typ: typeHash, value: maps.Clone(hash)}
}

func newTypeSetObject(set map[string]struct{}) *redisObject {
	return &redisObject{typ: typeSet, value: maps.Clone(set)}
}

func newTypeZsetObject() *redisObject {
	return &redisObject{
		typ:   typeZset,
		value: &sortedSet{scores: make(map[string]float64)},
	}
}