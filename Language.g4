grammar Language;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
PAREN_LEFT: '(';
PAREN_RIGHT: ')';
BRACE_LEFT: '{';
BRACE_RIGHT: '}';
COMMA: ',';
SEMICOLON: ';';
CASE: 'case';
OF: 'of';
ARROW: '->';
INT: [0-9]+;
CHAR: '\''.'\'';
STRING: '"'.'"';
VARID: [a-z][a-zA-Z0-9_]*;
CONID: [A-Z][a-zA-Z0-9_]*;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : exp EOF;

exp
    : fexp # EFun
    | CASE exp OF BRACE_LEFT alts BRACE_RIGHT # ECase
    | e1=exp op=(MUL|DIV) e2=exp # EMulDiv
    | e1=exp op=(ADD|SUB) e2=exp # EAddSub
    ;

fexp
    : fexp aexp # FApp
    | aexp # FArg
    ;

aexp
    : VARID # Var
    | CONID # Constr
    | literal # Lit
    | PAREN_LEFT e=exp PAREN_RIGHT # ParenExp
    | PAREN_LEFT exp (COMMA exp)+ PAREN_RIGHT # Tuple
    ;

alts
    : alt (SEMICOLON alt)*  # Alternatives
    ;

alt
    : pat ARROW exp # Alternative
    ;

pat
    : apat # PatArg
    | CONID apat+ # PatCon
    ;

apat
    : VARID # ApatVar
    | CONID # ApatCon
    | literal # ApatLit
    ;

literal
    : INT # Int
    | CHAR # Char
    | STRING # String
    ;
