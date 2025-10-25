import * as React from 'react';
import { PureComponent } from 'react';
import { ClusterStatsItemState } from '../../types/clusterStatsItemState';
import { getDataSourceItemStyles } from './DataSourceItem.styles';
import { Button, HorizontalGroup, Spinner, ConfirmModal, Select, VerticalGroup } from '@grafana/ui';
import { toast } from 'react-toastify';
import type { SelectableValue } from '@grafana/runtime';
import { config, getBackendSrv } from '@grafana/runtime';

interface Props {
  dataSourceItem: any;
  theme: any;
  onDelete: (uid: string) => void;
}

export class DataSourceItem extends PureComponent<Props, ClusterStatsItemState> {
  user = config.bootData?.user;
  isServerAdmin = !!this.user?.isGrafanaAdmin;
  isOrgAdmin = this.user?.orgRole === 'Admin';

  state: ClusterStatsItemState = {
    cluster_name: '',
    cluster_uuid: '',
    status: '',
    versions: ['-'],
    numberOfIndices: 0,
    numberOfShards: 0,
    numberOfUnassignedShards: 0,
    docsCount: '0',
    usedStorage: '0',
    totalStorage: '0',
    totalNodes: 0,
    dataNodes: 0,
    dataHotNodes: 0,
    dataWarmNodes: 0,
    dataColdNodes: 0,
    monitorName: '',
    isOpenDialog: false,
    isLoading: false,
    isServerAdmin: this.isServerAdmin,
  };

  monitorOptions: Array<SelectableValue<string>> = [
    {
      label: 'Elastic Stack monitoring',
      value: 'stack-monitoring',
      description: 'Monitor your Elastic Stack',
    },
    {
      label: 'Elasticsearch Index overview',
      value: 'index-overview',
      description: 'View index statistics',
    },
    {
      label: 'Elasticsearch Shards overview',
      value: 'shards-overview',
      description: 'View shard distribution',
    },
    {
      label: 'Elasticsearch ingest pipelines overview',
      value: 'ingest-pipelines-overview',
      description: 'Monitor ingest pipelines',
    },
    {
      label: 'Logstash overview',
      value: 'logstash-overview',
      description: 'Monitor Logstash instances',
    },
    {
      label: 'Tasks analytics',
      value: 'tasks-analytics',
      description: 'Analyze running tasks',
    },
    {
      label: 'Elasticsearch ML Jobs Analytics',
      value: 'ml-jobs-analytics',
      description: 'Monitor ML jobs',
    },
  ];

  formatBytes(bytes: number, decimals = 2) {
    if (bytes === 0) {
      return '0';
    }
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
  }

  get uid() {
    return this.props.dataSourceItem.name.split('--').slice(2).join('--');
  }

  async componentDidMount() {
    try {
      await getBackendSrv()
        .get(
          `/api/datasources/proxy/uid/${this.props.dataSourceItem.uid}/_cluster/stats?filter_path=cluster_uuid,cluster_name,status,nodes.versions,indices.count,indices.shards.total,indices.docs.count,indices.store.size_in_bytes,nodes.fs.total_in_bytes,nodes.count.total,nodes.count.data,nodes.count.data_hot,nodes.count.data_warm,nodes.count.data_cold`
        )
        .then((dataSources: any) => {
          if (!dataSources) {
            console.debug('Error in the catch block:', dataSources);
            throw new Error('No data sources found');
          }
          this.setState({
            cluster_uuid: dataSources.cluster_uuid,
            cluster_name: dataSources.cluster_name,
            status: dataSources.status,
            versions: dataSources.nodes.versions,
            numberOfIndices: dataSources.indices.count,
            numberOfShards: dataSources.indices.shards.total,
            numberOfUnassignedShards: 0,
            docsCount: this.nFormatter(dataSources.indices.docs.count),
            usedStorage: this.formatBytes(dataSources.indices.store.size_in_bytes),
            totalStorage: this.formatBytes(dataSources.nodes.fs.total_in_bytes),
            totalNodes: dataSources.nodes.count.total,
            dataNodes: dataSources.nodes.count.data,
            dataHotNodes: dataSources.nodes.count.data_hot,
            dataWarmNodes: dataSources.nodes.count.data_warm,
            dataColdNodes: dataSources.nodes.count.data_cold,
          });
        })
        .catch((e) => {
          let regex = new RegExp(/Elasticsearch-direct-prod--(.*)--(.*)/g);
          const uid: string = this.props.dataSourceItem.name;
          const matches = regex.exec(uid);
          this.setState({
            cluster_name: matches ? matches[1] : '',
            cluster_uuid: matches ? matches[2] : '',
            status: 'ERROR',
            versions: ['-'],
            numberOfIndices: 0,
            numberOfShards: 0,
            numberOfUnassignedShards: 0,
            docsCount: '0',
            usedStorage: '0',
            totalStorage: '0',
            totalNodes: 0,
            dataNodes: 0,
            dataHotNodes: 0,
            dataWarmNodes: 0,
            dataColdNodes: 0,
          });
        });
    } catch (e) {
      console.debug('Error', e);
    }
    if (this.state.status !== 'ERROR') {
      await getBackendSrv()
        .get(
          `/api/datasources/proxy/uid/${this.props.dataSourceItem.uid}/_cluster/health?filter_path=unassigned_shards`
        )
        .then((health: any) => {
          this.setState({
            numberOfUnassignedShards: health.unassigned_shards,
          });
        });
    }
  }

  nFormatter(num: number): string {
    if (num >= 1000000000000) {
      return (num / 1000000000000).toFixed(1).replace(/\.0$/, '') + 'T';
    }
    if (num >= 1000000000) {
      return (num / 1000000000).toFixed(1).replace(/\.0$/, '') + 'B';
    }
    if (num >= 1000000) {
      return (num / 1000000).toFixed(1).replace(/\.0$/, '') + 'M';
    }
    if (num >= 1000) {
      return (num / 1000).toFixed(1).replace(/\.0$/, '') + 'K';
    }
    return num.toString();
  }

  onDelete = () => {
    this.setState({ isOpenDialog: true });
  };

  onTest = () => {
    this.setState({ isLoading: true });
    this.componentDidMount().then(() => {
      this.setState({ isLoading: false });
    });
  };

  onDeleteCluster = async () => {
    const backendSrv = getBackendSrv();
    this.setState({ isLoading: true });
    try {
      let dataSources: any = await backendSrv.get(`/api/datasources`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
      });
      if (dataSources) {
        this.setState({ isLoading: false });
      }

      for (const dataSource of dataSources) {
        if (dataSource.name.endsWith(this.state.cluster_uuid)) {
          try {
            await backendSrv.delete(`/api/datasources/name/${dataSource.name}`);
          } catch (deleteError) {
            console.debug('Delete data sources error: ', deleteError);
            this.setState({ isLoading: false });
          }
        }
      }

      try {
        await backendSrv.delete(
          `api/plugins/dbeast-dbeastmonitor-app/resources/delete_cluster/${this.state.cluster_uuid}`,
          {
            headers: {
              'Content-Type': 'application/json',
              Accept: 'application/json',
            },
          }
        );
      } catch (error: any) {
        toast.error(`${error.message}`, {
          position: toast.POSITION.BOTTOM_RIGHT,
          autoClose: false,
          closeButton: true,
          hideProgressBar: true,
          draggable: false,
        });
        this.setState({ isLoading: false });
      }

      this.setState({ isLoading: false });
      window.location.reload();
    } catch (getError) {
      console.debug('Cluster delete error: ', getError);
      this.setState({ isLoading: false });
    }
  };

  handleDelete(isOnYes: boolean): void {
    this.setState({ isOpenDialog: false });
    if (isOnYes) {
      this.onDeleteCluster().then(() => {});
    }
  }

  handleMonitorSelect = (value: SelectableValue<string>) => {
    if (value.value) {
      const dashboardUrls: Record<string, string> = {
        'stack-monitoring': `/d/elastic-stack-monitoring-dashboard/elastic-stack-monitoring-dashboard?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'index-overview': `/d/elasticsearch-index-overview/elasticsearch-index-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'shards-overview': `/d/elasticsearch-shards-overview-dashboard/elasticsearch-shards-overview-dashboard?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'ingest-pipelines-overview': `/d/elasticsearch-ingest-pipelines-overview/elasticsearch-ingest-pipelines-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'logstash-overview': `/d/logstash-overview/logstash-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'tasks-analytics': `/d/elasticsearch-tasks-analytics/elasticsearch-tasks-analytics?orgId=1&refresh=1m&var-cluster_ds=${this.uid}`,
        'ml-jobs-analytics': `/d/ml-jobs-analytics-dashboard/ml-jobs-analytics-dashboard?orgId=1&var-cluster_ds=${this.uid}`,
      };

      const url = dashboardUrls[value.value];
      if (url) {
        window.open(url, '_blank', 'noopener,noreferrer');
      }
    }
  };

  renderStatItem = (label: string, value: string | number) => {
    const theme = this.props.theme;
    const styles = getDataSourceItemStyles(theme);
    return (
      <div className={styles.statItem}>
        <span className={styles.statLabel}>{label}</span>
        <span className={styles.statValue}>{value ?? '0'}</span>
      </div>
    );
  };

  render() {
    const theme = this.props.theme;
    const styles = getDataSourceItemStyles(theme);

    return (
      <div className={styles.positionRelative}>
        {this.state.isLoading && (
          <div className={styles.spinnerOverlay}>
            <Spinner size={50} />
          </div>
        )}

        <div className={styles.formGroup}>
          <header>
            <div className={styles.headerCluster}>
              <h3>{this.state.cluster_name}</h3>
              <p>{this.state.cluster_uuid}</p>
            </div>
            <div className={styles.statusBadge}>
              <span className={this.state.status.toUpperCase()}>{this.state.status.toUpperCase()}</span>
            </div>
          </header>

          <div className={styles.divider} />

          <VerticalGroup spacing="md">
            <div className={styles.statsGrid}>
              {this.state.versions && this.renderStatItem('Version', this.state.versions)}
              {this.renderStatItem('Used storage', this.state.usedStorage)}
              {this.renderStatItem('Total storage', this.state.totalStorage)}
              {this.renderStatItem('Docs count', this.state.docsCount)}
              {this.renderStatItem('Total nodes', this.state.totalNodes)}
              {this.renderStatItem('Data nodes', this.state.dataNodes)}
              {this.renderStatItem('Hot nodes', this.state.dataHotNodes)}
              {this.renderStatItem('Warm nodes', this.state.dataWarmNodes)}
              {this.renderStatItem('Cold nodes', this.state.dataColdNodes)}
              {this.renderStatItem('Indices', this.state.numberOfIndices)}
              {this.renderStatItem('Total shards', this.state.numberOfShards)}
              {this.renderStatItem('Unassigned shards', this.state.numberOfUnassignedShards)}
            </div>
          </VerticalGroup>

          <div className={styles.divider} />

          <footer>
            <HorizontalGroup justify="space-between" spacing="md">
              <HorizontalGroup spacing="sm">
                {this.state.isServerAdmin && (
                  <Button variant="secondary" onClick={this.onDelete}>
                    Delete
                  </Button>
                )}
                <Button variant="secondary" onClick={() => this.onTest()}>
                  Test
                </Button>
              </HorizontalGroup>

              <div className={styles.selectWrapper}>
                <Select
                  options={this.monitorOptions}
                  placeholder="Monitor type"
                  onChange={this.handleMonitorSelect}
                  width={40}
                />
              </div>
            </HorizontalGroup>
          </footer>

          <ConfirmModal
            isOpen={this.state.isOpenDialog}
            title="Delete Cluster"
            body="Are you sure you want to delete this cluster?"
            confirmText="Yes"
            dismissText="No"
            onConfirm={() => this.handleDelete(true)}
            onDismiss={() => this.handleDelete(false)}
          />
        </div>
      </div>
    );
  }
}