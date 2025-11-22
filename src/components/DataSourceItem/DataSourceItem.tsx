import * as React from 'react';
import { PureComponent } from 'react';
import { config, getBackendSrv } from '@grafana/runtime';
import { ClusterStatsItemState } from '../../types/clusterStatsItemState';
import { getDataSourceItemStyles } from './DataSourceItem.styles';
import { Button, Spinner, Modal, HorizontalGroup, Select as GrafanaSelect } from '@grafana/ui';
import { toast } from 'react-toastify';
import { config } from '@grafana/runtime';
import { SelectableValue } from '@grafana/data';

interface Props {
  dataSourceItem: any;
  theme: any;
  onDelete: (uid: string) => void;
}

export class DataSourceItem extends PureComponent<Props, ClusterStatsItemState> {
  user = config.bootData?.user;
   isServerAdmin = !!this.user?.isGrafanaAdmin;       // Grafana server admin
   isOrgAdmin = this.user?.orgRole === 'Admin';       // Org admin (Admin/Editor/Viewer)


  // hide page, show “no permission”, etc.

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
  label = 'Monitor type';

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
            if (!dataSources ) {
              console.debug(
                  "Error in the catch block:",
                  dataSources)
              throw new Error('No data sources found');
            }
            const {} = dataSources;
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
    }
    catch (e) {
      console.debug("Error", e);
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
    this.setState({ isLoading: true })
    this.componentDidMount().then(() => {
      this.setState({ isLoading: false })
    });
  };
  onDeleteCluster = async () => {
    const backendSrv = getBackendSrv();
    this.setState({ isLoading: true })
    try {
      let dataSources: any = await backendSrv.get(`/api/datasources`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
      });
      if (dataSources){
        this.setState({ isLoading: false })
      }

      for (const dataSource of dataSources) {
        if (dataSource.name.endsWith(this.state.cluster_uuid)) {
          try {
            await backendSrv.delete(`/api/datasources/name/${dataSource.name}`);
          } catch (deleteError) {
            console.debug('Delete data sources error: ', deleteError);
            this.setState({ isLoading: false })
          }
        }
      }

      try {
        await backendSrv.delete(`api/plugins/dbeast-dbeastmonitor-app/resources/delete_cluster/${this.state.cluster_uuid}`, {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        });
      } catch (error: any) {
        toast.error(`${error.message}`, {
          position: toast.POSITION.BOTTOM_RIGHT,
          autoClose: false,
          closeButton: true,
          hideProgressBar: true,
          draggable: false,
        });
        this.setState({ isLoading: false })
      }

      this.setState({ isLoading: false })
      window.location.reload();
    } catch (getError) {
      console.debug('Cluster delete error: ', getError);
      this.setState({ isLoading: false })
    }
  };

  handleDelete(isOnYes: boolean): void {
    this.setState({ isOpenDialog: false });
    if (isOnYes) {
      this.onDeleteCluster().then(() => {});
    }
  }

  render() {
    const theme = this.props.theme;
    const styles = getDataSourceItemStyles(theme);

    return (
      <div className={styles.positionRelative}>
        {this.state.isLoading && (
          <div className={styles.spinnerOverlay}>
            <Spinner />
          </div>
        )}

        <div className={styles.formGroup}>
          <header>
            <div className={`col ${styles.headerCluster}`}>
              <h3>{this.state.cluster_name}</h3>
              <p>{this.state.cluster_uuid}</p>
            </div>
            <div className={`actions col ${styles.actions}`}>
              <span className={this.state.status.toUpperCase()}>{this.state.status.toUpperCase()}</span>
            </div>
          </header>

          <div className={styles.divider}>
            <div style={{ borderTop: '1px solid rgba(204, 204, 220, 0.12)', margin: '8px 0' }} />
          </div>

          <div className="grid-container">
            <div className="col">
              <div className={styles.listItem}>
                <div className='grid'>
                  <div>
                    {this.state.versions ? (
                      <div>
                        <span className="label">Version</span>
                        <div>{this.state.versions ?? '0'}</div>
                      </div>
                    ) : null}

                    <div>
                      <span className="label">Used storage</span>
                      <div>{this.state.usedStorage ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Total storage</span>
                      <div>{this.state.totalStorage ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Docs count</span>
                      <div>{this.state.docsCount ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Total nodes</span>
                      <div>{this.state.totalNodes ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Data nodes</span>
                      <div>{this.state.dataNodes ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label"> Hot nodes</span>
                      <div>{this.state.dataHotNodes ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label"> Warm nodes</span>
                      <div>{this.state.dataWarmNodes ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label"> Cold nodes</span>
                      <div>{this.state.dataColdNodes ?? '0'}</div>
                    </div>

                    <div>
                      <span className="label">Indices</span>
                      <div>{this.state.numberOfIndices ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Total shards</span>
                      <div>{this.state.numberOfShards ?? '0'}</div>
                    </div>
                    <div>
                      <span className="label">Unassigned shards</span>
                      <div>{this.state.numberOfUnassignedShards ?? '0'}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="col"></div>
          </div>

          <div className={styles.divider}>
            <div style={{ borderTop: '1px solid rgba(204, 204, 220, 0.12)', margin: '8px 0' }} />
          </div>

          <footer>
            <div className={styles.stack}>
              <HorizontalGroup spacing="sm">
                <div className={styles.buttons}>

                  {this.state.isServerAdmin && (
                  <Button variant="secondary" className="btn" onClick={this.onDelete}>
                    Delete
                  </Button>
                  )}

                  <Button variant="secondary" className="btn" onClick={() => this.onTest()}>
                    Test
                  </Button>
                </div>
                <div className={styles.select} >
                  <GrafanaSelect
                    options={[
                      { label: 'Elastic Stack monitoring', value: 'stack-monitoring', href: `/d/elastic-stack-monitoring-dashboard/elastic-stack-monitoring-dashboard?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Elasticsearch Index overview', value: 'index-overview', href: `/d/elasticsearch-index-overview/elasticsearch-index-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Elasticsearch Shards overview', value: 'shards-overview', href: `/d/elasticsearch-shards-overview-dashboard/elasticsearch-shards-overview-dashboard?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Elasticsearch ingest pipelines overview', value: 'ingest-pipelines-overview', href: `/d/elasticsearch-ingest-pipelines-overview/elasticsearch-ingest-pipelines-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Logstash overview', value: 'logstash-overview', href: `/d/logstash-overview/logstash-overview?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Tasks analytics', value: 'tasks-analytics', href: `/d/elasticsearch-tasks-analytics/elasticsearch-tasks-analytics?orgId=1&refresh=1m&var-cluster_ds=${this.uid}` },
                      { label: 'Elasticsearch ML Jobs Analytics', value: 'ml-jobs-analytics', href: `/d/ml-jobs-analytics-dashboard/ml-jobs-analytics-dashboard?orgId=1&var-cluster_ds=${this.uid}` },
                    ]}
                    value={this.state.monitorName ? { label: this.state.monitorName.split('-').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' '), value: this.state.monitorName } : null}
                    placeholder="Monitor type"
                    onChange={(option: SelectableValue) => {
                      if (option?.href) {
                        window.open(option.href, '_blank');
                      }
                    }}
                  />
                </div>
              </HorizontalGroup>
            </div>
          </footer>

          <div className={styles.dialog}>
            <Modal
              isOpen={this.state?.isOpenDialog!}
              title="Are you sure you want to delete this cluster?"
              onDismiss={() => this.handleDelete(false)}
            >
              <Modal.ButtonRow>
                <Button variant="secondary" onClick={() => this.handleDelete(false)} className="btn-error">
                  No
                </Button>
                <Button variant="destructive" onClick={() => this.handleDelete(true)}>
                  Yes
                </Button>
              </Modal.ButtonRow>
            </Modal>
          </div>
        </div>
      </div>
    );
  }
}
