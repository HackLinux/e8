Walk language, using suffix `.wlk`

I would like to implement a full Go language under `e8`.

However, as a type safe language, its runtime is too large. To be type safe, there is no pointer arithmetic. So the runtime has to handle all the pointers, has to have garbage collection implemented, and once it has garbage collection, the behavior becomes unpredictable.

The first language I need here is a language that can be used as a system programming language.

One thing good about faithfully implement the go language would be that I can immediately adopt the go std library when it is done.

However, the gap is nevertheless too wide for me. I only wrote a toy compiler when I was in college. I have no extensive experience of programming language. I even have no experience of writing parser and tokenizers. In fact, this has been an area that I do not completely understand. Parsing always seems to be a easy and straightforward task, where there should be parsing programs that are easy to understand. However, parser and grammar design always seems too complicated and hard to understand. Even for the parsing library for go (`go/parser`), it seems that there are still spaces for code refactoring.

So let's just design a C like language with consts, variables, functions, and structs.

It would be really cool to support interface, but I also feel that it will require some complexity on the runtime. So I think I will just put that aside for now.

And let's just go on and support pointer arithmetic, so we will have arrays and nul terminated strings immediately. This will basically be C language with Go grammar. It has namespaces, packages, struct methods, name visibility but no macros. It sounds like a pretty cool language for system programming.

And let's go on and write a parser.