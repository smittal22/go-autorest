package azure

// Copyright 2017 Microsoft Corporation
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	// EnvironmentFilepathName captures the name of the environment variable containing the path to the file
	// to be used while populating the Azure Environment.
	EnvironmentFilepathName = "AZURE_ENVIRONMENT_FILEPATH"

	// NotAvailable is used for endpoints and resource IDs that are not available for a given cloud.
	NotAvailable = "N/A"
)

var environments = map[string]Environment{
	"AZURECHINACLOUD":        ChinaCloud,
	"AZUREGERMANCLOUD":       GermanCloud,
	"AZURECLOUD":             PublicCloud,
	"AZUREPUBLICCLOUD":       PublicCloud,
	"AZUREUSGOVERNMENT":      USGovernmentCloud,
	"AZUREUSGOVERNMENTCLOUD": USGovernmentCloud, //TODO: deprecate
	"AZUREFRENCHCLOUD":       AzureFrenchCloud,
}

// ResourceIdentifier contains a set of Azure resource IDs.
type ResourceIdentifier struct {
	Graph               string `json:"graph"`
	KeyVault            string `json:"keyVault"`
	Datalake            string `json:"datalake"`
	Batch               string `json:"batch"`
	OperationalInsights string `json:"operationalInsights"`
	OSSRDBMS            string `json:"ossRDBMS"`
	Storage             string `json:"storage"`
	Synapse             string `json:"synapse"`
	ServiceBus          string `json:"serviceBus"`
	SQLDatabase         string `json:"sqlDatabase"`
	CosmosDB            string `json:"cosmosDB"`
	ManagedHSM          string `json:"managedHSM"`
	MicrosoftGraph      string `json:"microsoftGraph"`
}

// Environment represents a set of endpoints for each of Azure's Clouds.
type Environment struct {
	Name                         string             `json:"name"`
	ManagementPortalURL          string             `json:"managementPortalURL"`
	PublishSettingsURL           string             `json:"publishSettingsURL"`
	ServiceManagementEndpoint    string             `json:"serviceManagementEndpoint"`
	ResourceManagerEndpoint      string             `json:"resourceManagerEndpoint"`
	ActiveDirectoryEndpoint      string             `json:"activeDirectoryEndpoint"`
	GalleryEndpoint              string             `json:"galleryEndpoint"`
	KeyVaultEndpoint             string             `json:"keyVaultEndpoint"`
	ManagedHSMEndpoint           string             `json:"managedHSMEndpoint"`
	GraphEndpoint                string             `json:"graphEndpoint"`
	ServiceBusEndpoint           string             `json:"serviceBusEndpoint"`
	BatchManagementEndpoint      string             `json:"batchManagementEndpoint"`
	MicrosoftGraphEndpoint       string             `json:"microsoftGraphEndpoint"`
	StorageEndpointSuffix        string             `json:"storageEndpointSuffix"`
	CosmosDBDNSSuffix            string             `json:"cosmosDBDNSSuffix"`
	MariaDBDNSSuffix             string             `json:"mariaDBDNSSuffix"`
	MySQLDatabaseDNSSuffix       string             `json:"mySqlDatabaseDNSSuffix"`
	PostgresqlDatabaseDNSSuffix  string             `json:"postgresqlDatabaseDNSSuffix"`
	SQLDatabaseDNSSuffix         string             `json:"sqlDatabaseDNSSuffix"`
	TrafficManagerDNSSuffix      string             `json:"trafficManagerDNSSuffix"`
	KeyVaultDNSSuffix            string             `json:"keyVaultDNSSuffix"`
	ManagedHSMDNSSuffix          string             `json:"managedHSMDNSSuffix"`
	ServiceBusEndpointSuffix     string             `json:"serviceBusEndpointSuffix"`
	ServiceManagementVMDNSSuffix string             `json:"serviceManagementVMDNSSuffix"`
	ResourceManagerVMDNSSuffix   string             `json:"resourceManagerVMDNSSuffix"`
	ContainerRegistryDNSSuffix   string             `json:"containerRegistryDNSSuffix"`
	TokenAudience                string             `json:"tokenAudience"`
	APIManagementHostNameSuffix  string             `json:"apiManagementHostNameSuffix"`
	SynapseEndpointSuffix        string             `json:"synapseEndpointSuffix"`
	DatalakeSuffix               string             `json:"datalakeSuffix"`
	ResourceIdentifiers          ResourceIdentifier `json:"resourceIdentifiers"`
}

var (
	// PublicCloud is the default public Azure cloud environment
	PublicCloud = Environment{
		Name:                         "AzurePublicCloud",
		ManagementPortalURL:          "https://manage.windowsazure.com/",
		PublishSettingsURL:           "https://manage.windowsazure.com/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.windows.net/",
		ResourceManagerEndpoint:      "https://management.azure.com/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.com/",
		GalleryEndpoint:              "https://gallery.azure.com/",
		KeyVaultEndpoint:             "https://vault.azure.net/",
		ManagedHSMEndpoint:           "https://managedhsm.azure.net/",
		GraphEndpoint:                "https://graph.windows.net/",
		ServiceBusEndpoint:           "https://servicebus.windows.net/",
		BatchManagementEndpoint:      "https://batch.core.windows.net/",
		MicrosoftGraphEndpoint:       "https://graph.microsoft.com/",
		StorageEndpointSuffix:        "core.windows.net",
		CosmosDBDNSSuffix:            "documents.azure.com",
		MariaDBDNSSuffix:             "mariadb.database.azure.com",
		MySQLDatabaseDNSSuffix:       "mysql.database.azure.com",
		PostgresqlDatabaseDNSSuffix:  "postgres.database.azure.com",
		SQLDatabaseDNSSuffix:         "database.windows.net",
		TrafficManagerDNSSuffix:      "trafficmanager.net",
		KeyVaultDNSSuffix:            "vault.azure.net",
		ManagedHSMDNSSuffix:          "managedhsm.azure.net",
		ServiceBusEndpointSuffix:     "servicebus.windows.net",
		ServiceManagementVMDNSSuffix: "cloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.com",
		ContainerRegistryDNSSuffix:   "azurecr.io",
		TokenAudience:                "https://management.azure.com/",
		APIManagementHostNameSuffix:  "azure-api.net",
		SynapseEndpointSuffix:        "dev.azuresynapse.net",
		DatalakeSuffix:               "azuredatalakestore.net",
		ResourceIdentifiers: ResourceIdentifier{
			Graph:               "https://graph.windows.net/",
			KeyVault:            "https://vault.azure.net",
			Datalake:            "https://datalake.azure.net/",
			Batch:               "https://batch.core.windows.net/",
			OperationalInsights: "https://api.loganalytics.io",
			OSSRDBMS:            "https://ossrdbms-aad.database.windows.net",
			Storage:             "https://storage.azure.com/",
			Synapse:             "https://dev.azuresynapse.net",
			ServiceBus:          "https://servicebus.azure.net/",
			SQLDatabase:         "https://database.windows.net/",
			CosmosDB:            "https://cosmos.azure.com",
			ManagedHSM:          "https://managedhsm.azure.net",
			MicrosoftGraph:      "https://graph.microsoft.com/",
		},
	}

	// USGovernmentCloud is the cloud environment for the US Government
	USGovernmentCloud = Environment{
		Name:                         "AzureUSGovernmentCloud",
		ManagementPortalURL:          "https://manage.windowsazure.us/",
		PublishSettingsURL:           "https://manage.windowsazure.us/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:      "https://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.us/",
		GalleryEndpoint:              "https://gallery.usgovcloudapi.net/",
		KeyVaultEndpoint:             "https://vault.usgovcloudapi.net/",
		ManagedHSMEndpoint:           NotAvailable,
		GraphEndpoint:                "https://graph.windows.net/",
		ServiceBusEndpoint:           "https://servicebus.usgovcloudapi.net/",
		BatchManagementEndpoint:      "https://batch.core.usgovcloudapi.net/",
		MicrosoftGraphEndpoint:       "https://graph.microsoft.us/",
		StorageEndpointSuffix:        "core.usgovcloudapi.net",
		CosmosDBDNSSuffix:            "documents.azure.us",
		MariaDBDNSSuffix:             "mariadb.database.usgovcloudapi.net",
		MySQLDatabaseDNSSuffix:       "mysql.database.usgovcloudapi.net",
		PostgresqlDatabaseDNSSuffix:  "postgres.database.usgovcloudapi.net",
		SQLDatabaseDNSSuffix:         "database.usgovcloudapi.net",
		TrafficManagerDNSSuffix:      "usgovtrafficmanager.net",
		KeyVaultDNSSuffix:            "vault.usgovcloudapi.net",
		ManagedHSMDNSSuffix:          NotAvailable,
		ServiceBusEndpointSuffix:     "servicebus.usgovcloudapi.net",
		ServiceManagementVMDNSSuffix: "usgovcloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.usgovcloudapi.net",
		ContainerRegistryDNSSuffix:   "azurecr.us",
		TokenAudience:                "https://management.usgovcloudapi.net/",
		APIManagementHostNameSuffix:  "azure-api.us",
		SynapseEndpointSuffix:        "dev.azuresynapse.usgovcloudapi.net",
		DatalakeSuffix:               NotAvailable,
		ResourceIdentifiers: ResourceIdentifier{
			Graph:               "https://graph.windows.net/",
			KeyVault:            "https://vault.usgovcloudapi.net",
			Datalake:            NotAvailable,
			Batch:               "https://batch.core.usgovcloudapi.net/",
			OperationalInsights: "https://api.loganalytics.us",
			OSSRDBMS:            "https://ossrdbms-aad.database.usgovcloudapi.net",
			Storage:             "https://storage.azure.com/",
			Synapse:             "https://dev.azuresynapse.usgovcloudapi.net",
			ServiceBus:          "https://servicebus.azure.net/",
			SQLDatabase:         "https://database.usgovcloudapi.net/",
			CosmosDB:            "https://cosmos.azure.com",
			ManagedHSM:          NotAvailable,
			MicrosoftGraph:      "https://graph.microsoft.us/",
		},
	}

	// ChinaCloud is the cloud environment operated in China
	ChinaCloud = Environment{
		Name:                         "AzureChinaCloud",
		ManagementPortalURL:          "https://manage.chinacloudapi.com/",
		PublishSettingsURL:           "https://manage.chinacloudapi.com/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:      "https://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:      "https://login.chinacloudapi.cn/",
		GalleryEndpoint:              "https://gallery.chinacloudapi.cn/",
		KeyVaultEndpoint:             "https://vault.azure.cn/",
		ManagedHSMEndpoint:           NotAvailable,
		GraphEndpoint:                "https://graph.chinacloudapi.cn/",
		ServiceBusEndpoint:           "https://servicebus.chinacloudapi.cn/",
		BatchManagementEndpoint:      "https://batch.chinacloudapi.cn/",
		MicrosoftGraphEndpoint:       "https://microsoftgraph.chinacloudapi.cn/",
		StorageEndpointSuffix:        "core.chinacloudapi.cn",
		CosmosDBDNSSuffix:            "documents.azure.cn",
		MariaDBDNSSuffix:             "mariadb.database.chinacloudapi.cn",
		MySQLDatabaseDNSSuffix:       "mysql.database.chinacloudapi.cn",
		PostgresqlDatabaseDNSSuffix:  "postgres.database.chinacloudapi.cn",
		SQLDatabaseDNSSuffix:         "database.chinacloudapi.cn",
		TrafficManagerDNSSuffix:      "trafficmanager.cn",
		KeyVaultDNSSuffix:            "vault.azure.cn",
		ManagedHSMDNSSuffix:          NotAvailable,
		ServiceBusEndpointSuffix:     "servicebus.chinacloudapi.cn",
		ServiceManagementVMDNSSuffix: "chinacloudapp.cn",
		ResourceManagerVMDNSSuffix:   "cloudapp.chinacloudapi.cn",
		ContainerRegistryDNSSuffix:   "azurecr.cn",
		TokenAudience:                "https://management.chinacloudapi.cn/",
		APIManagementHostNameSuffix:  "azure-api.cn",
		SynapseEndpointSuffix:        "dev.azuresynapse.azure.cn",
		DatalakeSuffix:               NotAvailable,
		ResourceIdentifiers: ResourceIdentifier{
			Graph:               "https://graph.chinacloudapi.cn/",
			KeyVault:            "https://vault.azure.cn",
			Datalake:            NotAvailable,
			Batch:               "https://batch.chinacloudapi.cn/",
			OperationalInsights: NotAvailable,
			OSSRDBMS:            "https://ossrdbms-aad.database.chinacloudapi.cn",
			Storage:             "https://storage.azure.com/",
			Synapse:             "https://dev.azuresynapse.net",
			ServiceBus:          "https://servicebus.azure.net/",
			SQLDatabase:         "https://database.chinacloudapi.cn/",
			CosmosDB:            "https://cosmos.azure.com",
			ManagedHSM:          NotAvailable,
			MicrosoftGraph:      "https://microsoftgraph.chinacloudapi.cn",
		},
	}

	// GermanCloud is the cloud environment operated in Germany
	GermanCloud = Environment{
		Name:                         "AzureGermanCloud",
		ManagementPortalURL:          "http://portal.microsoftazure.de/",
		PublishSettingsURL:           "https://manage.microsoftazure.de/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.cloudapi.de/",
		ResourceManagerEndpoint:      "https://management.microsoftazure.de/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.de/",
		GalleryEndpoint:              "https://gallery.cloudapi.de/",
		KeyVaultEndpoint:             "https://vault.microsoftazure.de/",
		ManagedHSMEndpoint:           NotAvailable,
		GraphEndpoint:                "https://graph.cloudapi.de/",
		ServiceBusEndpoint:           "https://servicebus.cloudapi.de/",
		BatchManagementEndpoint:      "https://batch.cloudapi.de/",
		MicrosoftGraphEndpoint:       NotAvailable,
		StorageEndpointSuffix:        "core.cloudapi.de",
		CosmosDBDNSSuffix:            "documents.microsoftazure.de",
		MariaDBDNSSuffix:             "mariadb.database.cloudapi.de",
		MySQLDatabaseDNSSuffix:       "mysql.database.cloudapi.de",
		PostgresqlDatabaseDNSSuffix:  "postgres.database.cloudapi.de",
		SQLDatabaseDNSSuffix:         "database.cloudapi.de",
		TrafficManagerDNSSuffix:      "azuretrafficmanager.de",
		KeyVaultDNSSuffix:            "vault.microsoftazure.de",
		ManagedHSMDNSSuffix:          NotAvailable,
		ServiceBusEndpointSuffix:     "servicebus.cloudapi.de",
		ServiceManagementVMDNSSuffix: "azurecloudapp.de",
		ResourceManagerVMDNSSuffix:   "cloudapp.microsoftazure.de",
		ContainerRegistryDNSSuffix:   NotAvailable,
		TokenAudience:                "https://management.microsoftazure.de/",
		APIManagementHostNameSuffix:  NotAvailable,
		SynapseEndpointSuffix:        NotAvailable,
		DatalakeSuffix:               NotAvailable,
		ResourceIdentifiers: ResourceIdentifier{
			Graph:               "https://graph.cloudapi.de/",
			KeyVault:            "https://vault.microsoftazure.de",
			Datalake:            NotAvailable,
			Batch:               "https://batch.cloudapi.de/",
			OperationalInsights: NotAvailable,
			OSSRDBMS:            "https://ossrdbms-aad.database.cloudapi.de",
			Storage:             "https://storage.azure.com/",
			Synapse:             NotAvailable,
			ServiceBus:          "https://servicebus.azure.net/",
			SQLDatabase:         "https://database.cloudapi.de/",
			CosmosDB:            "https://cosmos.azure.com",
			ManagedHSM:          NotAvailable,
			MicrosoftGraph:      NotAvailable,
		},
	}

	// AzureFrenchCloud is the cloud environment operated in France
	AzureFrenchCloud = Environment{
		Name:                         "AzureFrenchCloud",
		ManagementPortalURL:          "",
		PublishSettingsURL:           "",
		ServiceManagementEndpoint:    "",
		ResourceManagerEndpoint:      "https://management.sovcloud-api.fr/",
		ActiveDirectoryEndpoint:      "",
		GalleryEndpoint:              "",
		KeyVaultEndpoint:             "https://vault.sovcloud-api.fr/",
		ManagedHSMEndpoint:           NotAvailable,
		GraphEndpoint:                "https://graph.svc.sovcloud.fr/",
		ServiceBusEndpoint:           "",
		BatchManagementEndpoint:      "",
		MicrosoftGraphEndpoint:       NotAvailable,
		StorageEndpointSuffix:        "core.sovcloud-api.fr",
		CosmosDBDNSSuffix:            "",
		MariaDBDNSSuffix:             "",
		MySQLDatabaseDNSSuffix:       "",
		PostgresqlDatabaseDNSSuffix:  "",
		SQLDatabaseDNSSuffix:         "",
		TrafficManagerDNSSuffix:      "",
		KeyVaultDNSSuffix:            "vault.sovcloud-api.fr",
		ManagedHSMDNSSuffix:          NotAvailable,
		ServiceBusEndpointSuffix:     "",
		ServiceManagementVMDNSSuffix: "",
		ResourceManagerVMDNSSuffix:   "",
		ContainerRegistryDNSSuffix:   NotAvailable,
		TokenAudience:                "",
		APIManagementHostNameSuffix:  NotAvailable,
		SynapseEndpointSuffix:        NotAvailable,
		DatalakeSuffix:               NotAvailable,
		ResourceIdentifiers: ResourceIdentifier{
			Graph:               "",
			KeyVault:            "https://vault.sovcloud-api.fr",
			Datalake:            NotAvailable,
			Batch:               "",
			OperationalInsights: NotAvailable,
			OSSRDBMS:            "",
			Storage:             "",
			Synapse:             NotAvailable,
			ServiceBus:          "",
			SQLDatabase:         "",
			CosmosDB:            "",
			ManagedHSM:          NotAvailable,
			MicrosoftGraph:      NotAvailable,
		},
	}
)

// EnvironmentFromName returns an Environment based on the common name specified.
func EnvironmentFromName(name string) (Environment, error) {
	// IMPORTANT
	// As per @radhikagupta5:
	// This is technical debt, fundamentally here because Kubernetes is not currently accepting
	// contributions to the providers. Once that is an option, the provider should be updated to
	// directly call `EnvironmentFromFile`. Until then, we rely on dispatching Azure Stack environment creation
	// from this method based on the name that is provided to us.
	if strings.EqualFold(name, "AZURESTACKCLOUD") {
		return EnvironmentFromFile(os.Getenv(EnvironmentFilepathName))
	}

	name = strings.ToUpper(name)
	env, ok := environments[name]
	if !ok {
		return env, fmt.Errorf("autorest/azure: There is no cloud environment matching the name %q", name)
	}

	return env, nil
}

// EnvironmentFromFile loads an Environment from a configuration file available on disk.
// This function is particularly useful in the Hybrid Cloud model, where one must define their own
// endpoints.
func EnvironmentFromFile(location string) (unmarshaled Environment, err error) {
	fileContents, err := os.ReadFile(location)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileContents, &unmarshaled)

	return
}

// SetEnvironment updates the environment map with the specified values.
func SetEnvironment(name string, env Environment) {
	environments[strings.ToUpper(name)] = env
}
