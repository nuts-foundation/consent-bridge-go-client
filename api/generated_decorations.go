package api

import (
	core "github.com/nuts-foundation/nuts-go-core"
	"log"
)

func (i *Identifier) UnmarshalJSON(bytes []byte) error {
	partyID := core.PartyID{}
	if err := partyID.UnmarshalJSON(bytes); err != nil {
		return err
	}
	*i = Identifier(partyID.String())
	return nil
}

func (i Identifier) PartyID() core.PartyID {
	id, err := core.ParsePartyID(string(i))
	if err != nil {
		log.Fatalf("should never happen: invalid PartyID: %s", i)
	}
	return id
}
