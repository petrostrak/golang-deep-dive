#include <stdio.h>
#include "usedByC.h"

// gcc -o willUseGo willUseGo.c ./usedByC.o
int main(int argc, char **argv) {
    GoInt x = 11;
    GoInt y = 24;

    printf("About to call a Go function!\n");
    PrintMessage();

    // The p variable is needed for getting an integet value from a Go function...
    GoInt p = Multiply(x, y);

    // ...which is converted to a C integer using the (int) p notation
    printf("Product: %d\n", (int)p);
    printf("It worked!\n");
    return 0;
}