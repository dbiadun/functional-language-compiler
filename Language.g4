grammar Language;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
ASSIGN: '=';
LT: '<';
GT: '>';
LTE: '<=';
GTE: '>=';
EQ: '==';
NEQ: '/=';
AND: '&&';
OR: '||';
PARENS: '()';
PAREN_LEFT: '(';
PAREN_RIGHT: ')';
BRACE_LEFT: '{';
BRACE_RIGHT: '}';
COMMA: ',';
SEMICOLON: ';';
VERTICAL_BAR: '|';
DOUBLE_COLON: '::';
IMPORT: 'import';
DATA: 'data';
DO: 'do';
CASE: 'case';
OF: 'of';
ARROW_RIGHT: '->';
ARROW_LEFT: '<-';
DEC: [0-9]+;
HEX: '0x' [0-9a-fA-F]+;
CHAR: '\''.'\'';
STRING: '"'.*?'"';
VARID: [a-z][a-zA-Z0-9_]*;
CONID: [A-Z][a-zA-Z0-9_]*;
WHITESPACE: [ \r\n\t]+ -> skip;
COMMENT: '{-' .*? '-}' -> channel (HIDDEN);
LINE_COMMENT: '--' ~[\r\n]* -> channel (HIDDEN);

// Rules
start : topdecls EOF;

topdecls
    : topdecl? (SEMICOLON topdecl)* SEMICOLON? # TopDeclsList
    ;

topdecl
    : IMPORT STRING # ImportTopDecl
    | DATA simpletype ASSIGN constrs # DataTopDecl
    | decl # FunTopDecl
    ;

simpletype
    : CONID VARID* # DataType
    ;

constrs
    : constrdef (VERTICAL_BAR constrdef)* # ConstrList
    ;

constrdef
    : CONID atype* # ConstrType
    ;

decl
    : gendecl # FunTypeDecl
    | funlhs rhs # FunDecl
    | pat rhs # VarDecl
    ;

gendecl
    : vars DOUBLE_COLON type # TypeSignature
    ;

vars
    : VARID (COMMA VARID)* # VarList
    ;

type
    : btype (ARROW_RIGHT btype)* # FunType
    ;

btype
    : atype+ # TypeApp
    ;

atype
    : CONID # ConType
    | VARID # VarType
    | PAREN_LEFT type (COMMA type)+ PAREN_RIGHT # TupleType
    | PAREN_LEFT type PAREN_RIGHT # ParenType
    | PARENS # UnitType
    ;

funlhs
    : VARID apat+ # DeclLhs
    ;

rhs
    : ASSIGN exp # DeclExp
    ;

exp
    : fexp # EFun
    | DO BRACE_LEFT stmts BRACE_RIGHT # EDo
    | CASE exp OF BRACE_LEFT alts BRACE_RIGHT # ECase
    | e1=exp op=(MUL|DIV) e2=exp # EMulDiv
    | e1=exp op=(ADD|SUB) e2=exp # EAddSub
    | e1=exp op=(LT|GT|LTE|GTE|EQ|NEQ) e2=exp # EComp
    | e1=exp op=(AND|OR) e2=exp # ELogical
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

stmts
    : stmt* exp SEMICOLON? # StmtsList
    ;

stmt
    : exp SEMICOLON # SExp
    | pat ARROW_LEFT exp SEMICOLON # SAssign
    ;

alts
    : alt (SEMICOLON alt)* SEMICOLON?  # Alternatives
    ;

alt
    : pat ARROW_RIGHT exp # Alternative
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
    : integer # Int
    | CHAR # Char
    | STRING # String
    ;

integer
    : DEC # Dec
    | HEX # Hex
    ;
