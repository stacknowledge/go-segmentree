package segment

// Tree is the structure that holds tree data,
// lazy propagation helper and initial tree elements size
type Tree struct {
	nodes     []int // nodes array
	lazyNodes []int // auxiliar nodes to lazy propagation
	size      int   // tree inital elements
}

// NewTree constructs a segmentation tree based on given integer elements
func NewTree(elements []int) *Tree {
	tree := Tree{
		make([]int, len(elements)*4), // max nodes it can generate per given size
		make([]int, len(elements)*4), // max nodes it can generate per given size
		len(elements) - 1,            // initial tree elements, 0 is considered
	}

	tree.compose(1, 0, tree.size, elements)

	return &tree
}

// Query is an public interface to query the segment tree
// given a left and right range
func (t *Tree) Query(left, right int) int {
	return t.search(1, 0, t.size, left, right)
}

// Update is an public interface to update the segment tree
// given a value to update between left and right range
func (t *Tree) Update(left, right, value int) {
	t.apply(1, 0, t.size, left, right, value)
}

func (t *Tree) compose(node, start, end int, elements []int) {
	if start > end {
		return // out of range control
	}

	if start == end {
		t.nodes[node] = elements[start] // leaf node found
		return
	}

	root := (start + end) / 2

	t.compose(node*2, start, root, elements)            // compose left node
	t.compose(node*2+1, 1+root, end, elements)          // compose right node
	t.nodes[node] = t.nodes[node*2] + t.nodes[node*2+1] // compose root node with the result of the left and right nodes
}

func (t *Tree) search(node, start, end, left, right int) int {
	if start > end || start > right || end < left {
		return 0 // out of range control
	}

	if t.lazyNodes[node] != 0 {
		t.updateLazyNodes(node, start, end) // while searching if lazy node found, update it right away
	}

	if start >= left && end <= right {
		return t.nodes[node] // found result in range of [left, right]
	}

	root := (start + end) / 2
	leftChildQuery := t.search(node*2, start, root, left, right)    // searches on left node
	rightChildQuery := t.search(1+node*2, 1+root, end, left, right) // searches on right node
	return leftChildQuery + rightChildQuery                         // returns the sum of the left and right search query
}

func (t *Tree) apply(node, start, end, left, right, value int) {
	if t.lazyNodes[node] != 0 { // if lazy node found, update it right away
		t.updateLazyNodes(node, start, end)
	}

	if start > end || start > right || end < left {
		return // out of range control
	}

	if start >= left && end <= right {
		t.nodes[node] += value

		if start != end { // not leaf node
			t.lazyNodes[node*2] += value   // mark left node as lazy
			t.lazyNodes[node*2+1] += value // mark right node as lazy
		}

		return
	}

	root := (start + end) / 2
	t.apply(node*2, start, root, left, right, value)    // update left node
	t.apply(1+node*2, 1+root, end, left, right, value)  // update right node
	t.nodes[node] = t.nodes[node*2] + t.nodes[node*2+1] // update root node with sum of left and right nodes
}

func (t *Tree) updateLazyNodes(node, start, end int) {
	t.nodes[node] += t.lazyNodes[node] //update node with the lazy value

	if start != end { // not leaf node
		t.lazyNodes[node*2] += t.lazyNodes[node]   // mark left node as lazy
		t.lazyNodes[node*2+1] += t.lazyNodes[node] // mark right node as lazy
	}

	t.lazyNodes[node] = 0 // mark it as handled
}
