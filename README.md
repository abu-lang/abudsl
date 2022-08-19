# AbU DSL
A Domain Specific Language (DSL) for the IoT.

`abudsl` is a distributed language merging the programming simplicity prerogative of Event Condition Action (ECA) rules with a powerful (decentralized) communication and coordination mechanism based on attributes. The language exploits *Attribute-based memory Updates* and it has its roots in the *AbU calculus*[^1].

[^1]: Marino Miculan and Michele Pasqua. "A Calculus for Attribute-Based Memory Updates". In Antonio Cerone and Peter Ölveczky, editors, Proceedings of the 18th international colloquium on theoretical aspects of computing, ICTAC 2021, volume 12819 of Lecture Notes in Computer Science. Springer, 2021.

### Brief introduction

In `abudsl`, IoT devices (comprising sensors, actuators and internal variables) are programmed by means of *rules* following the Event Condition Action programming style:
> **on** *event* **if** *condition* **do** *action*

This reads as: when *event* occurs, if *condition* is verified, then execute *action*. The event can be a change in a sensor or the modification of an actuator, while an action is a list of updates of variables and actuators.

The peculiarity of `abudsl` is that such rules are *distributed*, in the sense that they can act on a distributed network of IoT devices. Another key point is *decentralization*. Indeed, in `abudsl`, devices do not have a global knowledge about the network: communication and coordination is performed by means of an attribute-based interaction. In this respect, `abudsl` exploits *Attribute-based memory Updates*: an attributed-based interaction mechanism which is decentralized and that fits neatly within the ECA paradigm. Indeed, in `abudsl` a device can perform an action on itself (the current device) or on an external (remote) device. But, from the programmer point of view, both types of action are just *memory updates*.

## Basic syntax
In the following, we gently present the grammar for the (basic) syntax of the DSL. Have a look at the [abudsl-grammar](/docs/abudsl-grammar.md) for the complete syntax of `abudsl`. In the [examples](/examples) folder it is possible to find some `abudsl` coding examples.

> **File extension:** the official extension for `abudsl` program source files is `.abu`.

An `abudsl` program consists in a non-empty list of (IoT) devices, equipped with sensors, actuators and internal variables, followed by a list of ECA rules acting on these devices.

### Devices
A ***device*** is of the form:
> *DeviceId* `:` *Description* `{` <br>
      &emsp;&emsp;*ResourceDeclaration* <br>
    &emsp;**[** `where` *BooleanExpression* **]** <br>
  `}`

where *DeviceId* is the name of the device (an alphanumeric string); *Description* is a quoted string describing the device functionality; and *ResourceDeclaration* is a non-empty list of resource (sensors, actuators and internal variables) declarations. A resource can be physical or logical.

A ***physical resource*** declaration can be of the forms:
> ***Input*** &nbsp;&nbsp; `physical input` *Type* *ResourceId* <br>
  ***Output*** &nbsp;&nbsp; `physical output` *Type* *ResourceId* `=` *Expression*

***Input*** physical resources are used to model sensors; while ***Output*** pyshical resources are used to model actuators. The first, are supposed to be read-only; while the latter are supposed to be write-only.

A ***logical resource*** declaration is of the form:
> `logical` *Type* *ResourceId* `=` *Expression*

and it is used to model internal device variables. This kind of resource does not have read/write constraints. Note that, logical and physical output resources have to be declared with an initialization *Expression*, while physical input resources do not.

Each resource is declared with a name *ResourceId* (an alphanumeric string) and a type *Type*. In the basic syntax of `abudsl`, we have the following ***types***:
- `boolean`, for boolean resources like `true` or `false`
- `integer`, for integer resources like `42` or `-42`
- `decimal`, for decimal resources like `3.14` or `-3.14`
- `string`, for (quoted) string resources like `"sTr1nG"`

> **Strings format:** strings can contain spaces and special characters, like `_` (underscore), `\` (backslash), `#` (octothorpe) or `'` (single quote); but they cannot contain the double quote symbol `"`.

Finally, a device can be equipped with an (optional) invariant, introduced after the keyword `where`. The invariant is a boolean expression that the device have to fulfill during execution (no updates violating the invariant are allowed).

Here is a device full example (self-explanatory):
```
hvac : "An HVAC control system" {
    # Resources declaration.
    physical output boolean heating = false
    physical output boolean conditioning = false
    logical integer temperature = 0
    logical integer humidity = 0
    physical input boolean airButton
    logical string node = "hvac"
  where
    # Device invariant.
    not (conditioning and heating)
}
```
where the lines after the keyword `#` are comments.

#### Expressions
In the basic syntax of `abudsl`, an *Expression* can be a *BooleanExpression*, a *NumericExpression* or a *StringExpression*. The definition of expressions is standard and it comprises: boolean operators, like `not` (negation), `and` (conjunction), `or` (disjunction); arithmetic and string operators, like `abs` (absolute value), `+` (addition), `-` (subtraction), `*` (multiplication), `/` (division), `%` (modulo), `::` (concatenation); and comparison operators, like `==` (equal), `!=` (not equal), `<` (less than), `<=` (less than or equal), `>` (greater than), `>=` (greater than or equal). The standard operators composition priority can be overridden by using left `(` and right `)` round brackets.

The detailed grammar of expressions can be found [here](/docs/abudsl-grammar.md#syntax-for-expressions-and-conditions).

#### Referencing rules
As said at the beginning, devices are programmed by means of ECA rules acting on them. A rule can act on multiple devices and a device can be influenced by multiple rules. Hence, each device is suffixed by a list of *RuleId*, namely rule names, after the (optional) keyword `has`. For instance:
```
hvac : "An HVAC control system" {
    # Resources declaration.
    ...
    logical string node = "hvac"
  where
    # Device invariant.
    not (conditioning and heating)
} has cool warm dry stopAir
```
means that the device `hvac` can be affected by the ECA rules named `cool`, `warm`, `dry` and `stopAir`. Note that, the list of rules is optional since a device may not have specific rules acting on it (for instance, when an actuator can only be changed by external devices but its does not impact any other device).

### ECA rules
An ECA ***rule*** is of the form:
> `rule` *RuleId* <br>
      &emsp;&emsp;`on` *Event* <br>
      &emsp;&emsp;**(** *Task* **)<sup>+</sup>**

where *RuleId* is the name of the rule; *Event* is a non-empty space-separated list of resources on which the rule is waiting for changes; and **(** *Task* **)<sup>+</sup>** is a non-empty list of tasks that may be activated when a resource in *Event* changes. Rule names and resources are alphanumeric strings. For instance:
```
rule dry
    on humidity temperature
```
is a rule named `dry` that is waiting for changes in the resources `humidity` and `temperature`.

An ECA rule ***task*** is of the form:
> `for` **[** `all` **]** *Condition* <br>
      &emsp;&emsp;`do` *Action*

where *Condition* is a boolean expression and *Action* is a list of semicolon-separated list of resource assignments. When *Condition* is satisfied, then the assignments in *Action* are performed. For instance:
```
for (2 + 0.5 * temperature < humidity and 38 - temperature < humidity)
    do conditioning = true
```
is a task that turns on the conditioning system (doing `conditionig = true`) when the humidity is above a given threshold (namely when the condition after `for` is true).

The full code of the `dry` rule is then the following:
```
rule dry
    on humidity temperature
    for (2 + 0.5 * temperature < humidity and 38 - temperature < humidity)
        do conditioning = true
```

#### External tasks

ECA rule tasks may act on external devices, by using the (optional) modifier `all`. In this case, the condition and the action in the task may reference resources on external devices, by prefixing them with the `ext.` keyword. For instance:
```
rule notifyTemp
    on temperature
    for all (ext.node == "hvac")
        do ext.temperature = this.temperature
```
is a rule that updates the temperature of external devices with the temperature value of the current device (doing `ext.temperature = this.temperature`). External devices may be filtered. Indeed, only the devices with `node == "hvac"` are affected by the update. The use of the keyword `this.` to indicate a resource on the current device is optional.

#### Special rules
To easy the programming of ECA rules, `abudsl` provides the following ***rule abstractions***.
>***Default Rule*** <br>
`rule` *RuleId* <br>
    &emsp;&emsp;`on` *Event* `default` *Action* <br>
    &emsp;&emsp;**(** *Task* **)***

>***IfElse Rule*** <br>
`rule` *RuleId* <br>
    &emsp;&emsp;`on` *Event* <br>
    &emsp;&emsp;`for` **[** `all` **]** *Condition* <br>
        &emsp;&emsp;&emsp;&emsp;`do` *Action* <br>
        &emsp;&emsp;&emsp;&emsp;`owise` *Action*

>***Let Rule*** <br>
`rule` *RuleId* <br>
    &emsp;&emsp;`on` *Event* <br>
    &emsp;&emsp;&emsp;`let` *LetDeclaration* `in` <br>
    &emsp;&emsp;**(** *Task* **)<sup>+</sup>**

In a ***Default*** rule the assignments in *Action* are always executed when *Event* happens, independently from tasks condition. In a ***IfElse*** rule the action after `do` is performed when *Condition* is true, while the action after `owise` is performed when  *Condition* is false. Finally, in a ***Let*** rule the substitutions in *LetDeclaration* are applied inside the non-empty list of tasks **(** *Task* **)<sup>+</sup>**. In particular, *LetDeclaration* is a semicolon-separated list of substitutions from expressions to resources. For instance, the rule
```
rule stupidCalculatorLet 
    on x y
      let sum := (x + y); diff := (x - y) in
    for (sum > 0)
        do result = sum * diff
```
is equivalent to the following:
```
rule stupidCalculator
    on x y
    for ((x + y) > 0)
        do result = (x + y) * (x - y)
```

### Comments
Inline comments start with a `#`:
```
# This is an inline comment.
```
while multi-line comments are enclosed between `\@` and `@\`:
```
\@
    This is a multi-
    line comment.
@\
```
