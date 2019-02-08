// Package awscdkgo is an example of what a CDK for Go could look like
// generated from JSII.
//
// This module creates both an interface and struct for each JSII class in
// order to support subclasses, (e.g. Foo extends Bar, so Foo must be usable
// everywhere Bar is usable). Subtyping in Go requires the usage of Interfaces.
// To support this all usages of a JSII class, (e.g. property, parameter, or
// return type) must use the interface form of the JSII class. Generated
// constructors for the JSII classes will return a struct pointer that
// satisfies its associated interface. Types are generated in the pattern of,
// <JSIIClassName> as the struct, and <JSIIClassName>_Iface for the interface.
//
// The generated structs wrap the underlying JSII kernel calls. Subclasses is
// obtained by embedding the extended class within the subclass's Go struct.
// The struct pointer for the extended class is the type embedded. The
// interfaces generated for each JSII embeds the "_Iface" form of the
// extended class.
//
// To extend a JSII class in Go, a customer must initialize the base JSII type,
// and provide the appropriate method overrides. Method overrides in this
// prototype are an addition constructor parameter of an interface type for the
// required methods (aka pure virtual, aka abstract). Additional methods can be
// overridden as all JSII methods and properties are virtual.
//
// Benefits of Generating both Interface and Structs
//
// Customers receives improved documentation via godoc. Godoc does expand
// interface methods like it does for structs. Using structs as the type
// returned to the user from constructs provides users with a more useful godoc
// Table of Contents.
//
// Interfaces can include private members ensuring customers do not implement a
// JSII class's interface, because the interface types can be generated to
// include private methods that only can be satisfied by the generated struct.
// This "feature" may be too restrictive, and prevent users from implementing
// interfaces for simple types, (e.g not constructs).
//
// Pitfalls of this prototype
//
// Customers must know which form of a type they need to use in their code. All
// JSII generated types would always take the interface form of a class, but
// customers writing their own code could easily get confused when to use the
// interface vs struct. In addition generating both interface increases the
// documentation surface area.
//
// The FailProps and StateProps types in awsstepfunctions pkg are exported
// structs instead of interfaces like other JSII classes. It is not clear if
// JSII for Go has enough information to make this determination of when a
// struct can safely be used versus an interface.
package awscdkgo
