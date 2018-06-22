#!/usr/bin/env bash

echo running int slice benchmark
go test -bench BenchmarkIntSlice -benchmem > ./benchmarks/int_slice.txt
echo running int8 slice benchmark
go test -bench BenchmarkInt8Slice -benchmem > ./benchmarks/int8_slice.txt
echo running int16 slice benchmark
go test -bench BenchmarkInt16Slice -benchmem > ./benchmarks/int16_slice.txt
echo running int32 slice benchmark
go test -bench BenchmarkInt32Slice -benchmem > ./benchmarks/int32_slice.txt
echo running int64 slice benchmark
go test -bench BenchmarkInt64Slice -benchmem > ./benchmarks/int64_slice.txt
echo running uint slice benchmark
go test -bench BenchmarkUIntSlice -benchmem > ./benchmarks/uint_slice.txt
echo running uint8 slice benchmark
go test -bench BenchmarkUInt8Slice -benchmem > ./benchmarks/uint8_slice.txt
echo running uint16 slice benchmark
go test -bench BenchmarkUInt16Slice -benchmem > ./benchmarks/uint16_slice.txt
echo running uint32 slice benchmark
go test -bench BenchmarkUInt32Slice -benchmem > ./benchmarks/uint32_slice.txt
echo running uint64 slice benchmark
go test -bench BenchmarkUInt64Slice -benchmem > ./benchmarks/uint64_slice.txt
echo running uintptr slice benchmark
go test -bench BenchmarkUIntPtrSlice -benchmem > ./benchmarks/uintptr_slice.txt
echo running float32 slice benchmark
go test -bench BenchmarkFloat32Slice -benchmem > ./benchmarks/float32_slice.txt
echo running float64 slice benchmark
go test -bench BenchmarkFloat64Slice -benchmem > ./benchmarks/float64_slice.txt
echo running string slice benchmark
go test -bench BenchmarkStringSlice -benchmem > ./benchmarks/string_slice.txt
