### Using container/heap
The `container/heap` package implements a heap, which is a tree where the value of each node of the tree is the smallest element in its subtree. Heap does not support numerical values only.

However, in order to implement a heap tree in Go, you will have to develop a way to tell which of two elements is smaller that the other on your own. In such cases, Go uses `interfaces` because they allow you to define such a behavior.

This means that the `container/heap` package is more advanced that the other two packages found in `container`, and that you will have to define some things before being able to use the functionality of the `container/heap` package. Strictly speaking, the `container/heap` package requires that you implement `container/heap.Interface` which is defined as follows:

```
type Interface interface{
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1
}
```

`sort.Interface` requires that you implement the `Len()`, `Less()` and `Swap()` functions, which makes perfect sense because you cannot perform any kind of sorting without being able to swap two elements, being able to calculate a value for the things that you want to sort and being able to tell which element between two elements is bigger that the other based on the value that you calculated previously.