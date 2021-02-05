// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	CacheClusterID     *string `json:"cacheClusterID,omitempty"`
	KMSKeyID           *string `json:"kmsKeyID,omitempty"`
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`
	// +kubebuilder:validation:Required
	SnapshotName       *string `json:"snapshotName"`
	SourceSnapshotName *string `json:"sourceSnapshotName,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot
type SnapshotStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions                  []*ackv1alpha1.Condition `json:"conditions"`
	AutoMinorVersionUpgrade     *bool                    `json:"autoMinorVersionUpgrade,omitempty"`
	AutomaticFailover           *string                  `json:"automaticFailover,omitempty"`
	CacheClusterCreateTime      *metav1.Time             `json:"cacheClusterCreateTime,omitempty"`
	CacheNodeType               *string                  `json:"cacheNodeType,omitempty"`
	CacheParameterGroupName     *string                  `json:"cacheParameterGroupName,omitempty"`
	CacheSubnetGroupName        *string                  `json:"cacheSubnetGroupName,omitempty"`
	Engine                      *string                  `json:"engine,omitempty"`
	EngineVersion               *string                  `json:"engineVersion,omitempty"`
	NodeSnapshots               []*NodeSnapshot          `json:"nodeSnapshots,omitempty"`
	NumCacheNodes               *int64                   `json:"numCacheNodes,omitempty"`
	NumNodeGroups               *int64                   `json:"numNodeGroups,omitempty"`
	Port                        *int64                   `json:"port,omitempty"`
	PreferredAvailabilityZone   *string                  `json:"preferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow  *string                  `json:"preferredMaintenanceWindow,omitempty"`
	PreferredOutpostARN         *string                  `json:"preferredOutpostARN,omitempty"`
	ReplicationGroupDescription *string                  `json:"replicationGroupDescription,omitempty"`
	SnapshotRetentionLimit      *int64                   `json:"snapshotRetentionLimit,omitempty"`
	SnapshotSource              *string                  `json:"snapshotSource,omitempty"`
	SnapshotStatus              *string                  `json:"snapshotStatus,omitempty"`
	SnapshotWindow              *string                  `json:"snapshotWindow,omitempty"`
	TopicARN                    *string                  `json:"topicARN,omitempty"`
	VPCID                       *string                  `json:"vpcID,omitempty"`
}

// Snapshot is the Schema for the Snapshots API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// SnapshotList contains a list of Snapshot
// +kubebuilder:object:root=true
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
