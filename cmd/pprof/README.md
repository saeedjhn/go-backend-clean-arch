requirement:
-> sudo apt install graphviz

Run /cmd/pprof/main.go
> make pprof

Request
Show Types of profiles available:
-> http://localhost:PORT/debug/pprof

Show goroutine usage for project with graphviz
> make pprof
> make tool/pprof/goroutine
OR:
> curl http://localhost:PORT/debug/pprof/goroutine --output fileName.ExtensionName
> go tool pprof -http=:PORT2 fileName.ExtensionName

Serving web UI on http://localhost:PORT2/
go to http://localhost:PORT2/ui/
