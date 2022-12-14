# Full syntax of AbU DSL
In the following, we denote with **(** *exp* **)*** zero or more repetitions of *exp*, while with **(** *exp* **)<sup>+</sup>** one or more repetitions of *exp*. Furthermore, **[** *exp* **]** represents an optional occurrence of *exp*, while *exp1* **|** *exp2* stands for either *exp1* or *exp2*.

## Syntax for programs, devices and ECA rules
>*Program* &nbsp;**::=**&nbsp; **(** *TypeDeclaration* **)*** **(** *Device* **[** `has` *RuleIdList* **] )<sup>+</sup>** **(** *ECARule* **)*** <br>
*TypeDeclaration* &nbsp;**::=**&nbsp; `define` *CompoundType* `as {` *FieldDeclaration* `}` <br>
*FieldDeclaration* &nbsp;**::=**&nbsp; **(** *ResourceId* `:` **(** `physical input` *PrimitiveType* **|** `physical output` *PrimitiveType* **|** `logical` *PrimitiveType* **) )<sup>+</sup>** <br>
*Device* &nbsp;**::=**&nbsp; *DeviceId* `:` *Description* `{` *ResourceDeclaration* **[** `where` *BooleanExpression* **]** `}` <br>
*ResourceDeclaration* &nbsp;**::=**&nbsp; **(** *PhysicalResource* **|** *LogicalResource* **|** *CompoundResource* **)<sup>+</sup>** <br>
*PhysicalResource* &nbsp;**::=**&nbsp; `physical output` *PrimitiveType* *ResourceId* `=` *Value* **|** `physical input` *PrimitiveType* *ResourceId* <br>
*LogicalResource* &nbsp;**::=**&nbsp; `logical` *PrimitiveType* *ResourceId* `=` *Value* <br>
*CompoundResource* &nbsp;**::=**&nbsp; *CompoundType* *ResourceId* `= (` **[** *Constructor* **]** `)` <br>
*Constructor* &nbsp;**::=**&nbsp; *ResourceId* `=` *Value* **(** `,` *ResourceId* `=` *Value* **)***  <br>
*RuleIdList* &nbsp;**::=**&nbsp; **(** *RuleId* **)<sup>+</sup>** <br>
*ECARule* &nbsp;**::=**&nbsp; `rule` *RuleId* `on` *Event* **(** *Task* **)<sup>+</sup>** <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `default` *Action* **(** *Task* **)*** <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `for` **[** `all` **]** *Condition* `do` *Action* `owise` *Action* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `rule` *RuleId* `on` *Event* `let` *LetDeclaration* `in` **(** *Task* **)<sup>+</sup>** <br>
*Event* &nbsp;**::=**&nbsp; **(** *ResourceId* **|** *ResourceId `[` *ResourceId* `]` **)<sup>+</sup>** <br>
*Task* &nbsp;**::=**&nbsp; `for` **[** `all` **]** *Condition* `do` *Action* <br>
*Action* &nbsp;**::=**&nbsp; *Assignment* **(** `,` *Assignment* **)** * <br>
*Assignment* &nbsp;**::=**&nbsp; *LocalResourceAccess* `=` *Expression* **|** *RemoteResourceAccess* `=` *Expression* <br>
*LocalResourceAccess* &nbsp;**::=**&nbsp; **[** `this.` **]** *ResourceId* **|**  **[** `this.` **]** *ResourceId* `[` ResourceId `]` <br>
*RemoteResourceAccess* &nbsp;**::=**&nbsp; `ext.` *ResourceId* **|** `ext.` *ResourceId* `[` ResourceId `]` <br>
*LetDeclaration* &nbsp;**::=**&nbsp; *ResourceId* `:=` *Expression* **( `;`** *ResourceId* `:=` *Expression* **)** * <br>

## Syntax for expressions and conditions
>*Expression* &nbsp;**::=**&nbsp; *BooleanExpression* **|** *NonBooleanExpression* <br>
*Condition* &nbsp;**::=**&nbsp; *BooleanExpression* <br>
*BooleanExpression* &nbsp;**::=**&nbsp; *BooleanValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *LocalResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *RemoteResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *ForeignFunction* <br>
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
  &emsp;&emsp;&emsp; **|**&nbsp; *LocalResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *RemoteResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *ForeignFunction* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `(` *NumericExpression* `)` <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `absint` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; `absdec` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `+` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `-` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `*` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `/` *NumericExpression* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *NumericExpression* `%` *NumericExpression* <br>
*StringExpression* &nbsp;**::=**&nbsp; *StringValue* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *LocalResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *RemoteResourceAccess* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *ForeignFunction* <br>
  &emsp;&emsp;&emsp; **|**&nbsp; *StringExpression* `::` *StringExpression* <br>
*ForeignFunction* &nbsp;**::=**&nbsp; `foreign (` *StringValue* `)` **|** `foreign (` *StringValue* **(** `,` *Param* **)<sup>+</sup>** `)` <br>
*Param* &nbsp;**::=**&nbsp; *Value* **|** *LocalResourceAccess* <br>


## Syntax for values, types and identifiers
>*ResourceId* &nbsp;**::=**&nbsp; *Identifier* <br>
*Identifier* &nbsp;**::=**&nbsp; *Character* **(** *Character* **|** *Digit* **)*** <br>
*Character* &nbsp;**::=**&nbsp; `a` **|** `b` ... **|** `z` **|** `A` **|** `B` ... **|** `Z` <br>
*SpecialCharacter* &nbsp;**::=**&nbsp; ` `&nbsp;**|** `!` **|** `#` **|** ... **|** `~` <br>
*Digit* &nbsp;**::=**&nbsp; `0` **|** `1` **|** ... **|** `9` <br>
*Value* &nbsp;**::=**&nbsp; *BooleanValue* **|** *NumericValue* **|** *StringValue* <br>
*BooleanValue* &nbsp;**::=**&nbsp; `true` **|** `false` <br>
*NumericValue* &nbsp;**::=**&nbsp; *IntegerValue* **|** *DecimalValue* <br>
*IntegerValue* &nbsp;**::=**&nbsp; **[** `-` **]** **(** *Digit* **)<sup>+</sup>** <br>
*DecimalValue* &nbsp;**::=**&nbsp; *IntegerValue* `.` **(** *Digit* **)<sup>+</sup>** <br>
*StringValue* &nbsp;**::=**&nbsp; `"` **(** *Character* **|** *SpecialCharacter* **|** *Digit* **)*** `"` <br>
*Type* &nbsp;**::=**&nbsp; *PrimitiveType* **|** *CompoundType* <br>
*PrimitiveType* &nbsp;**::=**&nbsp; `boolean` **|** `integer` **|** `decimal` **|** `string` <br>
*CompoundType* &nbsp;**::=**&nbsp; *Identifier* <br>
*DeviceId* &nbsp;**::=**&nbsp; *Identifier* <br>
*Description* &nbsp;**::=**&nbsp; *StringValue* <br>
*RuleId* &nbsp;**::=**&nbsp; *Identifier* <br>

## Comments
Inline comments keyword: `#` <br>
Multi-line comments start delimiter: `\@` <br>
Multi-line comments end delimiter: `@\`
