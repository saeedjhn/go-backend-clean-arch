## Here’s a comprehensive list of go test commands with explanations and examples, including common scenarios and their use cases:

### Basic Commands:

``` 
go test # Executes all test functions (*_test.go files) in the current directory.

go test -v # Displays detailed output for each test, including names, results, and durations.
go test -v > results.txt # Redirects the test output to a file.
go test -v -failfast # it will stop at the first test that fails, aborting the running test suite/files

go test ./mypackage(relative_path) # Runs tests in the specified package only.
go test ./pkg/numbers/numbers_test

go test -json [package] # Converts the output of the tests to JSON format
go test -json > result.json

go test -run <regex> # Runs only tests whose names match the given regular expression.
go test -run TestFunctionName
go test -run ^Test_Validate.*$
go test -v -count=1 -run="sub-test" # Selectively running subtests
go test -v -count=1 -run="TestOlder"
go test -v -count=1 -run="TestOlder/FirstOlderThanSecond"

go test -count=1 # Forces re-execution of tests, bypassing the cache.

go test ./... # Runs all tests in the current directory and its subdirectories.

go test -timeout <duration> # Specifies the maximum allowed duration for each test.
go test -timeout 30s

go test -list <regex> #  Lists test functions or benchmarks matching the regex without executing them.
go test -list Test 
go test -list Test_Validate

go test -- <custom-flags> # Allows passing custom flags defined in your test files.
go test -- -myFlag=value

go test -race #  Detects race conditions in your code.
```

### Coverage-Related Commands:

```
go test -cover # Displays the percentage of code covered by tests.

go test -coverprofile=<file> # Saves coverage data to a file for detailed analysis.
go test -coverprofile=coverage.out

go tool cover -html=<file> # View Coverage in Browser

# Opens a detailed HTML report of coverage in your default browser. 
go tool cover -html=coverage.out
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Parallel Execution:

```
go test -parallel <n> # Sets the maximum number of tests to run concurrently.
go test -parallel 4

go test -cpu=number # It limits the number of operating system threads that can execute user-level Go code simultaneously
go test -cpu 1
go test -cpu=1,2
go test -cpu=1,2,4
```