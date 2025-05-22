import React, { PureComponent } from 'react';
import { getBackendSrv } from '@grafana/runtime';
import { ClusterStatsItemState, MonitorState } from '../../types/clusterStatsItemState';
import './data-source-item.scss';
import {
  Dialog,
  DialogActions,
  DialogTitle,
  Divider,
  FormControl,
  List,
  ListItem,
  ListItemText,
  MenuItem,
  Select,
  SelectChangeEvent,
  Stack,
} from '@mui/material';
import { Button, ButtonVariant, Spinner, stylesFactory, useTheme } from '@grafana/ui';
import classNames from 'classnames';
import { toast } from 'react-toastify';

interface Props {
  dataSourceItem: any;
  theme: any;
  onDelete: (uid: string) => void;
}

const MyComponent = (_: any) => {
  // Get the current theme and its properties using the useTheme hook
  const theme = useTheme();
  const styles = getStyles(theme);

  return (
    <div style={styles.container}>
      <div>Hello, This is My Component!</div>
      {/* Add more components and content as needed */}
    </div>
  );
};
export default MyComponent;

const getStyles = stylesFactory((theme) => ({
  container: {
    backgroundColor: theme.colors.bg1,
    color: theme.colors.text,
    fontSize: 18,
    padding: 10,
  },
}));

export class DataSourceItem extends PureComponent<Props, ClusterStatsItemState> {
  monitorState: MonitorState = {
    monitorName: '',
  };
  loading = false;

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
  };
  outlined: ButtonVariant | undefined = 'destructive';
  listItem: HTMLElement | undefined = undefined;
  label = 'Monitor type';
  clusterMonitoring: HTMLLIElement | undefined = undefined;

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

  handleChange = (event: SelectChangeEvent) => {
    const inputString = this.props.dataSourceItem.name;
    const uid = inputString.split('--').slice(2).join('--');

    this.setState({
      monitorName: event.target.value as string,
    });
    switch (event.target.value as string) {
      case 'stack-monitoring':
        window.open(
          `/d/elastic-stack-monitoring-dashboard/elastic-stack-monitoring-dashboard?orgId=1&refresh=1m&var-cluster_uid=${uid}`,
          '_blank'
        );
        this.setState({
          monitorName: '',
        });

        break;
      case 'logstash-overview':
        window.open(`/d/logstash-overview/logstash-overview?orgId=1&refresh=1m&var-cluster_uid=${uid}`, '_blank');

        this.setState({
          monitorName: '',
        });
        break;
      case 'index-overview':
        window.open(
          `/d/elasticsearch-index-overview/elasticsearch-index-overview?orgId=1&refresh=1m&var-cluster_uid=${uid}`,
          '_blank'
        );

        this.setState({
          monitorName: '',
        });
        break;
      case 'shards-overview':
        window.open(
          `/d/elasticsearch-shards-overview-dashboard/elasticsearch-shards-overview-dashboard?orgId=1&refresh=1m&var-cluster_uid=${uid}`,
          '_blank'
        );

        this.setState({
          monitorName: '',
        });
        break;
      case 'ingest-pipelines-overview':
        window.open(
          `/d/elasticsearch-ingest-pipelines-overview/elasticsearch-ingest-pipelines-overview?orgId=1&refresh=1m&var-cluster_uid=${uid}`,
          '_blank'
        );

        this.setState({
          monitorName: '',
        });
        break;
      case 'tasks-analytics':
        window.open(
          `/d/elasticsearch-tasks-analytics/elasticsearch-tasks-analytics?orgId=1&refresh=1m&var-cluster_uid=${uid}`,
          '_blank'
        );

        this.setState({
          monitorName: '',
        });
        break;
      case 'ml-jobs-analytics':
        window.open(
          `/d/ml-jobs-analytics-dashboard/ml-jobs-analytics-dashboard?orgId=1&var-cluster_uid=${uid}`,
          '_blank'
        );
        this.setState({
          monitorName: '',
        });
        break;
    }
  };

  async componentDidMount() {
    // console.log('Props: ', this.props.theme)
    await getBackendSrv()
      .get(
        `/api/datasources/proxy/uid/${this.props.dataSourceItem.uid}/_cluster/stats?filter_path=cluster_uuid,cluster_name,status,nodes.versions,indices.count,indices.shards.total,indices.docs.count,indices.store.size_in_bytes,nodes.fs.total_in_bytes,nodes.count.total,nodes.count.data,nodes.count.data_hot,nodes.count.data_warm,nodes.count.data_cold`
      )
      .then((dataSources: any) => {
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
        let regex = new RegExp(/Elasticsearch-direct-prod-(.*)--(.*)/g);
        const uid: string = this.props.dataSourceItem.uid;
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
    this.componentDidMount().then(() => {});
  };
  onDeleteCluster = async () => {
    const backendSrv = getBackendSrv();
    this.loading = true;
    try {
      let dataSources: any = await backendSrv.get(`/api/datasources`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
      });

      for (const dataSource of dataSources) {
        if (dataSource.uid.endsWith(this.state.cluster_uuid)) {
          try {
            await backendSrv.delete(`/api/datasources/uid/${dataSource.uid}`);
          } catch (deleteError) {
            console.error('Delete data sources error: ', deleteError);
            this.loading = false;
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
        console.error('Delete from backend error: ', error);
        toast.error(`${error.message}`, {
          position: toast.POSITION.BOTTOM_RIGHT,
          autoClose: false,
          closeButton: true,
          hideProgressBar: true,
          draggable: false,
        });
        this.loading = false;
      }
      // this.props.onDelete(this.state.cluster_uuid)
      // this.componentDidMount().then(() => {
      // });
      this.loading = false;
      window.location.reload();
    } catch (getError) {
      console.error('Cluster delete error: ', getError);
      this.loading = false;
    }
  };

  handleDelete(isOnYes: boolean): void {
    this.setState({ isOpenDialog: false });
    if (isOnYes) {
      this.onDeleteCluster().then(() => {});
    }
  }

  render() {
    return (
      <>
        {this.loading && <Spinner></Spinner>}

        <div
          className={classNames({
            form_group: true,
            isLight: this.props.theme.isLight,
          })}
        >
          <header>
            <div className="col header-cluster">
              <h3>{this.state.cluster_name}</h3>
              <p>{this.state.cluster_uuid}</p>
            </div>
            <div className="actions col">
              <span className={this.state.status.toUpperCase()}>{this.state.status.toUpperCase()}</span>
            </div>
          </header>

          <Divider light />

          <div className="grid-container">
            <div className="col">
              <List>
                <ListItem>
                  {this.state.versions ? (
                    <div>
                      <span className="label">Version</span>
                      <ListItemText primary={this.state.versions ?? '0'} />
                    </div>
                  ) : null}

                  <div>
                    <span className="label">Used storage</span>
                    <ListItemText primary={this.state.usedStorage ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Total storage</span>
                    <ListItemText primary={this.state.totalStorage ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Docs count</span>
                    <ListItemText primary={this.state.docsCount ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Total nodes</span>
                    <ListItemText primary={this.state.totalNodes ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Data nodes</span>
                    <ListItemText primary={this.state.dataNodes ?? '0'} />
                  </div>
                  <div>
                    <span className="label"> Hot nodes</span>
                    <ListItemText primary={this.state.dataHotNodes ?? '0'} />
                  </div>
                  <div>
                    <span className="label"> Warm nodes</span>
                    <ListItemText primary={this.state.dataWarmNodes ?? '0'} />
                  </div>
                  <div>
                    <span className="label"> Cold nodes</span>
                    <ListItemText primary={this.state.dataColdNodes ?? '0'} />
                  </div>

                  <div>
                    <span className="label">Indices</span>
                    <ListItemText primary={this.state.numberOfIndices ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Total shards</span>
                    <ListItemText primary={this.state.numberOfShards ?? '0'} />
                  </div>
                  <div>
                    <span className="label">Unassigned shards</span>
                    <ListItemText primary={this.state.numberOfUnassignedShards ?? '0'} />
                  </div>
                </ListItem>
              </List>
            </div>
            <div className="col"></div>
          </div>
          <Divider light />
          <footer>
            <Stack spacing={2} direction="row">
              <Button variant="secondary" onClick={this.onDelete}>
                Delete
              </Button>
              <Button variant="secondary" onClick={() => this.onTest()}>
                Test
              </Button>
              <FormControl fullWidth id="select">
                <Select
                  labelId="demo-simple-select-label"
                  id="demo-simple-select"
                  value={this.state.monitorName! ? this.state.monitorName : 'Monitor type'}
                  onChange={this.handleChange}
                  renderValue={(value) => {
                    let text = value.split('-').join(' ');
                    text = text.charAt(0).toUpperCase() + text.slice(1);
                    return text ?? 'Monitor type';
                  }}
                >
                  <MenuItem value={'stack-monitoring'}>Elastic Stack monitoring</MenuItem>
                  <MenuItem value={'index-overview'}>Elasticsearch Index overview</MenuItem>
                  <MenuItem value={'shards-overview'}>Elasticsearch Shards overview</MenuItem>
                  <MenuItem value={'ingest-pipelines-overview'}>Elasticsearch ingest pipelines overview</MenuItem>
                  <MenuItem value={'logstash-overview'}>Logstash overview</MenuItem>
                  <MenuItem value={'tasks-analytics'}>Tasks analytics</MenuItem>
                  <MenuItem value={'ml-jobs-analytics'}>Elasticsearch ML Jobs Analytics</MenuItem>
                </Select>
              </FormControl>
            </Stack>
          </footer>

          <Dialog
            open={this.state?.isOpenDialog!}
            onClose={() => this.handleDelete(false)}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
          >
            <DialogTitle id="alert-dialog-title">{'Are you sure you want to delete this cluster?'}</DialogTitle>

            <DialogActions>
              <Button variant="destructive" onClick={() => this.handleDelete(false)} className="btn-error">
                No
              </Button>
              <Button variant="primary" onClick={() => this.handleDelete(true)} autoFocus>
                Yes
              </Button>
            </DialogActions>
          </Dialog>
        </div>
      </>
    );
  }
}
