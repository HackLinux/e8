[![BuildStatus](https://travis-ci.org/e8vm/e8.png?branch=master)](https://travis-ci.org/e8vm/e8)

# E8VM

This project is work in progress. I want to build:

- E8: a simulator that simulates a virtual architecture with a MIPS-like
instruction set.
- Leaf: a compiler that compiles a C/C++ like programming language with
Go-like syntax, targeting E8.
- An operating system that runs in E8 written in Leaf.
- Some example small programs in Leaf.
- A Leaf compiler that is written in Leaf.

## Why? Why? Why?

### Why E8?

The goal of the entire project is not to have something that can run as soon
as possible; we already have things that can run well, and I don't want to compete.

Instead, I want to answer a question in my mind.

The designed mechanisms and principles for a working computer (including the
architecture, the OS, the tool chain for building running programs) often seems simple
and straight-forward at a very high level, yet real working
systems are often extremely clunky and complicated. As a result,
on real systems, it is very
hard to even play with some simple research ideas. Even if you know how it
works at a high level, it is often hard to tell how it really works or even
if it really works.

I often feel very uncomfortable about this situation. Does it have to be that complicated? Given that we already
learned the lessons on designing ISAs, OSs and compilers, and also
lessons on good
software engineering, if we have the chance to design the ISA and the OS
from a clean slate, can we build a working (though simulated) computer stack that
is just simple? Can we build an operating system where its source code is
very well modularized? Can we have
a working simulated virtual machine where every piece of code is
easily comprehensible and every experiment is easily repeatable?

### Why not x86 or ARM or simply MIPS? Why not JVM, LuaVM, LLVM byte code?

To have repeatable experiment results, I would like an architecture that I can reason about the complexity of a
program by simply counting cycles. x86 and ARM are too complicated. E8 is
actually very similar to MIPS, but slightly different and simpler in my
opinion. Other language VMs have too many handy but high-level features
that are hard to reason about its running time.

### Why not NaCl? Why not PNaCl?

I am open to the idea of porting a subset of x86 or ARM, but it is
a huge amount of work comparing with porting E8. I would rather first
complete the project just for E8 first. I will keep retargeting in mind
when writing the compiler.

Having the code to be able to run in Chrome would be pretty nice.
However, besides Google, I don't see other open-source project that compiles to
NaCl or PNaCl. Neither I see documents that talk about how to write
compilers for NaCl or PNaCl.

### Why Leaf? Why a new language? Why not use C and Linux? Why not use tinycc?

I need a low-level programming language to write the OS for E8.
At the same time, to have the OS comprehensible, I want to write it in
a "modern" language that at least has built-in
supports for namespace, type methods and interface, and for
the source code to be easily comprehensible, I hope the language
does not have macros or templates or other features that
can arbitrarily polymorphs the source code with dark magic.

C and C++ are hence ruled out.

I considered referencing TinyCC's source code.
It is actually not very well modularized, like it uses a lot of global
variables.

TinyCC but it does not compile Linux source code. Even Clang
does not really compile Linux source code. Only GNU tools compile
Linux. GNU produces probably the worst tool chain in terms of
comprehensibility.

### Why not Python? Why not D? Why not Lua? Why not Rust? Why not Go?

Most languages are not simple. Python and D are very complicated languages
that have complex syntax and numbers of features, which makes writing/porting
a compiler a very hard job.

Lua is simple, but it is designed to be a scripting language.

I feel Rust is not stable and somehow too ambitious. I think it is a
good platform for testing cool system language ideas. However, it is not
stable, yet already pretty complex. Many open-source projects have
similar problems: people keep adding features that they like without much
discipline, and makes the result product too complex to use in the end.
Adding a new feature is easy; removing a bad one is hard.

Also, Rust uses LLVM as the back-end, which I dislike (see "Why not LLVM?")

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
clang or Rust to E8 using LLVM. I am not sure, but might worth a try.

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
