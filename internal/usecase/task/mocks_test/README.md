Installation
Once you have installed Go, install the mockgen tool.

Note: If you have not done so already be sure to add $GOPATH/bin to your PATH.

To get the latest released version use:

Go version < 1.16
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
Go 1.16+
go install github.com/golang/mock/mockgen@v1.6.0
If you use mockgen in your CI pipeline, it may be more appropriate to fixate on a specific mockgen version. You should try to keep the library in sync with the version of mockgen used to generate your mocks.

Running mockgen
mockgen has two modes of operation: source and reflect.

Source mode
Source mode generates mock interfaces from a source file. It is enabled by using the -source flag. Other flags that may be useful in this mode are -imports and -aux_files.

Example:

mockgen -source=foo.go [other options]

*********************************************************************************
run mockgen for task service:
> taskservice git:(master) ✗ mockgen -source=usecase.go -destination=mocks/usecase_mock.go -package=mocks_test
notice: Change package name: mocks_test -> mockstest
 
> run service_mock_test.go
*********************************************************************************
