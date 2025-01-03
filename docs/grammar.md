# Grammar

The following is a grammar for the subset of the C programming
language implemented by the Fusion Compiler. The grammar is presented in the EBNF form.

## Non-standard features

For simplicity, the operation for writing to the standard output is represented as
a non-standard keyword. This greatly reduces the number of C language features needed
to be implemented in order to express output operations. Even though inclusion of such a
keyword deviates from the C language specification the syntax remains very similar. For details
see the `<printf_statement>` non-terminal and the regular expression for the string literal token.

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
<program> ::= <function_list>

<function_list> ::= <function_definition>
                  | <function_definition> <function_list>
```

### Function definition

```
<function_definition> ::= 'int' <id> '(' <parameter_list_opt> ')' '{' <declaration_list_opt> <statement_list_opt> '}'

<parameter_list_opt> ::= <parameter_list>
                       | ε

<parameter_list> ::= <parameter_declaration>
                   | <parameter_declaration> ',' <parameter_list>

<parameter_declaration> ::= 'int' <id>
```

### Declarations

```
<declaration_list_opt> ::= <declaration_list>
                         | ε

<declaration_list> ::= <declaration>
                     | <declaration> <declaration_list>

<declaration> ::= 'int' <init_declarator_list> ';'

<init_declarator_list> ::= <init_declarator>
                         | <init_declarator> ',' <init_declarator_list>

<init_declarator> ::= <id>
                    | <id> '=' <expression>
```

### Statements

```
<statement_list_opt> ::= <statement_list>
                       | ε

<statement_list> ::= <statement>
                   | <statement> <statement_list>

<statement> ::= <compound_statement>
              | <expression_statement>
              | <selection_statement>
              | <iteration_statement>
              | <jump_statement>
			  | <printf_statement>

<compound_statement> ::= '{' <declaration_list_opt> <statement_list_opt> '}'

<expression_statement> ::= <expression_opt> ';'

<expression_opt> ::= <expression>
                   | ε

<selection_statement> ::= 'if' '(' <expression> ')' <statement> <else_clause_opt>

<else_clause_opt> ::= 'else' <statement>
                    | ε

<iteration_statement> ::= 'while' '(' <expression> ')' <statement>

<jump_statement> ::= 'return' <expression_opt> ';'

<printf_statement> ::= 'printf' '(' <str> (',' <expression>)* ')' ';'
```

### Expressions

```
<expression> ::= <assignment_expression>
              | <expression> ',' <assignment_expression>

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
                      | <postfix_expression> '(' <argument_expression_list_opt> ')'

<primary_expression> ::= <id>
                       | <constant>
                       | '(' <expression> ')'

<constant> ::= <num>
			 | <str>

<argument_expression_list_opt> ::= <argument_expression_list>
                                 | ε

<argument_expression_list> ::= <assignment_expression>
                            | <argument_expression_list> ',' <assignment_expression>
```
