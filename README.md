# goply
A pure go lexer and parser generator library.
goply stands for **g**o **p**arser **l**ex **y**acc.

It was inspired by [rply](https://github.com/alex/rply)
and David Beazley's [PLY](https://github.com/dabeaz/ply)
both of which are excellent parser and lexer generator
libraries for [python](https://www.python.org/).

For those of you not familiar with the *nix tools _lex_ and _yacc_, 
_lex_ is the lexer generator and _yacc_ is the parser generator.
So having a _parser_ in front of them in the acronym probably doesn't
make any sense, but I liked the name goply and hence stuck with it,
redefining the _p_ from _python_ to _parser_, though _yacc_ itself
is the parser generator and the acronym expansion does not
make sense, so ignore it ;)  

## Making a lexer
Making a lexer with goply is as simple as it gets,
only a single line is required:
> lex := goply.NewLexer(source)

where source is a string containing the source contents.

Now for the lexer to actually work, lexical rules need to be
added. A lexical rule is a mapping from a token to a _regular
expression_ or _regex_. A token identifies the type or class of the match that will
performed using the _regex_.

Say for the lexer to recognise numbers we can add the
following rule:
> lex.AddRule("\<number\>", "[0-9]+")

The first argument is the _token type_ and the second argument
is the _regex_. Note that we could have also done
the following:
> lex.AddRule("NUMBER", "[0-9]+")

or even the following:
> lex.AddRule("INTEGER","[0-9]+")

Hence, the first argument is only to give an unique name to
the class of _regex_ the lexer will try match. In this
case its an integer value.

These _token type's_ will actually be used by the parser in order
to denote terminal tokens in the grammar.

## Lexer Gotcha's
* The lexer uses [regexp package](https://golang.org/pkg/regexp/) from the go standard library. 
* The lexer will try to match the first rule first and then the second and so on.

Say you are trying to develop a programming language and
it has a keyword called _var_ and also variable names. So you
the lexer rules as follows:
> lex.AddRule("\<identifier\>","[A-Za-z_][A-Za-z0-9_]+") // your typical variable names  
lex.AddRule("\<var_keyword\>","var") // var keyword

The lexer would end up giving you tokens that does **not** have
_var_ at all, because _var_ is also a valid _\<identifier\>_ !!

So you must add general case rules like _\<identifier\>_ after
you add rules for all the specific cases.

* The lexer will **ignore** redeclaration's of the same _regex_ rule
with **different** _token type_

* The lexer will **discard** all previous declaration's
if a redeclaration is done with the **same** _token type_.

* The lexer **will** panic if it cannot compile the regular
expression given.

## Making a parser
The parser is yet to be implemented. see [issue #2](https://github.com/nayas360/goply/issues/2)

## License
This project is licensed under MIT License - see [LICENSE](LICENSE) for a copy.