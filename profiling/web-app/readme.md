#### Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool
- go build -o myserver main.go
- go tool pprof -seconds 30 myserver http://localhost:8989/debug/pprof/profile
- curl localhost:8080

### profiling via benchmark - 
- go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
- go tool pprof profile.out

### profiling via test (same as benchmark) - 
- go test -benchmem -memprofile memprofile.out -cpuprofile profile.out