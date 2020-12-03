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
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "number", Type: field.TypeString},
		{Name: "expired", Type: field.TypeTime},
		{Name: "user_card", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cards_users_card",
				Columns: []*schema.Column{CardsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "count", Type: field.TypeInt, Nullable: true},
		{Name: "code", Type: field.TypeInt, Nullable: true},
		{Name: "index", Type: field.TypeInt, Nullable: true},
		{Name: "min", Type: field.TypeInt, Nullable: true},
		{Name: "max", Type: field.TypeInt, Nullable: true},
		{Name: "range", Type: field.TypeInt, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "log", Type: field.TypeString, Nullable: true, Size: 200},
		{Name: "username", Type: field.TypeString, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt},
		{Name: "node_next", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "node_children", Type: field.TypeInt, Nullable: true},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "nodes_nodes_next",
				Columns: []*schema.Column{NodesColumns[2]},

				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "nodes_nodes_children",
				Columns: []*schema.Column{NodesColumns[3]},

				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_pets", Type: field.TypeInt, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "pets_users_pets",
				Columns: []*schema.Column{PetsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt, Nullable: true},
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
		{Name: "user_spouse", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_users_spouse",
				Columns: []*schema.Column{UsersColumns[12]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "group_users_group_id",
				Columns: []*schema.Column{GroupUsersColumns[0]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "group_users_user_id",
				Columns: []*schema.Column{GroupUsersColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
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
		NodesTable,
		PetsTable,
		UsersTable,
		GroupUsersTable,
		UserFriendsTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	NodesTable.ForeignKeys[0].RefTable = NodesTable
	NodesTable.ForeignKeys[1].RefTable = NodesTable
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}
