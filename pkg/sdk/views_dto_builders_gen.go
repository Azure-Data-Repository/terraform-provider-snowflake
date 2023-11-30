// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateViewRequest(
	name SchemaObjectIdentifier,
	sql string,
) *CreateViewRequest {
	s := CreateViewRequest{}
	s.name = name
	s.sql = sql
	return &s
}

func (s *CreateViewRequest) WithOrReplace(OrReplace *bool) *CreateViewRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateViewRequest) WithSecure(Secure *bool) *CreateViewRequest {
	s.Secure = Secure
	return s
}

func (s *CreateViewRequest) WithTemporary(Temporary *bool) *CreateViewRequest {
	s.Temporary = Temporary
	return s
}

func (s *CreateViewRequest) WithRecursive(Recursive *bool) *CreateViewRequest {
	s.Recursive = Recursive
	return s
}

func (s *CreateViewRequest) WithIfNotExists(IfNotExists *bool) *CreateViewRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateViewRequest) WithColumns(Columns []ViewColumnRequest) *CreateViewRequest {
	s.Columns = Columns
	return s
}

func (s *CreateViewRequest) WithColumnsMaskingPolicies(ColumnsMaskingPolicies []ViewColumnMaskingPolicyRequest) *CreateViewRequest {
	s.ColumnsMaskingPolicies = ColumnsMaskingPolicies
	return s
}

func (s *CreateViewRequest) WithCopyGrants(CopyGrants *bool) *CreateViewRequest {
	s.CopyGrants = CopyGrants
	return s
}

func (s *CreateViewRequest) WithComment(Comment *string) *CreateViewRequest {
	s.Comment = Comment
	return s
}

func (s *CreateViewRequest) WithRowAccessPolicy(RowAccessPolicy *ViewRowAccessPolicyRequest) *CreateViewRequest {
	s.RowAccessPolicy = RowAccessPolicy
	return s
}

func (s *CreateViewRequest) WithTag(Tag []TagAssociation) *CreateViewRequest {
	s.Tag = Tag
	return s
}

func NewViewColumnRequest(
	Name string,
) *ViewColumnRequest {
	s := ViewColumnRequest{}
	s.Name = Name
	return &s
}

func (s *ViewColumnRequest) WithComment(Comment *string) *ViewColumnRequest {
	s.Comment = Comment
	return s
}

func NewViewColumnMaskingPolicyRequest(
	Name string,
	MaskingPolicy SchemaObjectIdentifier,
) *ViewColumnMaskingPolicyRequest {
	s := ViewColumnMaskingPolicyRequest{}
	s.Name = Name
	s.MaskingPolicy = MaskingPolicy
	return &s
}

func (s *ViewColumnMaskingPolicyRequest) WithUsing(Using []string) *ViewColumnMaskingPolicyRequest {
	s.Using = Using
	return s
}

func (s *ViewColumnMaskingPolicyRequest) WithTag(Tag []TagAssociation) *ViewColumnMaskingPolicyRequest {
	s.Tag = Tag
	return s
}

func NewViewRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
	On []string,
) *ViewRowAccessPolicyRequest {
	s := ViewRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	s.On = On
	return &s
}

func NewAlterViewRequest(
	name SchemaObjectIdentifier,
) *AlterViewRequest {
	s := AlterViewRequest{}
	s.name = name
	return &s
}

func (s *AlterViewRequest) WithIfExists(IfExists *bool) *AlterViewRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterViewRequest) WithRenameTo(RenameTo *SchemaObjectIdentifier) *AlterViewRequest {
	s.RenameTo = RenameTo
	return s
}

func (s *AlterViewRequest) WithSetComment(SetComment *string) *AlterViewRequest {
	s.SetComment = SetComment
	return s
}

func (s *AlterViewRequest) WithUnsetComment(UnsetComment *bool) *AlterViewRequest {
	s.UnsetComment = UnsetComment
	return s
}

func (s *AlterViewRequest) WithSetSecure(SetSecure *bool) *AlterViewRequest {
	s.SetSecure = SetSecure
	return s
}

func (s *AlterViewRequest) WithSetChangeTracking(SetChangeTracking *bool) *AlterViewRequest {
	s.SetChangeTracking = SetChangeTracking
	return s
}

func (s *AlterViewRequest) WithUnsetSecure(UnsetSecure *bool) *AlterViewRequest {
	s.UnsetSecure = UnsetSecure
	return s
}

func (s *AlterViewRequest) WithSetTags(SetTags []TagAssociation) *AlterViewRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterViewRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterViewRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterViewRequest) WithAddRowAccessPolicy(AddRowAccessPolicy *ViewAddRowAccessPolicyRequest) *AlterViewRequest {
	s.AddRowAccessPolicy = AddRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropRowAccessPolicy(DropRowAccessPolicy *ViewDropRowAccessPolicyRequest) *AlterViewRequest {
	s.DropRowAccessPolicy = DropRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropAndAddRowAccessPolicy(DropAndAddRowAccessPolicy *ViewDropAndAddRowAccessPolicyRequest) *AlterViewRequest {
	s.DropAndAddRowAccessPolicy = DropAndAddRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropAllRowAccessPolicies(DropAllRowAccessPolicies *bool) *AlterViewRequest {
	s.DropAllRowAccessPolicies = DropAllRowAccessPolicies
	return s
}

func (s *AlterViewRequest) WithSetMaskingPolicyOnColumn(SetMaskingPolicyOnColumn *ViewSetColumnMaskingPolicyRequest) *AlterViewRequest {
	s.SetMaskingPolicyOnColumn = SetMaskingPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithUnsetMaskingPolicyOnColumn(UnsetMaskingPolicyOnColumn *ViewUnsetColumnMaskingPolicyRequest) *AlterViewRequest {
	s.UnsetMaskingPolicyOnColumn = UnsetMaskingPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithSetTagsOnColumn(SetTagsOnColumn *ViewSetColumnTagsRequest) *AlterViewRequest {
	s.SetTagsOnColumn = SetTagsOnColumn
	return s
}

func (s *AlterViewRequest) WithUnsetTagsOnColumn(UnsetTagsOnColumn *ViewUnsetColumnTagsRequest) *AlterViewRequest {
	s.UnsetTagsOnColumn = UnsetTagsOnColumn
	return s
}

func NewViewAddRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
	On []string,
) *ViewAddRowAccessPolicyRequest {
	s := ViewAddRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	s.On = On
	return &s
}

func NewViewDropRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
) *ViewDropRowAccessPolicyRequest {
	s := ViewDropRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	return &s
}

func NewViewDropAndAddRowAccessPolicyRequest(
	Drop ViewDropRowAccessPolicyRequest,
	Add ViewAddRowAccessPolicyRequest,
) *ViewDropAndAddRowAccessPolicyRequest {
	s := ViewDropAndAddRowAccessPolicyRequest{}
	s.Drop = Drop
	s.Add = Add
	return &s
}

func NewViewSetColumnMaskingPolicyRequest(
	Name string,
	MaskingPolicy SchemaObjectIdentifier,
) *ViewSetColumnMaskingPolicyRequest {
	s := ViewSetColumnMaskingPolicyRequest{}
	s.Name = Name
	s.MaskingPolicy = MaskingPolicy
	return &s
}

func (s *ViewSetColumnMaskingPolicyRequest) WithUsing(Using []string) *ViewSetColumnMaskingPolicyRequest {
	s.Using = Using
	return s
}

func (s *ViewSetColumnMaskingPolicyRequest) WithForce(Force *bool) *ViewSetColumnMaskingPolicyRequest {
	s.Force = Force
	return s
}

func NewViewUnsetColumnMaskingPolicyRequest(
	Name string,
) *ViewUnsetColumnMaskingPolicyRequest {
	s := ViewUnsetColumnMaskingPolicyRequest{}
	s.Name = Name
	return &s
}

func NewViewSetColumnTagsRequest(
	Name string,
	SetTags []TagAssociation,
) *ViewSetColumnTagsRequest {
	s := ViewSetColumnTagsRequest{}
	s.Name = Name
	s.SetTags = SetTags
	return &s
}

func NewViewUnsetColumnTagsRequest(
	Name string,
	UnsetTags []ObjectIdentifier,
) *ViewUnsetColumnTagsRequest {
	s := ViewUnsetColumnTagsRequest{}
	s.Name = Name
	s.UnsetTags = UnsetTags
	return &s
}

func NewDropViewRequest(
	name SchemaObjectIdentifier,
) *DropViewRequest {
	s := DropViewRequest{}
	s.name = name
	return &s
}

func (s *DropViewRequest) WithIfExists(IfExists *bool) *DropViewRequest {
	s.IfExists = IfExists
	return s
}

func NewShowViewRequest() *ShowViewRequest {
	return &ShowViewRequest{}
}

func (s *ShowViewRequest) WithTerse(Terse *bool) *ShowViewRequest {
	s.Terse = Terse
	return s
}

func (s *ShowViewRequest) WithLike(Like *Like) *ShowViewRequest {
	s.Like = Like
	return s
}

func (s *ShowViewRequest) WithIn(In *In) *ShowViewRequest {
	s.In = In
	return s
}

func (s *ShowViewRequest) WithStartsWith(StartsWith *string) *ShowViewRequest {
	s.StartsWith = StartsWith
	return s
}

func (s *ShowViewRequest) WithLimit(Limit *LimitFrom) *ShowViewRequest {
	s.Limit = Limit
	return s
}

func NewDescribeViewRequest(
	name SchemaObjectIdentifier,
) *DescribeViewRequest {
	s := DescribeViewRequest{}
	s.name = name
	return &s
}