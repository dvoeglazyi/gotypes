### gotypes

Some simple types based on builtin Golang types
that implement interfaces for working with DB (Scan / Value) and JSON (Marshal / Unmarshal).

#### NullUint
Simplified sql.NullInt64 (but unsigned): not struct, based on builtin uint type.
```
ni := sql.NullInt64{Int64: 32, Valid: true}
// Corresponds to
nu := gotypes.NullUint(32)

ni := sql.NullInt64{Int64: 0, Valid: true}
// Corresponds to
nu := gotypes.NullUint(0)

ni := sql.NullInt64{Int64: 32, Valid: false}
// Corresponds to
nu := gotypes.NullUint(0)
```

#### NullString
Simplified sql.NullString: not struct, based on builtin string type.
```
ni := sql.NullString{String: "example", Valid: true}
// Corresponds to
nu := gotypes.NullString("example")

ni := sql.NullInt64{String: "", Valid: true}
// Corresponds to
nu := gotypes.NullString("")

ni := sql.NullInt64{String: "example", Valid: false}
// Corresponds to
nu := gotypes.NullString("")
```

#### Base64
Type for simply decoding / encoding Base64 in JSON-structs.
```
type Something struct {
    Base gotypes.Base64 `json:"base"`
}

something := Something{Base: gotypes.Base64("decode me please")}

jsonString, err := json.Marshal(something)
if err != nil {
    return err
}
// Base is Base64 encoded string.
fmt.Println(string(jsonString))
// Output:
// {"base":"ZGVjb2RlIG1lIHBsZWFzZQ=="}

var anything Something
if err := json.Unmarshal(jsonString, &anything); err != nil {
    panic(err)
}
// Base has value decoded from Base64 string.
fmt.Println(string(anything.Base))
// Output:
// decode me please
```