grammar C;

// Lexer Rules

// Keywords
INT: 'int';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
RETURN: 'return';
PRINTF: 'printf';

// Operators
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
ASSIGN: '=';
EQ: '==';
NEQ: '!=';
LT: '<';
GT: '>';
LE: '<=';
GE: '>=';
AND: '&&';
OR: '||';

// Separators
SEMI: ';';
COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';

// Identifiers and literals
ID: [a-zA-Z_][a-zA-Z0-9_]*;
NUM: '0' | [1-9][0-9]*;
STR: '"' ( '\\' . | ~["\\] )* '"' ;

// Whitespace and comments
WS: [ \t\r\n]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;


// Parser Rules

program
    : functionDefinition+ EOF
    ;

functionDefinition
    : INT ID LPAREN parameterList? RPAREN compoundStatement
    ;

parameterList
    : parameterDeclaration (COMMA parameterDeclaration)*
    ;

parameterDeclaration
    : INT ID
    ;

declarationList
    : declaration+
    ;

declaration
    : INT initDeclaratorList SEMI
    ;

initDeclaratorList
    : initDeclarator (COMMA initDeclarator)*
    ;

initDeclarator
    : ID (ASSIGN expression)?
    ;

statementList
    : statement+
    ;

statement
    : compoundStatement
    | expressionStatement
    | selectionStatement
    | iterationStatement
    | jumpStatement
    | printfStatement
    ;

compoundStatement
    : LBRACE declarationList? statementList? RBRACE
    ;

expressionStatement
    : expression? SEMI
    ;

selectionStatement
    : IF LPAREN expression RPAREN statement (ELSE statement)?
    ;

iterationStatement
    : WHILE LPAREN expression RPAREN statement
    ;

jumpStatement
    : RETURN expression? SEMI
    ;

printfStatement
    : PRINTF LPAREN STR (COMMA expression)* RPAREN SEMI
    ;

expression
    : assignmentExpression
    ;

assignmentExpression
    : conditionalExpression
    | unaryExpression ASSIGN assignmentExpression
    ;

conditionalExpression
    : logicalOrExpression
    ;

logicalOrExpression
    : logicalAndExpression (OR logicalAndExpression)*
    ;

logicalAndExpression
    : equalityExpression (AND equalityExpression)*
    ;

equalityExpression
    : relationalExpression ((EQ | NEQ) relationalExpression)*
    ;

relationalExpression
    : additiveExpression ((LT | GT | LE | GE) additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*
    ;

multiplicativeExpression
    : unaryExpression ((MULT | DIV) unaryExpression)*
    ;

unaryExpression
    : postfixExpression
    | PLUS unaryExpression
    | MINUS unaryExpression
    ;

postfixExpression
    : primaryExpression
    | postfixExpression LPAREN argumentExpressionList? RPAREN
    ;

primaryExpression
    : ID
    | constant
    | LPAREN expression RPAREN
    ;

constant
    : NUM
    | STR
    ;

argumentExpressionList
    : assignmentExpression (COMMA assignmentExpression)*
    ;
