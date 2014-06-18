[![Build Status](https://travis-ci.org/h8liu/e8.png?branch=master)](https://travis-ci.org/h8liu/e8)

**Road Plan**

- Build a MIPS-like very simple virtual machine (*mostly done*)
- Build an assembler (*mostly done*)
- Build a SSA/TCA-like language `sand`. This serves as the intermediate language for the
  more fancy compiler. It will also enable some simple optimization (like constant
  propagation and register allocation).
- Build a compiler for a Go-like system language `stay` (*work in progress*)
- Port the assembler and compiler to `stay`.
- Implement the VM in Javascript. (*mostly done*, see [here](https://e8vm.net/pjs) )
- Write a small OS in `stay`, so that is runs in the browser.
- Launch it online.
- Add file system support.

**TODO list**

- Compiler side
  - Package parsing and package aliasing
  - Package namespace building, types
  - Consts
  - Global variables
  - Simple functions and linking
  - Function code generation
  - Simple optimizations
- VM side
  - Event support
  - More sample programs in ASM
  - Web visualization, stepping demo and debugger
  - A simple console with some built-in programs
  - User VM support
  - Fancier console
  - Simple canvas support
  - Tic-tac-toe

**Why**

- This could be an online platform for ACM training
- This could be an online arena/platform for game AI fighting
- This could be used for collaborative programming
- This encourages users to collaborate by assembling modules via clearly
  defined interfaces and test them with testcases.
- Ideally, all the modules and interfaces could be easily reused, since they
  would likely be small, and hence easy for human to read, understand, modify
  and maintain.

See more on my [Motivation](https://e8vm.net/p/wiki/Motivation) page.

**Design**

- [RISC](https://e8vm.net/p/wiki/RISC-Specification): The MIPS-like
  simple instructions set that `e8` uses only has less than 40 instructions,
  which means `e8` CPU will be very easy to port. In fact, I have already
  ported the core to Javascript: [`e8js`](https://e8vm.net/pjs).
- [System Page](https://e8vm.net/p/wiki/Page-0:-System-page): all IOs
  in `e8` will be memory mapped, so there is no need for special instructions
  like `in` and `out`. Basic system functionality will be mapped to Page 0.
  Future fancy hardware will be mapped to the following small-id pages in the
  address space.
- [OS-related features](https://e8vm.net/p/wiki/Interrupts-and-Operating-System): 
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
  [More thoughts here](https://e8vm.net/p/wiki/Thoughts-on-Language).

**Install**

`$ go get e8vm.net/p`

This should grab all `e8` into you `$GOPATH` folder.

**Documents**

Incomplete now. However, you can try type `make doc` under this project folder, 
and visit `http://localhost:8000` for the API document.
