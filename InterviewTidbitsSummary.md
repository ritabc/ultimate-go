# Terms
* Alignment
* Padding
* Stack
    - as we make function calls, we progress DOWN
    - as we return, we come back UP
    - Memory below active frame is invalid and garbage collected
* Heap: more global storage than stack, Used when need for:
    - sharing value across goroutine boundaries
    - Value that needs to persist to higher frame
    - When we don't know size during compile time
* Value Semantics: sharing data by copies
    - benefits of isolation and immutability: objects only are mutated in expected places
    - cost of multiple copies aross the program, and need for duplicate updating
* Pointer Semantics: sharing data by address
    - cost of dangerous mutating: must always be aware when other threads have access to the data
    - benefit of efficiency
* '*' means declaring pointer variable
* '&' means Sharing
* Escape Analysis - whether values are allocated to the Heap
    - stacks are self-cleaning, meaning everything below active frame on stack is cleaned
    - Hopefully values stay on the stack, and ARE NOT escaped to the Heap (b/c then we don't have to worry about explicitly gc'ing them)
    - return &u tells us value is shared up the call stack. If it stayed on the stack, we'd lose access to &u's underlying value, thus it's put on the heap automatically
* Don't start life of variable as pointer, as in dont do: u := &user{name: "Bill"}
* Goroutines, and thus stacks (1-1 relationship?), must be small so program can be fast
* Constants of kind vs of type
    - of-kind constants aren't initialized with type, and can be implicitly converted, and are not limited by precision (as in, can handle more precise numbers)
    - of-type are initialized with a type and cannot be converted implicitly
* Kind Promotion: type promotes over kind, floats promote over ints
* Prefetchers & Hardware (software that walk through memory at predictable stride) work well with arrays (which are contiguously blocked in memory). They can pick up on data access and start bringing in cache lines way ahead of when they're needed
* For range value vs. pointer semantics:
    - Value: for i, v := range slice {    } // v is copy of element
    - Pointer: for i := range slice {} // no copy is made
* Append is able to do mutations in isolation without causing side effects (mutates w/o using pointers)
* Common Go causes of memory leak?
    - extra goroutines being created
    - failing to terminate them 
    - extra maps and extra keys
    - append not being used properly, as in data = append(bill, newItem) (We could be holding reference to old backing arrays and not replacing them)
    - Closing APIs
* Append behaves differently based on whether len == capacity of the slice it's appending to