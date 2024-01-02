# orderedmap

*forked from iancoleman/orderedmap v0.3.0*

A golang data type equivalent to python's collections.OrderedDict

Retains order of keys in maps

Can be JSON serialized / deserialized

# Usage

[example](example)

# Caveats

* OrderedMap only takes strings for the key, as per [the JSON spec](http://json.org/).

# Tests

```
go test
```

# Alternatives

None of the alternatives offer JSON serialization.

* [cevaris/ordered_map](https://github.com/cevaris/ordered_map)
* [mantyr/iterator](https://github.com/mantyr/iterator)
