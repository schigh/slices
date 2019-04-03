[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.org/schigh/slices.svg?branch=master)](https://travis-ci.org/schigh/slices)
[![codecov](https://codecov.io/gh/schigh/slices/branch/master/graph/badge.svg?token=hhqA1l88kx)](https://codecov.io/gh/schigh/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/schigh/slices)](https://goreportcard.com/report/github.com/schigh/slices)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/schigh/slices)

> Some great and powerful gopher once said you shouldn't put this kind of stuff in a library because it's so simple.  Until you have to write a contains method for the thousandth time.

![slices](slices_small.png)

Slices bolts on some generics-ish functionality to native Go data types when represented as a slice.  Currently, the following types are supported:

| Go Slice Type | Slices Type    | xtra |
| ------------- | -------------- | ---- |
| `[]int`       | `IntSlice`     | [benchmark](/benchmarks/int_slice.txt)     |
| `[]int8`      | `Int8Slice`    | [benchmark](/benchmarks/int8_slice.txt)     |
| `[]int16`     | `Int16Slice`   | [benchmark](/benchmarks/int16_slice.txt)     |
| `[]int32`     | `Int32Slice`   | [benchmark](/benchmarks/int32_slice.txt)     |
| `[]int64`     | `Int64Slice`   | [benchmark](/benchmarks/int64_slice.txt)     |
| `[]uint`      | `UIntSlice`    | [benchmark](/benchmarks/uint_slice.txt)     |
| `[]uint8`     | `UInt8Slice`   | [benchmark](/benchmarks/uint8_slice.txt)     |
| `[]uint16`    | `UInt16Slice`  | [benchmark](/benchmarks/uint16_slice.txt)     |
| `[]uint32`    | `UInt32Slice`  | [benchmark](/benchmarks/uint32_slice.txt)     |
| `[]uint64`    | `UInt64Slice`  | [benchmark](/benchmarks/uint64_slice.txt)     |
| `[]uintptr`   | `UIntPtrSlice` | [benchmark](/benchmarks/intptr_slice.txt)     |
| `[]float32`   | `Float32Slice` | [benchmark](/benchmarks/float32_slice.txt)     |
| `[]float64`   | `Float64Slice` | [benchmark](/benchmarks/float64_slice.txt)     |
| `[]string`    | `StringSlice`  | [benchmark](/benchmarks/string_slice.txt)     |

For the above types, the following operations are supported (_x_ is the type in the slice):

| Function               | Description                                                  |
| ---------------------- | ------------------------------------------------------------ |
| IndexOf(_x_)           | Get the first index of the searched element in the slice.  If the element is not found, this function returns -1 |
| Contains(_x_)          | Test to see if slice contains element _x_                    |
| Unique()               | Returns unique items in slice.  The first occurence of an element is the one that is kept |
| SortAsc()              | Sort the slice in ascending order                            |
| SortDesc()             | Sort the slice in descending order                           |
| Reverse()              | Reverses the slice                                           |
| Filter(func(_x_) bool) | Applies a filter function to every item in the slice and return all items where the filter returns true |
| Map(func(x) x)	| Iterate over the slice and replace the current value with the computed value |
| Each(func(x))  | Iterate over the slice (non-mutating) |
| TryEach(func(_x_) error) (int, error) | Iterate over the slice, and halt if an error is returned from user func.  Return index of the failed member and the caught error |
| IfEach(func(_x_) bool) (int, bool) | Iterate over the slice, and halt if false is returned from user func.  Return the index of the element that caused the func to return false, and a bool that is true if every member of the slice returned true with the function applied.  If all elements return true, the index returned is `-1` |
| Chunk(size) | Splits the slice into chunks of a specified size |
| Value() | Returns the native type slice value |

#### Some examples...
```go
package main

import (
	"fmt"
	
	"github.com/schigh/slices"
)

func main() {
	orig := []string{"foo", "bar", "baz", "fizz", "buzz"}
	startsWithF := slices.StringSlice(orig).Filter(func(s string) bool {
		return len(s) > 0 && s[0] == 'f'
	}).Value()
	
	fmt.Println(startsWithF)
	// ["foo", "fizz"]
}
```
```go
package main

import (
	"fmt"
	
	"github.com/schigh/slices"
)

func main() {
	orig := []int{9,1,6,2,3,4,3,4,5,7,8,9,0}
	unique := slices.IntSlice(orig).Unique().SortDesc().Value()
	
	fmt.Println(unique)
	// [9,8,7,6,5,4,3,2,1,0]
}
```

Check out the GoDoc for more information.

## Slice Generator

The slice generator uses `go:generate` to add slice functionality to your existing structs. You may choose which features you'd like to add by setting them in the generate command.  For example:

```
//go:generate slices User map filter each
```

Will generate the `Map`, `Filter`, and `Each` functionality (see below) on a User struct's slice type.  You could also just say you want everything:

```
//go:generate slices User all
```

This will generate all functions produced by the tool.

### Generator Options

#### Closures
Some generated functions take as their receiver a function that operates on every member of the slice.  By default, these receivers use a function where each member is passed by reference.  For example, the generated `Filter` function on a `User` struct:

```go
// Filter evaluates every element in the slice, and returns all User 
// instances where the eval function returns true
func (slice UserSlice) Filter(f func(*User) bool) UserSlice {
	out := make([]User, 0, len(slice))
	for _, v := range slice {
		if f(&v) {
			out = append(out, v)
		}
	}

	return UserSlice(out)
}
```

The user-defined function receives a pointer to `User` by default.

If instead you want your slice functions to operate by value, you can modify your generator tags like so:

```
//go:generate slices User filter(byvalue)
```

This would produce the following function:

```go
// Filter evaluates every element in the slice, and returns all User 
// instances where the eval function returns true
func (slice UserSlice) Filter(f func(User) bool) UserSlice {
	out := make([]User, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return UserSlice(out)
}
```

#### Pointer Slices
You can also generate pointer slices by prepending an asterisk to your struct name in the `go generate` directive, like so:

```
//go:generate slices *User all
```

This will generate a type called `UserPtrSlice`, which is an alias of `[]*User`.  All slice functions take a pointer receiver, for example:

```go
func (slice UserPtrSlice) Map(f func(*User) *User) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
```

