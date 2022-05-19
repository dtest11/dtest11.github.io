// runNext： 最近的 ： 255
// localQueue: 先进先出 0...254
// globalQueue： 先进先出
/**
sysmon: 不需要P 直接运行在M
https://xiaomi-info.github.io/2019/11/27/golang-compiler-plan9/

go tool compile -S -N -l main.go

GOOS=linux GOARCH=amd64 go tool compile -S main.go >> main.asm

https://github.com/go-internals-cn/go-internals/blob/master/chapter2_interfaces/README.md
https://xargin.com/go-and-plan9-asm/
https://xargin.com/plan9-assembly/
FP: Frame pointer: arguments and locals.
PC: Program counter: jumps and branches.
SB: Static base pointer: global symbols.
SP: Stack pointer: the highest address within the local stack frame.
