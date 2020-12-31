// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

.globl _GoObjc_Fst
// void GoObjc_Fst(uintptr f)
_GoObjc_Fst:
	mov  4(%esp), %eax
	fstl (%eax)
	ret

.globl _GoObjc_CallTargetFrameSetup
// uintptr GoObjc_CallTargetFrameSetup()
_GoObjc_CallTargetFrameSetup:
	pushl %ebp
	movl  %esp, %ebp

	pushl %esp
	call _goMethodCallEntryPoint
	popl  %ecx

	cmpl  $0, %ecx
	je    done
	fldl  (%ecx)

done:
	leave
	ret