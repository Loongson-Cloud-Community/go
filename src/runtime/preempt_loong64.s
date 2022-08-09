// Code generated by mkpreempt.go; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	MOVV R1, -472(R3)
	SUBV $472, R3
	MOVV R4, 8(R3)
	MOVV R5, 16(R3)
	MOVV R6, 24(R3)
	MOVV R7, 32(R3)
	MOVV R8, 40(R3)
	MOVV R9, 48(R3)
	MOVV R10, 56(R3)
	MOVV R11, 64(R3)
	MOVV R12, 72(R3)
	MOVV R13, 80(R3)
	MOVV R14, 88(R3)
	MOVV R15, 96(R3)
	MOVV R16, 104(R3)
	MOVV R17, 112(R3)
	MOVV R18, 120(R3)
	MOVV R19, 128(R3)
	MOVV R20, 136(R3)
	MOVV R21, 144(R3)
	MOVV R23, 152(R3)
	MOVV R24, 160(R3)
	MOVV R25, 168(R3)
	MOVV R26, 176(R3)
	MOVV R27, 184(R3)
	MOVV R28, 192(R3)
	MOVV R29, 200(R3)
	MOVV RSB, 208(R3)
	#ifndef GOLOONG64_softfloat
	MOVD F0, 216(R3)
	MOVD F1, 224(R3)
	MOVD F2, 232(R3)
	MOVD F3, 240(R3)
	MOVD F4, 248(R3)
	MOVD F5, 256(R3)
	MOVD F6, 264(R3)
	MOVD F7, 272(R3)
	MOVD F8, 280(R3)
	MOVD F9, 288(R3)
	MOVD F10, 296(R3)
	MOVD F11, 304(R3)
	MOVD F12, 312(R3)
	MOVD F13, 320(R3)
	MOVD F14, 328(R3)
	MOVD F15, 336(R3)
	MOVD F16, 344(R3)
	MOVD F17, 352(R3)
	MOVD F18, 360(R3)
	MOVD F19, 368(R3)
	MOVD F20, 376(R3)
	MOVD F21, 384(R3)
	MOVD F22, 392(R3)
	MOVD F23, 400(R3)
	MOVD F24, 408(R3)
	MOVD F25, 416(R3)
	MOVD F26, 424(R3)
	MOVD F27, 432(R3)
	MOVD F28, 440(R3)
	MOVD F29, 448(R3)
	MOVD F30, 456(R3)
	MOVD F31, 464(R3)
	#endif
	CALL ·asyncPreempt2(SB)
	#ifndef GOLOONG64_softfloat
	MOVD 464(R3), F31
	MOVD 456(R3), F30
	MOVD 448(R3), F29
	MOVD 440(R3), F28
	MOVD 432(R3), F27
	MOVD 424(R3), F26
	MOVD 416(R3), F25
	MOVD 408(R3), F24
	MOVD 400(R3), F23
	MOVD 392(R3), F22
	MOVD 384(R3), F21
	MOVD 376(R3), F20
	MOVD 368(R3), F19
	MOVD 360(R3), F18
	MOVD 352(R3), F17
	MOVD 344(R3), F16
	MOVD 336(R3), F15
	MOVD 328(R3), F14
	MOVD 320(R3), F13
	MOVD 312(R3), F12
	MOVD 304(R3), F11
	MOVD 296(R3), F10
	MOVD 288(R3), F9
	MOVD 280(R3), F8
	MOVD 272(R3), F7
	MOVD 264(R3), F6
	MOVD 256(R3), F5
	MOVD 248(R3), F4
	MOVD 240(R3), F3
	MOVD 232(R3), F2
	MOVD 224(R3), F1
	MOVD 216(R3), F0
	#endif
	MOVV 208(R3), RSB
	MOVV 200(R3), R29
	MOVV 192(R3), R28
	MOVV 184(R3), R27
	MOVV 176(R3), R26
	MOVV 168(R3), R25
	MOVV 160(R3), R24
	MOVV 152(R3), R23
	MOVV 144(R3), R21
	MOVV 136(R3), R20
	MOVV 128(R3), R19
	MOVV 120(R3), R18
	MOVV 112(R3), R17
	MOVV 104(R3), R16
	MOVV 96(R3), R15
	MOVV 88(R3), R14
	MOVV 80(R3), R13
	MOVV 72(R3), R12
	MOVV 64(R3), R11
	MOVV 56(R3), R10
	MOVV 48(R3), R9
	MOVV 40(R3), R8
	MOVV 32(R3), R7
	MOVV 24(R3), R6
	MOVV 16(R3), R5
	MOVV 8(R3), R4
	MOVV 472(R3), R1
	MOVV (R3), R30
	ADDV $480, R3
	JMP (R30)
