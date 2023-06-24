%{
package parser

import(
	"github.com/Hilson-Alex/Butterfly/src/compiler/checker"
)

%}


%union{
	content 		string
//	currentToken 	token
	parsed 			string
//	tokType			checker.BFType
	scope 			*checker.BFScope
	result			*string
}

%type <parsed>
expression logicOrComparison logicExpression comparison logicValue
valueAccess properties varAddr addrOrProperty mathExpression mathItem
baseValue baseConstValue value declarations declaration constDeclaration
constType varDeclaration varAssign typeDef typeNote optAssign arrNotation
arrValue typeDef mapType eventDeclaration eventOrValuesDeclaration
eventOrValuesDeclarations eventParameter messageContent commands optElse
optReassign mapValue commandBlock command loop for while doWhile whileClause
forAssign reassignment optReassign updateValue ifStatement eventShare
eventName messageParameter messageValue switchStatement cases case caseValue
messageDeclaration messageItem itemValue anyType optTypeCoercion
compoundValue
moduleBody

%token <content>
// VALUES
FLOAT_NUMBER UNSIGNED_NUMBER SIGNED_NUMBER TEXT

// RESERVED KEYWORDS
BOOLEAN TYPE ANY CONST LET DO WHILE FOR IF ELSE SWITCH CASE DEFAULT

// EVENT KEYWORDS
ON SHARE FINISH MODULE MESSAGE

// SYMBOLS
DELIMITER INCREMENT NOT LOGIC ARITHMETIC COMPARATOR ASSIGN EXP_ASSIGN COMMA COLON DOT

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

module: MODULE IDENTIFIER { createModule($<scope>2, $2) } moduleBody {
	saveModule($4, $<result>1)
} | error;

moduleBody:
	OP_CURLY
		declarations
		eventDeclaration
		eventOrValuesDeclarations
	CL_CURLY
	{
		$$ = blJoin($2,$3,$4)
	};

declarations: /*empty*/ 		{$$ = ""}
	| declarations declaration 	{$$ = blJoin($1,$2)};

declaration: constDeclaration DELIMITER {$$ = $1}
 	| varDeclaration DELIMITER  		{$$ = $1}
 	| error DELIMITER			 		{$$ = ""};

eventOrValuesDeclarations: /*empty*/ {$$ = ""}
	| eventOrValuesDeclarations eventOrValuesDeclaration {$$ = blJoin($1,$2)};

eventOrValuesDeclaration: declaration {$$ = $1}
	| eventDeclaration {$$ = $1};

// Consts

constDeclaration: CONST IDENTIFIER constType ASSIGN baseConstValue {
//	addConst($<scope>$,$2,$5,$3)
	$$ = wsJoin($1,$2,$3,$4,$5)
};

constType: /*empty*/ 	{$$ = ""}
	| COLON TYPE 		{$$ = getType($2)};

baseConstValue: FLOAT_NUMBER 	{$$ = $1}
	| UNSIGNED_NUMBER 			{$$ = $1}
	| SIGNED_NUMBER 			{$$ = $1}
	| TEXT 						{$$ = $1}
	| BOOLEAN 					{$$ = $1};

// Variables

varDeclaration: LET IDENTIFIER varAssign {
	$$ = wsJoin("var",$2,$3)
};

varAssign: typeNote optAssign 	{$$ = wsJoin($1, $2)}
	| ASSIGN value 				{$$ = wsJoin($1, $2)};

optAssign: /*empty*/ 	{$$ = ""}
	| ASSIGN value 		{$$ = wsJoin($1, $2)};

typeNote: COLON typeDef {$$ = $2};

typeDef: anyType arrNotation {$$ = $2 + getType($1)} | mapType {$$ = $1} ;

anyType: TYPE {$$ = $1} | ANY {$$ = "interface{}"};

arrNotation: /*empty*/ 									{$$ = ""}
	| arrNotation OP_SQUARE UNSIGNED_NUMBER CL_SQUARE 	{$$ = concat($1,"[",$3,"]")};

mapType: OP_SQUARE TYPE typeNote CL_SQUARE arrNotation {
	$$ = concat($5,"map[", getType($2),"]",$3)
};

value: typeDef OP_CURLY compoundValue CL_CURLY 	{$$ = concat($1,"{", $3, "}")}
	| expression 							{$$ = $1}
	| OP_SQUARE arrValue CL_SQUARE 		 	{$$ = concat("[]interface{}","{",$2,"}")};

compoundValue: mapValue {$$ = $1} | arrValue {$$ = $1};

mapValue: baseValue COLON value 			{$$ = concat($1,":",$3)}
	| mapValue COMMA baseValue COLON value	{$$ = concat($1,", ",$3,":",$5)};

arrValue: /*empty*/ {$$ = ""}
	| value {$$ = $1}
	| arrValue COMMA value {$$ = concat($1,",",$3)};

// Expression

expression: logicOrComparison	{$$ = $1}
	| mathExpression			{$$ = $1};

logicExpression: logicOrComparison LOGIC logicOrComparison
	{$$ = wsJoin($1,$2,$3)}
	| NOT logicValue %prec NEGATED
	{$$ = concat("!(",$2,")")}
	| NOT OP_PARENT logicExpression CL_PARENT %prec NEGATED
	{$$ = concat("!(",$3,")")}
	| OP_PARENT logicExpression CL_PARENT
	{$$ = concat("(",$2,")")}
	| logicValue
	{$$ = $1};

logicOrComparison: logicExpression	{$$ = $1}
	| comparison					{$$ = $1};

logicValue: BOOLEAN {$$ = $1}
	| valueAccess 	{$$ = $1};

comparison: mathItem COMPARATOR mathItem {$$ = wsJoin($1,$2,$3)};

mathExpression: mathItem ARITHMETIC mathItem 	{$$ = concat($1,$2,$3)}
	| mathItem SIGNED_NUMBER 					{$$ = concat($1,$2)}
	| OP_PARENT mathExpression CL_PARENT 		{$$ = concat("(",$2,")")}
	| FLOAT_NUMBER 								{$$ = $1}
	| UNSIGNED_NUMBER 							{$$ = $1}
	| SIGNED_NUMBER								{$$ = $1}
	| TEXT 										{$$ = $1};

mathItem: mathExpression 	{$$ = $1}
	| valueAccess 			{$$ = $1};

baseValue: baseConstValue 	{$$ = $1}
	| valueAccess 			{$$ = $1};

valueAccess: IDENTIFIER addrOrProperty {$$ = concat($1,$2)};

addrOrProperty: varAddr 					{$$ = $1}
	| properties optTypeCoercion varAddr	{$$ = concat($1,$2,$3)};

optTypeCoercion: /* empty*/ {$$ = ""}
	| DOT OP_PARENT typeDef CL_PARENT {$$ = concat(".(",getType($3),")")};

properties: DOT IDENTIFIER {$$ = concat("[\"",$2,"\"]")}
	| properties DOT IDENTIFIER {$$ = concat($1,".(map[string]interface{})", concat("[\"",$3,"\"]"))};

varAddr: /*empty*/ {$$ = ""}
	| varAddr OP_SQUARE baseValue CL_SQUARE
	{$$ = concat($1,"[",$3,"])")};

// Events

eventDeclaration: ON TEXT OP_PARENT eventParameter CL_PARENT commandBlock {
	$$ = eventListen($2, $4, $6, $<scope>2)
};

eventParameter: /*empty*/ {$$ = "_"}
	| IDENTIFIER {$$ = $1};

eventShare: SHARE eventName OP_PARENT messageParameter CL_PARENT DELIMITER {
	$$ = eventShare($2, $4, $<scope>1)
};

eventName: TEXT {$$ = $1} | valueAccess {$$ = $1};

messageParameter: IDENTIFIER {$$ = $1} | messageValue {$$ = $1};

messageDeclaration: MESSAGE IDENTIFIER ASSIGN messageValue DELIMITER {
	$$ = wsJoin("var",$2,$3,$4)
};

messageValue: OP_CURLY messageContent CL_CURLY {$$ = concat("map[string]interface{}{",$2,"}")};

messageContent: /*empty*/ {$$ = ""}
	| messageItem {$$ = $1}
	| messageContent COMMA messageItem {$$ = concat($1,",",$3)};

messageItem: IDENTIFIER COLON itemValue {
	$$ = concat("\"",$1,"\": ",$3)
};

itemValue: value {$$ = $1} | messageValue {$$ = $1};

// Commands

commandBlock: OP_CURLY commands CL_CURLY {$$ = blJoin("{",$2,"}")} | error CL_CURLY {$$ = "{}"};

commands: /*empty*/ {$$ = ""}
	| commands command {$$ = blJoin($1,$2)};

command: declaration 			{$$ = $1}
	| loop 						{$$ = $1}
	| ifStatement 				{$$ = $1}
	| switchStatement			{$$ = $1}
	| eventShare 				{$$ = $1}
	| messageDeclaration 		{$$ = $1}
	| reassignment DELIMITER 	{$$ = $1}
	| FINISH DELIMITER 			{$$ = "return"};

// Loops

loop: for 		{$$ = $1}
	| while 	{$$ = $1}
	| doWhile	{$$ = $1};

for: FOR OP_PARENT forAssign DELIMITER logicOrComparison DELIMITER optReassign CL_PARENT commandBlock {
	$$ = concat("for ",$3,";",$5,";",$7," ",$9)
};

forAssign: LET IDENTIFIER ASSIGN value {$$ = wsJoin($2,":=",$4)}
	| optReassign {$$ = $1};

while: whileClause commandBlock {$$ = wsJoin("for",$1,$2)};

doWhile: DO commandBlock whileClause DELIMITER {
	$$ = concat("for {",$2," if !(",$3,") { break }}")
};

whileClause: WHILE OP_PARENT logicOrComparison CL_PARENT {$$ = $3};

// Conditionals

ifStatement: IF OP_PARENT logicOrComparison CL_PARENT commandBlock optElse {
	$$ = wsJoin("if",$3,$5,$6)
};

optElse: /*empty*/ {$$ = ""}
	| ELSE ifStatement {$$ = "else " + $2}| ELSE commandBlock {$$ = "else " + $2};

switchStatement: SWITCH OP_PARENT valueAccess CL_PARENT OP_CURLY cases CL_CURLY {
	$$ = wsJoin("switch",$3,"{\n",$6,"\n}")
};

cases: case {$$ = $1} | cases case {$$ = blJoin($1,$2)};

case: caseValue COLON commands {
	$$ = concat($1,":",$3)
};

caseValue: CASE value 	{$$ = wsJoin("case",$2)}
	| DEFAULT 			{$$ = "default"};

// Reassign

optReassign: /*empty*/ {$$ = ""}
	| reassignment {$$ = $1};

reassignment: valueAccess updateValue {$$ = $1 + $2};

updateValue: ASSIGN value {$$ = $1 + $2}
	| EXP_ASSIGN value {$$ = $1 + $2}
	| INCREMENT {$$ = $1};

%%