# Butterfly Grammar

This document is a definition of the Butterfly syntax specified by a variation of the 
[EBNF Notation Used to define the XML Grammar](https://www.w3.org/TR/REC-xml/#sec-notation)
that combine the EBNF notation with regular Expressions.

## Notation

The initial symbol of the Butterfly grammar is the module rule.

The grammar rules define a symbol as follows:

    symbol ::= Expression


To build the expressions, some operators can be used to match the valid strings:


 - `"string", 'string'`: matches a literal string inside the quotes;
 - `[a-z]`: matches any character on the specified range;
 - `[^a-z]`: matches any character outside the specified range;
 - `[abc]`: matches any of the enumerated characters;
 - `[^abc]`: matches any not enumerated character;
 - `(expression)`: treats the *expression* as a unit;
 - `A?`: optional A; A or nothing;
 - `A*`: zero or more sequential occurrences of A;
 - `A+`: one or more sequential occurrences of A;
 - `A B`: A followed by B;
 - `A | B`: matches A or B;

## Grammar

Module ::= "module" Identifier ModuleBody;

Identifier ::= \[a-zA-Z_]\[a-zA-Z0-9_]*;

ModuleBody ::= "{"
    Declaration*
    EventIdentifier
    (Declaration | EventIdentifier)*
"}";

Declaration ::= (ConstIdentifier | VarIdentifier) ";";

ConstIdentifier ::= "const" Identifier (":"  Type)? "=" BaseConstValue;

BaseConstValue ::= FloatNumber | UnsignedNumber | SignedNumber | Text | Boolean;

Boolean ::= "true" | "false";

UnsignedNumber ::= \[0-9]+;

SignedNumber ::= "-" UnsignedNumber;

FloatNumber ::= (SignedNumber | UnsignedNumber) "." UnsignedNumber;

Text ::= '"' (\[^"\n\] | '\' ('"' | \[A-Za-z]) | '\\')* '"';

VarIdentifier ::= "let" Identifier (TypeNote ( "=" Value)? | "=" Value);

TypeNote ::= ":" (AnyType ArrNotation | MapType);

Type ::= "bool" | "byte" | "int" | "uint" | "float" | "string";

AnyType ::= Type | "any";

ArrNotation ::= ("\[" UnsignedNumber "]")*;

MapType ::= "\[" Type TypeNote "]" ArrNotation;

Value ::= (AnyType ArrNotation | MapType) "{" (MapValue | ArrValue) "}" | Expression | "\[" ArrValue "]";

MapValue ::= BaseValue ":" Value ("," BaseValue ":" Value)*;

ArrValue ::= (Value ("," Value)*)?;

Expression ::= LogicOrComparison | MathExpression | Text ("+" Text)*;

LogicOrComparison ::= LogicExpression | Comparison;

LogicExpression ::= LogicOrComparison LogicOp LogicOrComparison
    | "(" LogicExpression ")"
    | LogicValue
    | "!" LogicValue
    | "!" "(" LogicExpression ")";

LogicOp ::= "&&" | "||";

LogicValue ::= Boolean | ValueAccess;

Comparison ::= MathItem Comparator MathItem;

Comparator ::= \[=!><]"=" | ">" | "<";

MathItem ::= MathExpression | ValueAccess;

MathExpression ::= MathItem ArithmeticOp MathItem
    | MathItem SignedNumber 					
    | "(" MathExpression ")" 		
    | FloatNumber 								
    | UnsignedNumber 							
    | SignedNumber;

ArithmeticOp ::= "+" | "\" | "-" | "*" | "/" | "%";

BaseValue ::= BaseConstValue | ValueAccess;

ValueAccess ::= Identifier Properties? ("\[" BaseValue "]")*;

Properties ::= "." Identifier ("." Identifier)* ("." "(" Type ")")?;

EventIdentifier ::= "on" Text "(" Identifier? ")" CommandBlock;

EventShare ::= "share" (Text | ValueAccess) "(" (Identifier | MessageValue) ")" ";";

MessageIdentifier ::= "message" Identifier "=" MessageValue ";";

MessageValue ::= "{" (MessageItem ("," MessageItem)* )? "}";

MessageItem ::= Identifier ":" (Value | MessageValue);

CommandBlock ::= "{" Command* "}";

Command ::= Declaration 
    | Loop 
    | IfStatement 
    | SwitchStatement 
    | EventShare 				
    | MessageIdentifier 		
    | Reassignment ";" 	
    | "finish" ";";

Loop ::= For 		
    | While 	
    | DoWhile;

For ::= "for" "(" ("let" Identifier "=" Value | Reassignment)? ";" LogicOrComparison ";" Reassignment? ")" CommandBlock;

While ::= WhileClause CommandBlock ;

DoWhile ::= "do" CommandBlock WhileClause ";";

WhileClause ::= "while" "(" LogicOrComparison ")";

IfStatement ::= "if" "(" LogicOrComparison ")" CommandBlock ("else" (IfStatement | CommandBlock))?;

SwitchStatement ::= "switch" "(" ValueAccess ")" "{" Case+ "}";

Case ::= ("case" Value | "default") ":" Command*;

Reassignment ::= ValueAccess (("=" | ExpAssign) Value | Increment);

ExpAssign ::= "+=" | "\=" | "-=" | "*=" | "/=" | "%=";

Increment ::= "++" | "--";