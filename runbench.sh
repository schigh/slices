#!/usr/bin/env bash

echo running int slice benchmark
go test -bench BenchmarkIntSlice > ./benchmarks/int_slice.txt
echo running int8 slice benchmark
go test -bench BenchmarkInt8Slice > ./benchmarks/int8_slice.txt
echo running int16 slice benchmark
go test -bench BenchmarkInt16Slice > ./benchmarks/int16_slice.txt
echo running int32 slice benchmark
go test -bench BenchmarkInt32Slice > ./benchmarks/int32_slice.txt
echo running int64 slice benchmark
go test -bench BenchmarkInt64Slice > ./benchmarks/int64_slice.txt
echo running uint slice benchmark
go test -bench BenchmarkUIntSlice > ./benchmarks/uint_slice.txt
echo running uint8 slice benchmark
go test -bench BenchmarkUInt8Slice > ./benchmarks/uint8_slice.txt
echo running uint16 slice benchmark
go test -bench BenchmarkUInt16Slice > ./benchmarks/uint16_slice.txt
echo running uint32 slice benchmark
go test -bench BenchmarkUInt32Slice > ./benchmarks/uint32_slice.txt
echo running uint64 slice benchmark
go test -bench BenchmarkUInt64Slice > ./benchmarks/uint64_slice.txt
echo running uintptr slice benchmark
go test -bench BenchmarkUIntPtrSlice > ./benchmarks/uintptr_slice.txt
echo running float32 slice benchmark
go test -bench BenchmarkFloat32Slice > ./benchmarks/float32_slice.txt
echo running float64 slice benchmark
go test -bench BenchmarkFloat64Slice > ./benchmarks/float64_slice.txt
echo running string slice benchmark
go test -bench BenchmarkStringSlice > ./benchmarks/string_slice.txt
