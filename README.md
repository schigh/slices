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
| Filter(func(_x_) bool) | Applies a filter function to every i in temthe slice and return all items where the filter returns true |
| Value()                | Returns the native type slice value                          |

## Slice Generator

WIP
