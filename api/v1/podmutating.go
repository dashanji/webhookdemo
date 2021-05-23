/*
Copyright 2018 The Kubernetes Authors.
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

package v1

import (
	"context"
	"encoding/json"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-pod,mutating=true,failurePolicy=fail,groups=app.rjscy.cn,resources=pods,verbs=create;update,versions=v1,name=mapp.kb.io

// podMutate
type podMutate struct {
	Client  client.Client
	decoder *admission.Decoder
}

// podMutate
func (p *podMutate) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}
	podMutateLog := setupLog.WithName("podMutate")

	err := p.decoder.Decode(req, pod)
	if err != nil {
		podMutateLog.Error(err, "failed decoder pod")
		return admission.Errored(http.StatusBadRequest, err)
	}
	needAddContainer := corev1.Container{
		Name:  "sidecar-nginx",
		Image: "nginx:1.12.2",
	}
	pod.Spec.Containers = append(pod.Spec.Containers, needAddContainer)
	podMutateLog.Info("will inject some concepts")
	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		podMutateLog.Error(err, "failed marshal pod")
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

func (p *podMutate) InjectDecoder(d *admission.Decoder) error {
	p.decoder = d
	return nil
}
