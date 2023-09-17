# MyTime

**MyTime** is a Go package that extends the functionality of the standard `time` package, providing convenient functions for working with date and time in the ISO 8601 format, including milliseconds.

## Installation

To use **MyTime** in your Go project, you can simply run:

```bash
go get github.com/heyitsfranky/MyTime@latest
```

## Usage

Here's a basic example of how to use some of the functionalities of **MyTime**:
```go
package main

import (
    "fmt"
    "github.com/heyitsfranky/MyTime"
)

func main() {
    // Check if a date and time string is correctly formatted.
    input := "2023-09-16T12:34:56.789Z"
    isFormatCorrect := myTime.IsDateTimeFormatCorrect(input)
    fmt.Println("Is Format Correct:", isFormatCorrect)

    // Get the current date and time with milliseconds.
    currentDateTime := myTime.GetCurrentDateTimeWithMilliseconds()
    fmt.Println("Current Date and Time with Milliseconds:", currentDateTime)

    // And many more...
}
```

## License

This package is distributed under the MIT License.
Feel free to contribute or report issues on GitHub.

Happy coding!
