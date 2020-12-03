package main

import (
	"testing"
	"time"
)

func TestO2OTwoTypes(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	a8m, err := client.User.Create().SetAge(30).SetName("Mashraki").Save(ctx)
	if err != nil {
		t.Fatalf("creating user failed %v", err)
	}
	t.Logf("user is %v", a8m)

	card1, err := client.Card.Create().SetOwner(a8m).SetNumber("1020").SetExpired(time.Now().Add(time.Minute)).Save(ctx)
	if err != nil {
		t.Fatalf("creating card failed %v", err)
	}
	t.Logf("card is %v", card1)

	card2, err := a8m.QueryCard().Only(ctx)
	if err != nil {
		t.Fatalf("querying card failed %v", err)
	}
	t.Logf("a8m card is %v", card2)

	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		t.Fatalf("querying owner failed %v", err)
	}
	t.Logf("card owner is %v", owner)
}
