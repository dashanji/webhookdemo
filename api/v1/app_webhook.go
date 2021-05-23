/*
Copyright 2021 dsj.

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
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var applog = logf.Log.WithName("app-resource")

func (r *App) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-app-rjscy-cn-v1-app,mutating=true,failurePolicy=fail,groups=app.rjscy.cn,resources=apps,verbs=create;update,versions=v1,name=mapp.kb.io

var _ webhook.Defaulter = &App{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *App) Default() {
	applog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	var cns []core.Container
	cns = r.Spec.Deploy.Template.Spec.Containers

	container := core.Container{
		Name:  "sidecar-nginx",
		Image: "nginx:1.12.2",
	}

	cns = append(cns, container)
	r.Spec.Deploy.Template.Spec.Containers = cns
	applog.Info("Default end!")
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-app-rjscy-cn-v1-app,mutating=false,failurePolicy=fail,groups=app.rjscy.cn,resources=apps,versions=v1,name=vapp.kb.io

var _ webhook.Validator = &App{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateCreate() error {
	applog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	/*applog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	var cns []core.Container
	cns = r.Spec.Deploy.Template.Spec.Containers

	container := core.Container{
		Name:  "sidecar-nginx",
		Image: "nginx:1.12.2",
	}

	cns = append(cns, container)
	r.Spec.Deploy.Template.Spec.Containers = cns */
	applog.Info("ValidateCreate end!")
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateUpdate(old runtime.Object) error {
	applog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *App) ValidateDelete() error {
	applog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
