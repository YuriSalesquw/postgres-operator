package config

import (
	"encoding/json"
	"strings"
	"time"

	"fmt"

	"github.com/zalando-incubator/postgres-operator/pkg/spec"
)

// CRD describes CustomResourceDefinition specific configuration parameters
type CRD struct {
	ReadyWaitInterval time.Duration `name:"ready_wait_interval" default:"4s"`
	ReadyWaitTimeout  time.Duration `name:"ready_wait_timeout" default:"30s"`
	ResyncPeriod      time.Duration `name:"resync_period" default:"5m"`
}

// Resources describes kubernetes resource specific configuration parameters
type Resources struct {
	ResourceCheckInterval   time.Duration     `name:"resource_check_interval" default:"3s"`
	ResourceCheckTimeout    time.Duration     `name:"resource_check_timeout" default:"10m"`
	PodLabelWaitTimeout     time.Duration     `name:"pod_label_wait_timeout" default:"10m"`
	PodDeletionWaitTimeout  time.Duration     `name:"pod_deletion_wait_timeout" default:"10m"`
	PodTerminateGracePeriod time.Duration     `name:"pod_terminate_grace_period" default:"5m"`
	ClusterLabels           map[string]string `name:"cluster_labels" default:"application:spilo"`
	ClusterNameLabel        string            `name:"cluster_name_label" default:"cluster-name"`
	PodRoleLabel            string            `name:"pod_role_label" default:"spilo-role"`
	PodToleration           map[string]string `name:"toleration" default:""`
	DefaultCPURequest       string            `name:"default_cpu_request" default:"100m"`
	DefaultMemoryRequest    string            `name:"default_memory_request" default:"100Mi"`
	DefaultCPULimit         string            `name:"default_cpu_limit" default:"3"`
	DefaultMemoryLimit      string            `name:"default_memory_limit" default:"1Gi"`
	PodEnvironmentConfigMap string            `name:"pod_environment_configmap" default:""`
	NodeReadinessLabel      map[string]string `name:"node_readiness_label" default:""`
	MaxInstances            int32             `name:"max_instances" default:"-1"`
	MinInstances            int32             `name:"min_instances" default:"-1"`
}

// Auth describes authentication specific configuration parameters
type Auth struct {
	SecretNameTemplate            stringTemplate      `name:"secret_name_template" default:"{username}.{cluster}.credentials.{tprkind}.{tprgroup}"`
	PamRoleName                   string              `name:"pam_role_name" default:"zalandos"`
	PamConfiguration              string              `name:"pam_configuration" default:"https://info.example.com/oauth2/tokeninfo?access_token= uid realm=/employees"`
	TeamsAPIUrl                   string              `name:"teams_api_url" default:"https://teams.example.com/api/"`
	OAuthTokenSecretName          spec.NamespacedName `name:"oauth_token_secret_name" default:"postgresql-operator"`
	InfrastructureRolesSecretName spec.NamespacedName `name:"infrastructure_roles_secret_name"`
	SuperUsername                 string              `name:"super_username" default:"postgres"`
	ReplicationUsername           string              `name:"replication_username" default:"standby"`
}

// Scalyr holds the configuration for the Scalyr Agent sidecar for log shipping:
type Scalyr struct {
	ScalyrAPIKey        string `name:"scalyr_api_key" default:""`
	ScalyrImage         string `name:"scalyr_image" default:""`
	ScalyrServerURL     string `name:"scalyr_server_url" default:"https://upload.eu.scalyr.com"`
	ScalyrCPURequest    string `name:"scalyr_cpu_request" default:"100m"`
	ScalyrMemoryRequest string `name:"scalyr_memory_request" default:"50Mi"`
	ScalyrCPULimit      string `name:"scalyr_cpu_limit" default:"1"`
	ScalyrMemoryLimit   string `name:"scalyr_memory_limit" default:"1Gi"`
}

// Config describes operator config
type Config struct {
	CRD
	Resources
	Auth
	Scalyr
	WatchedNamespace          string            `name:"watched_namespace"` // special values: "*" means 'watch all namespaces', the empty string "" means 'watch a namespace where operator is deployed to'
	EtcdHost                  string            `name:"etcd_host" default:"etcd-client.default.svc.cluster.local:2379"`
	DockerImage               string            `name:"docker_image" default:"registry.opensource.zalan.do/acid/spiloprivate-9.6:1.2-p4"`
	ServiceAccountName        string            `name:"service_account_name" default:"operator"`
	DbHostedZone              string            `name:"db_hosted_zone" default:"db.example.com"`
	EtcdScope                 string            `name:"etcd_scope" default:"service"`
	WALES3Bucket              string            `name:"wal_s3_bucket"`
	KubeIAMRole               string            `name:"kube_iam_role"`
	DebugLogging              bool              `name:"debug_logging" default:"true"`
	EnableDBAccess            bool              `name:"enable_database_access" default:"true"`
	EnableTeamsAPI            bool              `name:"enable_teams_api" default:"true"`
	EnableTeamSuperuser       bool              `name:"enable_team_superuser" default:"false"`
	TeamAdminRole             string            `name:"team_admin_role" default:"admin"`
	EnableMasterLoadBalancer  bool              `name:"enable_master_load_balancer" default:"false"`
	EnableReplicaLoadBalancer bool              `name:"enable_replica_load_balancer" default:"false"`
	MasterDNSNameFormat       stringTemplate    `name:"master_dns_name_format" default:"{cluster}.{team}.{hostedzone}"`
	ReplicaDNSNameFormat      stringTemplate    `name:"replica_dns_name_format" default:"{cluster}-repl.{team}.{hostedzone}"`
	PDBNameFormat             stringTemplate    `name:"pdb_name_format" default:"postgres-{cluster}-pdb"`
	Workers                   uint32            `name:"workers" default:"4"`
	APIPort                   int               `name:"api_port" default:"8080"`
	RingLogLines              int               `name:"ring_log_lines" default:"100"`
	ClusterHistoryEntries     int               `name:"cluster_history_entries" default:"1000"`
	TeamAPIRoleConfiguration  map[string]string `name:"team_api_role_configuration" default:"log_statement:all"`
	PodTerminateGracePeriod   time.Duration     `name:"pod_terminate_grace_period" default:"5m"`
	ProtectedRoles            []string          `name:"protected_role_names" default:"admin"`
}

// MustMarshal marshals the config or panics
func (c Config) MustMarshal() string {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// NewFromMap creates Config from the map
func NewFromMap(m map[string]string) *Config {
	cfg := Config{}
	fields, _ := structFields(&cfg)

	for _, structField := range fields {
		key := strings.ToLower(structField.Name)
		value, ok := m[key]
		if !ok && structField.Default != "" {
			value = structField.Default
		}

		if value == "" {
			continue
		}
		err := processField(value, structField.Field)
		if err != nil {
			panic(err)
		}
	}
	if err := validate(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}

// Copy creates a copy of the config
func Copy(c *Config) Config {
	cfg := *c

	cfg.ClusterLabels = make(map[string]string, len(c.ClusterLabels))
	for k, v := range c.ClusterLabels {
		cfg.ClusterLabels[k] = v
	}

	return cfg
}

func validate(cfg *Config) (err error) {
	if cfg.MinInstances > 0 && cfg.MaxInstances > 0 && cfg.MinInstances > cfg.MaxInstances {
		err = fmt.Errorf("minimum number of instances %d is set higher than the maximum number %d",
			cfg.MinInstances, cfg.MaxInstances)
	}
	if cfg.Workers == 0 {
		err = fmt.Errorf("number of workers should be higher than 0")
	}
	return
}
