# mcachesimple
Key/Value Store Caching System.

*This project is developed for personal use - You can use it at your own risk and purpose !!!*

**Example Usage: **

```go
package main

import (
  "fmt"
  "github.com/authapon/mcachesimple"
  "errors"
)

func main() {
  c := mcachesimple.New()
  c.SetWriteFunc(writer)
  c.SetRemoveFunc(remover)
  c.SetReadFunc(reader)

  if data, err := c.Get("a1"); err != nil {
    fmt.Printf("error a1")
  } else {
    fmt.Printf("Got a1 - %s\n", data.(string))
  }

  if data, err := c.Get("a2"); err != nil {
    fmt.Printf("error a2")
  } else {
    fmt.Printf("Got a2 - %s\n", data.(string))
  }

  fmt.Printf("Set a2 - dddddddd\n")
  c.Set("a2","dddddddd")
  if data, err := c.Get("a2"); err != nil {
    fmt.Printf("error a2")
  } else {
    fmt.Printf("Got a2 - %s\n", data.(string))
  }  

  fmt.Printf("Set a3 - eeeeeeee\n")
  c.Set("a3","eeeeeeee")
  if data, err := c.Get("a3"); err != nil {
    fmt.Printf("error a3")
  } else {
    fmt.Printf("Got a3 - %s\n", data.(string))
  } 

  c.Remove("a2")
  if data, err := c.Get("a2"); err != nil {
    fmt.Printf("error a2")
  } else {
    fmt.Printf("Got a2 - %s\n", data.(string))
  } 
 
}

func reader(key string) (interface{}, error) {
  if key == "a1" {
    return "aaaaaaaa", nil
  }

  if key == "a2" {
    return "bbbbbbb", nil
  }

  return nil, errors.New("No data to read!")
}

func writer(key string, value interface{}) {
  if key == "a2" {
    fmt.Printf("Writing a2 = %s\n", value.(string))
    return
  }

  if key == "a3" {
    fmt.Printf("Writing a3 = %s\n", value.(string))
    return
  }

  fmt.Printf("error writing!\n")
}

func remover(key string) {
  fmt.Printf("removing %s\n", key)
}
```

**Output:**

Got a1 - aaaaaaaa  
Got a2 - bbbbbbb  
Set a2 - dddddddd  
Writing a2 = dddddddd  
Got a2 - dddddddd  
Set a3 - eeeeeeee  
Writing a3 = eeeeeeee  
Got a3 - eeeeeeee  
removing a2  
Got a2 - bbbbbbb  
