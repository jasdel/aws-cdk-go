我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# JSII & CDK for Go Design

## Abstract

[JSII] provides the tools for software to invoke and interact with with a
Javascript runtime. JSII provides this with models generated from exported
Javascript, and a RPC protocol. The JSII model defines the JSII types that
language binding must code generated JSII proxies for. The code generated
proxies wrap calls to the JSII runtime via a JSII RPC client. The JSII RPC
client must be implemented by the language binding for communicating with the
JSII runtime through the JSII RPC.

[CDK] is the only Javascript tool with JSII cross language bindings. This may
change in the future, but JSII was purpose built for the CDK usecase. The JSII
for Go implementation should be able to support an other Javascript tool with
JSII cross language bindings.

## JSII proxy model

JSII is a JSON document IDL defining classes and interfaces that language
bindings are generated from. JSII modeled types are classes and interfaces.
These types are built on subtyping inheritance, and polymorphic Object Oriented
concepts. Binding languages are required to use these concepts in their
language, or if the language does not have these concepts, simulate them to the
best means possible.

### Simulating polymorphism

The Go programming language is type system is not based on object oriented and
polymorphic principles. Instead, Go is a procedural language of types which may
contain methods, and duck-typed interfaces. In addition, concepts such as class
subtyping are not directly supported by the language's type system. These
concepts can be simulated in go, but this comes at a unexpected and
non-idiomatic experience cost for the customer.

#### Class inheritance

All JSII classes and datatype interfaces must be rendered into two Go types, an
interface and struct. The two types are required to facilitate subtyping of
JSII types and class inheritance. Go structs do not support subtyping. The
struct provides the concrete implementation of the JSII class. The interface
provides support for subtyping the struct's value as its extended type, or
implemented interfaces.

#### Subtyping

JSII requires language bindings to support subtyping. Subtyping is the concept
where a subtype implements or extend another base type, and that subtype can be
used anywhere the base type is accepted, (e.g. `Foo` extends `Bar`, so `Foo`
must be usable as a parameter anywhere `Bar` is).

Given Go's type system, only interfaces can satisfy this requirement. In order
for a Go type to be used as a subtype of another Go type, the target type
must be defined as an interface. There are no alternatives to this requirement.
This requirement, combined with class inheritance and polymorphism requirements
place significant limitations and restrictions how Go proxy types can be
defined.

## Type and Naming Convention

JSII types are defined with a Javascript first approach. This means that all
JSII types are based heavily on Javascript type system. To translate these JSII
types into Go proxy types, care must be taken to ensure the generated Go proxy
types are forward compatible with the underling Javascript types. In addition,
the JSII for Go generation must take care to limit or reduce the risk of
identifier name clashing. Due to the differences in language philosophy between
Javascript and Go. The Go proxy types must generate additional identifiers into
the namespace that risk clashing with JSII types. The additional identifier
names generated must also be forward compatible with future additions made to
the underling Javascript types.

### Enums Type

JSII Unions are defined as collections of immutable string values with the enum
type as the namespace. Go does not have direct support for enumerations, but
does have common idiom for defining a collection of immutable values of scalar
types, (e.g. string, float, bool, etc). The Go idiom for enums does not have
namespacing though. This means enum values are global to the Go package they
are defined within.

Without the namespace the enum value identifier name have a significant risk to
name clashing with other JSII types. To address this risk the JSII for Go
generator must render enum value's name with the namespace included. This
technique is the idiomatic naming convention of enum values in Go.

There are two different ways the JSII for Go could define enum value identifier
names, EnumValue_EnumType or EnumValueEnumType. The enum value's identifier
name would or would not use a underscore(_) separating the enum value and enum
type identifier sections. With the underscore there is very little to no risk
of identifier name clashing, but this is not an idiomatic style. Go idiomatic
naming does not include underscores(_).

```go
type Color string

const(
	Red Color                Color = "Red"
	Harlequin_Color          Color = "Harlequin"
	TropicalRainForest_Color Color = "Tropical rain forest"
)
```

### Interface
JSII interfaces are similar to Go's duck-typed interfaces. The JSII interface
is a collection of methods and property getter/setters that can be satisfied by
a JSII class.

JSII interfaces identifier names can be rendered as Go types without any
modifications.


```go
type IConstruct interface{
	GetFoo() string
	SomeMethod(a string)
}
```

**Note**: JSII interfaces, and JSII datatype interfaces are very different.
JSII datatype interfaces are similar to JSII classes with the exception that
properties are exported struct members, regardless of the `immutability` flag.

#### Extending other interfaces

A JSII interface that extends another interface when rendered in Go must the
extended interface be an anonymous member of the Go sub interface in addition
to the methods of the sub interface. 

It is invalid in Go for an embedded interface to include a method with the same
name as a method in the interface where it is embedded. Two sibling interfaces
may define the same method as long all definitions are the same.

For example the `IFoo` JSII interface extends `IBar` JSII interface. The
following is the resulting Go types.

```go
type IBar interface{ /* */ }

type IFoo interface{
	IBar
	/* */
}
```

### Properties

JSII language bindings are proxies for the underlying JSII RPC class, Go
structs representing the JSII class must only export methods 

The only exception to this is JSII datatype interface. Datatype interfaces are
equivalent to structs or typed dictionaries where the customer provides the
member values directly. The datatype interfaces are not JSII proxies.

#### Optional

JSII properties and method parameters can be decorated as optional. This
decoration means that a value for the parameter or property is not needed. In
Javascript this means a `undef` value is allowed. Parameters decorated with
optional in JSII for Go must use a pointer type when rendering the parameter.

Users of JSII for Go would leave properties unset, and use `nil` for method
parameters that are optional.

#### Getter

If a JSII class, datatype interface, or interface has a property that decorated
with `immutable` a getter method must be generated on the type for property.
The method must be prefixed with `Get` to distinguish it from the member name,
if any.

```go
type StatePropsIface interface {
	GetComment() *string
}

type StateProps struct {
	Comment *string
}

func (s *StateProps) GetComment() *string { return s.Comment }
```

#### Setter
If a JSII class, datatype interface, or interface has a property that not
decorated with `immutable` a setter method must be generated on the type for
the property. The method must be prefixed with `Set` to distinguish it from
the member name, if any.

```go
type StatePropsIface interface {
	GetComment() *string
	SetComment() *string
}

type StateProps struct {
	Comment *string
}

func (s *StateProps) GetComment() *string { return s.Comment }
func (s *StateProps) SetComment(v string) { return s.Comment }
```

### Class

JSII classes support polymorphic inheritance and subtyping. The JSII for Go
must render both a Go interface and struct for every one JSII class in order to
satisfy this requirement. Go structs do not support subtyping. The struct
provides the concrete implementation of the JSII class. The interface provides
support for subtyping the Struct's value as its extended type, or implemented
interfaces.

To ensure subtyping is supported by all Go JSII proxy types, all occurrences of
the JSII class in property type and method input/output parameters must use the
Go interface type.

Even if the JSII class does not extend another JSII class, the JSII for Go
generator must always generate both Go interface and struct types. The reason
for this is to ensure the JSII for Go generated code will be forward
compatible, and it ensures customers will be able to write custom Go types that
extend JSII classes.

The generated Go interface must embed all extended JSII class's Go interfaces
as anonymous fields. Likewise the Go struct for a JSII class must include the
extended JSII class's Go struct as anonymous field. This pattern ensures the Go
types correctly satisfy the simulated polymorphic requirements.

The Go interface for a JSII class must include a private method in its
definition to prevent users of the JSII for Go from providing custom
implementations for the JSII proxy. Users of the JSII for Go will still be able
to extend JSII classes, but would not be able to provide custom
implementations. This design requirement ensures that user code will not be
broken when JSII classes add methods or properties to the underlying Javascript
classes. 

In order to prevent name classing the Go interface must be suffixed with an
identifier that separates it from the Go struct. For example, an `Iface` suffix
provides this separation.

```go
// Fail provides the subtyping interfaces for JSII Fail class.
type FailIface interface {
	StateIface

	faiPrivate()
}

// Fail is a JSII class.
type Fail struct {
	*State

	base jsii.Base
}

func (*Fail) fail_private() {}
```

Removing the private method from the Go interface would allow users to provide
custom implementations for JSII classes, and would remove the need for the Go
struct to be generated for each JSII class. While this would remove the overall
type confusion it opens the user up to breaking changes or unexpected behavior.
If the Go struct is remove the interface's identifier name does not need to be
suffixed with `Iface`.

#### Class constructor

**TODO**: fill out this section with reasoning about why the three constructors are needed.

* `New<ClassName>` - Public constructor the user would use to create a new instance of this JSII proxy type.
* `InternalNew<ClassName>AsBaseClass` - Internally used by the JSII for Go generated proxy types to initialize extended JSII classes.
* `Extend<ClassName>` - Public constructor for user to associate their custom Go type's methods as overriding methods defined by the JSII proxy type. This allows customers to override JSII class methods.

#### Class static methods and constant properties

JSII classes may be defined with static methods and constant properties. Both
of these translate to package global Go functions. Similar to JSII enums, JSII
static methods and constant properties must be rendered with the JSII class's
name included in the identifier name of the Go function. 

The static methods and constant property Go functions could suffix the JSII
class name to the method name, or separate the two names with an underscore(_).
Similar to JSII enums, using the underscore separator ensures the resulting
identifier name will always be forward compatible, and not clash with other
JSII identifiers, but at the cost of not being idiomatic Go types.

```go
func ClassName_StaticMethodName(...) T { ... }
```

### Datatype interface

JSII datatype interfaces are a collection of property members on a datatype.
This correlates closely to Go's struct type. JSII datatype interfaces are very
different from JSII interfaces not labeled as datatype. JSII datatype
interfaces are a collection of properties that can be extended with other JSII
datatype interfaces. 

The datatype interface's properties can be decorated as `abstract`. This
decoration has no meaningful value. All datatype interface properties in Go
must be represented as struct members.

Similar to JSII classes, datatype interface must generate two Go types, an
interface, and struct for each JSII datatype interface. Like JSII classes two
types are generated to support subtyping. JSII datatype interfaces are always
rendered as two Go type, even when a datatype interface does not extend another
datatype interface. This is done to ensure forward compatibility if a JSII
datatype interface were to extend another datatype interface in the future.

```go
type StatePropsIface interface {
	GetComment() *string
	GetInputPath() *string
}
type StateProps struct {
	Comment *string
	InputPath *string
}
```

Unlike the JSII class the JSII datatype interface's Go interface does not need
to define a private method to restrict implementation. There is no reason to
limit the custom implementation of the datatype interface's proxy as this type
is only a collection of properties.

### Promise
Promise do not exist in Go. In Go synchronous API function are idiomatic APIs.
JSII functions retuning a  would be wrapped within a synchronous Go
function that wait for the async JSII function to be complete. The synchronous
go function would then return the result or error. This is the same pattern
used for Java and .Net JSII bindings.

Unlike most JSII functions, functions returning Promises will return both the
functions modeled return type, and an error.

```go
func AnAsyncJSIIFunction() (T, error) { /* ... */ }
```

### Inline Anonymous Unions
JSII Unions are defined on a type's reference, not as their own type.
Implementation of Unions in Go would require a type to be defined, and named.
JSII Unions are anonymous and the JSII for Go would have to derive a
sustainable name for the union that does not clash with other types in the JSII
model.

Unions should not be supported at all, or if they are required to be supported,
JSII Unions must translate to empty `interface{}`. There is no other
sustainable way anonymous unions could be implemented in Go.

### JSII collections

JSII includes both maps and arrays. Both of these types translate easily to Go
types. An JSII array is a Go slice `[]T`. A JSII map is a Go map,
`map[string]T`. JSII a map must only have a key of a `string` type. `T` is the
JSII modeled `elementType`.

### JSII primitives

#### Date

JSII Date would be satisfied by Go's `time.Time` type.

#### String

JSII String would be satisfied by Go's `string` type.

#### Number

JSII number does not distinguish the difference between integer and float. This
either requires JSII for Go to treat all numbers as `float64`, or define a new
`jsii.Number` type that "wraps" Go number types into the singular JSII Number
type. 

Typescript numbers are float64 with a max integer size of 2^53 -1. This means
that Customers could want to provide a `int64` value that would lose precision.

#### JSON

The closet analogy to JSII's `json` `PrimitiveType` is a
`map[string]interface{}`. While the map does allow the user to retrieve
arbitrary fields from the map, the user is required to cast the retrieved value
to a Go type before it can be used. This process can intensive for users, and
an area for code bugs.

```go
v := failState.ToStateObject()
f, ok := v["foo"]
if !ok {
	// handle field not set
}
ft, ok := f.(jsii.Number)
if !ok {
	// handle field not a number
}

// Use ft as a jsii.Number type
```

#### Any

JSII any maps to Go's empty interface type. JSII for Go should replace all JSII
`any` types with `interface{}`

## JSII for Go Runtime

JSII requires that language bindings must ensure only a single instance of the
JSII runtime is executing per invocation of the app built with the JSII
language bindings. For JSII for Go this translates to a global instance of the
JSII client value that all Go proxy types generated from the JSII model will
reference.

### Error Handling

JSII has no method for defining errors, or exceptions that can be thrown.
Regardless if an exception is a recoverable exception or not. In addition to
not defining errors, the JSII RPC makes no distinction between RPC errors,
kernel runtime errors, and errors encountered with the underlying Javascript
code.

This limitation forces the JSII for Go into a poor customer experience for
exposing errors to customers. Either all Go JSII proxies return error in
addition to their modeled return type, or all JSII proxies panic if a JSII
exception is received. Without JSII including modeling or customizations on top
of the models there are no other options. Both of these solutions are poor
customer experiences with panic being the least bad since customers will not
need to check for error on every possible function, method, and property call.

Using JSII for Go outside of CDK's specific all errors are unrecoverable, and
singleton design is most likely unusable due to the lack of meaningful errors
defined within JSII.

### JSII RPC client

#### JSII runtime kernal NodeJS subprocess

The JSII runtime "kernel" is executed within a NodeJS subprocess. The JSII RPC
client must communicate with the NodeJS subprocess using stdin/stdout pipes.

### JSII for Go code generation

JSII's `pacmac` should be extended to generate Go code from the JSII model.
While JSII is a JSON document there are implicit complexities reading the
format that would incur a high maintenance  cost implementing a JSII for Go
generator in Go instead of using JSII's built in `pacmac`. The `pacmac` is a
typescript implementation reading JSII model and JSII language binding code
generation. There doesn't seem to be a strong reason why a Go `pacmac` should
be written instead of adding JSII for Go generation to the existing `pacmac`.


[JSII]: https://github.com/awslabs/jsii
[CDK]: https://github.com/awslabs/aws-cdk
