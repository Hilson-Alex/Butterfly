%{
package parser

%}


%union{

}

%token
// VALUES
FLOAT_NUMBER UNSIGNED_NUMBER SIGNED_NUMBER TEXT

// RESERVED KEYWORDS
BOOLEAN TYPE CONST LET DO WHILE FOR IF ELSE SWITCH CASE DEFAULT

// EVENT KEYWORDS
ON SHARE FINISH MODULE MESSAGE

// SYMBOLS
DELIMITER INCREMENT NOT LOGIC ARITHMETIC COMPARATOR ASSIGN COMMA COLON DOT

// GROUPERS
OP_CURLY CL_CURLY OP_PARENT CL_PARENT OP_SQUARE CL_SQUARE

IDENTIFIER

%left NEGATED
%left LOGIC
%nonassoc COMPARATOR
%left ARITHMETIC
%left SIGNED_NUMBER

%start module

%%

module: MODULE IDENTIFIER moduleBody;

moduleBody:
	OP_CURLY
		declarations
		eventDeclaration
		eventOrValuesDeclarations
	CL_CURLY;

declarations: /* empty*/ | declarations declaration;

declaration: constDeclaration DELIMITER | varDeclaration DELIMITER | error DELIMITER;

eventOrValuesDeclarations: /* empty*/ | eventOrValuesDeclarations eventOrValuesDeclaration;

eventOrValuesDeclaration: declaration | eventDeclaration;

// Consts

constDeclaration: CONST IDENTIFIER constType ASSIGN baseConstValue;

constType: /*empty*/ | COLON TYPE;

baseConstValue: FLOAT_NUMBER | UNSIGNED_NUMBER | SIGNED_NUMBER | TEXT | BOOLEAN;

// Variables

varDeclaration: LET IDENTIFIER varAssign;

varAssign: typeNote optAssign | ASSIGN value;

optAssign: /*empty*/ | ASSIGN value;

typeNote: COLON typeDef;

typeDef: TYPE arrNotation | mapType;

arrNotation: /*empty*/ | arrNotation OP_SQUARE UNSIGNED_NUMBER CL_SQUARE;

mapType: OP_SQUARE TYPE COLON typeDef CL_SQUARE;

value: mapType OP_CURLY mapValue CL_CURLY | expression | OP_SQUARE arrValue CL_SQUARE;

mapValue: baseValue COLON value | mapValue COMMA baseValue COLON value;

arrValue: /*empty*/ | value | arrValue COMMA value;

// Expression

expression: logicOrComparison | mathExpression;

logicExpression: logicOrComparison LOGIC logicOrComparison
	| NOT logicValue %prec NEGATED
	| NOT OP_PARENT logicExpression CL_PARENT %prec NEGATED
	| OP_PARENT logicExpression CL_PARENT
	| logicValue;

logicOrComparison: logicExpression | comparison;

logicValue: BOOLEAN | valueAccess;

comparison: mathItem COMPARATOR mathItem;

mathExpression: mathItem ARITHMETIC mathItem
	| mathItem SIGNED_NUMBER
	| OP_PARENT mathExpression CL_PARENT
	| FLOAT_NUMBER
	| UNSIGNED_NUMBER
        | SIGNED_NUMBER
        | TEXT;

mathItem: mathExpression | valueAccess;

baseValue: baseConstValue | valueAccess;

valueAccess: IDENTIFIER addrOrProperty;

addrOrProperty: varAddr | properties;

properties: DOT IDENTIFIER | properties DOT IDENTIFIER;

varAddr: /*empty*/ | varAddr OP_SQUARE baseValue CL_SQUARE;

// Events

eventDeclaration: ON TEXT OP_PARENT eventParameter CL_PARENT commandBlock;

eventParameter: /*empty*/ | IDENTIFIER;

eventShare: SHARE eventName OP_PARENT messageParameter CL_PARENT DELIMITER ;

eventName: TEXT | valueAccess;

messageParameter: IDENTIFIER | messageValue;

messageDeclaration: MESSAGE IDENTIFIER ASSIGN messageValue DELIMITER;

messageValue: OP_CURLY messageContent CL_CURLY;

messageContent: /*empty*/ | messageItem | messageContent COMMA messageItem;

messageItem: TEXT COLON itemValue;

itemValue: value | messageValue;

// Commands

commandBlock: OP_CURLY commands CL_CURLY ;

commands: /*empty*/ | commands command;

command: declaration
	| loop
	| ifStatement
	| switchStatement
	| eventShare
	| messageDeclaration
	| reassignment DELIMITER
	| FINISH DELIMITER;

// Loops

loop: for | while | doWhile;

for: FOR OP_PARENT forAssign DELIMITER logicOrComparison DELIMITER optReassign CL_PARENT commandBlock;

forAssign: varDeclaration | optReassign;

while: whileClause commandBlock;

doWhile: DO commandBlock whileClause DELIMITER;

whileClause: WHILE OP_PARENT logicOrComparison CL_PARENT;

// Conditionals

ifStatement: IF OP_PARENT logicOrComparison CL_PARENT commandBlock optElse;

optElse: /*empty*/ | ELSE ifStatement | ELSE commandBlock

switchStatement: SWITCH OP_PARENT valueAccess CL_PARENT OP_CURLY cases CL_CURLY;

cases: case | cases case;

case: caseValue COLON commands;

caseValue: CASE value | DEFAULT;

// Reassign

optReassign: /*empty*/ | reassignment;

reassignment: valueAccess updateValue;

updateValue: ASSIGN value | INCREMENT;

%%