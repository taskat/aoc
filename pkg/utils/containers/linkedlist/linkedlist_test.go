package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ll := New[int]()
	assert.NotNil(t, ll)
	assert.Nil(t, ll.first)
	assert.Nil(t, ll.last)
	assert.Equal(t, 0, ll.length)
}

func TestFromSlice(t *testing.T) {
	type testCase[T any] struct {
		testName string
		values   []T
		expected *LinkedList[T]
	}
	singleElementList := New[int]()
	singleElementList.InsertFirst(1)
	multipleElementsList := New[int]()
	multipleElementsList.InsertFirst(3)
	multipleElementsList.InsertFirst(2)
	multipleElementsList.InsertFirst(1)
	testCases := []testCase[int]{
		{"Empty slice", []int{}, New[int]()},
		{"Single element", []int{1}, singleElementList},
		{"Multiple elements", []int{1, 2, 3}, multipleElementsList},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, FromSlice(tc.values))
		})
	}
}

func TestClear(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Empty list", New[int]()},
		{"Non-empty list", FromSlice([]int{1, 2, 3})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.Clear()
			assert.Nil(t, tc.ll.first)
			assert.Nil(t, tc.ll.last)
			assert.Equal(t, 0, tc.ll.length)
		})
	}
}

func TestForEach(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), []int{}},
		{"Single element", FromSlice([]int{1}), []int{1}},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			values := []int{}
			tc.ll.ForEach(func(node Node[int]) { values = append(values, node.Value()) })
			assert.Equal(t, tc.expected, values)
		})
	}
}

func TestForEach_i(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), []int{}},
		{"Single element", FromSlice([]int{1}), []int{1}},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			values := []int{}
			tc.ll.ForEach_i(func(node Node[int], index int) { values = append(values, node.Value()) })
			assert.Equal(t, tc.expected, values)
		})
	}
}

func TestGet(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		ll            *LinkedList[T]
		index         int
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), 0, 1},
		{"Multiple elements", FromSlice([]int{1, 2, 3, 4}), 1, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedValue, tc.ll.Get(tc.index))
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Empty list", New[int](), 0, 0},
		{"Out of bounds", FromSlice([]int{1}), 1, 0},
		{"Negative index", FromSlice([]int{1}), -1, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.Get(tc.index) })
		})
	}
}

func TestGetNode(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		ll            *LinkedList[T]
		index         int
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), 0, 1},
		{"Multiple elements", FromSlice([]int{1, 2, 3, 4}), 1, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			node := tc.ll.GetNode(tc.index)
			assert.Equal(t, tc.expectedValue, node.Value())
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Empty list", New[int](), 0, 0},
		{"Out of bounds", FromSlice([]int{1}), 1, 0},
		{"Negative index", FromSlice([]int{1}), -1, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.GetNode(tc.index) })
		})
	}
}

func TestInsert(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		value    T
		index    int
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), 1, 0, FromSlice([]int{1})},
		{"Insert at beginning", FromSlice([]int{2}), 1, 0, FromSlice([]int{1, 2})},
		{"Insert at end", FromSlice([]int{1}), 2, 1, FromSlice([]int{1, 2})},
		{"Insert in middle", FromSlice([]int{1, 3}), 2, 1, FromSlice([]int{1, 2, 3})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.Insert(tc.index, tc.value)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Out of bounds", New[int](), 1, 1, FromSlice([]int{1})},
		{"Negative index", New[int](), 1, -1, FromSlice([]int{1})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.Insert(tc.index, tc.value) })
		})
	}
}

func TestInsertFirst(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		value    T
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), 1, FromSlice([]int{1})},
		{"Non-empty list", FromSlice([]int{2}), 1, FromSlice([]int{1, 2})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.InsertFirst(tc.value)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
}

func TestInsertLast(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		value    T
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), 1, FromSlice([]int{1})},
		{"Non-empty list", FromSlice([]int{1}), 2, FromSlice([]int{1, 2})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.InsertLast(tc.value)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
}

func TestLength(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected int
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), 0},
		{"Single element", FromSlice([]int{1}), 1},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), 3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.ll.Length())
		})
	}
}

func TestRemove(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		index    int
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), 0, New[int]()},
		{"Multiple elements - beginning", FromSlice([]int{1, 2}), 0, FromSlice([]int{2})},
		{"Multiple elements - end", FromSlice([]int{1, 2}), 1, FromSlice([]int{1})},
		{"Multiple elements - middle", FromSlice([]int{1, 2, 3}), 1, FromSlice([]int{1, 3})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.Remove(tc.index)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Empty list", New[int](), 0, New[int]()},
		{"Out of bounds", FromSlice([]int{1}), 1, FromSlice([]int{1})},
		{"Negative index", FromSlice([]int{1}), -1, FromSlice([]int{1})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.Remove(tc.index) })
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), New[int]()},
		{"Multiple elements", FromSlice([]int{1, 2}), FromSlice([]int{2})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.RemoveFirst()
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
}

func TestRemoveLast(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), New[int]()},
		{"Multiple elements", FromSlice([]int{1, 2}), FromSlice([]int{1})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.RemoveLast()
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
}

func TestReplace(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		value    T
		index    int
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), 2, 0, FromSlice([]int{2})},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), 4, 1, FromSlice([]int{1, 4, 3})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.Replace(tc.index, tc.value)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Empty list", New[int](), 1, 0, New[int]()},
		{"Out of bounds", FromSlice([]int{1}), 2, 1, FromSlice([]int{1})},
		{"Negative index", FromSlice([]int{1}), 2, -1, FromSlice([]int{1})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.Replace(tc.value, tc.index) })
		})
	}
}

func TestSet(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		value    T
		index    int
		expected *LinkedList[T]
	}
	testCases := []testCase[int]{
		{"Single element", FromSlice([]int{1}), 2, 0, FromSlice([]int{2})},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), 4, 1, FromSlice([]int{1, 4, 3})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.ll.Set(tc.index, tc.value)
			assert.Equal(t, tc.expected, tc.ll)
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Empty list", New[int](), 1, 0, New[int]()},
		{"Out of bounds", FromSlice([]int{1}), 2, 1, FromSlice([]int{1})},
		{"Negative index", FromSlice([]int{1}), 2, -1, FromSlice([]int{1})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { tc.ll.Set(tc.value, tc.index) })
		})
	}
}

func TestLinkedListString(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected string
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), ""},
		{"Single element", FromSlice([]int{1}), "1"},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), "1 -> 2 -> 3"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.ll.String())
		})
	}
}

func TestToSlice(t *testing.T) {
	type testCase[T any] struct {
		testName string
		ll       *LinkedList[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"Empty list", New[int](), []int{}},
		{"Single element", FromSlice([]int{1}), []int{1}},
		{"Multiple elements", FromSlice([]int{1, 2, 3}), []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.ll.ToSlice())
		})
	}
}
