// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// APIConverter is an autogenerated mock type for the APIConverter type
type APIConverter struct {
	mock.Mock
}

// MultipleInputFromGraphQL provides a mock function with given fields: in
func (_m *APIConverter) MultipleInputFromGraphQL(in []*graphql.APIDefinitionInput) ([]*model.APIDefinitionInput, []*model.SpecInput, error) {
	ret := _m.Called(in)

	var r0 []*model.APIDefinitionInput
	if rf, ok := ret.Get(0).(func([]*graphql.APIDefinitionInput) []*model.APIDefinitionInput); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.APIDefinitionInput)
		}
	}

	var r1 []*model.SpecInput
	if rf, ok := ret.Get(1).(func([]*graphql.APIDefinitionInput) []*model.SpecInput); ok {
		r1 = rf(in)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*model.SpecInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]*graphql.APIDefinitionInput) error); ok {
		r2 = rf(in)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MultipleToGraphQL provides a mock function with given fields: in, specs
func (_m *APIConverter) MultipleToGraphQL(in []*model.APIDefinition, specs []*model.Spec) ([]*graphql.APIDefinition, error) {
	ret := _m.Called(in, specs)

	var r0 []*graphql.APIDefinition
	if rf, ok := ret.Get(0).(func([]*model.APIDefinition, []*model.Spec) []*graphql.APIDefinition); ok {
		r0 = rf(in, specs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*graphql.APIDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*model.APIDefinition, []*model.Spec) error); ok {
		r1 = rf(in, specs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGraphQL provides a mock function with given fields: in, spec
func (_m *APIConverter) ToGraphQL(in *model.APIDefinition, spec *model.Spec) (*graphql.APIDefinition, error) {
	ret := _m.Called(in, spec)

	var r0 *graphql.APIDefinition
	if rf, ok := ret.Get(0).(func(*model.APIDefinition, *model.Spec) *graphql.APIDefinition); ok {
		r0 = rf(in, spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.APIDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.APIDefinition, *model.Spec) error); ok {
		r1 = rf(in, spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
