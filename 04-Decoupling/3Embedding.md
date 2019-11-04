# Embedding
```go
type user struct {
    name string
    email string
}
func (u *user) notify() {
    fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email))
}
```
* notify method that uses pointer semantics
## NOT Embedding
```go
type admin struct {
    person user
    level string
}
func main() {
    ad := admin{
        person: user{
            name: "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }
    // We can access the notify
    ad.person.notify()
}
```
# Inner Vs. Outer Type 
* inner type: user
* outer type: admin
* inner/outer relationship: Inner type promotion. Anything related to inner type can be promoted up to outer type

## Embedding
```go
type admin struct {
    user
    level string
}
func main() {
    ad := admin{
        user: user{
            name: "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }
    // We can access the notify
    ad.user.notify()
    // Can also access inner type's method which is promoted
    ad.notify()
}
```
## Add complexity - Add interface
```go
type user struct {
    name string
    email string
}
func (u *user) notify() {
    fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email))
}
type notifier interface{
    notify()
}
type admin struct {
    user
    level string
}
func sendNotification(n notifier) {

}
func main() {
    ad := admin{
        user: user{
            name: "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }
    // Send the admin user a notification
    // Embedded inner type's implementation of the interface is "promoted" to the outer type
    sendNotification(&ad)
```
* Add polymorphic function sendNotification(), and b/c of inner type promotion, can pass the outer type value (admin) to a function that takes a concrete data that satisfies the interface notifier
## Add another layer:
```go
func (a *admin) notify {

}
```
* In this case, there is no inner type promotion - outer type's (`sendNotification(&ad)`) implementation will override the inner type
    * `ad.user.notify()` can still be used, 
    * but `ad.notify()` will call admin's method