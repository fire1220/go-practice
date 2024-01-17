#include "textflag.h"

// go-go1.16.14/src/runtime/internal/atomic/asm_amd64.s

// uint32 Xadd(uint32 volatile *val, int32 delta)
// Atomically:
//	*val += delta;
//	return *val;
// 注释：内存原子操作，两个参数交换后相加，返回第1个参数的旧值
TEXT ·Xadd(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX   // 注释：参数1：(ptr)是个指针
	MOVL	delta+8(FP), AX // 注释：参数2：(delta)需要相加的数据
	MOVL	AX, CX          // 注释：(备份delta数据)把AX值复制到CX寄存器里（把参数2(delta)放到CX里）
	LOCK                    // 注释：同步锁，保证寄存器在CPU缓存中的数据是一致的
	XADDL	AX, 0(BX)       // 注释： 注释：(原子操作,两值调换相加),结果：0(BX) = AX + 0(BX)； AX = 0(BX),AX是ptr指针里的旧值
	ADDL	CX, AX          // 注释：【AX = delta旧值(CX) + ptr旧值(AX)】把0(BX)内存数据原子操作时的前数据拿出来和CX相加（CX就是函数的第二个参数），防止再次读取内存破坏原子性
	MOVL	AX, ret+16(FP)  // 注释：返回原子相加后的值（是针对内存的原子操作）
	RET                     // 注释：函数退出
