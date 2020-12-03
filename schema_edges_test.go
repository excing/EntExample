package main

import (
	"ent_example/ent"
	"ent_example/ent/node"
	"ent_example/ent/user"
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

func TestO2OBidirectional(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	a8m, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
	if err != nil {
		t.Fatalf("creating a8m user failed %v", err)
	}

	nati, err := client.User.Create().SetAge(28).SetName("nati").SetSpouse(a8m).Save(ctx)
	if err != nil {
		t.Fatalf("creating nati user failed %v", err)
	}

	spouse := nati.QuerySpouse().OnlyX(ctx)
	t.Logf("nati spouse is %v", spouse)

	spouse = a8m.QuerySpouse().OnlyX(ctx)
	t.Logf("a8m spouse is %v", spouse)

	count := client.User.Query().Where(user.HasSpouse()).CountX(ctx)
	t.Logf("has spouse count %d", count)

	spouse = client.User.Query().Where(user.HasSpouseWith(user.Name("a8m"))).OnlyX(ctx)
	t.Logf("a8m spouse is %v", spouse)
}

func TestO2MTwoTypes(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	padro, err := client.Pet.Create().SetName("padro").Save(ctx)
	if err != nil {
		t.Fatalf("creating padro pet failed %v", err)
	}

	lola, err := client.Pet.Create().SetName("lola").Save(ctx)
	if err != nil {
		t.Fatalf("creating lola pet failed %v", err)
	}

	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddPets(padro, lola).Save(ctx)
	if err != nil {
		t.Fatalf("creating a8m user failed %v", err)
	}

	t.Logf("User created: %v", a8m)

	owner := padro.QueryOwner().OnlyX(ctx)
	t.Logf("padro owner is %v", owner)

	count := padro.QueryOwner().QueryPets().CountX(ctx)
	t.Logf("padro owner has pet count is %d", count)
}

func TestO2MSameType(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	root, err := client.Node.Create().SetValue(2).Save(ctx)
	if err != nil {
		t.Fatalf("creating root failed %v", err)
	}

	n1 := client.Node.Create().SetValue(1).SetParent(root).SaveX(ctx)
	n4 := client.Node.Create().SetValue(4).SetParent(root).SaveX(ctx)
	n3 := client.Node.Create().SetValue(3).SetParent(n4).SaveX(ctx)
	n5 := client.Node.Create().SetValue(5).SetParent(n4).SaveX(ctx)

	t.Logf("Tree leafs %v", []int{n1.Value, n3.Value, n5.Value})

	ints := client.Node.Query().
		Where(node.Not(node.HasChildren())).
		Order(ent.Asc(node.FieldValue)).
		GroupBy(node.FieldValue).IntsX(ctx)
	t.Logf("all leafs(without children) %v", ints)

	orphan := client.Node.Query().
		Where(node.Not(node.HasParent())).
		OnlyX(ctx)
	t.Logf("all leafs(without parent) %v", orphan)
}
