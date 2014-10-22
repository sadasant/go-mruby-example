go-mruby-example
================

This is an example repository that could be useful to teach the
following topics:

-   How to use mruby internals.
-   How to compile Ruby code using mruby internals.
-   How to run Ruby code using mruby internals.
-   How to use cgo.
-   How to build a Go application using mruby internals.
-   How to compile and run ruby code from Go using mruby internals.
-   How to add mrbgems to the mruby build.
-   How to test mrbgems from Go.

The repository structure is as simple as possible and the
abstractions used are minimal to focus on displaying the process used
in each method. The methods are:

-   `func Compile(code string) ([]byte, error) {`: Compiles some
    provided source code into a resulting `[]byte`, we also catch
    compilation errors.
-   `func RunSource(code string) error {` Runs a given source code,
    catching and returning encountered errors.
-   `func RunBytecode(bin []byte) error {` Runs a given bytecode,
    catching and returning encountered errors.

Instead of making it an executable application, this repository only
provides some go tests which cover the functionalities of the methods
mentioned above.

To build `mruby` and run the tests afterwards, you can run `make`.

Other useful repositories:

-   <https://github.com/mitchellh/go-mruby>
-   <https://github.com/olivere/mruby-go>
-   <https://github.com/mattn/go-mruby>
