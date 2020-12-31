// Copyright 2012 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "asmregs_amd64.h"

.globl _VariadicCall
_VariadicCall:
.globl _VariadicCallFloat
_VariadicCallFloat:
.globl _VariadicCallDouble
_VariadicCallDouble:
	pushq %rbp
	movq  %rsp, %rbp

	movq  %rdi, %rax

	movq  REGS_RDI(%rax), %rdi
	movq  REGS_RSI(%rax), %rsi
	movq  REGS_RDX(%rax), %rdx
	movq  REGS_RCX(%rax), %rcx
	movq  REGS_R8(%rax), %r8
	movq  REGS_R9(%rax), %r9

	movsd  REGS_XMM0(%rax), %xmm0
	movsd  REGS_XMM1(%rax), %xmm1
	movsd  REGS_XMM2(%rax), %xmm2
	movsd  REGS_XMM3(%rax), %xmm3
	movsd  REGS_XMM4(%rax), %xmm4
	movsd  REGS_XMM5(%rax), %xmm5
	movsd  REGS_XMM6(%rax), %xmm6
	movsd  REGS_XMM7(%rax), %xmm7

	movq   REGS_NMEMORY(%rax), %r11
	movq   REGS_MEMORY(%rax), %r10

	cmpq  $0, %r11
	je     memdone
	imulq  $8, %r11
	addq   %r10, %r11
mem:
	subq   $8, %r11
	pushq  (%r11)
	cmpq   %r11, %r10
	jne    mem
memdone:

	movq REGS_ADDR(%rax), %r10
	movq REGS_NFLOAT(%rax), %rax

	callq *%r10

	movq  %rbp, %rsp
	popq  %rbp
	retq
