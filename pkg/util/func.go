package util

import openfunction "github.com/openfunction/apis/core/v1beta2"

func BuildFinished(f *openfunction.Function) bool {
	if f.Status.Build == nil {
		return false
	}
	return f.Status.Build.State == openfunction.Succeeded || f.Status.Build.State == openfunction.Skipped
}
func ServingReady(f *openfunction.Function) bool {
	if f.Status.Serving == nil {
		return false
	}
	return f.Status.Serving.State == openfunction.Running

}
