# Full syntax of AbU DSL
In the following, we denote with **(** *exp* **)*** zero or more repetitions of *exp*, while with **(** *exp* **)<sup>+</sup>** one or more repetitions of *exp*. Furthermore, **[** *exp* **]** represents an optional occurrence of *exp*, while *exp1* **|** *exp2* stands for either *exp1* or *exp2*.

## Syntax for programs, devices and ECA rules
>*Program* &nbsp;**::=**&nbsp; **(** *Device* **[** `has` *RuleIdList* **] )<sup>+</sup>** **(** *ECARule* **)*** <br>
*Device* &nbsp;**::=**&nbsp; *DeviceId* `:` *Description* `{` *ResourceDeclaration* **[** `where` *BooleanExpression* **]** `}` <br>
*ResourceDeclaration* &nbsp;**::=**&nbsp; **(** *PhysicalResource* **|** *LogicalResource* **)<sup>+</sup>** <br>
*PhysicalResource* &nbsp;**::=**&nbsp; `physical output` *Type* *ResourceId* `=` *Expression* **|** `physical input` *Type* *ResourceId* <br>
*LogicalResource* &nbsp;**::=**&nbsp; `logical` *Type* *ResourceId* `=` *Expression* <br>
*RuleIdList* &nbsp;**::=**&nbsp; **(** *RuleId* **)<sup>+</sup>** <br>
*ECARule* &nbsp;**::=**&nbsp; `rule` *RuleId* `on` *Event* **(** *Task* **)<sup>+</sup>** <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `default` *Action* **(** *Task* **)*** <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `for` **[** `all` **]** *Condition* `do` *Action* `owise` *Action* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `let` *LetDeclaration* `in` **(** *Task* **)<sup>+</sup>** <br>
*Event* &nbsp;**::=**&nbsp; **(** *ResourceId* **)<sup>+</sup>** <br>
*Task* &nbsp;**::=**&nbsp; `for` **[** `all` **]** *Condition* `do` *Action* <br>
*Action* &nbsp;**::=**&nbsp; *Assignment* **(** `;` *Assignment* **)*** <br>
*Assignment* &nbsp;**::=**&nbsp; **[** `this.` **]** *ResourceId* `=` *Expression* **|** `ext.` *ResourceId* `=` *Expression* <br>
*LetDeclaration* &nbsp;**::=**&nbsp; *ResourceId* `:=` *Expression* **( `;`** *ResourceId* `:=` *Expression* **)***

## Syntax for expressions and conditions
>*Expression* &nbsp;**::=**&nbsp; *BooleanExpression* **|** *NonBooleanExpression* <br>
*BooleanExpression* &nbsp;**::=**&nbsp; *BooleanValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; **[** `this.` **]** *ResourceId* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `(` *BooleanExpression* `)` <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `not` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *BooleanExpression* `and` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *BooleanExpression* `or` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `==` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `!=` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `<` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `<=` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `>` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `>=` *NonBooleanExpression* <br>
*NonBooleanExpression* &nbsp;**::=**&nbsp; *NumericExpression* **|** *StringExpression* <br>
*NumericExpression* &nbsp;**::=**&nbsp; *NumericValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; **[** `this.` **]** *ResourceId* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `(` *NumericExpression* `)` <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `abs` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `+` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `-` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `*` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `/` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `%` *NumericExpression* <br>
*StringExpression* &nbsp;**::=**&nbsp; *StringValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *Identifier* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *StringExpression* `::` *StringExpression* <br>
*Condition* &nbsp;**::=**&nbsp; *BooleanValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; **[** `this.` **]** *ResourceId* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `ext.` *ResourceId* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `(` *BooleanExpression* `)` <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `not` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *BooleanExpression* `and` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *BooleanExpression* `or` *BooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `==` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `!=` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `<` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `<=` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `>` *NonBooleanExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NonBooleanExpression* `>=` *NonBooleanExpression* <br>

## Syntax for values, types and identifiers
>*ResourceId* &nbsp;**::=**&nbsp; *Identifier* <br>
*Identifier* &nbsp;**::=**&nbsp; *Character* **(** *Character* **|** *Digit* **)*** <br>
*Character* &nbsp;**::=**&nbsp; `a` **|** `b` ... **|** `z` **|** `A` **|** `B` ... **|** `Z` <br>
*SpecialCharacter* &nbsp;**::=**&nbsp; ` `&nbsp;**|** `!` **|** `#` **|** ... **|** `~` <br>
*Digit* &nbsp;**::=**&nbsp; `0` **|** `1` **|** ... **|** `9` <br>
*BooleanValue* &nbsp;**::=**&nbsp; `true` **|** `false` <br>
*NumericValue* &nbsp;**::=**&nbsp; *IntegerValue* **|** *DecimalValue* <br>
*IntegerValue* &nbsp;**::=**&nbsp; **[** `-` **]** **(** *Digit* **)<sup>+</sup>** <br>
*DecimalValue* &nbsp;**::=**&nbsp; *IntegerValue* `.` **(** *Digit* **)<sup>+</sup>** <br>
*StringValue* &nbsp;**::=**&nbsp; `"` **(** *Character* **|** *SpecialCharacter* **|** *Digit* **)*** `"` <br>
*Type* &nbsp;**::=**&nbsp; `boolean` **|** `integer` **|** `decimal` **|** `string` <br>
*DeviceId* &nbsp;**::=**&nbsp; *Identifier* <br>
*Description* &nbsp;**::=**&nbsp; *StringValue* <br>
*RuleId* &nbsp;**::=**&nbsp; *Identifier* <br>

## Comments
Inline comments keyword: `#` <br>
Multi-line comments start delimiter: `\@` <br>
Multi-line comments end delimiter: `@\`
