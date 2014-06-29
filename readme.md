[![BuildStatus](https://travis-ci.org/e8vm/e8.png?branch=master)](https://travis-ci.org/e8vm/e8)

# E8VM

This project is work in progress. I want to build:

- E8: a simulator that simulates a virtual architecture with a MIPS-like
instruction set.
- Leaf: a compiler that compiles a C/C++ like programming language with
Go-like syntax, targeting E8.
- An operating system that runs in E8 written in Leaf.
- Some example small programs in Leaf.
- Rewrite the leaf compiler in leaf.

## Why? Why? Why?

### Why E8? Why not x86 or ARM or simply MIPS? Why not JVM, LuaVM, LLVM?

I want an architecture that I can reason about the complexity of a
program by simply counting cycles. x86 and ARM are too complicate. E8 is
actually very similar to MIPS, but slightly different. Other language VMs
have too many handy but high-level features that are hard to reason about
its running time.

### Why not NaCl? Why not PNaCl?

I am open to the idea of porting a subset of x86 or ARM, but it is
a huge amount of work comparing with porting E8. I would rather first
complete the project just for E8 first. I will keep retargeting in mind
when writing the compiler.

Having the code be able to run in Chrome would be pretty nice.
However, besides Google, I don't see other open-source projects compiles to
NaCl or PNaCl. Neither I see documents that talk about how to write
compilers for NaCl or PNaCl.

### Why leaf? Why a new language? Why not use Linux? Why not use tinycc?

I need a low-level programming language to write stuff (like an OS) for E8.

I don't like C/C++. It does not provide the right language syntax for writing
comprehensible code.

The goal of the entire project is not to have something that can run as soon
as possible; we already have things that can run well. I won't compete with
stuff that already works in the real world.

However, I do have a question in my mind.

The designed mechanisms and principles for a working computer (including the
architecture, the OS, the tool chain for building programs) often seems simple
and straight-forward at a very high level, yet real working
systems are often extremely clunky and complicated. On these systems, it is very
hard to even play with some simple research ideas.

My question: Does it have to be that complicated? Given that we already
learned the lessons on designing RISC, OS and compilers, and also on good
software engineering, if we have the chance to design the RISC and the OS
from a clean slate, can we build a working (simulated) computer stack that
is just simple? Can we build an operating system where the source code is
well modularized and written in a "modern" language that has built-in
supports for at least namespace, type methods and interface? Can we have
a working simulated VM that every piece of code is easily comprehensible and
every experiment is easily repeatable?

For that, I feel like we need a new language.

### Why not Lua? Why not Python? Why not D? Why not Rust? Why not Go?

Most languages are not simple. Python and D are very complicated languages
that have complex syntax and numbers of features, which makes writing/porting
a compiler a very hard job.

Lua is simple, but it is designed to be a scripting language.

I like Go very much. Go is simple, relatively speaking. However, Go is a
type-safe language. It has a large runtime, and a lot of Go features
heavily relies on the runtime (like garbage collection). This does not very
fit the need to write an operating system. I know there exists projects
that writes in high-level languages that has a large run-time, but I still
feel that the OS developer should naturally have the ability to manage
the memory layout on their own.

Go sounds like a very good choice for writing user-level programs for E8 in
the future. However, as a Google open sourced project, the source code of the
Go language compiler, written in C, is not easily comprehensible. I don't know
how easy it is to port Go to E8.

### Why not LLVM?

Despite its success, I just don't want to dive myself into this monster:
written in C++, heavily uses STL and inheritance, not always well-documented
on how it works internally, and takes more than 1GB of memory just to compile
its source code.

It might be worth prototyping the project with LLVM if I am a programming
language PhD student and already familiar with LLVM, but as a Sysnet PhD,
I think writing my own compiler from scratch might be even faster in the end.

The E8 architecture is pretty simple, so it might be pretty easy to port
clang to E8 using LLVM. I am not sure though, but might worth a try.

## Install

```
$ go get e8vm.net/e8
$ go get e8vm.net/leaf
```

## Old Readme

**Why**

- This could be an online platform for ACM training
- This could be an online arena/platform for game AI fighting
- This could be used for collaborative programming
- This encourages users to collaborate by assembling modules via clearly
  defined interfaces and test them with testcases.
- Ideally, all the modules and interfaces could be easily reused, since they
  would likely be small, and hence easy for human to read, understand, modify
  and maintain.

See more on my [Motivation](https://e8vm.net/e8/wiki/Motivation) page.

**Design**

- [RISC](https://e8vm.net/e8/wiki/RISC-Specification): The MIPS-like
  simple instructions set that `e8` uses only has less than 40 instructions,
  which means `e8` CPU will be very easy to port. In fact, I have already
  ported the core to Javascript: [`e8js`](https://e8vm.net/e8js).
- [System Page](https://e8vm.net/e8/wiki/Page-0:-System-page): all IOs
  in `e8` will be memory mapped, so there is no need for special instructions
  like `in` and `out`. Basic system functionality will be mapped to Page 0.
  Future fancy hardware will be mapped to the following small-id pages in the
  address space.
- [OS-related features](https://e8vm.net/e8/wiki/Interrupts-and-Operating-System):
  `e8` will not have protection rings (e.g. kernel mode and user mode).
  Instead, it will use an approach similar to ARM's TrustZone, where there will
  be a previledged VM that can manipulate other child VMs' execution state and
  page tables. Inside a VM, there will be events, but there will be no
  interrupt handlers. A VM can suspend itself and wake up on an interrupt
  event, so a VM can be event driven, but code execution will not be forced to
  suspend and continue with another program counter value. A previledged VM,
  however, can simulate interrupt handling on its child VMs, if it is desired
  to.
- Language support: The project (simulator, assembler, compiler) is written in
  golang, and I plan to implement a subset of golang that compiles to `e8`, the
  assembler and the compiler will be ported to that subset language later in
  the future, and there will be an advanced language compiler that runs in `e8`
  where it can compile itself.
  [More thoughts here](https://e8vm.net/e8/wiki/Thoughts-on-Language).
