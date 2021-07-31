### The select keyword
The `select` statement in Go looks like a `switch` statement but for channels. In practice, this means that `select` allows a goroutine to wait on multiple communication operations. Therefore, the main benefit that you receive from `select` is that it gives you the power to work with multiple channels using a single `select` block. As a consequence, you can have nonbloacking operations on channels, provided that you have appropriate `select` blocks.

The biggest problem when using multiple channels and the `select` keyword is `deadlocks`.