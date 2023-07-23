# Go Bindings for LibVersion C Library

This project provides a Go programming language wrapper for the LibVersion C library. The LibVersion C library is a powerful tool for working with version numbers and handling version-related tasks in C applications. By creating a Go wrapper, we aim to enable Go developers to utilize the functionality of LibVersion within their Go programs seamlessly.

## Installation

Before using this Go wrapper, you need to ensure that the LibVersion C library is installed on your system. Instructions for installing the LibVersion C library can be found [here](https://github.com/repology/libversion).

Make sure to have LibVersion installed in one of the following locations:
 - `.` (current directory)
 - `./lib`
 - `/usr/local/lib`
 - `/usr/lib`
 - `/lib`
 - `/usr/local/lib64`
 - `/usr/lib64`
 - `/lib64`
 - `~/.local/lib`
 - `~/.local/lib64`

To install the Go wrapper, you can use `go get`:

```bash
go get github.com/saenai255/golibversion
```

## Usage

Import the `golibversion` package into your Go code:

```go
import "github.com/saenai255/golibversion"
```

Now, you can use various functions provided by the LibVersion C library through the Go wrapper.

```go
version1 := "1.2.3"
version2 := "1.3.0"

// Compare versions
result := golibversion.Compare(version1, version2)
if result == -1 {
    fmt.Printf("%s is older than %s\n", version1, version2)
} else if result == 1 {
    fmt.Printf("%s is newer than %s\n", version1, version2)
} else {
    fmt.Printf("%s and %s are the same version\n", version1, version2)
}

// There are also functions for comparing versions with different flags,
// such as: CompareWithFlags

// Perform other version-related operations using the available functions in the wrapper.
// See the documentation or source code for a full list of available functions.
```

## Contributions

Contributions to this project are welcome. If you find any issues or want to add new features, feel free to open a pull request or an issue on the [GitHub repository](https://github.com/saenai255/golibversion).

## License

This project is licensed under the [MIT License](https://github.com/saenai255/golibversion/blob/main/LICENSE). Please review the license terms before using this wrapper.