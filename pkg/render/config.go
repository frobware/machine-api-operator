package render

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// Kind is the TypeMeta.Kind for the OperatorConfig.
	Kind = "MachineAPIOperatorConfig"
	// APIVersion is the TypeMeta.APIVersion for the OperatorConfig.
	APIVersion = "v1"
)

// OperatorConfig contains configuration for MAO
type OperatorConfig struct {
	metav1.TypeMeta `json:",inline"`
	TargetNamespace string         `json:"targetNamespace"`
	APIServiceCA    string         `json:"apiServiceCA"`
	Provider        string         `json:"provider"`
	AWS             *AWSConfig     `json:"aws"`
	Libvirt         *LibvirtConfig `json:"libvirt"`
}

// LibvirtConfig contains specific config for Libvirt
type LibvirtConfig struct {
	URI         string `json:"uri"`
	NetworkName string `json:"networkName"`
	IPRange     string `json:"ipRange"`
	ClusterName string `json:"clusterName"`
	Replicas    string `json:"replicas"`
}

// AWSConfig contains specific config for AWS
type AWSConfig struct {
	ClusterName           string `json:"clusterName"`
	ClusterID             string `json:"clusterID"`
	Region                string `json:"region"`
	AvailabilityZone      string `json:"availabilityZone"`
	Image                 string `json:"image"`
	ReleaseChannel        string `json:"releaseChannel"`
	ContainerLinuxVersion string `json:"containerLinuxVersion"`
	Replicas              string `json:"replicas"`
	WithCreds             bool   `json:"withCreds"`
}
