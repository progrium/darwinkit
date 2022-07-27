// Copyright 2022 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "asmregs_arm64.h"

.align 8
.globl _VariadicCall
_VariadicCall:
.globl _VariadicCallFloat
_VariadicCallFloat:
.globl _VariadicCallDouble
_VariadicCallDouble:
	sub sp, sp, 16
	stp x29, x30, [sp]

	sub sp, sp, 16
	stp x0, x0, [sp]

	mov x29, sp

	mov x9, x0

	ldr x0, [x9, #REGS_X0]
	ldr x1, [x9, #REGS_X1]
	ldr x2, [x9, #REGS_X2]
	ldr x3, [x9, #REGS_X3]
	ldr x4, [x9, #REGS_X4]
	ldr x5, [x9, #REGS_X5]
	ldr x6, [x9, #REGS_X6]
	ldr x7, [x9, #REGS_X7]

	ldr q0, [x9, #REGS_Q0]
	ldr q1, [x9, #REGS_Q1]
	ldr q2, [x9, #REGS_Q2]
	ldr q3, [x9, #REGS_Q3]
	ldr q4, [x9, #REGS_Q4]
	ldr q5, [x9, #REGS_Q5]
	ldr q6, [x9, #REGS_Q6]
	ldr q7, [x9, #REGS_Q7]

	ldr x10, [x9, #REGS_ADDR]

	blr x10

	ldr x9, [sp]
	add sp, sp, 16

	str x0, [x9, #REGS_X0]
	str x1, [x9, #REGS_X1]
	str x2, [x9, #REGS_X2]
	str x3, [x9, #REGS_X3]
	str x4, [x9, #REGS_X4]
	str x5, [x9, #REGS_X5]
	str x6, [x9, #REGS_X6]
	str x7, [x9, #REGS_X7]

	str q0, [x9, #REGS_Q0]
	str q1, [x9, #REGS_Q1]
	str q2, [x9, #REGS_Q2]
	str q3, [x9, #REGS_Q3]
	str q4, [x9, #REGS_Q4]
	str q5, [x9, #REGS_Q5]
	str q6, [x9, #REGS_Q6]
	str q7, [x9, #REGS_Q7]

	ldp x29, x30, [sp]
	add sp, sp, 16

	ret