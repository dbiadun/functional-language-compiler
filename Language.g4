grammar Language;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expr EOF;

expr
   : e1=expr op=('*'|'/') e2=expr # EMulDiv
   | e1=expr op=('+'|'-') e2=expr # EAddSub
   | NUMBER                       # Number
   ;