<!--

go is a compiled language.

source code is converted into (native machine) code.

# essential go commands:

   # compile and run
   go run file.go

   # compile once, save executable somewhere outside the project folder.
   go build file.go

# install some go script
go install gopl.io/ch1/helloworld@latest
go: downloading gopl.io v0.0.0-20211004154805-1ae3ec64947b

# it goes to: $GOPATH/bin

# ==============
# missing or unused imports -- causes: compile error.


-->

```text

# packages in go:

a package should be in a directory
even/
   even.go
   math.go

# and all the files inside that directory must have the same package name.

# with exception of test files, which they can have different package names

/tests
   fuck_test.go
   package fuck_test

# any package level name, whether it is a variable or a func or const , it Cap (exported) , if (lowercase) private.


#==============


The Go toolchain has a rule: Import paths that do not start with . or .. are first looked up in the standard library and GOPATH, not in the current directory . So even/ is treated as a "package name" to search for, rather than "the folder even in the current directory". Adding ./, i.e., ./even/, clearly tells Go "this is a local filesystem path" .

In addition, your previous go test (without arguments) only tested the current directory (the package where main.go is located), so it showed learn [no test files]. To run tests for a specific package, you must specify its path.

#=============

Method	Purpose
t.Fail()	Marks the test as failed but continues execution.
t.FailNow()	Marks the test as failed and stops the current test.
t.Log()	Records a message in the test log.
t.Fatal()	Logs a message and immediately fails the current test.

One important distinction: FailNow stops the current test function, not the entire test suite. Go can continue running other tests.

```
