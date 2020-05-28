// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	common "github.com/kubeflow/katib/pkg/apis/controller/common/v1beta1"
	"github.com/kubeflow/katib/pkg/controller.v1beta1/consts"
)

func (e *Experiment) SetDefault() {
	e.setDefaultParallelTrialCount()
	e.setDefaultResumePolicy()
	e.setDefaultTrialTemplate()
	e.setDefaultMetricsCollector()
}

func (e *Experiment) setDefaultParallelTrialCount() {
	if e.Spec.ParallelTrialCount == nil {
		e.Spec.ParallelTrialCount = new(int32)
		*e.Spec.ParallelTrialCount = DefaultTrialParallelCount
	}
}

func (e *Experiment) setDefaultResumePolicy() {
	if e.Spec.ResumePolicy == "" {
		e.Spec.ResumePolicy = DefaultResumePolicy
	}
}

func (e *Experiment) setDefaultTrialTemplate() {
	t := e.Spec.TrialTemplate
	if t == nil {
		t = &TrialTemplate{
			Retain: true,
		}
	}
	if t.GoTemplate == nil {
		t.GoTemplate = &GoTemplate{}
	}
	if t.GoTemplate.RawTemplate == "" && t.GoTemplate.TemplateSpec == nil {
		t.GoTemplate.TemplateSpec = &TemplateSpec{
			ConfigMapNamespace: consts.DefaultKatibNamespace,
			ConfigMapName:      DefaultTrialConfigMapName,
			TemplatePath:       DefaultTrialTemplatePath,
		}
	}
	e.Spec.TrialTemplate = t
}

func (e *Experiment) setDefaultMetricsCollector() {
	if e.Spec.MetricsCollectorSpec == nil {
		e.Spec.MetricsCollectorSpec = &common.MetricsCollectorSpec{}
	}
	if e.Spec.MetricsCollectorSpec.Collector == nil {
		e.Spec.MetricsCollectorSpec.Collector = &common.CollectorSpec{
			Kind: common.StdOutCollector,
		}
	}
	switch e.Spec.MetricsCollectorSpec.Collector.Kind {
	case common.PrometheusMetricCollector:
		if e.Spec.MetricsCollectorSpec.Source == nil {
			e.Spec.MetricsCollectorSpec.Source = &common.SourceSpec{}
		}
		if e.Spec.MetricsCollectorSpec.Source.HttpGet == nil {
			e.Spec.MetricsCollectorSpec.Source.HttpGet = &v1.HTTPGetAction{}
		}
		if e.Spec.MetricsCollectorSpec.Source.HttpGet.Path == "" {
			e.Spec.MetricsCollectorSpec.Source.HttpGet.Path = common.DefaultPrometheusPath
		}
		if e.Spec.MetricsCollectorSpec.Source.HttpGet.Port.String() == "0" {
			e.Spec.MetricsCollectorSpec.Source.HttpGet.Port = intstr.FromInt(common.DefaultPrometheusPort)
		}
	case common.FileCollector:
		if e.Spec.MetricsCollectorSpec.Source == nil {
			e.Spec.MetricsCollectorSpec.Source = &common.SourceSpec{}
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath == nil {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath = &common.FileSystemPath{}
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Kind == "" {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Kind = common.FileKind
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Path == "" {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Path = common.DefaultFilePath
		}
	case common.TfEventCollector:
		if e.Spec.MetricsCollectorSpec.Source == nil {
			e.Spec.MetricsCollectorSpec.Source = &common.SourceSpec{}
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath == nil {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath = &common.FileSystemPath{}
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Kind == "" {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Kind = common.DirectoryKind
		}
		if e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Path == "" {
			e.Spec.MetricsCollectorSpec.Source.FileSystemPath.Path = common.DefaultTensorflowEventDirPath
		}
	}
}
