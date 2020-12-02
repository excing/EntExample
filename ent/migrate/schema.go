// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// BlobsColumns holds the columns for the "blobs" table.
	BlobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
	}
	// BlobsTable holds the schema information for the "blobs" table.
	BlobsTable = &schema.Table{
		Name:        "blobs",
		Columns:     BlobsColumns,
		PrimaryKey:  []*schema.Column{BlobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "model", Type: field.TypeString},
		{Name: "registered_at", Type: field.TypeTime},
		{Name: "user_cars", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cars_users_cars",
				Columns: []*schema.Column{CarsColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amout", Type: field.TypeFloat64},
		{Name: "name", Type: field.TypeString, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:        "cards",
		Columns:     CardsColumns,
		PrimaryKey:  []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "count", Type: field.TypeInt},
		{Name: "code", Type: field.TypeInt},
		{Name: "index", Type: field.TypeInt},
		{Name: "min", Type: field.TypeInt},
		{Name: "max", Type: field.TypeInt},
		{Name: "range", Type: field.TypeInt},
		{Name: "note", Type: field.TypeString},
		{Name: "log", Type: field.TypeString, Size: 200},
		{Name: "username", Type: field.TypeString},
		{Name: "user_groups", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "groups_users_groups",
				Columns: []*schema.Column{GroupsColumns[12]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:        "pets",
		Columns:     PetsColumns,
		PrimaryKey:  []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "rank", Type: field.TypeFloat64, Nullable: true},
		{Name: "active", Type: field.TypeBool},
		{Name: "old_name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeJSON, Nullable: true},
		{Name: "strings", Type: field.TypeJSON, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"on", "off"}},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "group_users", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_groups_users",
				Columns: []*schema.Column{UsersColumns[13]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_friends_user_id",
				Columns: []*schema.Column{UserFriendsColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_friends_friend_id",
				Columns: []*schema.Column{UserFriendsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlobsTable,
		CarsTable,
		CardsTable,
		GroupsTable,
		PetsTable,
		UsersTable,
		UserFriendsTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}
