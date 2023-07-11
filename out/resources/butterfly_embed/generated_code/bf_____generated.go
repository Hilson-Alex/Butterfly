/*

This package represents where the generated go code of the Butterfly
compilation process will be stored before they're compiled down to
machine code by the Go Compiler.

Butterfly uses golang (The Go Programming Language) as intermediate
code for the final compilation process where an executable will be
generated.

By generating golang code, Butterfly can use the Go Compiler as a
backend for the Butterfly Compiler, reducing the effort of making
an experimental language.

So. As the final step of the Butterfly Language Compilation Process,
the runtime code is copied to a temporary folder and the generated code
injected in this package in a way that it can generate a working
executable.


Please do not write any piece of code in this package to not cause
conflicts with any generated code. This file exists for documentation
and place keeping only.

*/

package generated_code
