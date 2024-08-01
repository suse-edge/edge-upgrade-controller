package upgrade

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

const (
	ChartNamespace = "kube-system"
)

func ChartNamespacedName(chart string) types.NamespacedName {
	return types.NamespacedName{
		Name:      chart,
		Namespace: ChartNamespace,
	}
}

type HelmChartState int

const (
	ChartStateUnknown HelmChartState = iota
	ChartStateNotInstalled
	ChartStateVersionAlreadyInstalled
	ChartStateInProgress
	ChartStateFailed
	ChartStateSucceeded
)

func (s HelmChartState) Message() string {
	switch s {
	case ChartStateUnknown:
		return "Chart state is unknown"
	case ChartStateNotInstalled:
		return "Chart is not installed"
	case ChartStateVersionAlreadyInstalled:
		return "Chart version is already installed"
	case ChartStateInProgress:
		return "Chart upgrade is in progress"
	case ChartStateFailed:
		return "Chart upgrade failed"
	case ChartStateSucceeded:
		return "Chart upgrade succeeded"
	default:
		return ""
	}
}

func (s HelmChartState) FormattedMessage(chart string) string {
	switch s {
	case ChartStateUnknown:
		return fmt.Sprintf("State of chart %s is unknown", chart)
	case ChartStateNotInstalled:
		return fmt.Sprintf("Chart %s is not installed", chart)
	case ChartStateVersionAlreadyInstalled:
		return fmt.Sprintf("Specified version of chart %s is already installed", chart)
	case ChartStateInProgress:
		return fmt.Sprintf("Chart %s upgrade is in progress", chart)
	case ChartStateFailed:
		return fmt.Sprintf("Chart %s upgrade failed", chart)
	case ChartStateSucceeded:
		return fmt.Sprintf("Chart %s upgrade succeeded", chart)
	default:
		return ""
	}
}
