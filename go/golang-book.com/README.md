# Go

## Resource

[golang-book.com](https://golang-book.com/books/intro)

## CLI

1. `godoc <package> [function]`

   Function: The name says it all.

2. `cd /where/to/package && go install`

   Function: Compile our package and create a linkable object file.

3. `godoc /where/to/package FunctionName`

   Function: Generate documentation for our package (with a specific function).

4. `godoc -http=":6060"`

   Function: Create a local web documentation. Accessible via http://localhost:6060/pkg/.
