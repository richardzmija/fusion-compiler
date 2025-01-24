# Grammar

The following is a grammar for the subset of the C programming
language implemented by the Fusion Compiler. The grammar is presented in the EBNF form.

## Non-standard features

For simplicity, the operation for writing to the standard output is represented as
a non-standard keyword. This greatly reduces the number of C language features needed
to be implemented in order to express output operations. Even though inclusion of such a
keyword deviates from the C language specification the syntax remains very similar. For details
see the `<printf_statement>` non-terminal and the regular expression for the string literal token.

Since the declarations are not supported by this subset for simplicity the order of function definitions
in a source file doesn't matter. This means that if a function is called before being defined the compiler
will not raise semantic errors.
## Tokens:

- Keywords: `int`, `if`, `else`, `while`, `return`, `printf`
- Operators: `+`, `-`, `*`, `/`, `=`, `==`, `!=`, `<`, `>`, `<=`, `>=`, `&&`, `||`
- Separators: `;`, `,`, `(`, `)`, `{`, `}`
- Identifiers: `id` (variable and function names)
- Constants: `num` (integer literals), `str` (string literals)

Tokens that are not fixed character sequences are defined using regular expressions according to the C language specification.

## Regular expressions for tokens:

The following are the regular expressions for the tokens which require them in PCRE syntax:

- `id`: [a-zA-Z_][a-zA-Z0-9_]*
- `num`: 0|[1-9][0-9]*
- `str`: "(\\n|[^"\\])*"

## Grammar:

### Program structure

```
<program> ::= <function_definition> { <function_definition> }
```

### Function definition

```
<function_definition> ::= 'int' <id> '(' [ <parameter_list> ] ')' '{' [ <declaration_list> ] [ <statement_list> ] '}'

<parameter_list> ::= <parameter_declaration> {',' <parameter_declaration> }

<parameter_declaration> ::= 'int' <id>
```

### Declarations

```
<declaration_list> ::= { <declaration> }

<declaration> ::= 'int' <init_declarator_list> ';'

<init_declarator_list> ::= <init_declarator> {',' <init_declarator> }

<init_declarator> ::= <id>
                    | <id> '=' <expression>
```

### Statements

```
<statement_list> ::= { <statement> }

<statement> ::= <compound_statement>
              | <expression_statement>
              | <selection_statement>
              | <iteration_statement>
              | <jump_statement>
			  | <printf_statement>

<compound_statement> ::= '{' [ <declaration_list> ] [ <statement_list> ] '}'

<expression_statement> ::= [ <expression> ] ';'

<selection_statement> ::= 'if' '(' <expression> ')' <statement> [ 'else' <statement> ]

<iteration_statement> ::= 'while' '(' <expression> ')' <statement>

<jump_statement> ::= 'return' [ <expression> ] ';'

<printf_statement> ::= 'printf' '(' <str> { ',' <expression> } ')' ';'
```

### Expressions

```
<expression> ::= <assignment_expression>

<assignment_expression> ::= <conditional_expression>
                          | <unary_expression> '=' <assignment_expression>

<conditional_expression> ::= <logical_or_expression>

<logical_or_expression> ::= <logical_and_expression>
                         | <logical_or_expression> '||' <logical_and_expression>

<logical_and_expression> ::= <equality_expression>
                          | <logical_and_expression> '&&' <equality_expression>

<equality_expression> ::= <relational_expression>
                       | <equality_expression> '==' <relational_expression>
                       | <equality_expression> '!=' <relational_expression>

<relational_expression> ::= <additive_expression>
                         | <relational_expression> '<' <additive_expression>
                         | <relational_expression> '>' <additive_expression>
                         | <relational_expression> '<=' <additive_expression>
                         | <relational_expression> '>=' <additive_expression>

<additive_expression> ::= <multiplicative_expression>
                       | <additive_expression> '+' <multiplicative_expression>
                       | <additive_expression> '-' <multiplicative_expression>

<multiplicative_expression> ::= <unary_expression>
                             | <multiplicative_expression> '*' <unary_expression>
                             | <multiplicative_expression> '/' <unary_expression>

<unary_expression> ::= <postfix_expression>
                    | '-' <unary_expression>
                    | '+' <unary_expression>

<postfix_expression> ::= <primary_expression>
                      | <postfix_expression> '(' [ <argument_expression_list> ] ')'

<primary_expression> ::= <id>
                       | <constant>
                       | '(' <expression> ')'

<constant> ::= <num>
			 | <str>

<argument_expression_list> ::= <assignment_expression> { ',' <assignment_expression> }
```
