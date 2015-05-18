package cache

import (
	"reflect"
	"testing"
	"time"
)

func TestAscendGreaterOrEqual(t *testing.T) {
	tree := New()
	tree.InsertNoReplace(Int(4))
	tree.InsertNoReplace(Int(6))
	tree.InsertNoReplace(Int(1))
	tree.InsertNoReplace(Int(3))
	var ary []Item
	tree.AscendGreaterOrEqual(Int(-1), func(i Item, t *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected := []Item{Int(1), Int(3), Int(4), Int(6)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.AscendGreaterOrEqual(Int(3), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected = []Item{Int(3), Int(4), Int(6)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.AscendGreaterOrEqual(Int(2), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected = []Item{Int(3), Int(4), Int(6)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
}

func TestDescendLessOrEqual(t *testing.T) {
	tree := New()
	tree.InsertNoReplace(Int(4))
	tree.InsertNoReplace(Int(6))
	tree.InsertNoReplace(Int(1))
	tree.InsertNoReplace(Int(3))
	var ary []Item
	tree.DescendLessOrEqual(Int(10), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected := []Item{Int(6), Int(4), Int(3), Int(1)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.DescendLessOrEqual(Int(4), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected = []Item{Int(4), Int(3), Int(1)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.DescendLessOrEqual(Int(5), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected = []Item{Int(4), Int(3), Int(1)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
}

func TestAscendGreaterOrEqualForStruct(t *testing.T) {
	tree := New()

	p1 := Person{
		ID:    1,
		Name:  "Person1",
		Birth: time.Now(),
	}
	p2 := Person{
		ID:    2,
		Name:  "Person2",
		Birth: p1.Birth,
	}
	p3 := Person{
		ID:    3,
		Name:  "Person3",
		Birth: p1.Birth,
	}
	p4 := Person{
		ID:    4,
		Name:  p1.Name,
		Birth: p1.Birth,
	}
	p5 := Person{
		ID:    5,
		Name:  "Person5",
		Birth: time.Now(),
	}
	tree.InsertNoReplace(Person(p1))
	tree.InsertNoReplace(Person(p2))
	tree.InsertNoReplace(Person(p3))
	tree.InsertNoReplace(Person(p4))
	tree.InsertNoReplace(Person(p5))

	var ary []Item
	tree.AscendGreaterOrEqual(Person(p4), func(i Item, t *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected := []Item{Person(p1), Person(p4), Person(p2), Person(p3), Person(p5)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}

	ary = nil
	tree.AscendGreaterOrEqual(Person(p3), func(i Item, tree *LLRB) bool {
		ary = append(ary, i)
		return true
	})
	expected = []Item{Person(p3), Person(p5)}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
}
