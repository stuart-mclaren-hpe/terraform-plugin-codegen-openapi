// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oas

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

func (s *OASSchema) BuildIntegerResource(name string, computability schema.ComputedOptionalRequired) (*resource.Attribute, error) {
	return &resource.Attribute{
		Name: name,
		Int64: &resource.Int64Attribute{
			ComputedOptionalRequired: computability,
			Description:              s.GetDescription(),
		},
	}, nil
}

func (s *OASSchema) BuildIntegerDataSource(name string, computability schema.ComputedOptionalRequired) (*datasource.Attribute, error) {
	return &datasource.Attribute{
		Name: name,
		Int64: &datasource.Int64Attribute{
			ComputedOptionalRequired: computability,
			Description:              s.GetDescription(),
		},
	}, nil
}

func (s *OASSchema) BuildIntegerElementType() (schema.ElementType, error) {
	return schema.ElementType{
		Int64: &schema.Int64Type{},
	}, nil
}