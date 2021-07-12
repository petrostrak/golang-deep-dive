### Binary Tree
A binary tree is a data structure where underneath each node there exists at most two other nodes. This means that a node can be connected to one, two, or no other nodes. The `root of a tree` is the first node of the tree. The `depth of a tree`, which is also called the `height of a tree`, is defined as the longest path from the root to a node, whereas the `depth of a node` is the number of edges from the node to the root of the tree. A leaf is a node with no children.

A tree is considered `balanced` when the longest length from the root node to a leaf is at most one more than the shortest such length. An `unbalanced tree` is a tree that is not balanced. Balancing a tree might be a difficult and slow operation, so it is better to keep your tree balanced from the beginning rather that trying to balance it after you have created it, especially when your tree has a large number of nodes.

### Advantages of binaty trees
You cannot beat a tree when you need to represent hierarchical data. For that eason, trees are extensively used when the compiler of a programming language parses a computer program.

Additionally, trees are `ordered` by design, which means that you do not have to make any special effort to order them. Putting an element into its correct place keeps them ordered. However, deleting an element from a tree is not always trivial becuse of the way that trees are constructed.

If a binary tree is balanced, its search, insert and delete operations take about `log(n)` steps, where `n` is the total number of elements that the tree holds. Additionally, the height of a balanced binary tree is approximately `log2(n)`, which means that a balanced tree with 10000 elements has a height of about 13, and that is remarkably small.

Similarly, the height of a balanced tree with 100000 elements will be about 17, and the height of a balanced tree with 1000000 elements will be about 20. In other words, putting a significantly large number of elements into a balanced binary tree does not change the speed of the tree in an extreme way. Stated differently, you can reach any node of a tree with 1000000 nodes in less than 20 steps.

A major disadvantage of binary trees is that the shape of the tree depends on the order in which its elements were inserted. If the keys of a tree are long and complex, then inserting or searching for an element might be slow due to the large number of comparisons required. Finally, if a tree is not balanced, then the performance of the tree will be upredictable.