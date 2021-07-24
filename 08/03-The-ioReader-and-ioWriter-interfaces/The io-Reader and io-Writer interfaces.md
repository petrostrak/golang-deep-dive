### The io.Reader and io.Writer interfaces
Compliance with the `io.Reader` interface requires the implementation of the `Read()` method, whereas if you want to satisfy the `io.Writer` interface guidelines, you will need to implement the `Write()` method.

### Buffered and unbuffered file input and output
Buffered file input and output happens when there is a buffer for temporarily storing data before reading data or writing data. Thus, instead of reading a file byte to byte, you read many bytes at once. You put the data in a buffer and wait for someone to read it in the desired way. Unbuffered file input and output happens when there is no buffer to temporarily store data before actually reading or writing it.

The next question that you might ask is how to decide when to use buffered and when to use unbuffered file input and output. When dealing with critical data, unbuffered file input and output is generally a better choice because buffered reads might result in out-of-date data and buffered writes might result in data loss when the power of your computer is interrupted.

### The bufio package
As the name suggests, the `bufio` package is about buffered input and output. However, the `bufio` package still uses (internally) the `io.Reader` and `io.Writer` objects, respectively.

### Reading text files
A text file is the most common kind of file that you can find on a UNIX system. In this section, you will be shown how to read text files in three ways: line by line, word by word and character by character.