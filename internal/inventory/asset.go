// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package inventory

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/samber/lo"
)

// AssetType is used to build the document index.
type AssetType string

const (
	CategoryAccessManagement       AssetType = "Access Management"
	CategoryAccount                AssetType = "Account"
	CategoryContainerRegistry      AssetType = "Container Registry"
	CategoryContainerService       AssetType = "Container Service"
	CategoryDatabase               AssetType = "Database"
	CategoryFaaS                   AssetType = "FaaS"
	CategoryFileSystemService      AssetType = "File System Service"
	CategoryFirewall               AssetType = "Firewall"
	CategoryGateway                AssetType = "Gateway"
	CategoryGroup                  AssetType = "Group"
	CategoryHost                   AssetType = "Host"
	CategoryIdentity               AssetType = "Identity"
	CategoryInfrastructure         AssetType = "Infrastructure"
	CategoryLoadBalancer           AssetType = "Load Balancer"
	CategoryMessagingService       AssetType = "Messaging Service"
	CategoryNetworking             AssetType = "Networking"
	CategoryOrchestrator           AssetType = "Orchestrator"
	CategoryOrganization           AssetType = "Organization"
	CategoryPrivateEndpoint        AssetType = "Private Endpoint"
	CategoryServiceAccount         AssetType = "Service Account"
	CategoryServiceUsageTechnology AssetType = "Service Usage Technology"
	CategorySnapshot               AssetType = "Snapshot"
	CategoryStorageBucket          AssetType = "Storage Bucket"
	CategorySubnet                 AssetType = "Subnet"
	CategoryVolume                 AssetType = "Volume"
	CategoryWebService             AssetType = "Web Service"
)

// AssetSubType is used to build the document index.
type AssetSubType string

const (
	AwsCloudProvider   = "aws"
	AzureCloudProvider = "azure"
	GcpCloudProvider   = "gcp"
)

// AssetClassification holds the taxonomy of an asset
type AssetClassification struct {
	Category AssetType    `json:"type"`
	Type     AssetSubType `json:"sub_type"`
}

// AssetClassifications below are used to generate
// 'internal/inventory/ASSETS.md'. Please keep formatting consistent.
var (
	// AWS
	AssetClassificationAwsEc2Instance              = AssetClassification{CategoryHost, "AWS EC2 Instance"}
	AssetClassificationAwsElbV1                    = AssetClassification{CategoryLoadBalancer, "AWS Elastic Load Balancer"}
	AssetClassificationAwsElbV2                    = AssetClassification{CategoryLoadBalancer, "AWS Elastic Load Balancer v2"}
	AssetClassificationAwsIamPolicy                = AssetClassification{CategoryAccessManagement, "AWS IAM Policy"}
	AssetClassificationAwsIamRole                  = AssetClassification{CategoryServiceAccount, "AWS IAM Role"}
	AssetClassificationAwsIamUser                  = AssetClassification{CategoryIdentity, "AWS IAM User"}
	AssetClassificationAwsLambdaEventSourceMapping = AssetClassification{CategoryFaaS, "AWS Lambda Event Source Mapping"}
	AssetClassificationAwsLambdaFunction           = AssetClassification{CategoryFaaS, "AWS Lambda Function"}
	AssetClassificationAwsLambdaLayer              = AssetClassification{CategoryFaaS, "AWS Lambda Layer"}
	AssetClassificationAwsInternetGateway          = AssetClassification{CategoryGateway, "AWS Internet Gateway"}
	AssetClassificationAwsNatGateway               = AssetClassification{CategoryGateway, "AWS NAT Gateway"}
	AssetClassificationAwsNetworkAcl               = AssetClassification{CategoryNetworking, "AWS EC2 Network ACL"}
	AssetClassificationAwsNetworkInterface         = AssetClassification{CategoryNetworking, "AWS EC2 Network Interface"}
	AssetClassificationAwsSecurityGroup            = AssetClassification{CategoryFirewall, "AWS EC2 Security Group"}
	AssetClassificationAwsSubnet                   = AssetClassification{CategoryNetworking, "AWS EC2 Subnet"}
	AssetClassificationAwsTransitGateway           = AssetClassification{CategoryGateway, "AWS Transit Gateway"}
	AssetClassificationAwsTransitGatewayAttachment = AssetClassification{CategoryGateway, "AWS Transit Gateway Attachment"}
	AssetClassificationAwsVpcPeeringConnection     = AssetClassification{CategoryNetworking, "AWS VPC Peering Connection"}
	AssetClassificationAwsVpc                      = AssetClassification{CategoryNetworking, "AWS VPC"}
	AssetClassificationAwsRds                      = AssetClassification{CategoryDatabase, "AWS RDS Instance"}
	AssetClassificationAwsS3Bucket                 = AssetClassification{CategoryStorageBucket, "AWS S3 Bucket"}
	AssetClassificationAwsSnsTopic                 = AssetClassification{CategoryMessagingService, "AWS SNS Topic"}

	// Azure
	AssetClassificationAzureAppService           = AssetClassification{CategoryWebService, "Azure App Service"}
	AssetClassificationAzureContainerRegistry    = AssetClassification{CategoryContainerRegistry, "Azure Container Registry"}
	AssetClassificationAzureCosmosDBAccount      = AssetClassification{CategoryInfrastructure, "Azure Cosmos DB Account"}
	AssetClassificationAzureCosmosDBSQLDatabase  = AssetClassification{CategoryInfrastructure, "Azure Cosmos DB SQL Database"}
	AssetClassificationAzureDisk                 = AssetClassification{CategoryVolume, "Azure Disk"}
	AssetClassificationAzureElasticPool          = AssetClassification{CategoryDatabase, "Azure Elastic Pool"}
	AssetClassificationAzureEntraGroup           = AssetClassification{CategoryGroup, "Azure Microsoft Entra ID Group"}
	AssetClassificationAzureEntraUser            = AssetClassification{CategoryIdentity, "Azure Microsoft Entra ID User"}
	AssetClassificationAzureResourceGroup        = AssetClassification{CategoryAccessManagement, "Azure Resource Group"}
	AssetClassificationAzureRoleDefinition       = AssetClassification{CategoryAccessManagement, "Azure RoleDefinition"}
	AssetClassificationAzureSQLDatabase          = AssetClassification{CategoryDatabase, "Azure SQL Database"}
	AssetClassificationAzureSQLServer            = AssetClassification{CategoryDatabase, "Azure SQL Server"}
	AssetClassificationAzureServicePrincipal     = AssetClassification{CategoryServiceAccount, "Azure Principal"}
	AssetClassificationAzureSnapshot             = AssetClassification{CategorySnapshot, "Azure Snapshot"}
	AssetClassificationAzureStorageAccount       = AssetClassification{CategoryPrivateEndpoint, "Azure Storage Account"}
	AssetClassificationAzureStorageBlobContainer = AssetClassification{CategoryStorageBucket, "Azure Storage Blob Container"}
	AssetClassificationAzureStorageBlobService   = AssetClassification{CategoryServiceUsageTechnology, "Azure Storage Blob Service"}
	AssetClassificationAzureStorageFileService   = AssetClassification{CategoryFileSystemService, "Azure Storage File Service"}
	AssetClassificationAzureStorageFileShare     = AssetClassification{CategoryFileSystemService, "Azure Storage File Share"}
	AssetClassificationAzureStorageQueue         = AssetClassification{CategoryMessagingService, "Azure Storage Queue"}
	AssetClassificationAzureStorageQueueService  = AssetClassification{CategoryMessagingService, "Azure Storage Queue Service"}
	AssetClassificationAzureStorageTable         = AssetClassification{CategoryDatabase, "Azure Storage Table"}
	AssetClassificationAzureStorageTableService  = AssetClassification{CategoryServiceUsageTechnology, "Azure Storage Table Service"}
	AssetClassificationAzureSubscription         = AssetClassification{CategoryAccessManagement, "Azure Subscription"}
	AssetClassificationAzureTenant               = AssetClassification{CategoryAccessManagement, "Azure Tenant"}
	AssetClassificationAzureVirtualMachine       = AssetClassification{CategoryHost, "Azure Virtual Machine"}

	// GCP
	AssetClassificationGcpProject           = AssetClassification{CategoryAccount, "GCP Project"}
	AssetClassificationGcpOrganization      = AssetClassification{CategoryOrganization, "GCP Organization"}
	AssetClassificationGcpFolder            = AssetClassification{CategoryOrganization, "GCP Folder"}
	AssetClassificationGcpInstance          = AssetClassification{CategoryHost, "GCP Compute Instance"}
	AssetClassificationGcpBucket            = AssetClassification{CategoryStorageBucket, "GCP Bucket"}
	AssetClassificationGcpFirewall          = AssetClassification{CategoryFirewall, "GCP Firewall"}
	AssetClassificationGcpSubnet            = AssetClassification{CategorySubnet, "GCP Subnet"}
	AssetClassificationGcpServiceAccount    = AssetClassification{CategoryAccessManagement, "GCP Service Account"}
	AssetClassificationGcpServiceAccountKey = AssetClassification{CategoryAccessManagement, "GCP Service Account Key"}
	AssetClassificationGcpGkeCluster        = AssetClassification{CategoryOrchestrator, "GCP Kubernetes Engine (GKE) Cluster"}
	AssetClassificationGcpForwardingRule    = AssetClassification{CategoryLoadBalancer, "GCP Load Balancing Forwarding Rule"}
	AssetClassificationGcpIamRole           = AssetClassification{CategoryServiceUsageTechnology, "GCP IAM Role"}
	AssetClassificationGcpCloudFunction     = AssetClassification{CategoryFaaS, "GCP Cloud Function"}
	AssetClassificationGcpCloudRunService   = AssetClassification{CategoryContainerService, "GCP Cloud Run Service"}
	AssetClassificationGcpNetwork           = AssetClassification{CategoryNetworking, "GCP VPC Network"}
)

// AssetEvent holds the whole asset
type AssetEvent struct {
	Entity       Entity
	Event        Event
	Cloud        *Cloud
	Container    *Container
	Fass         *Fass
	Group        *Group
	Host         *Host
	Network      *Network
	Orchestrator *Orchestrator
	Organization *Organization
	URL          *URL
	User         *User
	Labels       map[string]string
	Tags         []string
}

type Organization struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Fass struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type URL struct {
	Full string `json:"full"`
}

// Entity contains the identifiers of the asset
type Entity struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Source *string `json:"source"`
	Raw    *any    `json:"raw"`
	AssetClassification

	// non exported fields
	relatedEntityId []string
}

type Event struct {
	Kind string `json:"kind"`
}

type Network struct {
	Name      string `json:"name,omitempty"`
	Direction string `json:"direction,omitempty"`
	Type      string `json:"type,omitempty"`
}

type Cloud struct {
	Provider         string `json:"provider,omitempty"`
	Region           string `json:"region,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	AccountID        string `json:"account.id,omitempty"`
	AccountName      string `json:"account.name,omitempty"`
	InstanceID       string `json:"instance.id,omitempty"`
	InstanceName     string `json:"instance.name,omitempty"`
	MachineType      string `json:"machine.type,omitempty"`
	ServiceName      string `json:"service.name,omitempty"`
	ProjectID        string `json:"project.id,omitempty"`
	ProjectName      string `json:"project.name,omitempty"`
}

type Group struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type Host struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Architecture string   `json:"architecture,omitempty"`
	Type         string   `json:"type,omitempty"`
	IP           string   `json:"ip,omitempty"`
	MacAddress   []string `json:"mac,omitempty"`
}

type User struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Email string   `json:"email,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type Orchestrator struct {
	ClusterID   string `json:"cluster.id,omitempty"`
	ClusterName string `json:"cluster.name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Container struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ImageName string `json:"image.name,omitempty"`
}

// AssetEnricher functional builder function
type AssetEnricher func(asset *AssetEvent)

func NewAssetEvent(c AssetClassification, id string, name string, enrichers ...AssetEnricher) AssetEvent {
	a := AssetEvent{
		Entity: Entity{
			Id:                  id,
			Name:                name,
			AssetClassification: c,
		},
		Event: Event{
			Kind: "asset",
		},
	}

	for _, enrich := range enrichers {
		enrich(&a)
	}

	return a
}

func WithRawAsset(raw any) AssetEnricher {
	return func(a *AssetEvent) {
		a.Entity.Raw = &raw
	}
}

func WithRelatedAssetIds(ids []string) AssetEnricher {
	return func(a *AssetEvent) {
		ids = lo.Filter(ids, func(id string, _ int) bool {
			return id != ""
		})

		if len(ids) == 0 {
			a.Entity.relatedEntityId = nil
			return
		}

		a.Entity.relatedEntityId = lo.Uniq(ids)
	}
}

func WithLabels(labels map[string]string) AssetEnricher {
	return func(a *AssetEvent) {
		if len(labels) == 0 {
			return
		}

		a.Labels = labels
	}
}

func WithLabelsFromAny(labels map[string]any) AssetEnricher {
	return func(a *AssetEvent) {
		if len(labels) == 0 {
			return
		}
		output := map[string]string{}
		if err := mapstructure.Decode(labels, &output); err != nil {
			return
		}
		a.Labels = output
	}
}

func WithNetwork(network Network) AssetEnricher {
	return func(a *AssetEvent) {
		a.Network = &network
	}
}

func WithCloud(cloud Cloud) AssetEnricher {
	return func(a *AssetEvent) {
		a.Cloud = &cloud
		a.Entity.Source = &cloud.Provider
	}
}

func WithGroup(group Group) AssetEnricher {
	return func(a *AssetEvent) {
		a.Group = &group
	}
}

func WithHost(host Host) AssetEnricher {
	return func(a *AssetEvent) {
		a.Host = &host
	}
}

func WithTags(tags []string) AssetEnricher {
	return func(a *AssetEvent) {
		if len(tags) == 0 {
			return
		}
		a.Tags = tags
	}
}

func WithUser(user User) AssetEnricher {
	return func(a *AssetEvent) {
		a.User = &user
	}
}

func EmptyEnricher() AssetEnricher {
	return func(_ *AssetEvent) {}
}

func WithOrganization(org Organization) AssetEnricher {
	return func(a *AssetEvent) {
		a.Organization = &org
	}
}

func WithFass(fass Fass) AssetEnricher {
	return func(a *AssetEvent) {
		a.Fass = &fass
	}
}

func WithURL(url URL) AssetEnricher {
	return func(a *AssetEvent) {
		a.URL = &url
	}
}

func WithOrchestrator(orchestrator Orchestrator) AssetEnricher {
	return func(a *AssetEvent) {
		a.Orchestrator = &orchestrator
	}
}

func WithContainer(container Container) AssetEnricher {
	return func(a *AssetEvent) {
		a.Container = &container
	}
}
