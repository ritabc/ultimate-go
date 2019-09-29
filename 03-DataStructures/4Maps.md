# Maps
* Maps are not usable in their zero value state
```go
type user struct {
    name string
    surname string
}

users := make(map[string]user)

users["roy"] = user{"Rob", "Roy"}

for key, value := range users {
    fmt.Println(key, value)
}
```
* Value semantics are used when assiging values into map 
* When iterating over map, copy of key and value are made
* Iterating over map produces random order
## Construct map with literal
* this is also possible
```go
users := map[string]user{
    "Roy": {"Rob", "Roy"},
    "Ford": {"Henry", "Ford"},
}
```

## delete() builtin function
* syntax is delete(map, key)
* deleting non-existent key doesn't error
* if using map as a cache, good practice to remove / clean at some point

## Reading with found variable vs. without
### With Found - Should always use
    1. First, do: `delete(users, "Roy")`
    1. Next, do `u, found := users["Roy"]`
    1. `found` will be `false`
### Without Found
    1. First, do: `delete(users, "Roy")`
    1. Next, do `u := users["Roy"]`
    1. `u` will now be equal to it's zero value

## Map key restrictions
* Keys must be comparable (or be able to be used in if statement) 
* Cannot be slice, for instance

## Sort Pkg
* a new pkg called "sort" 
```go
import (
    "fmt"
    "sort"
)
type user struct {
    name string
    surname string
}
users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
}
var keys []string
for key := range users {
    keys = append(keys, key)
}
sort.Strings(keys)
for _, key := range keys {
    fmt.Println(key, users[key])
}
```
* use to make map output not so random