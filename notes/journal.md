**2014.3.9**

So I have a pretty-much working but not very well tested lexer now.

Now I started to build my abstract syntax tree. It's pretty tedious.

And I also need to formalize the spec of my language.

I thought golang is pretty simple to implementing; I was wrong. Just parsing
the abstract syntax tree needs to take some extra care on programming. Well, I
write the lexer and parser all by hand, so that somehow makes things harder,
but not so hard. Anyway, golang is certainly simpler than C++, but it is not
that simple. Scheme might be the much simpler to implement.

And after that, if I follow the traditional way of writing a compiler. I also need
to write an intermediate representation. This IR probably won't need a parser,
rather, it will have some data structure and API that can use to easily construct
and scan over. So I also do not need to write another scanner/parser.

Maybe I really should have a deep look at how other SSAs (like go's and llvm's) are
implemented.

I think I probably should implement the SSA/TAC first. Rather than directly jump
to the entire language.

Another way might be first try to connect our RISC to LLVM backend, although
I am not sure if it is easy to do.

**2014.3.7**

Finished most of the part of the lexer. Now a program can be parsed into tokens.
The next step is to parse the tokens into ASTs, and once we have that, we can
start to emit SSAs.

I still remember the days back in college, where I tell myself I will never
write a compiler again (because it was too much tedious coding). Well, it is a 
long journey...

Anyway, at least we have a lexer now.

**2014.3.3**

Trying travis.yml

**2014.2.26**

Have a simple working assembler. TODO: add some constants, and string decl. like:

    ; data segments
    .string msg "Hello, world.\n\000'
    .uint8 endl '\n'
    .int16 start 0x2000
    .int32 magic 0x32323322

    .func main ; just a label that declares a label namespace, nothing special
    ; then our code here

