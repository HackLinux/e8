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

