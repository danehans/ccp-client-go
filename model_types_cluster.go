/*
 * localhost:31453
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ccp-client-go

type TypesCluster struct {
	Infra *TypesClusterInfra `json:"Infra"`
	AciProfileUuid string `json:"aci_profile_uuid"`
	CcpPrivateSshKey string `json:"ccp_private_ssh_key"`
	CcpPublicSshKey string `json:"ccp_public_ssh_key"`
	ClusterDashboardUrl string `json:"cluster_dashboard_url"`
	ClusterEnvUrl string `json:"cluster_env_url"`
	Deployer *TypesKubeadm `json:"deployer"`
	DeployerType string `json:"deployer_type"`
	Description string `json:"description"`
	HelmCharts []TypesHelmChart `json:"helm_charts"`
	IngressVipAddrId string `json:"ingress_vip_addr_id"`
	IngressVipPoolId string `json:"ingress_vip_pool_id"`
	IngressVips []string `json:"ingress_vips"`
	IsAdopt bool `json:"is_adopt"`
	IsControlCluster bool `json:"is_control_cluster"`
	KeepalivedVrid int32 `json:"keepalived_vrid"`
	KubernetesVersion string `json:"kubernetes_version"`
	Labels []TypesLabel `json:"labels"`
	MasterMacAddresses []string `json:"master_mac_addresses"`
	MasterVip string `json:"master_vip"`
	MasterVipAddrId string `json:"master_vip_addr_id"`
	Masters int32 `json:"masters"`
	Name string `json:"name"`
	NetworkPlugin *TypesNetworkPluginProfile `json:"network_plugin"`
	Nodes []TypesNode `json:"nodes"`
	NtpPools []string `json:"ntp_pools"`
	NtpServers []string `json:"ntp_servers"`
	ProviderClientConfigUuid string `json:"provider_client_config_uuid"`
	RegistriesInsecure []string `json:"registries_insecure"`
	RegistriesRootCa []string `json:"registries_root_ca"`
	RegistriesSelfSigned []string `json:"registries_self_signed"`
	SshKey string `json:"ssh_key"`
	SshPassword string `json:"ssh_password"`
	SshUser string `json:"ssh_user"`
	State string `json:"state"`
	Template string `json:"template"`
	Type_ *TypesClusterType `json:"type"`
	Uuid string `json:"uuid"`
	Workers int32 `json:"workers"`
}