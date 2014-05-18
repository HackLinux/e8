IR1 is intermediate representation 1.

In IR1, code segments are managed in packages.

A package contains four type of symbols

- Constants, which are all typed integers
- Types, which are definition of memory layouts
- Variables, which are preallocated objects on the heap
- Functions, which are code segments

A function contains labels and branching instructions.  It does not
have loops.  It is easy to directly translate the instructions into
ISAs like ARM, x86, MIPS, or e8.  It is not straightforward to
tranlate the instructions into asm.js (for now), becauase asm.js does
not have goto statements. For asm.js, we will probably have a
variation of IR1 which does not have goto statements.

Currently IR1 is a strict subset of golang. It is not part of the
design requirement, but it is a nice property to have.