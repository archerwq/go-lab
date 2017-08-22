// Go programs are organized into packages. A package is a collection of source files in the same directory
// that are compiled together. Functions, types, variables, and constants defined in one source file are
// visible to all other source files within the same package.

// A repository contains one or more modules. A module is a collection of related Go packages that are
// released together. A Go repository typically contains only one module, located at the root of the repository.
// A file named go.mod there declares the module path: the import path prefix for all packages within the module.
// The module contains the packages in the directory containing its go.mod file as well as subdirectories of
// that directory, up to the next subdirectory containing another go.mod file (if any).

// Note that you don't need to publish your code to a remote repository before you can build it. A module can be
// defined locally without belonging to a repository. However, it's a good habit to organize your code as if you
// will publish it someday.

// Each module's path not only serves as an import path prefix for its packages, but also indicates where the go
// command should look to download it. For example, in order to download the module golang.org/x/tools, the go
// command would consult the repository indicated by https://golang.org/x/tools (described more here).

// An import path is a string used to import a package. A package's import path is its module path joined with
// its subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in
// the directory cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard
// library do not have a module path prefix.

// When you run commands like go install, go build, or go run, the go command will automatically download
// the remote module and record its version in your go.mod file.

// Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated
// by the GOPATH environment variable. The downloaded contents for a given version of a module are shared
// among all other modules that require that version, so the go command marks those files and directories
// as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean.

// Commands like go install apply within the context of the module containing the current working directory.
// If the working directory is not within the example.com/user/hello module, go install may fail.

// The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set,
// binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory
// of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of
// the default GOPATH ($HOME/go or %USERPROFILE%\go).

package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
)

func main() {
	fmt.Println(">>>>>>", quote.Hello())
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(math.Pi)
}
