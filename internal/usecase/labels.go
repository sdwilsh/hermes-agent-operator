package usecase

import (
	"maps"

	agentsv1alpha1 "hermeum/hermes-agent-operator/api/v1alpha1"
)

const (
	labelName      = "app.kubernetes.io/name"
	labelInstance  = "app.kubernetes.io/instance"
	labelManagedBy = "app.kubernetes.io/managed-by"

	appNameValue   = "hermes-agent"
	managedByValue = "hermes-agent-operator"
)

func resourceLabels(ha *agentsv1alpha1.HermesAgent) map[string]string {
	return map[string]string{
		labelName:      appNameValue,
		labelInstance:  ha.Name,
		labelManagedBy: managedByValue,
	}
}

func selectorLabels(ha *agentsv1alpha1.HermesAgent) map[string]string {
	return map[string]string{
		labelName:     appNameValue,
		labelInstance: ha.Name,
	}
}

// podTemplateLabels returns the labels for the agent pod template.  It copies
// `spec.podLabels` first, then the operator-managed labels over the top.  An
// operator-managed label always wins, so an entry in `spec.podLabels` cannot
// shadow a key that the `StatefulSet` pod selector matches.
func podTemplateLabels(ha *agentsv1alpha1.HermesAgent) map[string]string {
	labels := make(map[string]string, len(ha.GetPodLabels())+len(resourceLabels(ha)))
	maps.Copy(labels, ha.GetPodLabels())
	maps.Copy(labels, resourceLabels(ha))
	return labels
}
