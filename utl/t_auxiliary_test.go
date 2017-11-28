// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utl

import (
	"testing"

	"github.com/CalOmnie/gosl/chk"
	"github.com/CalOmnie/gosl/io"
)

func TestBestsq01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Bestsq01")

	for i := 1; i <= 12; i++ {
		nrow, ncol := BestSquare(i)
		io.Pforan("nrow, ncol, nrow*ncol = %2d, %2d, %2d\n", nrow, ncol, nrow*ncol)
		if nrow*ncol != i {
			chk.Panic("BestSquare failed")
		}
	}
}

func TestAuxFuncs01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("AuxFuncs01. auxiliary functions")

	n := 1073741824 // 2³⁰
	io.Pf("n=%10v  IsPowerOfTwo=%v\n", n, IsPowerOfTwo(n))
	if !IsPowerOfTwo(n) {
		tst.Errorf("n=%d is power of 2\n", n)
		return
	}

	n = 1234567
	io.Pf("n=%10v  IsPowerOfTwo=%v\n", n, IsPowerOfTwo(n))
	if IsPowerOfTwo(n) {
		tst.Errorf("n=%d is not power of 2\n", n)
		return
	}

	n = 0
	io.Pf("n=%10v  IsPowerOfTwo=%v\n", n, IsPowerOfTwo(n))
	if IsPowerOfTwo(n) {
		tst.Errorf("n=%d is not power of 2\n", n)
		return
	}

	n = -2
	io.Pf("n=%10v  IsPowerOfTwo=%v\n", n, IsPowerOfTwo(n))
	if IsPowerOfTwo(n) {
		tst.Errorf("n=%d is not power of 2\n", n)
		return
	}

	n = 1 // 2⁰
	io.Pf("n=%10v  IsPowerOfTwo=%v\n", n, IsPowerOfTwo(n))
	if !IsPowerOfTwo(n) {
		tst.Errorf("n=%d is power of 2\n", n)
		return
	}

	a, b := 123.0, 456.0
	io.Pf("a=%v b=%v\n", a, b)
	Swap(&a, &b)
	io.Pf("a=%v b=%v\n", a, b)
	if a == 123 || b == 456 {
		tst.Errorf("Swap failed\n")
		return
	}

	c := -1
	if Iabs(c) != 1 {
		tst.Errorf("Iabs failed\n")
		return
	}

	c = 1
	if Iabs(c) != 1 {
		tst.Errorf("Iabs failed\n")
		return
	}
}

func TestMinMax01(tst *testing.T) {

	//chk.Verbose = true
	chk.PrintTitle("MinMax01")

	if Imin(1, 2) != 1 {
		tst.Errorf("Imin() failed\n")
		return
	}

	if Imin(2, 1) != 1 {
		tst.Errorf("Imin() failed\n")
		return
	}

	if Imax(1, 2) != 2 {
		tst.Errorf("Imax() failed\n")
		return
	}

	if Imax(2, 1) != 2 {
		tst.Errorf("Imax() failed\n")
		return
	}

	if Min(1, 2) != 1 {
		tst.Errorf("Min() failed\n")
		return
	}

	if Min(2, 1) != 1 {
		tst.Errorf("Min() failed\n")
		return
	}

	if Max(1, 2) != 2 {
		tst.Errorf("Max() failed\n")
		return
	}

	if Max(2, 1) != 2 {
		tst.Errorf("Max() failed\n")
		return
	}
}
