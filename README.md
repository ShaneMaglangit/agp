# agp

[![GoDoc](https://pkg.go.dev/badge/github.com/shanemaglangit/agp?utm_source=godoc)](https://godoc.org/github.com/shanemaglangit/agp)

Package `agp` is a gene parsing package for Axie Infinity.

The name agp stands for "Axie Gene Parse" which decodes the hex representation of an Axie's gene into a human readable format.

---

* [Install](#install)
* [Usage](#examples)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/shanemaglangit/agp
```

## Examples

To get started, you'll first need to get the gene of an Axie in hex. You may use the [Axie Infinity GraphQL endpoint](https://axie-graphql.web.app/) to get this detail. For this example, let's use the hex `0x11c642400a028ca14a428c20cc011080c61180a0820180604233082`

Let us first parse this hex into a GeneBinGroup object. `ParseHex()` first converts the given hex into its binary format. It thens divides these binary bits into their own respective groups, each representing a certain attribute of the Axie's gene.

```go
gbg, err := ParseHex("0x11c642400a028ca14a428c20cc011080c61180a0820180604233082")
```

Once we generated the GeneBinGroup, we can then use decode this object into human readable format using `Decode()`

```go
genes, err := Decode(gbg)
```

The generated output should look like this

```go
&Genes{
  Class:    Beast,
  Region:   Global,
  Tag:      NoTag,
  BodySkin: DefBodySkin,
  Pattern:  PatternGene{"000001", "000111", "000110"},
  Color:    ColorGene{"f0c66e", "ffec51", "f0c66e"},
  Eyes: Part{
    D:  PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"},
    R1: PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"},
    R2: PartGene{"eyes-blossom", Plant, "", Eyes, "Blossom"},
  },
  Ears: Part{
    D:  PartGene{"ears-lotus", Plant, "", Ears, "Lotus"},
    R1: PartGene{"ears-nut-cracker", Beast, "", Ears, "Nut Cracker"},
    R2: PartGene{"ears-inkling", Aquatic, "", Ears, "Inkling"},
  },
  Horn: Part{
    D:  PartGene{"horn-rose-bud", Plant, "", Horn, "Rose Bud"},
    R1: PartGene{"horn-caterpillars", Bug, "", Horn, "Caterpillars"},
    R2: PartGene{"horn-dual-blade", Beast, "", Horn, "Dual Blade"},
  },
  Mouth: Part{
    D:  PartGene{"mouth-tiny-turtle", Reptile, "", Mouth, "Tiny Turtle"},
    R1: PartGene{"mouth-piranha", Aquatic, "", Mouth, "Piranha"},
    R2: PartGene{"mouth-serious", Plant, "", Mouth, "Serious"},
  },
  Back: Part{
    D:  PartGene{"back-balloon", Bird, "", Back, "Balloon"},
    R1: PartGene{"back-jaguar", Beast, "", Back, "Jaguar"},
    R2: PartGene{"back-jaguar", Beast, "", Back, "Jaguar"},
  },
  Tail: Part{
    D:  PartGene{"tail-ant", Bug, "", Tail, "Ant"},
    R1: PartGene{"tail-hot-butt", Plant, "", Tail, "Hot Butt"},
    R2: PartGene{"tail-swallow", Bird, "", Tail, "Swallow"},
  },
}
```
