package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = OwnerWithName{}
var _ basetypes.StringValuable = OwnerWithNameValue{}

type OwnerWithName struct {
	basetypes.StringType
}

type OwnerWithNameValue struct {
	basetypes.StringValue
}

// ValueFromTerraform returns a Value given a tftypes.Value. This is
// meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (o OwnerWithName) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := o.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := o.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to OwnerWithNameValue: %v", diags)
	}

	return stringValuable, nil
}

// ValueType should return the attr.Value type returned by
// ValueFromTerraform. The returned attr.Value can be any null, unknown,
// or known value for the type, as this is intended for type detection
// and improving error diagnostics.
func (o OwnerWithName) ValueType(_ context.Context) attr.Value {
	return OwnerWithNameValue{}
}

// Equal should return true if the Type is considered equivalent to the
// Type passed as an argument.
//
// Most types should verify the associated Type is exactly equal to prevent
// potential data consistency issues. For example:
//
//   - basetypes.Number is inequal to basetypes.Int64 or basetypes.Float64
//   - basetypes.String is inequal to a custom Go type that embeds it
func (o OwnerWithName) Equal(oo attr.Type) bool {
	other, ok := oo.(*OwnerWithName)

	if !ok {
		return false
	}

	return o.StringType.Equal(other.StringType)
}

// String should return a human-friendly version of the Type.
func (o OwnerWithName) String() string {
	return "OwnerWithName"
}

// ValueFromString should convert the String to a StringValuable type.
func (o OwnerWithName) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return OwnerWithNameValue{
		StringValue: in,
	}, nil
}

func (o OwnerWithNameValue) Owner() string {
	var splits = strings.Split(o.StringValue.ValueString(), "/")
	if len(splits) < 2 {
		return ""
	}

	return splits[0]
}

func (o OwnerWithNameValue) Name() string {
	var splits = strings.Split(o.StringValue.ValueString(), "/")
	if len(splits) < 2 {
		return ""
	}

	return splits[1]
}
