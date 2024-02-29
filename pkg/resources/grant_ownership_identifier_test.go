package resources

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGrantOwnershipId(t *testing.T) {
	testCases := []struct {
		Name       string
		Identifier string
		Expected   GrantOwnershipId
		Error      string
	}{
		{
			Name:       "grant ownership on database to account role",
			Identifier: `ToAccountRole|"account-role"|COPY|OnObject|DATABASE|"database-name"`,
			Expected: GrantOwnershipId{
				GrantOwnershipTargetRoleKind: ToAccountGrantOwnershipTargetRoleKind,
				AccountRoleName:              sdk.NewAccountObjectIdentifier("account-role"),
				OutboundPrivilegesBehavior:   sdk.Pointer(CopyOutboundPrivilegesBehavior),
				Kind:                         OnObjectGrantOwnershipKind,
				Data: &OnObjectGrantOwnershipData{
					ObjectType: sdk.ObjectTypeDatabase,
					ObjectName: sdk.NewAccountObjectIdentifier("database-name"),
				},
			},
		},
		// TODO:
		//{
		//	Name:       "grant ownership on schema to account role",
		//	Identifier: `ToAccountRole|"account-role"|COPY|OnObject|SCHEMA|"database-name"."schema-name"`,
		//	Expected: GrantOwnershipId{
		//		GrantOwnershipTargetRoleKind: ToAccountGrantOwnershipTargetRoleKind,
		//		AccountRoleName:              sdk.NewAccountObjectIdentifier("account-role"),
		//		OutboundPrivilegesBehavior:   sdk.Pointer(CopyOutboundPrivilegesBehavior),
		//		Kind:                         OnObjectGrantOwnershipKind,
		//		Data: &OnObjectGrantOwnershipData{
		//			ObjectType: sdk.ObjectTypeSchema,
		//			ObjectName: sdk.NewDatabaseObjectIdentifier("database-name", "schema-name"),
		//		},
		//	},
		//},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			id, err := ParseGrantOwnershipId(tt.Identifier)
			if tt.Error == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, id)
			} else {
				assert.ErrorContains(t, err, tt.Error)
			}
		})
	}
}

//func TestGrantOwnershipIdString(t *testing.T) {
//	testCases := []struct {
//		Name       string
//		Identifier GrantPrivilegesToAccountRoleId
//		Expected   string
//		Error      string
//	}{
//		{
//			Name: "grant account role on account",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: true,
//				AllPrivileges:   true,
//				Kind:            OnAccountAccountRoleGrantKind,
//				AlwaysApply:     true,
//				Data:            new(OnAccountGrantData),
//			},
//			Expected: `"account-role"|true|true|ALL|OnAccount`,
//		},
//		{
//			Name: "grant account role on account object",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: true,
//				AllPrivileges:   true,
//				Kind:            OnAccountObjectAccountRoleGrantKind,
//				AlwaysApply:     true,
//				Data: &OnAccountObjectGrantData{
//					ObjectType: sdk.ObjectTypeDatabase,
//					ObjectName: sdk.NewAccountObjectIdentifier("database-name"),
//				},
//			},
//			Expected: `"account-role"|true|true|ALL|OnAccountObject|DATABASE|"database-name"`,
//		},
//		{
//			Name: "grant account role on schema on schema",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaAccountRoleGrantKind,
//				Data: &OnSchemaGrantData{
//					Kind:       OnSchemaSchemaGrantKind,
//					SchemaName: sdk.Pointer(sdk.NewDatabaseObjectIdentifier("database-name", "schema-name")),
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchema|OnSchema|"database-name"."schema-name"`,
//		},
//		{
//			Name: "grant account role on all schemas in database",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaAccountRoleGrantKind,
//				Data: &OnSchemaGrantData{
//					Kind:         OnAllSchemasInDatabaseSchemaGrantKind,
//					DatabaseName: sdk.Pointer(sdk.NewAccountObjectIdentifier("database-name")),
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchema|OnAllSchemasInDatabase|"database-name"`,
//		},
//		{
//			Name: "grant account role on future schemas in database",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaAccountRoleGrantKind,
//				Data: &OnSchemaGrantData{
//					Kind:         OnFutureSchemasInDatabaseSchemaGrantKind,
//					DatabaseName: sdk.Pointer(sdk.NewAccountObjectIdentifier("database-name")),
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchema|OnFutureSchemasInDatabase|"database-name"`,
//		},
//		{
//			Name: "grant account role on schema object on object",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaObjectAccountRoleGrantKind,
//				Data: &OnSchemaObjectGrantData{
//					Kind: OnObjectSchemaObjectGrantKind,
//					Object: &sdk.Object{
//						ObjectType: sdk.ObjectTypeTable,
//						Name:       sdk.NewSchemaObjectIdentifier("database-name", "schema-name", "table-name"),
//					},
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchemaObject|OnObject|TABLE|"database-name"."schema-name"."table-name"`,
//		},
//		{
//			Name: "grant account role on schema object on all tables in database",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaObjectAccountRoleGrantKind,
//				Data: &OnSchemaObjectGrantData{
//					Kind: OnAllSchemaObjectGrantKind,
//					OnAllOrFuture: &BulkOperationGrantData{
//						ObjectNamePlural: sdk.PluralObjectTypeTables,
//						Kind:             InDatabaseBulkOperationGrantKind,
//						Database:         sdk.Pointer(sdk.NewAccountObjectIdentifier("database-name")),
//					},
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchemaObject|OnAll|TABLES|InDatabase|"database-name"`,
//		},
//		{
//			Name: "grant account role on schema object on all tables in schema",
//			Identifier: GrantPrivilegesToAccountRoleId{
//				RoleName:        sdk.NewAccountObjectIdentifier("account-role"),
//				WithGrantOption: false,
//				Privileges:      []string{"CREATE SCHEMA", "USAGE", "MONITOR"},
//				Kind:            OnSchemaObjectAccountRoleGrantKind,
//				Data: &OnSchemaObjectGrantData{
//					Kind: OnAllSchemaObjectGrantKind,
//					OnAllOrFuture: &BulkOperationGrantData{
//						ObjectNamePlural: sdk.PluralObjectTypeTables,
//						Kind:             InSchemaBulkOperationGrantKind,
//						Schema:           sdk.Pointer(sdk.NewDatabaseObjectIdentifier("database-name", "schema-name")),
//					},
//				},
//			},
//			Expected: `"account-role"|false|false|CREATE SCHEMA,USAGE,MONITOR|OnSchemaObject|OnAll|TABLES|InSchema|"database-name"."schema-name"`,
//		},
//	}
//
//	for _, tt := range testCases {
//		tt := tt
//		t.Run(tt.Name, func(t *testing.T) {
//			assert.Equal(t, tt.Expected, tt.Identifier.String())
//		})
//	}
//}