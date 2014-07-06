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

## Links

- [FAQ](http://e8vm.net/faq.html)
- [Motivation](https://github.com/e8vm/e8/wiki/Motivation)

## Install

```
$ go get e8vm.net/e8
```

## Design Thoughts

- [RISC](https://github.com/e8vm/e8/wiki/RISC-Specification): The MIPS-like
  simple instructions set that `e8` uses only has less than 40 instructions,
  which means `e8` CPU will be very easy to port. In fact, I have already
  ported the core to Javascript: [`e8js`](https://github.com/h8liu/e8js).
- [System Page](https://github.com/e8vm/e8/wiki/Page-0:-System-page): all IOs
  in `e8` will be memory mapped, so there is no need for special instructions
  like `in` and `out`. Basic system functionality will be mapped to Page 0.
  Future fancy hardware will be mapped to the following small-id pages in the
  address space.
- [OS-related features](https://github.com/e8vm/e8/wiki/Interrupts-and-Operating-System):
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
  [More thoughts here](https://github.com/e8vm/e8/wiki/Thoughts-on-Language).
