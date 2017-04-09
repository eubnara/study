import the entire "foo" and "bar" modules;
module foo from "foo";
module bar from "bar";

console.log(
  bar.hello("rhino")
)
foo.awesome()
