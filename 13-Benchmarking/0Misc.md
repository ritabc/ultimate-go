# Miscellaneous Benchmarking Notes
## From 3.2.1 Arrays Part 1 - Mechanical Sympathy
* ideally, have machine be idle on everything besides benchmark
* Function must start with Benchmark and take *testing.B
```go
func BenchmarkLinkListTraverse(b *testing.B) {
    var a int

    for i := 0; i < b.N; i++ {
        a = LinkedListTraverse()
    }
}
```
* go compiler will re-compile the code, and has the ability to throw away dead code
    - if LinkedListTraverse() returns something but doesn't get used, the Benchmark could decide not to run it it the returned value doesn't get aassigned anywhere
    - So, set result to var a int