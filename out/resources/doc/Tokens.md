# Language Tokens

List of the Language Tokens to easy the understanding of the language, and it's syntax.


### Value Tokens

These are tokens that represents a value or an alias to any value.

- **UNSIGNED_NUMBER:** A Whole number. An integer number without negative values. Ex. `0` or `45`.

- **SIGNED_NUMBER:** A negative integer. Ex: `-5` or `-40`.

- **FLOAT_NUMBER:** Any number positive or negative that has a decimal point. Ex: `6.7736` or `-0.233`.

- **TEXT:** A text surrounded by double quotes (" ") that can contains any symbol inside. Ex: `"Hello World"`
  - As usual, escaped chars and escaped double quotes are accepted. 
  Just use a backslash to escape anything, including the backslash itself.

- **IDENTIFIER:** A name for variable or constants. It must start with a letter or underscore followed by 
any quantity of numbers, letters or underscores and not be a reserved keyword. Ex: `a`, `fj99`, `_45g` or `ex_Ample`.


### Reserved keywords

These tokens are words that have a special meaning on the language.

- **BOOLEAN:** `true` or `false`. They are also a boolean value, and so, they can be associated to an identifier.

- **TYPE:** `bool`, `byte`, `int`, `uint`, `float` or `string`. Used to type variables or constants.

- **ANY:** `any` is a special type that can only be used on variables. It can represent any value. 

- **CONST:** `const` keyword for constants declaration.

- **LET:** `let` keyword for variables declaration.

- **DO:** `do`. Used in do-while loops.

- **WHILE:** `while`. Can be used in both while or do-while loops.

- **FOR:** `for`. Used in for loops.

- **IF:** `if`. For if statements.

- **ELSE:** `else`. For else clauses. 

- **SWITCH:** `switch`. Used in switch statements.

- **CASE:** `case`. the keyword for a case on the switch statement. Like in golang, the cases don't need a break;

- **DEFAULT:** `default`. A special switch case that executes when none of the other cases match; 


#### Event keywords

These are keywords used to handle the language events.

- **ON:** `on`. Keyword to declare a response for an event.

- **SHARE:** `share`. Keyword to dispatch an event.

- **FINISH:** `finish`. Ends an event. Works like the return keyword on void methods in other languages.

- **MODULE:** `module`. Defines a butterfly module.

- **MESSAGE:** `message`. Declares a message for an event.


### Symbols

Special symbols for the language, used for operations or segregate values. 

- **DELIMITER:** a semicolon (`;`). Used to end commands.

- **INCREMENT:** `++` or `--`.

- **NOT:** An exclamation mark (`!`). Used To negate logic sentences.

- **LOGIC:** Logic operators for logical and (`&&`) and for logical or (`||`).

- **ARITHMETIC:** Arithmetic operators for addition (`+`), subtraction (`-`), multiplication (`*`) and division (`/`).

- **COMPARATOR:** Comparison operators, like equals (`==`), not equals (`!=`), greater than (`>`) lower than (`<`), 
grater or equals (`>=`) and lower or equals (`<=`).

- **ASSIGN:** The assign token (`=`) assigns a value to a constant or variable. 

- **EXP_ASSIGN:** The expression assign token is used on reassignments, when any arithmetic token
  is placed before the assign token to apply the operation before assignment.
  Ex: `value += 45` (as a shortcut to `value = value + 45`).

- **COMMA:** The comma separator (`,`).

- **COLON:** A Colon (`:`).


### Groupers

These Tokens are used to contain conditions, values, commands, etc...

- **OP_CURLY:** An opening curly bracket (`{`).

- **CL_CURLY:** A closing curly bracket (`}`).

- **OP_PARENT:** An opening parenthesis (`(`).

- **CL_PARENT:** A closing parenthesis (`)`).

- **OP_SQUARE:** An opening square bracket (`[`).

- **CL_SQUARE:** A closing square bracket (`]`).