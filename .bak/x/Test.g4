// Test.g4
grammar Test;

WHITESPACE: (' ' | '\r' | '\n' | '\t')+ -> skip;

LB: '(';
RB: ')';
PLUS: '+';
MINUS: '-';

start : exp+ EOF;

exp : cmd=(PLUS|MINUS) expr | sub ;

sub : LB expr RB ;

expr : exp0 (';' exp0)* (';')?;

exp0 : exp1 | exp2;

exp1 : 'A' | 'B' | 'C';

exp2 : 'D' | 'E' | 'F';
