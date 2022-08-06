package rb

import (
	"golang.org/x/exp/constraints"
)

type node[TKey constraints.Ordered, TValue any] struct {
	color  Color
	key    TKey
	value  TValue
	parent *node[TKey, TValue]
	left   *node[TKey, TValue]
	right  *node[TKey, TValue]
}

func (nd *node[TKey, TValue]) search(key TKey) (TValue, bool) {
	var zero TValue

	if nd == nil {
		return zero, false
	}

	if nd.key == key {
		return nd.value, true
	}

	if key < nd.key {
		if nd.left == nil {
			return zero, false
		}

		return nd.left.search(key)
	}

	if key > nd.key {
		if nd.right == nil {
			return zero, false
		}

		return nd.right.search(key)
	}

	return zero, false
}

func (nd *node[TKey, TValue]) traverse(visit func(currentKey TKey, currentValue TValue) error) error {
	if nd.left != nil {
		err := nd.left.traverse(visit)

		if err != nil {
			return err
		}
	}

	if nd != nil {
		err := visit(nd.key, nd.value)

		if err != nil {
			return err
		}
	}

	if nd.right != nil {
		err := nd.right.traverse(visit)

		if err != nil {
			return err
		}
	}

	return nil
}
