# keypresses
This is a lib that provide you checking if any key is pressed

### Installation
Install this package with command `go get github.com/Numenorean/keypresses`

### Finding virtualKeyCodes
Just go to this page https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes

### Usage
# Example of getting key state (even if window inactive)
```go
package main

import "github.com/Numenorean/keypresses"

func main() {
  for true {
    // 0x50 is virtualKeyCode, char "P" in human format
    fmt.Println(keypresses.IsKeyPressed(0x50))
    // Sleeping to prevent 100% CPU usage
    time.Sleep(1 * time.Microsecond)
  }
}
```

# Example of getting key state only if window is activate
```go
package main

import "github.com/Numenorean/keypresses"

func main() {
  for true {
    // "false" argument means that to get key state, window should be active
    // "true" argument means that to get key state, window might not be active. The same as an IsKeyPressed function
    fmt.Println(keypresses.IsKeyPressedGlobal(0x50, false))
    // Sleeping to prevent 100% CPU usage
    time.Sleep(1 * time.Microsecond)
  }
}
```
