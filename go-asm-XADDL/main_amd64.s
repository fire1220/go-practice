
#include "textflag.h"

// uint32 Xadd(uint32 volatile *val, int32 delta)
// Atomically:
//	*val += delta;
//	return *val;
// 注释：内存原子操作，把两数之和返回
TEXT ·Xadd(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX   // 注释：第一个参数是指针
	MOVL	delta+8(FP), AX // 注释：第二个参数是数值
	MOVL	AX, CX          // 注释：备份第二个参数到CX里
	LOCK                    // 注释：同步锁，保证寄存器在CPU缓存中的数据是一致的
	XADDL	AX, 0(BX)       // 注释：原子操作内存，0(BX)值：AX+0(BX)；AX值：0(BX)，内存0(BX)是原子相加
	ADDL	CX, AX          // 注释：AX = CX + AX ；把0(BX)内存数据原子操作时的前数据拿出来和CX相加（CX就是函数的第二个参数），防止再次读取内存破坏原子性
	MOVL	AX, ret+16(FP)  // 注释：返回原子相加后的值（是针对内存的原子操作）
	RET                     // 注释：函数退出
