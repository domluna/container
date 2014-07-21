package lru

import "testing"

func TestCacheGet(t *testing.T) {
	c := New(10)
	c.Add("Hello", 1)
	c.Add("blah", 2)

	v, _ := c.Get("blah")
	if v != 2 {
		t.Error("expected 2, got %v", v)
	}

	_, ok := c.Get("not in there")
	if ok {
		t.Error("expected false, got %t", ok)
	}
}

func TestCacheEvict(t *testing.T) {
	c := New(2)
	c.Add("first", 1)
	c.Add("second", 2)
	c.Add("third", 42)

	_, ok := c.Get("first")
	if ok {
		t.Error("expected false, got %t", ok)
	}

	v, _ := c.Get("second")
	if v != 2 {
		t.Error("expected 2, got %v", v)
	}

	v, _ = c.Get("third")
	if v != 42 {
		t.Error("expected 42, got %v", v)
	}

	if c.Len() != 2 {
		t.Error("expected Len to be 2, got %v", c.Len())
	}
}
