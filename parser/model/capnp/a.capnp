using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("capnp");
$Go.import("foo/books");

struct CIStr {
	o @0 :Text;
	l @1 :Text;
}

struct ColumnInfo {
	id @0 :Int64;
	name @1 :CIStr;
	offset @2 :Int64;
	originDefaultValue @3 :Text;
	originDefaultValueBit @4 :Data;
	defaultIsExpr @5 :Bool;
	generatedExprString @6 :Text;
	dependences @7 :List(Text);
}
