package main

import "github.com/hashicorp/terraform-plugin-framework/types"

type Repository struct {
	FullName    OwnerWithNameValue `tfsdk:"full_name"`
	Description types.String       `tfsdk:"description"`
}

func (r Repository) Owner() string {
	return r.FullName.Owner()
}

func (r Repository) Name() string {
	return r.FullName.Name()
}
