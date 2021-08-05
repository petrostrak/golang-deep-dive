package main

import (
	"context"
	"fmt"
)

type aKey string

// searchKey retrieves a value from a context and checks whether that value exists
// or not.
func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {

	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")

	searchKey(ctx, myKey)
	searchKey(ctx, aKey("notThere"))

	// In this case, we declare that although we intend to use an operation context, we are not sure about
	// it yet - this is signified by the use of context.TODO() function.
	//
	// Remember that you should never pass a nil context - use the context.TODO() to create a suitable context
	// and remember that the context.TODO() should be used when you are not sure about the Context that you
	// want to use.
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))

}
