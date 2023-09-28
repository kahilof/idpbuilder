package v1alpha1

import (
	"fmt"

	"git.autodesk.com/forge-cd-services/idpbuilder/globals"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArgoPackageConfigSpec Allows for configuration of the ArgoCD Installation.
// If no fields are specified then the binary embedded resources will be used to intall ArgoCD.
type ArgoPackageConfigSpec struct {
	// Enabled controls whether to install ArgoCD.
	Enabled bool `json:"enabled,omitempty"`
}

// EmbeddedArgoApplicationsPackageConfigSpec Controls the installation of the embedded argo applications.
type EmbeddedArgoApplicationsPackageConfigSpec struct {
	// Enabled controls whether to install the embedded argo applications and the associated GitServer
	Enabled bool `json:"enabled,omitempty"`
}

type PackageConfigsSpec struct {
	Argo                     ArgoPackageConfigSpec                     `json:"argoPackageConfigs,omitempty"`
	EmbeddedArgoApplications EmbeddedArgoApplicationsPackageConfigSpec `json:"embeddedArgoApplicationsPackageConfigs,omitempty"`
}

type LocalbuildSpec struct {
	PackageConfigs PackageConfigsSpec `json:"packageConfigs,omitempty"`
}

type LocalbuildStatus struct {
	// ObservedGeneration is the 'Generation' of the Service that was last processed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	GitServerAvailable bool `json:"gitServerAvailable,omitempty"`
	ArgoAvailable      bool `json:"argoAvailable,omitempty"`
	ArgoAppsCreated    bool `json:"argoAppsCreated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=localbuilds,scope=Cluster
type Localbuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocalbuildSpec   `json:"spec,omitempty"`
	Status LocalbuildStatus `json:"status,omitempty"`
}

func (l *Localbuild) GetArgoProjectName() string {
	return fmt.Sprintf("%s-%s-gitserver", globals.ProjectName, l.Name)
}

func (l *Localbuild) GetArgoApplicationName(name string) string {
	return fmt.Sprintf("%s-%s-gitserver-%s", globals.ProjectName, l.Name, name)
}

// +kubebuilder:object:root=true
type LocalbuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Localbuild `json:"items"`
}