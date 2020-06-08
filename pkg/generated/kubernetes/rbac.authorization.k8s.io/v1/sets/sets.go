// Code generated by skv2. DO NOT EDIT.

package v1sets

import (
	. "k8s.io/api/rbac/v1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type RoleSet interface {
	Keys() sets.String
	List() []*Role
	Map() map[string]*Role
	Insert(role ...*Role)
	Equal(roleSet RoleSet) bool
	Has(role *Role) bool
	Delete(role *Role)
	Union(set RoleSet) RoleSet
	Difference(set RoleSet) RoleSet
	Intersection(set RoleSet) RoleSet
}

func makeGenericRoleSet(roleList []*Role) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range roleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type roleSet struct {
	set sksets.ResourceSet
}

func NewRoleSet(roleList ...*Role) RoleSet {
	return &roleSet{set: makeGenericRoleSet(roleList)}
}

func (s roleSet) Keys() sets.String {
	return s.set.Keys()
}

func (s roleSet) List() []*Role {
	var roleList []*Role
	for _, obj := range s.set.List() {
		roleList = append(roleList, obj.(*Role))
	}
	return roleList
}

func (s roleSet) Map() map[string]*Role {
	newMap := map[string]*Role{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*Role)
	}
	return newMap
}

func (s roleSet) Insert(
	roleList ...*Role,
) {
	for _, obj := range roleList {
		s.set.Insert(obj)
	}
}

func (s roleSet) Has(role *Role) bool {
	return s.set.Has(role)
}

func (s roleSet) Equal(
	roleSet RoleSet,
) bool {
	return s.set.Equal(makeGenericRoleSet(roleSet.List()))
}

func (s roleSet) Delete(Role *Role) {
	s.set.Delete(Role)
}

func (s roleSet) Union(set RoleSet) RoleSet {
	return NewRoleSet(append(s.List(), set.List()...)...)
}

func (s roleSet) Difference(set RoleSet) RoleSet {
	newSet := s.set.Difference(makeGenericRoleSet(set.List()))
	return roleSet{set: newSet}
}

func (s roleSet) Intersection(set RoleSet) RoleSet {
	newSet := s.set.Intersection(makeGenericRoleSet(set.List()))
	var roleList []*Role
	for _, obj := range newSet.List() {
		roleList = append(roleList, obj.(*Role))
	}
	return NewRoleSet(roleList...)
}

type RoleBindingSet interface {
	Keys() sets.String
	List() []*RoleBinding
	Map() map[string]*RoleBinding
	Insert(roleBinding ...*RoleBinding)
	Equal(roleBindingSet RoleBindingSet) bool
	Has(roleBinding *RoleBinding) bool
	Delete(roleBinding *RoleBinding)
	Union(set RoleBindingSet) RoleBindingSet
	Difference(set RoleBindingSet) RoleBindingSet
	Intersection(set RoleBindingSet) RoleBindingSet
}

func makeGenericRoleBindingSet(roleBindingList []*RoleBinding) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range roleBindingList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type roleBindingSet struct {
	set sksets.ResourceSet
}

func NewRoleBindingSet(roleBindingList ...*RoleBinding) RoleBindingSet {
	return &roleBindingSet{set: makeGenericRoleBindingSet(roleBindingList)}
}

func (s roleBindingSet) Keys() sets.String {
	return s.set.Keys()
}

func (s roleBindingSet) List() []*RoleBinding {
	var roleBindingList []*RoleBinding
	for _, obj := range s.set.List() {
		roleBindingList = append(roleBindingList, obj.(*RoleBinding))
	}
	return roleBindingList
}

func (s roleBindingSet) Map() map[string]*RoleBinding {
	newMap := map[string]*RoleBinding{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*RoleBinding)
	}
	return newMap
}

func (s roleBindingSet) Insert(
	roleBindingList ...*RoleBinding,
) {
	for _, obj := range roleBindingList {
		s.set.Insert(obj)
	}
}

func (s roleBindingSet) Has(roleBinding *RoleBinding) bool {
	return s.set.Has(roleBinding)
}

func (s roleBindingSet) Equal(
	roleBindingSet RoleBindingSet,
) bool {
	return s.set.Equal(makeGenericRoleBindingSet(roleBindingSet.List()))
}

func (s roleBindingSet) Delete(RoleBinding *RoleBinding) {
	s.set.Delete(RoleBinding)
}

func (s roleBindingSet) Union(set RoleBindingSet) RoleBindingSet {
	return NewRoleBindingSet(append(s.List(), set.List()...)...)
}

func (s roleBindingSet) Difference(set RoleBindingSet) RoleBindingSet {
	newSet := s.set.Difference(makeGenericRoleBindingSet(set.List()))
	return roleBindingSet{set: newSet}
}

func (s roleBindingSet) Intersection(set RoleBindingSet) RoleBindingSet {
	newSet := s.set.Intersection(makeGenericRoleBindingSet(set.List()))
	var roleBindingList []*RoleBinding
	for _, obj := range newSet.List() {
		roleBindingList = append(roleBindingList, obj.(*RoleBinding))
	}
	return NewRoleBindingSet(roleBindingList...)
}

type ClusterRoleSet interface {
	Keys() sets.String
	List() []*ClusterRole
	Map() map[string]*ClusterRole
	Insert(clusterRole ...*ClusterRole)
	Equal(clusterRoleSet ClusterRoleSet) bool
	Has(clusterRole *ClusterRole) bool
	Delete(clusterRole *ClusterRole)
	Union(set ClusterRoleSet) ClusterRoleSet
	Difference(set ClusterRoleSet) ClusterRoleSet
	Intersection(set ClusterRoleSet) ClusterRoleSet
}

func makeGenericClusterRoleSet(clusterRoleList []*ClusterRole) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range clusterRoleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterRoleSet struct {
	set sksets.ResourceSet
}

func NewClusterRoleSet(clusterRoleList ...*ClusterRole) ClusterRoleSet {
	return &clusterRoleSet{set: makeGenericClusterRoleSet(clusterRoleList)}
}

func (s clusterRoleSet) Keys() sets.String {
	return s.set.Keys()
}

func (s clusterRoleSet) List() []*ClusterRole {
	var clusterRoleList []*ClusterRole
	for _, obj := range s.set.List() {
		clusterRoleList = append(clusterRoleList, obj.(*ClusterRole))
	}
	return clusterRoleList
}

func (s clusterRoleSet) Map() map[string]*ClusterRole {
	newMap := map[string]*ClusterRole{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*ClusterRole)
	}
	return newMap
}

func (s clusterRoleSet) Insert(
	clusterRoleList ...*ClusterRole,
) {
	for _, obj := range clusterRoleList {
		s.set.Insert(obj)
	}
}

func (s clusterRoleSet) Has(clusterRole *ClusterRole) bool {
	return s.set.Has(clusterRole)
}

func (s clusterRoleSet) Equal(
	clusterRoleSet ClusterRoleSet,
) bool {
	return s.set.Equal(makeGenericClusterRoleSet(clusterRoleSet.List()))
}

func (s clusterRoleSet) Delete(ClusterRole *ClusterRole) {
	s.set.Delete(ClusterRole)
}

func (s clusterRoleSet) Union(set ClusterRoleSet) ClusterRoleSet {
	return NewClusterRoleSet(append(s.List(), set.List()...)...)
}

func (s clusterRoleSet) Difference(set ClusterRoleSet) ClusterRoleSet {
	newSet := s.set.Difference(makeGenericClusterRoleSet(set.List()))
	return clusterRoleSet{set: newSet}
}

func (s clusterRoleSet) Intersection(set ClusterRoleSet) ClusterRoleSet {
	newSet := s.set.Intersection(makeGenericClusterRoleSet(set.List()))
	var clusterRoleList []*ClusterRole
	for _, obj := range newSet.List() {
		clusterRoleList = append(clusterRoleList, obj.(*ClusterRole))
	}
	return NewClusterRoleSet(clusterRoleList...)
}

type ClusterRoleBindingSet interface {
	Keys() sets.String
	List() []*ClusterRoleBinding
	Map() map[string]*ClusterRoleBinding
	Insert(clusterRoleBinding ...*ClusterRoleBinding)
	Equal(clusterRoleBindingSet ClusterRoleBindingSet) bool
	Has(clusterRoleBinding *ClusterRoleBinding) bool
	Delete(clusterRoleBinding *ClusterRoleBinding)
	Union(set ClusterRoleBindingSet) ClusterRoleBindingSet
	Difference(set ClusterRoleBindingSet) ClusterRoleBindingSet
	Intersection(set ClusterRoleBindingSet) ClusterRoleBindingSet
}

func makeGenericClusterRoleBindingSet(clusterRoleBindingList []*ClusterRoleBinding) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range clusterRoleBindingList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterRoleBindingSet struct {
	set sksets.ResourceSet
}

func NewClusterRoleBindingSet(clusterRoleBindingList ...*ClusterRoleBinding) ClusterRoleBindingSet {
	return &clusterRoleBindingSet{set: makeGenericClusterRoleBindingSet(clusterRoleBindingList)}
}

func (s clusterRoleBindingSet) Keys() sets.String {
	return s.set.Keys()
}

func (s clusterRoleBindingSet) List() []*ClusterRoleBinding {
	var clusterRoleBindingList []*ClusterRoleBinding
	for _, obj := range s.set.List() {
		clusterRoleBindingList = append(clusterRoleBindingList, obj.(*ClusterRoleBinding))
	}
	return clusterRoleBindingList
}

func (s clusterRoleBindingSet) Map() map[string]*ClusterRoleBinding {
	newMap := map[string]*ClusterRoleBinding{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*ClusterRoleBinding)
	}
	return newMap
}

func (s clusterRoleBindingSet) Insert(
	clusterRoleBindingList ...*ClusterRoleBinding,
) {
	for _, obj := range clusterRoleBindingList {
		s.set.Insert(obj)
	}
}

func (s clusterRoleBindingSet) Has(clusterRoleBinding *ClusterRoleBinding) bool {
	return s.set.Has(clusterRoleBinding)
}

func (s clusterRoleBindingSet) Equal(
	clusterRoleBindingSet ClusterRoleBindingSet,
) bool {
	return s.set.Equal(makeGenericClusterRoleBindingSet(clusterRoleBindingSet.List()))
}

func (s clusterRoleBindingSet) Delete(ClusterRoleBinding *ClusterRoleBinding) {
	s.set.Delete(ClusterRoleBinding)
}

func (s clusterRoleBindingSet) Union(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	return NewClusterRoleBindingSet(append(s.List(), set.List()...)...)
}

func (s clusterRoleBindingSet) Difference(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	newSet := s.set.Difference(makeGenericClusterRoleBindingSet(set.List()))
	return clusterRoleBindingSet{set: newSet}
}

func (s clusterRoleBindingSet) Intersection(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	newSet := s.set.Intersection(makeGenericClusterRoleBindingSet(set.List()))
	var clusterRoleBindingList []*ClusterRoleBinding
	for _, obj := range newSet.List() {
		clusterRoleBindingList = append(clusterRoleBindingList, obj.(*ClusterRoleBinding))
	}
	return NewClusterRoleBindingSet(clusterRoleBindingList...)
}