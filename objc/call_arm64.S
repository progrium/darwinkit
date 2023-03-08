// Copyright (c) 2012-2022 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

.align 8
.globl _GoObjc_CallTargetFrameSetup
// uintptr GoObjc_CallTargetFrameSetup()
_GoObjc_CallTargetFrameSetup:
	sub sp, sp, 16
	stp x29, x30, [sp]

	sub sp, sp, 16
	stp x0, x0, [sp]

	mov x29, sp

	sub sp, sp, 16
	str q7, [sp]
	sub sp, sp, 16
	str q6, [sp]
	sub sp, sp, 16
	str q5, [sp]
	sub sp, sp, 16
	str q4, [sp]
	sub sp, sp, 16
	str q3, [sp]
	sub sp, sp, 16
	str q2, [sp]
	sub sp, sp, 16
	str q1, [sp]
	sub sp, sp, 16
	str q0, [sp]
	sub sp, sp, 16
	stp x6, x7, [sp]
	sub sp, sp, 16
	stp x4, x5, [sp]
	sub sp, sp, 16
	stp x2, x3, [sp]
	sub sp, sp, 16
	stp x0, x1, [sp]

	mov x0, sp

	bl _goMethodCallEntryPoint

	add sp, sp, 192 // Restore stack

	add sp, sp, 16 // Restore x0

	ldp x29, x30, [sp]
	add sp, sp, 16

	ret