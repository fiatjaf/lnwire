lnwire
======

Install with `go get github.com/fiatjaf/lnwire` or [download a binary](https://github.com/fiatjaf/lnwire/releases).

Use with `lnwire <hint for the type of the message you're trying to decode> <actual message you're trying to decode as hex>`, for example:

```bash
lnwire fee-insu 100c0000000000635265008868db916580a1f3e4395eb01763c31cea814b397cd6ba5
cc414ce211f471b0eac0de87ffafb889d01891a7171c773db5a6a6bd11a07aab9443c556a46f1b684e56fe28c0ab
6f1b372c1a6a246ae63f74f931e8365e15a089c68d619000000000009d6e800084200015fe71b370101002800000
000000003e800000001000000400000000076046700
{
  "kind": "lnwire.FailFeeInsufficient",
  "HtlcMsat": 1156299204327374947,
  "Update": {
    "kind": "lnwire.ChannelUpdate",
    "Signature": "008868db916580a1f3e4395eb01763c31cea814b397cd6ba5cc414ce211f471b0eac0de87f
fafb889d01891a7171c773db5a6a6bd11a07aab9443c556a46f1b6",
    "ChainHash": "84e56fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d619000000",
    "ShortChannelID": {
      "kind": "lnwire.ShortChannelID",
      "BlockHeight": 9,
      "TxIndex": 14084096,
      "TxPosition": 2114
    },
    "Timestamp": 90087,
    "MessageFlags": 27,
    "ChannelFlags": 55,
    "TimeLockDelta": 257,
    "HtlcMinimumMsat": 11258999068426240,
    "BaseFee": 65536000,
    "FeeRate": 65536,
    "HtlcMaximumMsat": 18014398509512196,
    "ExtraOpaqueData": "ZwA="
  }
}
```

Or:

```bash
lnwire channel-u 100c0000000000635265008868db916580a1f3e4395eb01763c31cea814b397
cd6ba5cc414ce211f471b0eac0de87ffafb889d01891a7171c773db5a6a6bd11a07aab9443c556a46f1b684e56fe
28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d619000000000009d6e800084200015fe71b370101002
800000000000003e800000001000000400000000076046700
{
  "kind": "lnwire.ChannelUpdate",
  "Signature": "100c0000000000635265008868db916580a1f3e4395eb01763c31cea814b397cd6ba5cc414ce
211f471b0eac0de87ffafb889d01891a7171c773db5a6a6bd11a",
  "ChainHash": "07aab9443c556a46f1b684e56fe28c0ab6f1b372c1a6a246ae63f74f931e8365",
  "ShortChannelID": {
    "kind": "lnwire.ShortChannelID",
    "BlockHeight": 14768648,
    "TxIndex": 10250454,
    "TxPosition": 6400
  },
  "Timestamp": 0,
  "MessageFlags": 9,
  "ChannelFlags": 214,
  "TimeLockDelta": 59392,
  "HtlcMinimumMsat": 595038106670275383,
  "BaseFee": 16842792,
  "FeeRate": 0,
  "HtlcMaximumMsat": 4294967296001,
  "ExtraOpaqueData": "AAAAQAAAAAB2BGcA"
}
```

## License

Public domain, except you can't use for shitcoins.
