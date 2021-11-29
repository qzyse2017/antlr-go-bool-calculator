
/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

grammar Cal;

userexpr: expr;
intexpr: MYINT | MYINT NUMCOMP MYINT;

expr : NAME 
    | intexpr
    | expr  AND expr
    | expr  OR   expr
    ;
MYINT : [0-9]+;
NAME : [A-Z]+ ;
OR: 'OR' | '|';
AND: 'AND' | '&';

NUMCOMP: '>'
|'<'
|'='
|'>='
|'<='
;

STRINGCOMP: 'CONTAINS'
| 'STRINGEQUAL'
|'STRINGSTARTSWITH'
;

WS: [ \t\n\r] + -> skip;