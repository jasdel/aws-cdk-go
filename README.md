# JSII for Go Design Prototypes

## Overview
Notable Design considerations that may lead to poor customer experience:
* **Forbidden features**:
	* Circular package/module references/dependencies.
	* Type self reference without pointer.
	* Exported Identifiers starting with "_" character.
	* Overridden methods with different types

* **Breaking Changes**:
	* Adding optional parameter to function or method
	* Removing any type, property, member, or function
	* Changing the name of an existing identifier

* **CDK Construct ToString Method**:
Go has a convention of to string helper methods being named "String" it may be
possible to add this customization for CDK constructs to use "String" instead
of "ToString" for better customer experience.

* **Identifier Name collisions**:
JSII for Go will need to add additional new identifiers to the generated
package from the JSII model. It is possible for these new identifiers added to
the namespace to collide with JSII modeled identifiers.
	* Package scoped:
		* `<TypeName>Iface` - Go interfaces for class & datatypes
		* `New<ClassName>` - constructor
		* `Extend<ClassName>` - constructor for custom Go constructs
		* `Internal<ClassName>AsBaseClass` - internal constructor for class hierarchy
		* `<ClassName>_<StaticMethodName>`
		* `<ClassName>_Get<StaticPropertyName>` - Static property getter
		* `<ClassName>_Set<StaticPropertyName>` - Static property setter
		* `<EnumValue><EnumType>` - Enum types
	* Type scoped (aka JSII Class/interface):
		* `Get<PropetyName>` - Property getter method
		* `Set<PropertyName>` - Property setter method

* **Optional Type References**:
Optional type references require Go to use pointer types for the type
reference. Pointer types are required in order to identify if the user provided
a value or not. Unlike TypeScript Go has no way to distinguish between `null`
and `undefined`.

Go customers will be annoyed with APIs using pointers for parameters when they
think the parameter type's zero value should be sufficient to identify unset,
(e.g. empty string). Customers will be required to used additional code to set
literal values to optional fields. Go does not support taking the address of a
literal value, (e.g. "abc123"). This is a common customer pain point with the
AWS SDK for Go.

See `Type References - Optional` for implementation and examples.

* **Unions**:
JSII for Go will either not support JSII unions at all, or will be forced to
use empty `interface{}` for the union property or parameter type.

See `Unions` for implementation and more details.

* **Class Constructors**:
**TODO**: Update mocks so New<ClassName>() should only be created for
non-abstract classes.  Customers should use Extend<ClassName>() to extend an
abstract class within their Go application..

JSII Classes can be labeled as "abstract". In TypeScript this means the class
cannot be initialized outside of a subclass, (e.g. super). How should JSII
language bindings allow customers to extend "abstract" classes with custom
Constructor functions are generated for each JSII class defined. If the class
is abstract a New<ClassName> constructor is not generated. The other two
constructors will always be generated.

See `Classes - Constructors` for implementation and details.

* **Abstract Methods and Properties**:
Abstract methods and properties of a class must be implemented by the subclass.
Abstract properties within JSII datatype interfaces are ignored.

* **JSII's `json` primitive**:
JSII defines an `json` primitive which equates to a dictionary based type,
(e.g. Javascript's object). The closet translation to Go would be the
`map[string]interface{}` pattern to represent arbitrary key value pairs. The Go
map requires users to do type assertions to extract the values of keys. This
involves significant level of effort on users to extract fields from the
returned value.

JSII for Go may provide a utility to easy working with the `json` value.

See `Types - Primitives - JSON` for implementation and examples.

## Enums

While Go does not have direct support for enums, a common pattern can be used
to support JSII enums. Go enum pattern injects the enums into the package
global namespace, similar to C enums. JSII for Go must convert the enum value
to a valid Go identifier. The type of the enum is suffixed to the enum value's
identifier to reduce the risk of name collisions within a package.

```go
type Color string

const (
	Red_Color                Color = "Red"
	Harlequin_Color          Color = "Harlequin"
	TropicalRainForest_Color Color = "Tropical rain forest"
)
```

## Type Reference

### Unions

JSII defines Union types as type references, not as a type like class, enum,
and interface. JSII for Go must render the union as a concrete Go type. Since a
JSII union can be an combination of any types, including primitives Go
interfaces may not be sufficient to wrap the type, while still providing strong
type assertions.

JSII for Go users may be required to wrap their value with the Union type for
the type reference. 

JSII for Go must derive some type name for the union. This name must never
change.

**TODO needs design**

### Optional
Optional type references in Go are represented as pointer types. Go has no,
`option` or `sum` type. Go is required to define optional type references as
pointer types since Go types are values, and their `unset` state is the zero
value of the type. (e.g. empty string).  Since the zero value of a type may be
valid, and JSII optional is defined as allowing `undefined`, Go type references
must be pointer types.

This is a common pain point for AWS SDK for Go customers for a few reasons:
* Working with pointers risks nil pointer dereferences, causing runtime panics.
* Using pointers for values generally lead to the value being allocated on the stack impacting runtime performance. 
* Go does not allow taking the address of a composite literal, (e.g. "abc123", true, 123, etc).

To ease this pain JSII for Go should create helper functions to assist users
using composite literal.

```go
// S returns a new pointer to string value passed in.
func S(v string) *string { /* ... */ }
// SVal returns a string value pointed to by the string pointer. Returns empty
// string if nil.
func SVal(v *string) string { /* ... */ }
// B returns a new pointer to bool value passed in.
func B(v bool) *bool { /* ... */ }
// BVal returns a bool value pointed to by the bool pointer. Returns false if
// nil.
func BVal(v *bool) bool { /* ... */ }
// ...
```

With calling functions any optional parameter, `nil` must be passed in by the
user for parameters they do not want to set.

```go
func Print(foo string, message *string) { /* ... */ }

// Unset optional field
Print("abc123", nil)

// Setting optional field.
Print("abc123", jsii.S("foo")) 
```

Optional type references can be used as both members of structs and parameters
of functions. 

```go
type MyClass struct {
	OptionalMember *string
	NotOptional string
}

// Unset optional field
c := MyClass{NoOptional: "abc123"}

// Setting optional field.
c := MyClass
```

### Promise
Promise do not exist in Go. In Go synchronous API function are idiomatic APIs.
JSII functions retuning a promise would be wrapped within a synchronous Go
function that wait for the async JSII function to be complete. The synchronous
go function would then return the result or error. JSII for Java and .Net use
this pattern as well.

Unlike most JSII functions, functions returning Promises will return both the
functions modeled return type, and an error.

```go
func AnAsyncJSIIFunction() (int, error) { /* ... */ }
```

## Types
### Documentation
JSII define documentation for function parameters. Go docs do not have special
formating for function parameters, (e.g JSDoc). This style of documentation is
counter to Go's. Either JSII for Go drops these doc comments or, some
acceptable formating is used.

```go
// Print prints a message.
// 
// Parameters:
// * message - Message to print
func Print(message string) { /* ... */ }
```

### Collection
#### Array
JSII Arrays map to Go slices, `[]T`, where T is the JSII modeled `elementtype` of the map.

#### Map
JSII Maps map to Go `map[string]T`, where T is the JSII modeled `elementtype` of the map.

### Primitives
#### Date
JSII Date would be satisfied by Go's `time.Time` type.

**TODO** Need to validate that `time.Time` meets all of the requirements of JSII Date.

#### String
JSII String would be satisfied by Go's `string` type.

#### Number
JSII number does not distinguish the difference between integer and float. This
either requires JSII for Go to treat all numbers as `float64`, or define a new
`jsii.Number` type that "wraps" Go number types into the singular JSII Number
type. 

Typescript numbers are float64 with a max integer size of 2^53 -1. This means
that Customers could want to provide a `int64` value that would lose precision.

**Prefer JSII split into integer and float**

#### JSON
The closet analogy to JSII's `json` `PrimitiveType` is a
`map[string]interface{}`. While the map does allow the user to retrieve
arbitrary fields from the map, the user is required to case the retrieved value
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
JSII any maps to Go's empty interface (`interface{}`) type.


## Functions
### Parameters
#### Variadic
#### Default

## Classes
### Constructor
JSII expects generated classes to call the base class's constructor (aka super)
with implicitly passing all parameters from sub class's constructor to base
class's constructor. This is by convention and is very fragile, and can lead to
code that will not compile if the constructor parameters do not match up.

### Abstract
Abstract class can never be instantiated. Go would need to enforce this as well
for classes decorated with abstract.

### Implemented Interfaces

These are the interfaces that a JSII class must implement. The Go type must
provide all methods defined within the interface. 

Go uses anonymous fields within to represent each implemented interface. This
require that the interface name and member/method names of the JSII class
must never clash.


### Properties (aka Members)

All members must be getter/setters so that the get/update can be forwarded
through the underlying JSII runtime.

* Getters for properties are the name of the property.
* Setters for properties are <Set><PropertyName>.

#### Immutable
JSII immutable members will only have a getter.

#### Protected
Protected is public in Go. 

#### Abstract (aka pure virtual)
Abstract members and methods must be implemented by subclasses.

#### Virtual Methods
All methods and properties in JSII are virtual and can be overridden by a
subclass. JSII for Go must provide a way when generating JSII classes for users
to provide overrides for every method and property.

#### Method invoke Scope
JSII methods must be invoked in the correct context. Some kind of callback
registration system is needed to register a overridden method to be called from
JSII kernel.

#### Static and Constants Property
JSII class constant and static methods/properties are equivalent to Go package
global accessor functions.  Class constant values in JSII model are package
global accessor functions in JSII for go.

```go
func ClassName_PropertyName(...) T { ... }
```

### Methods 
#### Initializer (aka Constructor)
#### Protected
#### Abstract
#### Static

## Interfaces

### Base

The class or interface that the interface "extends"

### Methods
#### Abstract

### Properties

Interfaces in Go cannot have properties. Go interfaces may only have methods. All Properties must be rendered as methods instead.

* **Getter**: Uses the `<PropetyName>` naming scheme.
* **Setter**: uses the `Set<propertyName>` naming scheme.
* **Member**: Uses the `<PropertyName>_` naming scheme. Adds trailing `_` to name preventing clash with method.

### Datatype
JSII interfaces which are datatypes translate to Go structs. JSII for Go is
able to translate the datatype interfaces to Go structs because the JSII types
are collection of data values.

JSII datatype interfaces' properties can include `abstract` and `immutable`. If
Go were to implement Go structs for these datatypes, the abstract and immutable 

In JSII an InterfaceType can extend other InterfaceTypes. Unlike Classes an
InterfaceType uses the field `interfaces` to model the `base` InterfaceTypes,
that the sub interface extends.

JSII interfaces of properties, e.g `datatype`, are collections properties with
accessors. To represent this hierarchy in Go the consumer methods of the
property interfaces must accept an interface form of the modeled JSII type.

Datatype properties must be suffixed with underscore(`_`) to prevent clashing
with the associated getter method that must be generated for the interface.

```go
type StatePropser interface {
	Comment() *string
	InputPath() *string
}
type StateProps struct {
	Comment_ *string
	InputPath_ *string
}

type FailPropser interface {
	StatePropser
	Error() string
	Comment() *string
}
type FailProps struct {
	StateProps
	Error_ string
	Comment_ *string
}

failState := NewFail(parentConstruct, id, FailProps{})
```

## Naming Design

All JSII classes must be rendered into two Go types, an interface and struct.
The two types are required to facilitate subtyping of JSII types. Go structs do
not support subtyping. The struct provides the concrete implementation of the
JSII class. The interface provides support for subtyping the Struct's value as
its extended type, or implemented interfaces.

### Class Constructor

Three "constructor methods are generated for every JSII class
* New<Name>: Creates a new JSII object for the class
* New<Name>_AsBaseClass: Initializes the type for use as a base class of another type. Used internally for simulating class inheritance of JSII classes.
* New<Name>_WithOverrides: Creates a new JSII object  with extended 

### How To JSII -> Go
* JSII Interface: `<InterfaceName>`, (e.g `IConstruct`) (non-datatype)
* JSII Class's Implementation: `<ClassName>_`, (e.g `State_`)
* JSII Class's Go interface: `<ClassName>`, (e.g `State`)
* JSII Class's Interface for overriding methods: `<ClassName>_Overrides interface{}`
* JSII Class extending base class
	* Embed extended class's Go interface in subsclass's Go interface and Go Struct.
* JSII Property
	* Datatype property member:  `<PropertyName>_ T`, (e.g. `Error_ string`)
	* Getter method (immutable): `<PropertyName>() T`, (e.g. `Error() string`)
	* Setter method (mutable): `Set<PropertyName)(v T)`, (e.g. `SetError(v string)`)
* JSII Class Constructor:
	* `New<ClassName>` - Constructor customer would use
	* `New<ClassName>_AsBaseClass` - internally used by generated classes for extended classes.
	* `New<ClassName>_WithOverrides` - for extending CDK construct with custom code.
* JSII Class Method: `<MethodName>`
* JSII Interface Method:  `<MethodName>`
* JSII Static Class Methods: `func <ClassName>_<MethodName>(...)`
* JSII Static Class Prop Getter: `func <ClassName>_<PropertyName>() T`
* JSII Static Class Prop Setter: `func <ClassName>_Set<PropertyName>(v T)`
* JSII enum:
	* Base type: `type <EnumTypeName> string`
	* Enum Value: `const <EnumValue>_<EnumTypeName> <EnumTypeName> = "<enum value>"` (grouped)
