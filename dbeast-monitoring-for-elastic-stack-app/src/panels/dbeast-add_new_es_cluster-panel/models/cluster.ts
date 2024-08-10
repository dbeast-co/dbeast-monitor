import {ConnectionSettings} from './connection-settings';

export interface CheckboxConfig {
    label: string;
    id: string;
    is_checked: boolean;
}

export interface Host {
    server_address: string;
    logstash_api_host: string;
    logstash_logs_folder?: string;
}

export interface Cluster {
    cluster_connection_settings: ConnectionSettings;
    logstash_configurations: {
        es_monitoring_configuration_files: CheckboxConfig[];
        logstash_monitoring_configuration_files: {
            configurations: CheckboxConfig[];
            hosts: Host[];
        };
    };
    monitoring_cluster_injection: MonitoringClusterInjection
}

export interface MonitoringClusterInjection {
    templates_injection: boolean,
    ilm_policies_injection: boolean,
    create_first_indices: boolean
}
