package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/lightningnetwork/lnd/lnwire"
)

const USAGE = `lnwire.

Usage:
  lnwire <type-hint> <hex-encoded-lightning-message>
  lnwire -h | --help
  lnwire --version
`

func main() {
	parser := &docopt.Parser{
		HelpHandler:  docopt.PrintHelpOnly,
		OptionsFirst: true,
	}
	opts, err := parser.ParseArgs(USAGE, os.Args[1:], "1.0")
	if err != nil {
		return
	}

	var raw []byte
	var matcher = func(d decodeable) bool { return true }

	if rawhex, ok := opts["<hex-encoded-lightning-message>"].(string); ok {
		raw, err = hex.DecodeString(rawhex)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error decoding hex: %s\n", err.Error())
			return
		}
	}

	if query, ok := opts["<type-hint>"].(string); ok {
		query = strings.ToLower(strings.ReplaceAll(query, "-", ""))

		matcher = func(d decodeable) bool {
			name := strings.ToLower(
				reflect.Indirect(reflect.ValueOf(d)).Type().String())
			if strings.Index(name, query) != -1 {
				return true
			}
			return false
		}
	}

	for _, msg := range []decodeable{
		&lnwire.FailAmountBelowMinimum{},
		&lnwire.FailChannelDisabled{},
		&lnwire.FailExpiryTooSoon{},
		&lnwire.FailFeeInsufficient{},
		&lnwire.FailFinalIncorrectCltvExpiry{},
		&lnwire.FailFinalIncorrectHtlcAmount{},
		&lnwire.FailIncorrectCltvExpiry{},
		&lnwire.FailIncorrectDetails{},
		&lnwire.FailInvalidOnionHmac{},
		&lnwire.FailInvalidOnionKey{},
		&lnwire.FailInvalidOnionVersion{},
		&lnwire.FailTemporaryChannelFailure{},
		&lnwire.AcceptChannel{},
		&lnwire.AnnounceSignatures{},
		&lnwire.ChannelAnnouncement{},
		&lnwire.ChannelUpdate{},
		&lnwire.ClosingSigned{},
		&lnwire.CommitSig{},
		&lnwire.FundingCreated{},
		&lnwire.FundingLocked{},
		&lnwire.FundingSigned{},
		&lnwire.GossipTimestampRange{},
		&lnwire.Init{},
		&lnwire.InvalidOnionPayload{},
		&lnwire.NodeAnnouncement{},
		&lnwire.OpenChannel{},
		&lnwire.Ping{},
		&lnwire.Pong{},
		&lnwire.QueryChannelRange{},
		&lnwire.QueryShortChanIDs{},
		&lnwire.ReplyChannelRange{},
		&lnwire.ReplyShortChanIDsEnd{},
		&lnwire.RevokeAndAck{},
		&lnwire.Shutdown{},
		&lnwire.UpdateAddHTLC{},
		&lnwire.UpdateFailHTLC{},
		&lnwire.UpdateFailMalformedHTLC{},
		&lnwire.UpdateFee{},
		&lnwire.UpdateFulfillHTLC{},
	} {
		if !matcher(msg) {
			continue
		}

		err := msg.Decode(bytes.NewReader(raw), 99)
		if err != nil {
			continue
		}
		fmt.Println(marshal(msg))
		return
	}

	fmt.Fprintf(os.Stderr, "couldn't decode to any lightning wire message format.")
}

type decodeable interface {
	Decode(r io.Reader, pver uint32) error
}
