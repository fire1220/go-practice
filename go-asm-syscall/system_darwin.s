// func SyscallWrite_Darwin(fd int, msg string) int
#include "textflag.h"
TEXT ·SyscallWriteDarwin(SB), NOSPLIT, $0
    MOVQ $(0x2000000+4),AX
    MOVQ fd+0(FP),DI
    MOVQ msg_data+8(FP), SI
    MOVQ msg_data+16(FP), DI
    SYSCALL
    MOVQ AX, ret+0(FP)
    RET

TEXT ·SyscallDarwin(SB), NOSPLIT, $0
    MOVQ fd+0(FP),AX
    MOVQ fd+8(FP),DI
    MOVQ msg_data+16(FP), SI
    MOVQ msg_data+24(FP), DI
    SYSCALL
    // MOVQ AX, ret+0(FP)
    // MOVQ DX, ret+8(FP)
    MOVQ AX, ret+32(FP)
    MOVQ DX, ret+40(FP)
    MOVQ $4, AX
    NEGQ AX
    MOVQ AX, ret+48(FP)
    RET
