package rb

import "golang.org/x/exp/constraints"

type Tree[TKey constraints.Ordered, TValue any] struct {
	root *node[TKey, TValue]
}

func (tree *Tree[TKey, TValue]) rotateLeft(x *node[TKey, TValue]) {
	if x == nil || x.right == nil {
		return
	}

	y := x.right
	x.right = y.left

	if x.right != nil {
		x.right.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (tree *Tree[TKey, TValue]) rotateRight(y *node[TKey, TValue]) {
	if y == nil || y.left == nil {
		return
	}

	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if x.parent == nil {
		tree.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	x.right = y
	y.parent = x
}

func (tree *Tree[TKey, TValue]) insertFixUp(z *node[TKey, TValue]) {
	for z != tree.root && z != tree.root.left && z != tree.root.right && z.color == ColorRed {
		var y *node[TKey, TValue]

		if z.parent != nil && z.parent.parent != nil && z.parent == z.parent.parent.left {
			y = z.parent.parent.right
		} else {
			y = z.parent.parent.left
		}

		if y == nil {
			z = z.parent.parent
		} else if y.color == ColorRed {
			y.color = ColorBlack
			z.parent.color = ColorBlack
			z.parent.parent.color = ColorRed
			z = z.parent.parent
		} else {
			if z.parent == z.parent.parent.left && z == z.parent.left {
				ch := z.parent.color
				z.parent.color = z.parent.parent.color
				z.parent.parent.color = ch

				tree.rotateRight(z.parent.parent)
			}

			if z.parent != nil && z.parent.parent != nil && z.parent == z.parent.parent.left && z == z.parent.right {
				ch := z.color
				z.color = z.parent.parent.color
				z.parent.parent.color = ch

				tree.rotateLeft(z.parent)
				tree.rotateRight(z.parent.parent)
			}

			if z.parent != nil && z.parent.parent != nil && z.parent == z.parent.parent.right && z == z.parent.right {
				ch := z.parent.color
				z.parent.color = z.parent.parent.color
				z.parent.parent.color = ch

				tree.rotateLeft(z.parent.parent)
			}

			if z.parent != nil && z.parent.parent != nil && z.parent == z.parent.parent.right && z == z.parent.left {
				ch := z.color
				z.color = z.parent.parent.color
				z.parent.parent.color = ch

				tree.rotateRight(z.parent)
				tree.rotateLeft(z.parent.parent)
			}
		}
	}

	tree.root.color = ColorBlack
}

func (tree *Tree[TKey, TValue]) insert(key TKey, value TValue) {
	z := &node[TKey, TValue]{
		key:   key,
		value: value,
	}

	if tree.root == nil {
		z.color = ColorBlack
		tree.root = z

		return
	}

	var y *node[TKey, TValue]
	x := tree.root

	for x != nil {
		y = x

		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y

	if z.key > y.key {
		y.right = z
	} else {
		y.left = z
	}

	z.color = ColorRed

	tree.insertFixUp(z)
}

func (tree *Tree[TKey, TValue]) Insert(key TKey, value TValue) {
	tree.insert(key, value)
}

func (tree *Tree[TKey, TValue]) Search(key TKey) (TValue, bool) {
	return tree.root.search(key)
}

func (tree *Tree[TKey, TValue]) Traverse(visit func(currentKey TKey, currentValue TValue) error) error {
	return tree.root.traverse(visit)
}

func NewTree[TKey constraints.Ordered, TValue any]() *Tree[TKey, TValue] {
	return &Tree[TKey, TValue]{}
}
