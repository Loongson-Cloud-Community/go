// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64

#include "go_asm.h"
#include "textflag.h"

#define	REGCTXT	R29

// memequal(a, b unsafe.Pointer, size uintptr) bool
TEXT runtime·memequal(SB),NOSPLIT|NOFRAME,$0-25
	MOVV	a+0(FP), R4
	MOVV	b+8(FP), R5
	BEQ	R4, R5, eq
	MOVV	size+16(FP), R6
	ADDV	R4, R6, R7
loop:
	BNE	R4, R7, test
	MOVV	$1, R4
	MOVB	R4, ret+24(FP)
	RET
test:
	MOVBU	(R4), R9
	ADDV	$1, R4
	MOVBU	(R5), R10
	ADDV	$1, R5
	BEQ	R9, R10, loop

	MOVB	R0, ret+24(FP)
	RET
eq:
	MOVV	$1, R4
	MOVB	R4, ret+24(FP)
	RET

// memequal_varlen(a, b unsafe.Pointer) bool
TEXT runtime·memequal_varlen(SB),NOSPLIT,$40-17
	MOVV	a+0(FP), R4
	MOVV	b+8(FP), R5
	BEQ	R4, R5, eq
	MOVV	8(REGCTXT), R6    // compiler stores size at offset 8 in the closure
	MOVV	R4, 8(R3)
	MOVV	R5, 16(R3)
	MOVV	R6, 24(R3)
	JAL	runtime·memequal(SB)
	MOVBU	32(R3), R4
	MOVB	R4, ret+16(FP)
	RET
eq:
	MOVV	$1, R4
	MOVB	R4, ret+16(FP)
	RET
