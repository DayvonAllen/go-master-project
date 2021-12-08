## Test Code
- `go test <file_path>`
- `go test -v <file_path>` - verbose mode
- `go test -cover <file_path>` - for code coverage
- `go test -coverprofile=coverage.out <path> && go tool cover -html=coverage.out` - show coverage in web page
- `go build -o <path> <desired-copied-path>` - builds the command line tool app and copies to desired location 
- `./...` - from root level will include all subdirectories
---