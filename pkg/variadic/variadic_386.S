// Copyright 2013 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "asmregs_386.h"

#ifndef __linux__
# define name(x) _ ## x
#else
# define name(x) x
#endif

.globl name(VariadicCall)
name(VariadicCall):
.globl name(VariadicCallFloat)
name(VariadicCallFloat):
.globl name(VariadicCallDouble)
name(VariadicCallDouble):
	pushl %ebp
	movl  %esp, %ebp

	movl  8(%esp), %eax

	// Align the stack to a 16-byte boundary
	movl %esp, %ecx
	andl $0xfffffff0, %ecx
	movl %ecx, %esp

	movl  REGS_NARGS(%eax), %ecx

	cmpl  $0, %ecx
	je    memdone
	imull $4, %ecx
	addl  %eax, %ecx
mem:
	subl   $4, %ecx
	pushl  (%ecx)
	cmpl   %ecx, %eax
	jne    mem
memdone:

	addl   $REGS_ADDR, %eax
	call   *(%eax)

	leave
	ret
