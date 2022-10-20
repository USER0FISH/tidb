# Serialization rework

- Author(s): [xhe](http://github.com/xhebox)
- Discussion PR: https://github.com/pingcap/tidb/pull/XXX
- Tracking Issue: https://github.com/pingcap/tidb/issues/XXX

## Table of Contents

* [Introduction](#introduction)
* [Motivation or Background](#motivation-or-background)
* [Detailed Design](#detailed-design)
* [Test Design](#test-design)
    * [Functional Tests](#functional-tests)
    * [Scenario Tests](#scenario-tests)
    * [Compatibility Tests](#compatibility-tests)
    * [Benchmark Tests](#benchmark-tests)
* [Impacts & Risks](#impacts--risks)
* [Investigation & Alternatives](#investigation--alternatives)
* [Unresolved Questions](#unresolved-questions)

## Introduction

Extended serialization formats.

## Motivation or Background

Serialization has grown to be very mature that there are more suitable formats for TiDB. Especially metadata of TiDB mostly will be store in json format, which is costy. The bench below shows a simple comparasion of some unmarshal implementations.

```
type TColumnInfo struct {
	ID                    int64           `json:"id"`
	Name                  CIStr           `json:"name"`
	Offset                int             `json:"offset"`
	OriginDefaultValue    interface{}     `json:"origin_default"`
	OriginDefaultValueBit []byte          `json:"origin_default_bit"`
	DefaultValue          interface{}     `json:"default"`
	DefaultValueBit       []byte          `json:"default_bit"`
	DefaultIsExpr         bool            `json:"default_is_expr"`
	GeneratedExprString   string          `json:"generated_expr_string"`
	Dependences           map[string]bool `json:"dependences"`
}

BenchmarkJsonUnmarshal-8    	  285547	      3567 ns/op	     720 B/op	      28 allocs/op
BenchmarkCapnpUnmarshal-8   	 2905576	       411.2 ns/op	     192 B/op	       3 allocs/op
BenchmarkProtoUnmarshal-8   	 1000000	      1060 ns/op	     504 B/op	      20 allocs/op
BenchmarkKMUnmarshal-8      	15301827	        90.12 ns/op	     256 B/op	       1 allocs/op
```

There are several modern strategies that makes serializing faster:

1. Pre-generated code: protobuf is a typical example. Unlike `encoding/json`, which will parse and resolve struct tags dynamically, protobuf code is pre-generated that it just knows how to decode/encode. It saves some allocations and the slow reflection logic.
2. Zero-copy: in the old good c language, we can cast `char[4]` as `int32_t`, so why don't we do the same for serialization? That is how flatbuffers/capnp raises. Such formats typically will only supports array and do not support dict structure.

We will eventually benifit from the faster speed, lower memory footprint, and possibly separated immutable/mutable metadata.

## Detailed Design

### Formats

The new serialization format will start with a byte indicating the wire format: `[type byte] [data]`. By inspecting the first byte of the data, we can tell how to deserialize.

There is only one exception: all existed entries in TiDB encoded to json are structs meaning their json will start with `{`. For old data, the formats is `'{' [json data without '{' prefix]`. So compatibility is not a problem for newer TiDB.

If users want to gain performance from nothing, it is fair to ask them to be responsible to upgrade all clusters before enabling the new formats.

### Zero copy 

Now theoritically we can apply protobuf and other similar formats. But we are not doing such big refactors just for speed. Frequent RPC has brought huge GC pressure. We expect less live pointers and smaller footprints. That leads to zero-copy solutions. In GO, it is done by `unsafe` tricks.

However, zero-copy has several obvious limitations:

1. Most type casting are done by assuming a fixed endianess. Considering useful CPUs are mostly targetting little endian, it is not a big deal.
2. Dict structures, i.e. `map[any]any` are not widely supported. One can not simply casting some bytes to a map possibly constructed by hash tables and trees.
3. `interface{}` is not widely supported. Some will provide tagged unions as helpers to solve the problem.

There are two possible solutions the these limitations:

1. An additional serialization code for zero-copy solutions: converting `[]TaggeUnion` into `[]interface`, and `[]KVEntry` into `map[K]V`.
2. Wrapped accessing: every field should have getters. Transformations are possibly done lazily, avoid parsing as much as possible.

Overhead of function calls are negligible. Both solutions should have similar performance. Moreover, we have getters for most of our metdata structs.

Zero-copy solutions typically have a `[]byte` buffer as the backend of unsafe tricks. It should be stored along with the deserialized structs wating to be GC. While this buffer can be used to construct `struct XXXViwer` with some getters, we can also use the same buffer for marshalling, if we only need one copy. Though it may be resized during marshalling, it is still reusing.

### Future work

It is common in TiDB that `model.FindXXXByName(data, name) is used instead of `data.FindXXXByName(name)`. I known that they are identical. But the latter will be more convenient for IDE/edtior autocompletion. It is very suitable to refactor the codebase to the latter along with the current proposal: `FindXXXByName(name)` is actually a conditional getter same as wrapped accessors.

But a big problem of wrapped accessors is that we should not give public fields, .e.g having an integer field `Count int` and a methods `GetCount() int`. We will easily write `x.Count` instead of `x.GetCount()` somehow. This brought us further: introducing setters and all fields become private, enforcing the usage of accessors.

Some previous work is done by https://github.com/pingcap/tidb/pull/33569. It is known that we can not tell if `FieldType` is unmarshalled from `model/[]byte`, or have been modified by upstreaming callers. Introducing mutable wrappers: `struct ImmutableXXX` and `strct MutableXXX` will solve the problem.

### Implementation

The proposal only works theoritically. And obviously it should be done incrementally to help reviewers. It can be devided into 3 steps: new formats & zero-copy with additional serilization -> getters with public fields -> setters/immutables/mutables.

The 3rd step is not included in the plan since it needs a lot of investigation of code. It is good to have, but it does not need to be done. 

To avoid `interface{}` overhead, we will likely stick to one new format only. I chose karmem.

Due to the limitation of zero-copy, we need to investigate every metatdata struct in the first step. To lazily unmarshal dict structures and interface, we will also need to investigate that in the second step.

While the code should be done in several parts, the investigation will be done together. Here is the investigate report.

1. `parser.model.CIStr`: special that only `CIStr.O` is marshalled, thus a special getter would be needed.
2. `parser.model.ColumnInfo`:

	1. `OriginDefaultValueBit` is a workaround for `OriginDefaultValue`. And `OriginDefaultValue` is promised to be a string. Thus we only need `[]byte` in the new format. 
	2. `DefaultValue` is similar to `OriginDefaultValue`.
	3. `Dependences` is `map[string]struct{}`. However, only `ddl/generated_column.go:393 checkExpressionIndexAutoIncrement` is using map access, all other usages are `for k := range map`. It is fair to just store a list and generate the map lazily.

3. ``


## Test Design

A brief description of how the implementation will be tested. Both the integration test and the unit test should be considered.

### Functional Tests

It's used to ensure the basic feature function works as expected. Both the integration test and the unit test should be considered.

### Scenario Tests

It's used to ensure this feature works as expected in some common scenarios.

### Compatibility Tests

A checklist to test compatibility:
- Compatibility with other features, like partition table, security & privilege, charset & collation, clustered index, async commit, etc.
- Compatibility with other internal components, like parser, DDL, planner, statistics, executor, etc.
- Compatibility with other external components, like PD, TiKV, TiFlash, BR, TiCDC, Dumpling, TiUP, K8s, etc.
- Upgrade compatibility
- Downgrade compatibility

### Benchmark Tests

The following two parts need to be measured:
- The performance of this feature under different parameters
- The performance influence on the online workload

## Impacts & Risks

Describe the potential impacts & risks of the design on overall performance, security, k8s, and other aspects. List all the risks or unknowns by far.

Please describe impacts and risks in two sections: Impacts could be positive or negative, and intentional. Risks are usually negative, unintentional, and may or may not happen. E.g., for performance, we might expect a new feature to improve latency by 10% (expected impact), there is a risk that latency in scenarios X and Y could degrade by 50%.

## Investigation & Alternatives

How do other systems solve this issue? What other designs have been considered and what is the rationale for not choosing them?

## Unresolved Questions

What parts of the design are still to be determined?
