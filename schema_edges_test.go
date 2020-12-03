package main

import (
	"ent_example/ent/node"
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

func TestO2OSameType(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	head, err := client.Node.Create().SetValue(1).Save(ctx)
	if err != nil {
		t.Fatalf("creating the head failed %v", err)
	}
	t.Logf("head is %v", head)

	curr := head
	// Generate the following linked-list: 1<->2<->3<->4<->5
	for i := 0; i < 4; i++ {
		curr, err = client.Node.Create().SetValue(curr.Value + 1).SetPrev(curr).Save(ctx)
		if err != nil {
			t.Fatalf("creating the curr %d failed %v", i, err)
		}
	}

	// Loop over the list and print it. `FirstX` panics if an error occur.
	for curr = head; curr != nil; curr = curr.QueryNext().FirstX(ctx) {
		t.Logf("node value is %d", curr.Value)
	}

	tail, err := client.Node.Query().Where(node.Not(node.HasNext())).Only(ctx)
	if err != nil {
		t.Fatalf("getting the tail of the list failed %v", err)
	}
	tail, err = tail.Update().SetNext(head).Save(ctx)
	if err != nil {
		t.Fatalf("updating next of the tail failed %v", err)
	}

	// check that the change actually applied:
	prev, err := head.QueryPrev().Only(ctx)
	if err != nil {
		t.Fatalf("getting head's prev failed %v", err)
	}
	t.Logf("head's prev is %v", prev)
	t.Logf("\n%v", prev.Value == tail.Value)
}
