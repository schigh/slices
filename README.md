[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.org/schigh/slices.svg?branch=master)](https://travis-ci.org/schigh/slices)
[![codecov](https://codecov.io/gh/schigh/slices/branch/master/graph/badge.svg?token=hhqA1l88kx)](https://codecov.io/gh/schigh/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/schigh/slices)](https://goreportcard.com/report/github.com/schigh/slices)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/schigh/slices)

![slices](slices_small.png)

Slices bolts on some generics-ish functionality to native Go data types when represented as a slice.  Currently, the following types are supported:

| Go Slice Type | Slices Type    |
| ------------- | -------------- |
| `[]int`       | `IntSlice`     |
| `[]int8`      | `Int8Slice`    |
| `[]int16`     | `Int16Slice`   |
| `[]int32`     | `Int32Slice`   |
| `[]int64`     | `Int64Slice`   |
| `[]uint`      | `UIntSlice`    |
| `[]uint8`     | `UInt8Slice`   |
| `[]uint16`    | `UInt16Slice`  |
| `[]uint32`    | `UInt32Slice`  |
| `[]uint64`    | `UInt64Slice`  |
| `[]uintptr`   | `UIntPtrSlice` |
| `[]float32`   | `Float32Slice` |
| `[]float64`   | `Float64Slice` |
| `[]string`    | `StringSlice`  |

For the above types, the following operations are supported (_x_ is the type in the slice):

| Function               | Description                                                  |
| ---------------------- | ------------------------------------------------------------ |
| IndexOf(_x_)           | Get the first index of the searched element in the slice.  If the element is not found, this function returns -1 |
| Contains(_x_)          | Test to see if slice contains element _x_                    |
| SortAsc()              | Sort the slice in ascending order                            |
| SortDesc()             | Sort the slice in descending order                           |
| Reverse()              | Reverses the slice                                           |
| Filter(func(_x_) bool) | Applies a filter function to every item in the slice and return all items where the filter returns true |
| Copy()				 | Returns a copy from the original backing array/slice.  This is useful if you no longer need to keep the original backing array hanging around in memory and no longer wish to reference it. For small backing arrays, this function doesn't really produce any benefit |
| Value()                | Returns the native type slice value                          |

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
	for _, instance := range slice {
		if f(&instance) {
			out = append(out, instance)
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
	for _, instance := range slice {
		if f(instance) {
			out = append(out, instance)
		}
	}

	return UserSlice(out)
}
```

#### Copies
WIP
