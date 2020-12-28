package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"reflect"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/lightningnetwork/lnd/lnwire"
)

func marshal(msg interface{}) string {
	var buf string

	val := reflect.Indirect(reflect.ValueOf(msg))
	t := val.Type()
	buf += `{"kind": "` + t.String() + `"`

	for i := 0; i < val.NumField(); i++ {
		name := t.Field(i)
		field := val.Field(i)

		var encodedValue string
		if field.Kind().String() == "struct" {
			encodedValue = marshal(field.Interface())
		} else {
			encodedValue = marshalValue(field.Interface())
		}
		buf += "," + `"` + name.Name + `":` + encodedValue
	}

	buf += "}"

	dst := &bytes.Buffer{}
	json.Indent(dst, []byte(buf), "", "  ")

	return dst.String()
}

func marshalValue(val interface{}) string {
	switch v := val.(type) {
	case lnwire.Sig:
		val = hex.EncodeToString(v[:])
	case chainhash.Hash:
		val = hex.EncodeToString(v[:])
	default:
		val = v
	}

	vjson, _ := json.Marshal(val)
	return string(vjson)
}
