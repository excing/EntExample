// https://entgo.io/docs/getting-started/
// Installation
// go get github.com/facebook/ent/cmd/ent
//
// Setup A Go Environment
// go mod init <project>
//
// Create Your First Schema
// ent init User
// and
// go generate ./ent

package main

import (
	"context"
	"ent_example/ent"
	"ent_example/ent/car"
	"ent_example/ent/enttest"
	"ent_example/ent/group"
	"ent_example/ent/migrate"
	"ent_example/ent/user"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateClient(t *testing.T) (context.Context, *ent.Client) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	// https://godoc.org/github.com/mattn/go-sqlite3#SQLiteDriver.Open
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, nil
	}
	return context.Background(), client
}

func TestCreateFirstEntity(t *testing.T) {
	_, client := CreateClient(t)
	defer client.Close()
}

func TestCreateUser(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)

	if err != nil {
		t.Fatalf("save entity failed %v", err)
	}

	t.Logf("entity is %v ", u)

	QueryUser(ctx, client, t)
}

func QueryUser(ctx context.Context, client *ent.Client, t *testing.T) {
	u, err := client.User.Query().Where(user.NameEQ("a8m")).Only(ctx)

	if err != nil {
		t.Fatalf("save entity failed %v", err)
	}

	t.Logf("entity name is %s, age is %d ", u.Name, u.Age)
}

func TestAddYourFirstEdge(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	// create a new car with model "Tesla".

	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		t.Fatalf("create tesla failed %v", err)
	}

	t.Logf("create %s successed", tesla.Model)

	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		t.Fatalf("create ford failed %v", err)
	}

	t.Logf("create %s successed", ford.Model)

	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddCars(tesla, ford).Save(ctx)
	if err != nil {
		t.Fatalf("add cars failed %v", err)
	}

	QueryCars(ctx, a8m, t)
}

func QueryCars(ctx context.Context, a8m *ent.User, t *testing.T) {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		t.Fatalf("query cars failed %v", err)
	}
	t.Logf("cars is %v", cars)

	ford, err := a8m.QueryCars().Where(car.ModelEQ("Ford")).Only(ctx)
	if err != nil {
		t.Fatalf("query ford failed %v", err)
	}
	t.Logf("ford is %v", ford)
}

func TestAddYourFirstInverseEdge(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	// create a new car with model "Tesla".

	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		t.Fatalf("create tesla failed %v", err)
	}

	t.Logf("create %s successed", tesla.Model)

	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		t.Fatalf("create ford failed %v", err)
	}

	t.Logf("create %s successed", ford.Model)

	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddCars(tesla, ford).Save(ctx)
	if err != nil {
		t.Fatalf("add cars failed %v", err)
	}

	QueryCarUsers(ctx, a8m, t)
}

func QueryCarUsers(ctx context.Context, a8m *ent.User, t *testing.T) {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		t.Fatalf("query cars failed %v", err)
	}

	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			t.Fatalf("query car %q owner %v", ca.Model, err)
		}
		t.Logf("car %q owner: %q", ca.Model, owner.Name)
	}
}

func TestCreateYourSecondEdge(t *testing.T) {
	ctx, client := CreateClient(t)
	defer client.Close()

	a8m, err := client.User.Create().SetAge(30).SetName("Ariel").Save(ctx)
	if err != nil {
		t.Fatalf("create Ariel failed %v", err)
	}

	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
	if err != nil {
		t.Fatalf("create Neta failed %v", err)
	}

	// Then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).SetOwner(a8m).Save(ctx)
	if err != nil {
		t.Fatalf("create Tesla failed %v", err)
	}
	_, err = client.Car.Create().SetModel("Mazda").SetRegisteredAt(time.Now()).SetOwner(a8m).Save(ctx)
	if err != nil {
		t.Fatalf("create Mazda failed %v", err)
	}
	_, err = client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).SetOwner(neta).Save(ctx)
	if err != nil {
		t.Fatalf("create Ford failed %v", err)
	}

	_, err = client.Group.Create().SetName("GitLab").AddUsers(neta, a8m).Save(ctx)
	if err != nil {
		t.Fatalf("create GitLab failed %v", err)
	}
	_, err = client.Group.Create().SetName("GitHub").AddUsers(a8m).Save(ctx)
	if err != nil {
		t.Fatalf("create GitHub failed %v", err)
	}
	t.Log("The graph was created successfully")

	QueryGithub(ctx, client, t)
	QueryArielCars(ctx, client, t)
	QueryGroupWithUsers(ctx, client, t)
}

func QueryGithub(ctx context.Context, client *ent.Client, t *testing.T) {
	cars, err := client.Group.Query().Where(group.Name("GitHub")).QueryUsers().QueryCars().All(ctx)
	if err != nil {
		t.Fatalf("failed getting cars: %v", err)
	}
	t.Logf("cars returned: %v", cars)
}

func QueryArielCars(ctx context.Context, client *ent.Client, t *testing.T) {
	a8m := client.User.Query().Where(user.HasCars(), user.Name("Ariel")).OnlyX(ctx)
	cars, err := a8m.QueryGroups().QueryUsers().QueryCars().Where(car.Not(car.ModelEQ("Mazda"))).All(ctx)
	if err != nil {
		t.Fatalf("failed getting Ariel cars: %v", err)
	}
	t.Logf("Ariel cars returned: %v", cars)
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client, t *testing.T) {
	groups, err := client.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		t.Fatalf("failed getting groups: %v", err)
	}
	t.Logf("groups returned: %v", groups)
}
