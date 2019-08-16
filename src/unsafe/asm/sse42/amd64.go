// +build amd64,!noavx,!nosimd
package sse42

//go:asm VPACKUSWB
func VPADDD_YMMu32_MASKmskw_YMMu32_YMMu32_AVX512() uint8
