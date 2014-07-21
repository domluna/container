package trie

import "testing"

var table1 = map[string]interface{}{
	"on":                2,
	"one":               10,
	"some other string": 55,
}

var table2 = map[string]interface{}{
	"blank1":  2,
	"one1":    10,
	"on1":     10,
	"blank55": 55,
	"o":       0,
}

func TestTrieBasics(t *testing.T) {
	trie := New()
	for k, v := range table1 {
		trie.Add(k, v)
	}

	for k, e := range table1 {
		v, _ := trie.Get(k)
		if v != e {
			t.Errorf("(Should be in Trie, key %v) expected %v, got %v", k, e, v)
		}
	}

	for k, _ := range table2 {
		_, ok := trie.Get(k)
		if ok {
			t.Errorf("(Should not be in Trie, key %v) expected %v, got %v", k, false, ok)
		}
	}
}

func TestTrieSet(t *testing.T) {
	k := "foo"
	trie := New()
	trie.Add(k, 5)
	trie.Add(k, 123)

	v, _ := trie.Get(k)
	if v != 123 {
		t.Errorf("(Should be in Trie, key %v) expected %v, got %v", k, 123, v)
	}
}

func TestTrieRemove(t *testing.T) {
	trie := New()
	trie.Add("foo", 5)
	trie.Add("fool", 123)

	trie.Remove("fool")

	_, ok := trie.Get("fool")
	if ok {
		t.Errorf("(Should not be in Trie, key fool) expected %v, got %t", false, ok)
	}

	v, _ := trie.Get("foo")
	if v != 5 {
		t.Errorf("(Should be in Trie, key foo) expected 5, got %v", v)
	}
}
