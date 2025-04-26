package collections

type Collection[T any] interface {
	// Size returns the number of elements in this collection.
	// If this collection contains more than math.MaxInt elements, returns Integer.MAX_VALUE.
	// Returns: the number of elements in this collection
	Size() int

	// IsEmpty returns true if this collection contains no elements.
	// Returns: true if this collection contains no elements
	IsEmpty() bool

	// Contains returns true if this collection contains the specified element.
	// More formally, returns true if and only if this collection contains at least one element e such that (item == e).
	//
	// Params: item – element whose presence in this collection is to be tested
	// Returns: true if this collection contains the specified element
	Contains(item T) bool

	// ToSlice returns a slice containing all the elements in this collection.
	// The returned slice will be "safe" in that no references to it are maintained by this collection.
	// The caller is thus free to modify the returned array.
	// This method acts as bridge between slice-based and collection-based APIs.
	// Returns: a slice containing all the elements in this collection
	ToSlice() []T

	// Add Ensures that this collection contains the specified element (optional operation).
	// Returns true if this collection changed as a result of the call.
	// Returns false if this collection does not permit duplicates and already contains the specified element.
	// Params: item – element whose presence in this collection is to be ensured
	//Returns: true if this collection changed as a result of the call
	Add(item T) bool

	// Remove removes a single instance of the specified element from this collection, if it is present (optional operation).
	// More formally, removes an element e such that (item == e), if this collection contains one or more such elements.
	// Params: item – element to be removed from this collection, if present
	// Returns: true if an element was removed as a result of this call
	Remove(item T) bool

	// ContainsAll returns true if this collection contains all the elements in the specified collection.
	// Params: c – collection to be checked for containment in this collection
	// Returns: true if this collection contains all the elements in the specified collection
	ContainsAll(c Collection[T]) bool

	// AddAll adds all the elements in the specified collection to this collection (optional operation).
	// The behavior of this operation is undefined if the specified collection is modified while the operation is in progress.
	// Params: c – collection containing elements to be added to this collection
	// Returns: true if this collection changed as a result of the call
	AddAll(c Collection[T]) bool

	// RemoveAll removes all of this collection's elements that are also contained in the specified collection (optional operation).
	// After this call returns, this collection will contain no elements in common with the specified collection.
	// Params: c – collection containing elements to be removed from this collection
	// Returns: true if this collection changed as a result of the call
	RemoveAll(c Collection[T]) bool

	// RetainAll retains only the elements in this collection that are contained in the specified collection (optional operation).
	// In other words, removes from this collection all of its elements that are not contained in the specified collection.
	// Params: c – collection containing elements to be retained in this collection
	// Returns: true if this collection changed as a result of the call
	RetainAll(c Collection[T]) bool

	// Clear removes all the elements from this collection (optional operation).
	// The collection will be empty after this method returns.
	Clear()

	// Equals compares the specified Collection with this Collection for equality.
	// Params: c – Collection to be compared for equality with this Collection
	// Returns: true if the specified object is equal to this collection
	Equals(c Collection[T]) bool
}

type List[T any] interface {
	Collection[T]
	// Get returns the element at the specified position in this list.
	// Params: index – index of the element to return
	// Returns: the element at the specified position in this list
	Get(index int) T

	// Set replaces the element at the specified position in this list with the specified element (optional operation).
	// Params:
	//		index – index of the element to replace
	//		item – element to be stored at the specified position
	// Returns: the element previously at the specified position
	Set(index int, item T) T

	// InsertAt inserts the specified element at the specified position in this list (optional operation).
	// Shifts the element currently at that position (if any) and any subsequent elements to the right (adds one to their indices).
	// Params:
	//		index – index at which the specified element is to be inserted
	//		item – element to be inserted
	InsertAt(index int, item T)

	// RemoveAt removes the element at the specified position in this list (optional operation).
	// Shifts any subsequent elements to the left (subtracts one from their indices). Returns the element that was removed from the list.
	// Params: index – the index of the element to be removed
	// Returns: the element previously at the specified position
	RemoveAt(index int) T

	// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
	// More formally, returns the lowest index i such that (item == Get(i)), or -1 if there is no such index.
	// Params: item – element to search for
	// Returns: the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element
	IndexOf(item T) int

	// LastIndexOf returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
	// More formally, returns the highest index i such that (item == Get(i)), or -1 if there is no such index.
	// Params: item – element to search for
	// Returns: the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element
	LastIndexOf(item T) int

	// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
	// The returned list is backed by this list, so non-structural changes in the returned list are reflected in this list, and vice versa.
	// Params:
	//		fromIndex – low endpoint (inclusive) of the subList
	//		toIndex – high endpoint (exclusive) of the subList
	// Returns: a view of the specified range within this list
	SubList(fromIndex, toIndex int) List[T]
}
