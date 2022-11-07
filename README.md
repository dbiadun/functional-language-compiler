# Functional language compiler

Compiler compiling a subset of Haskell for microcontrollers.

## Functionality

The compiler can compile slightly modified subset of Haskell for computers and microcontrollers.

For now it supports two targets (Linux and Adafruit Feather Sense).

## Project structure

- build/ - directory containing a compiler executable
- compiler/ - compiler code
- gmachinerun/ - runtime
- haskellPredefined/ - predefined Haskell types and functions
- parser/ - parser generated by antlr
- tests/ - compiler tests
  - functionalTests/
  - performanceTests/
  - libs/ - code used by multiple tests
- go.mod - go module config file
- Language.g4 - language grammar
- main.go - main compiler file
- Makefile - project makefile
- README.md - this file

## Dependencies

- Go 1.17
- TinyGo 0.25

## Compilation and execution

To build the compiler simply use `make` in the root directory.
It will compile the project and create `build/compiler` executable.

If you want to modify the language grammar you need to download `antlr 4.10.1` and run it in the project directory
on the modified Language.g4 file with the following flags: `-visitor -Dlanguage=Go -o parser`.
After that you can normally run make to build the compiler.

`build/compiler` will have two commands:
- `build/compiler build <path>` compiles a program at the location indicated by `<path>` and creates an executable.
- `build/compiler flash <path>` compiles a program at the location indicated by `<path>` and flashes it to the microcontroller.

Both provide some flags:
- `-target <target>` - tells the compiler for which target it should compile the code.
    The options are `linux` and `feather`. For the `build` command the default flag is set to `linux`.
    For the `flash` command the target must be provided.
- `-size` - tells the compiler to print the memory information of the compiled program.
    It prints, for example, the flash memory occupied by compiled program and the amount of ram reserved for stack.
- `-o <output_file>` - tells the compiler where to put the executable.
    The flag is supported only by the `build` command. The default value of the flag is `a`.

## Running tests

To compile and run functional tests use `make functional_tests`.

To compile and run performance tests use `make performance_tests`.

In both cases you can also choose a specific test to run by adding `test=<test_name>` to the command.
`<test_name>` is just the name of the test (without path and filename extension).
For example `make functional_tests test=fibonacci`.



