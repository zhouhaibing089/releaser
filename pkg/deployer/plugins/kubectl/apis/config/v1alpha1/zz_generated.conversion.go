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

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/ebay/releaser/pkg/deployer/plugins/kubectl/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*KubectlConfiguration)(nil), (*config.KubectlConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubectlConfiguration_To_config_KubectlConfiguration(a.(*KubectlConfiguration), b.(*config.KubectlConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubectlConfiguration)(nil), (*KubectlConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubectlConfiguration_To_v1alpha1_KubectlConfiguration(a.(*config.KubectlConfiguration), b.(*KubectlConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectReference)(nil), (*config.ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ObjectReference_To_config_ObjectReference(a.(*ObjectReference), b.(*config.ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ObjectReference)(nil), (*ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ObjectReference_To_v1alpha1_ObjectReference(a.(*config.ObjectReference), b.(*ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TemplateOptions)(nil), (*config.TemplateOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TemplateOptions_To_config_TemplateOptions(a.(*TemplateOptions), b.(*config.TemplateOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TemplateOptions)(nil), (*TemplateOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TemplateOptions_To_v1alpha1_TemplateOptions(a.(*config.TemplateOptions), b.(*TemplateOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*YTTOptions)(nil), (*config.YTTOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_YTTOptions_To_config_YTTOptions(a.(*YTTOptions), b.(*config.YTTOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.YTTOptions)(nil), (*YTTOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_YTTOptions_To_v1alpha1_YTTOptions(a.(*config.YTTOptions), b.(*YTTOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.PathOptions)(nil), (*PathOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PathOptions_To_v1alpha1_PathOptions(a.(*config.PathOptions), b.(*PathOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.PruneOptions)(nil), (*PruneOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PruneOptions_To_v1alpha1_PruneOptions(a.(*config.PruneOptions), b.(*PruneOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PathOptions)(nil), (*config.PathOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PathOptions_To_config_PathOptions(a.(*PathOptions), b.(*config.PathOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PruneOptions)(nil), (*config.PruneOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PruneOptions_To_config_PruneOptions(a.(*PruneOptions), b.(*config.PruneOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_KubectlConfiguration_To_config_KubectlConfiguration(in *KubectlConfiguration, out *config.KubectlConfiguration, s conversion.Scope) error {
	out.Command = in.Command
	if err := Convert_v1alpha1_PathOptions_To_config_PathOptions(&in.PathOptions, &out.PathOptions, s); err != nil {
		return err
	}
	out.Template = (*config.TemplateOptions)(unsafe.Pointer(in.Template))
	out.YTT = (*config.YTTOptions)(unsafe.Pointer(in.YTT))
	if in.Prune != nil {
		in, out := &in.Prune, &out.Prune
		*out = new(config.PruneOptions)
		if err := Convert_v1alpha1_PruneOptions_To_config_PruneOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Prune = nil
	}
	out.ExperimentalEnableClusterSize = in.ExperimentalEnableClusterSize
	out.FollowSymlink = in.FollowSymlink
	return nil
}

// Convert_v1alpha1_KubectlConfiguration_To_config_KubectlConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubectlConfiguration_To_config_KubectlConfiguration(in *KubectlConfiguration, out *config.KubectlConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubectlConfiguration_To_config_KubectlConfiguration(in, out, s)
}

func autoConvert_config_KubectlConfiguration_To_v1alpha1_KubectlConfiguration(in *config.KubectlConfiguration, out *KubectlConfiguration, s conversion.Scope) error {
	out.Command = in.Command
	if err := Convert_config_PathOptions_To_v1alpha1_PathOptions(&in.PathOptions, &out.PathOptions, s); err != nil {
		return err
	}
	out.Template = (*TemplateOptions)(unsafe.Pointer(in.Template))
	out.YTT = (*YTTOptions)(unsafe.Pointer(in.YTT))
	if in.Prune != nil {
		in, out := &in.Prune, &out.Prune
		*out = new(PruneOptions)
		if err := Convert_config_PruneOptions_To_v1alpha1_PruneOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Prune = nil
	}
	out.ExperimentalEnableClusterSize = in.ExperimentalEnableClusterSize
	out.FollowSymlink = in.FollowSymlink
	return nil
}

// Convert_config_KubectlConfiguration_To_v1alpha1_KubectlConfiguration is an autogenerated conversion function.
func Convert_config_KubectlConfiguration_To_v1alpha1_KubectlConfiguration(in *config.KubectlConfiguration, out *KubectlConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubectlConfiguration_To_v1alpha1_KubectlConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ObjectReference_To_config_ObjectReference(in *ObjectReference, out *config.ObjectReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_ObjectReference_To_config_ObjectReference is an autogenerated conversion function.
func Convert_v1alpha1_ObjectReference_To_config_ObjectReference(in *ObjectReference, out *config.ObjectReference, s conversion.Scope) error {
	return autoConvert_v1alpha1_ObjectReference_To_config_ObjectReference(in, out, s)
}

func autoConvert_config_ObjectReference_To_v1alpha1_ObjectReference(in *config.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_config_ObjectReference_To_v1alpha1_ObjectReference is an autogenerated conversion function.
func Convert_config_ObjectReference_To_v1alpha1_ObjectReference(in *config.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	return autoConvert_config_ObjectReference_To_v1alpha1_ObjectReference(in, out, s)
}

func autoConvert_v1alpha1_PathOptions_To_config_PathOptions(in *PathOptions, out *config.PathOptions, s conversion.Scope) error {
	// WARNING: in.Path requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_config_PathOptions_To_v1alpha1_PathOptions(in *config.PathOptions, out *PathOptions, s conversion.Scope) error {
	// WARNING: in.Paths requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_PruneOptions_To_config_PruneOptions(in *PruneOptions, out *config.PruneOptions, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	// WARNING: in.Whitelist requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_config_PruneOptions_To_v1alpha1_PruneOptions(in *config.PruneOptions, out *PruneOptions, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	// WARNING: in.SkipList requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_TemplateOptions_To_config_TemplateOptions(in *TemplateOptions, out *config.TemplateOptions, s conversion.Scope) error {
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	out.ConfigMapRefs = *(*[]config.ObjectReference)(unsafe.Pointer(&in.ConfigMapRefs))
	out.SecretRefs = *(*[]config.ObjectReference)(unsafe.Pointer(&in.SecretRefs))
	out.ValueType = config.ValueType(in.ValueType)
	return nil
}

// Convert_v1alpha1_TemplateOptions_To_config_TemplateOptions is an autogenerated conversion function.
func Convert_v1alpha1_TemplateOptions_To_config_TemplateOptions(in *TemplateOptions, out *config.TemplateOptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_TemplateOptions_To_config_TemplateOptions(in, out, s)
}

func autoConvert_config_TemplateOptions_To_v1alpha1_TemplateOptions(in *config.TemplateOptions, out *TemplateOptions, s conversion.Scope) error {
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	out.ConfigMapRefs = *(*[]ObjectReference)(unsafe.Pointer(&in.ConfigMapRefs))
	out.SecretRefs = *(*[]ObjectReference)(unsafe.Pointer(&in.SecretRefs))
	out.ValueType = ValueType(in.ValueType)
	return nil
}

// Convert_config_TemplateOptions_To_v1alpha1_TemplateOptions is an autogenerated conversion function.
func Convert_config_TemplateOptions_To_v1alpha1_TemplateOptions(in *config.TemplateOptions, out *TemplateOptions, s conversion.Scope) error {
	return autoConvert_config_TemplateOptions_To_v1alpha1_TemplateOptions(in, out, s)
}

func autoConvert_v1alpha1_YTTOptions_To_config_YTTOptions(in *YTTOptions, out *config.YTTOptions, s conversion.Scope) error {
	out.Overlay = in.Overlay
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	out.ConfigMapRefs = *(*[]config.ObjectReference)(unsafe.Pointer(&in.ConfigMapRefs))
	out.SecretRefs = *(*[]config.ObjectReference)(unsafe.Pointer(&in.SecretRefs))
	out.ValueType = config.ValueType(in.ValueType)
	return nil
}

// Convert_v1alpha1_YTTOptions_To_config_YTTOptions is an autogenerated conversion function.
func Convert_v1alpha1_YTTOptions_To_config_YTTOptions(in *YTTOptions, out *config.YTTOptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_YTTOptions_To_config_YTTOptions(in, out, s)
}

func autoConvert_config_YTTOptions_To_v1alpha1_YTTOptions(in *config.YTTOptions, out *YTTOptions, s conversion.Scope) error {
	out.Overlay = in.Overlay
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	out.ConfigMapRefs = *(*[]ObjectReference)(unsafe.Pointer(&in.ConfigMapRefs))
	out.SecretRefs = *(*[]ObjectReference)(unsafe.Pointer(&in.SecretRefs))
	out.ValueType = ValueType(in.ValueType)
	return nil
}

// Convert_config_YTTOptions_To_v1alpha1_YTTOptions is an autogenerated conversion function.
func Convert_config_YTTOptions_To_v1alpha1_YTTOptions(in *config.YTTOptions, out *YTTOptions, s conversion.Scope) error {
	return autoConvert_config_YTTOptions_To_v1alpha1_YTTOptions(in, out, s)
}
