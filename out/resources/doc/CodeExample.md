
# Writing a Program in Butterfly

The Butterfly files use the *.bf* extension to identify what files can be compiled. The specification of 
the Butterfly Programming Language can be found [here](./out/resources/doc/Grammar.md).

Butterfly is an event-driven language with modules that communicate with each other through events. A simple 
Hello World can be seen bellow:

```Butterfly
module Example {

  on "Start" () {
    share "Sys Println" ({
      text: "Hello World"
    });
  }

}
```

On the example is defined a module that responds the event `"Start"`, that is the first event called on the 
Butterfly Runtime. When the event is called, the designed code will run, trigger another event named `"Sys Println"`
while sends a message with the text to be printed.

## Basics

Butterfly is a *C-like* language, and so, most language structures works the same as C, Java and JavaScript,
making easy to learn for anyone who programs in those languages.

### Variables and Constants 

The variable and constant declaration for Butterfly follows the same pattern as TypeScript, allowing explicit and 
implicit typing:

```typescript
let explictType: int;
let implicitType = 0;
const constant = 10;
````

Please do note that the semicolons are **required** on Butterfly language.

Constants can only be declared as the language basic types:

- `bool`: binary type that can assume the value `true` or `false`; 
- `byte`: an unsigned integer that can assume values between 0 and 255;
- `uint`: positive integer numbers or 0;
- `int`: any integer positive or negative;
- `float`: float point numbers;
- `string`: text values surrounded by double quotes (");

Variables can be of any type listed above, but there are some types that only variables can assume:

- `any`: Any is an unknown type at compile time. Its use is not advised, but can be useful when working with 
[messages](#events-and-messages). When assigning a message field without asserting a type, the type any will 
be inferred;
- *array*: defined by adding `[i]` to a type, arrays can store multiple indexed values. The 'i' represents the 
length of the array and is defined by a variable or an unsigned integer.
  - Ex1. Defining an array for integers: `let a: int[3] = int[3]{0, 1, 2}`;
  - Ex2. Syntax sugar for an array of type `any`: `let a = [0, "a", 0.35]`. 

- *map*: defined as `[keyType : Valuetype]`, maps can also store multiple values, but each value is identified by 
a key. The key type follows the same restriction as the constants, and must be a basic type. The value type can be 
of any presented type, including map.
  - Ex: `let a: [string : int] = [string:int]{ "one": 1, "two": 2 }`.

### If, else and switch

For flow control, Butterfly has the if/else command that works as any other C-like language:

```javascript
if (condition) {
    // Do stuff here when condition is true
} else {
    // Do Stuff when condition is false
    // The else clause is optional
}
```

The Butterfly language also allow to chain the else clause with another if clause:

```javascript
if (condition) {
    // Do this
} else if (anotherCondition) {
    // Do this instead
} else {
    // Both conditions failed
}
```

The if chain can go on and on, but, instead of chain countless if clauses, you can use the `switch` command:

```javascript
switch (variable) {
  case "this": 
      // Do this
  case "that":
      // Do this instead
  case "those":
      // Do this, not the other two
  default:
      // All conditions failed
}
```

### Loops

Butterfly uses the same three loops from C: `while`, `do-while` and `for`.

While loops will repeat something while a condition is true:
```javascript
while (condition) {
    // will repeat this while condition is true
}
```

Do-while loops follows the same logic, but the code block will execute at least one time, while the `while` loop
can be ignored:

```javascript
while (false) {
    // will never be executed
}

do {
    // will execute once
} while (false);
```

For loops will iterate while the condition is true, like the `while` loop, but you can also instance a variable and
define what will change in each iteration:

```typescript
for (let i = 0; i < 10; i++) {
    // repeat this changing i in each iteration
}
```

### Events and Messages

Butterfly is an event-driven language, and so, all the modules communicate with each other by events and messages.

To define an event response, the keyword `on` is used defining the event that will be answered, the received message 
and the code that will be executed for the event:

```Butterfly
on "Event" (msg) {
    // code to be executed
}
```

Message follows the same syntax as a JS object, and, unlike variables and constants, can only be declared inside an
event response.

To assign some message values to typed variables, a type assertion can be done that will only be validated at execution 
time.

```Butterfly
module Example {

    let globalVar: int;
    // only variables and constants can be declared here
    
    on "Event" () {
        // here you can declare a message
        message msg = {
            example: 10,
        };
        globalVar = msg.example.(int); // type assertion
    }
```

Event responses can share new events to create an event chain, communicating with the event through the declared
messages:

```Butterfly
on "Event" () {
    message msg = {
        example: [10, 45, 12]
    };
    share "Another Event" (msg);
}

```

The `share` command can also share a text contained in a variable, allowing to share a received event:

```Butterfly
on "Event" (msg) {
    // Do things
    let callback = msg.callback.(string);
    share callback ({});
}
```

The keyword `finish` can be used to end a response early:

```Butterfly
on "Event" (msg) {
    let value: int = msg.value.(int);
    if (value > 10) {
        finish;
    }
    share "Event" ({value: value + 1});
}
```

## Concurrency

Fired events can have multiple answers on different modules. All the event responses will be triggered at once.
Each response will run on its own thread and receiving a copy of the shared message.

Now, consider we have these two modules on a directory:

`ModuleA.bf`: 
```
module ModuleA {

    on "Start" () {
        for (let i = 1; i <= 10; i++) {
            share "Sys Println" ({
                text: ["ModuleA says:", i]
            });
        }
    }
}
```

`ModuleB.bf`:
```
module ModuleB {

    on "Start" () {
        for (let i = 1; i <= 10; i++) {
            share "Sys Println" ({
                text: ["ModuleB says:", i*2]
            });
        }
    }
}
```

Both modules have a defined response to the `"Start"` event, and so, they will both run at the
same time when the program starts, printing the text on an unpredictable order.


A recursive approach can also be used to have the same result:

`index.bf`:
```
module Recurser {

    on "Start" () {
        share "Recurse" ({ iter: 1 });
    }
    
    on "Recurse" (msg) {
        let i = msg.iter.(int);
        if (i > 10) {
            // finishes this response here
            // much like a return on a void function
            finish;
        }
        share "Say" ({ value: i });
        share "Recurse"({ iter: i++ });  
    }
}
```

`ModuleA.bf`:
```
module ModuleA {

    on "Say" (msg) {
        share "Sys Println" ({
            text: ["ModuleA says:", msg.value]
        });
    }
}
```

`ModuleB.bf`:
```
module ModuleB {

    on "Say" (msg) {
        share "Sys Println" ({
            text: ["ModuleB says:", msg.value * 2]
        });
    }
}
```

With this approach, the code rely less on sequential code (loops) and uses more the events that are
the core of the Butterfly Language.

## Runtimes

The Butterfly has the capability to switch runtimes while compiling. Each Runtime can handle how the
modules subscribe to events and how those events are fired.

Till now, the Butterfly compiler has two runtimes:

- **Immediate Runtime**: This runtime Starts a thread for each response on the instant that an event
is fired. Here, each response is running simultaneously with themselves and with responses of
another events as well. There is no difference on this runtime between the loop or the recursive 
approach presented before; 

- **Queue Runtime**: This runtime adds the fired event to an event queue, that executes one event at time.
When the event is executed, then the threads for the responses are started, making the responses execution
simultaneously only between themselves, not with responses of other events. While using this runtime, the
recursive approach is recommended, as the loop approach locks the queue flow while running, making the 
recursion better for the queue flow rate. To use it use the `-runtime=queue` flag.