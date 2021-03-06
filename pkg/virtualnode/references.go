package virtualnode

import (
	appmesh "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
)

// ExtractVirtualServiceReferences extracts all virtualServiceReferences for this VirtualNode
func ExtractVirtualServiceReferences(vn *appmesh.VirtualNode) []appmesh.VirtualServiceReference {
	var vsRefs []appmesh.VirtualServiceReference
	for _, backend := range vn.Spec.Backends {
		if backend.VirtualService.VirtualServiceRef != nil {
			vsRefs = append(vsRefs, *backend.VirtualService.VirtualServiceRef)
		}
	}
	return vsRefs
}
