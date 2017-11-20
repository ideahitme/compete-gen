## Compete-gen

is a small utility to quickly generate code for programming competitions. Right now supports: 
- Cpp 
- Go

It creates two files (code.go/code.cpp and input) in the current directory. Input is piped to code.cpp/code.go
and should be read from stdin in the code.go/code.cpp (the boilerplate is already included)


### How to install

1. Have Go environment set up 
2. Clone this repository
3. Run `make`
4. If `$GOPATH/bin` is in your path, then `gen` binary should be available. 

### How to run

The utility can be run with:

1. `gen go` - creates code.go and input files in the current directory, if `input` file doesn't exist yet. 
Otherwise it will compile and execute the code with: `cat input | go run code.go`
2. `gen cpp` - creates code.cpp and input files in the current directory, if `input` file doesn't exist yet.
Otherwise it will compile and execute the code with: `g++ -pipe -std=c++11 code.cpp -o code && ./code < input`

