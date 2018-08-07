package solution

import (
	"testing"
)

func TestRandomizedCollection(t *testing.T) {
	func() {
		collection := Constructor()

		collection.Insert(1)
		collection.Insert(1)
		collection.Insert(2)
		collection.GetRandom()
		collection.Remove(1)
		collection.GetRandom()
	}()

	func() {
		collection := Constructor()

		collection.Insert(1)
		collection.Remove(1)
		collection.Insert(1)
	}()

	func() {
		collection := Constructor()

		collection.Insert(10)
		collection.Insert(10)
		collection.Insert(20)
		collection.Insert(20)
		collection.Insert(30)
		collection.Insert(30)
		collection.Remove(10)
		collection.Remove(10)
		collection.Remove(30)
		collection.Remove(30)
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
		collection.GetRandom()
	}()

	func() {
		collection := Constructor()

		collection.Insert(9)
		collection.Insert(9)
		collection.Insert(1)
		collection.Insert(1)
		collection.Insert(2)
		collection.Insert(1)
		collection.Remove(2)
		collection.Remove(1)
		collection.Remove(1)
		collection.Insert(9)
		collection.Remove(1)
	}()
}
