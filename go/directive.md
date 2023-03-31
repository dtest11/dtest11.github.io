### directive
hacking.md

. notinheap

go:notinheap applies to type declarations.
It indicates that a type must never be allocated from the GC'd heap or on the stack. 
Specifically, pointers to this type must always fail the runtime.inheap check. 
The type may be used for global variables, or for objects in unmanaged memory 
(e.g., allocated with sysAlloc, persistentalloc, fixalloc, 
or from a manually-managed span). Specifically:
