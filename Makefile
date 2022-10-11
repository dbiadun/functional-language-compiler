make_compiler:
	go build -o ./build/compiler

.PHONY : functional_tests

functional_tests:
	go build -o ./tests/test ./tests
	./tests/test functional
