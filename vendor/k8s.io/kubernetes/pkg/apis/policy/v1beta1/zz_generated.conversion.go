//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	policyv1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	policy "k8s.io/kubernetes/pkg/apis/policy"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*policyv1beta1.Eviction)(nil), (*policy.Eviction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Eviction_To_policy_Eviction(a.(*policyv1beta1.Eviction), b.(*policy.Eviction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policy.Eviction)(nil), (*policyv1beta1.Eviction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_policy_Eviction_To_v1beta1_Eviction(a.(*policy.Eviction), b.(*policyv1beta1.Eviction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policyv1beta1.PodDisruptionBudgetList)(nil), (*policy.PodDisruptionBudgetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(a.(*policyv1beta1.PodDisruptionBudgetList), b.(*policy.PodDisruptionBudgetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetList)(nil), (*policyv1beta1.PodDisruptionBudgetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(a.(*policy.PodDisruptionBudgetList), b.(*policyv1beta1.PodDisruptionBudgetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policyv1beta1.PodDisruptionBudgetSpec)(nil), (*policy.PodDisruptionBudgetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(a.(*policyv1beta1.PodDisruptionBudgetSpec), b.(*policy.PodDisruptionBudgetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetSpec)(nil), (*policyv1beta1.PodDisruptionBudgetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(a.(*policy.PodDisruptionBudgetSpec), b.(*policyv1beta1.PodDisruptionBudgetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policyv1beta1.PodDisruptionBudgetStatus)(nil), (*policy.PodDisruptionBudgetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(a.(*policyv1beta1.PodDisruptionBudgetStatus), b.(*policy.PodDisruptionBudgetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*policy.PodDisruptionBudgetStatus)(nil), (*policyv1beta1.PodDisruptionBudgetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(a.(*policy.PodDisruptionBudgetStatus), b.(*policyv1beta1.PodDisruptionBudgetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*policy.PodDisruptionBudget)(nil), (*policyv1beta1.PodDisruptionBudget)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(a.(*policy.PodDisruptionBudget), b.(*policyv1beta1.PodDisruptionBudget), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*policyv1beta1.PodDisruptionBudget)(nil), (*policy.PodDisruptionBudget)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(a.(*policyv1beta1.PodDisruptionBudget), b.(*policy.PodDisruptionBudget), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Eviction_To_policy_Eviction(in *policyv1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_v1beta1_Eviction_To_policy_Eviction is an autogenerated conversion function.
func Convert_v1beta1_Eviction_To_policy_Eviction(in *policyv1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	return autoConvert_v1beta1_Eviction_To_policy_Eviction(in, out, s)
}

func autoConvert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *policyv1beta1.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_policy_Eviction_To_v1beta1_Eviction is an autogenerated conversion function.
func Convert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *policyv1beta1.Eviction, s conversion.Scope) error {
	return autoConvert_policy_Eviction_To_v1beta1_Eviction(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *policyv1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *policyv1beta1.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *policyv1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]policy.PodDisruptionBudget, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *policyv1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *policyv1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]policyv1beta1.PodDisruptionBudget, len(*in))
		for i := range *in {
			if err := Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *policyv1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *policyv1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.UnhealthyPodEvictionPolicy = (*policy.UnhealthyPodEvictionPolicyType)(unsafe.Pointer(in.UnhealthyPodEvictionPolicy))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *policyv1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *policyv1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.UnhealthyPodEvictionPolicy = (*policyv1beta1.UnhealthyPodEvictionPolicyType)(unsafe.Pointer(in.UnhealthyPodEvictionPolicy))
	return nil
}

// Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *policyv1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *policyv1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.DisruptionsAllowed = in.DisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *policyv1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *policyv1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.DisruptionsAllowed = in.DisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *policyv1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in, out, s)
}
