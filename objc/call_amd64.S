// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

.globl _GoObjc_CallTargetFrameSetup
// uintptr GoObjc_CallTargetFrameSetup()
_GoObjc_CallTargetFrameSetup:
	pushq %rbp
	movq  %rsp, %rbp

	subq $8, %rsp
	movsd %xmm7, (%rsp)
	subq $8, %rsp
	movsd %xmm6, (%rsp)
	subq $8, %rsp
	movsd %xmm5, (%rsp)
	subq $8, %rsp
	movsd %xmm4, (%rsp)
	subq $8, %rsp
	movsd %xmm3, (%rsp)
	subq $8, %rsp
	movsd %xmm2, (%rsp)
	subq $8, %rsp
	movsd %xmm1, (%rsp)
	subq $8, %rsp
	movsd %xmm0, (%rsp)

	pushq %r9
	pushq %r8
	pushq %rcx
	pushq %rdx
	pushq %rsi
	pushq %rdi

	movq %rsp, %rdi

	callq _goMethodCallEntryPoint

	movsd 48(%rsp), %xmm0

	movq  %rbp, %rsp
	popq  %rbp
	retq