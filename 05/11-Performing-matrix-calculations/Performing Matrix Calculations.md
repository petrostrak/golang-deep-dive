### Performing Matrix Calculations
A `matrix` is an array with two dimensions. The easiest way to represent a matrix in Go is using a slice. However, if you know the diensions of your array in advance, an array will also do the job just fine. If both dimensions of a matrix are the same, then the matrix is called a `square matrix`.

There are some rules that can tell you whether you can perform a calculation between two matrices or not. The rules are the following:
* In order to `add` or `subtract` two matrices, they should have exactly the same dimentions.
* In order to `multiply` matrix `A` with matrix `B`, the number of columns of matrix `A` should be equal to the number of rows of matrix `B`. Otherwise, the multiplication of matrices `A` and `B` is impossible.
* In order to `divide` matrix `A` with matrix `B`, two conditions must be met. Firstly, you will need to be able to calculate the `inverse` of matrix `B` and secondly, you should be able to `multiply` matrix `A` with the inverse of matrix `B` according to the previous rule. Only square matrices can have an `inverse`.