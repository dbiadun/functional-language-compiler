make_compiler:
	go build -o ./build/compiler

.PHONY : functional_tests performance_tests

functional_tests:
	go build -o ./tests/test ./tests
	./tests/test functional $(test)

performance_tests:
	go build -o ./tests/test ./tests
	./tests/test performance $(test)
