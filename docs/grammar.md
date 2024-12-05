# Provisional grammar

The following is a provisional grammar for the subset of the C programming
language implemented by the compiler.


## Tokens:

- Keywords: `int`, `if`, `else`, `while`, `return`
- Operators: `+`, `-`, `*`, `/`, `=`, `==`, `!=`, `<`, `>`, `<=`, `>=`, `&&`, `||`
- Separators: `;`, `,`, `(`, `)`, `{`, `}`
- Identifiers: `id` (variable and function names)
- Constants: `num` (integer literals)


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

<compound_statement> ::= '{' <declaration_list_opt> <statement_list_opt> '}'

<expression_statement> ::= <expression_opt> ';'

<expression_opt> ::= <expression>
                   | ε

<selection_statement> ::= 'if' '(' <expression> ')' <statement> <else_clause_opt>

<else_clause_opt> ::= 'else' <statement>
                    | ε

<iteration_statement> ::= 'while' '(' <expression> ')' <statement>

<jump_statement> ::= 'return' <expression_opt> ';'

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

<argument_expression_list_opt> ::= <argument_expression_list>
                                 | ε

<argument_expression_list> ::= <assignment_expression>
                            | <argument_expression_list> ',' <assignment_expression>
```