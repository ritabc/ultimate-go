Exporting
* Encapsulation begins & ends at folder level
* match folder name & pkg name
* pkgs are imported by folder name
* any identifier starting with capital letter is exported, vice versa
* imports point from GOPATH to folder where pkg is:
    - `"github.com/ardanlabs/gotraining/.../counters"` to import counters pkg
    - it's already relative, don't do `"../github.com/.../counters"`

* Exported / Unexported are about identifiers. Whereas private/public distinction is about the data

## Can exported functions in go return unexported types? 
Yes but it's annoying so don't do it:
```go
type alertCounter int

func New(value int) alertCounter {
    return alertCounter(value)
}
```
* linter will complain

## Type level encapsulation
```go
package users
type User struct {
    Name string // exported field
    ID int // exported field

    password string // unexported field
}
```

# asd 
```go
package users

type user struct {
    Name string
    ID int
}

type Manager struct {
    Title string
    user
}

package main
```
* Once outside of this package and when using Manager, what fields do I have acess to? Ansser: Title, Name, Id
    * however, during construction, cannot initialize user value
    * Not getting encapuslation b/c all fields got promoted us anyways
    * This doesn't make sense to encapsulate like this (should usually also make user exported) except when marshalling (marshalling only respects exported fields)
        * When marshalling, types may be unexported but fields are exported 