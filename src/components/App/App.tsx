import React, { PureComponent } from 'react';
import { getBackendSrv } from '@grafana/runtime';
import { Alert } from '@grafana/ui';
import ClustersList from '../../pages/cluster-list-page/ClustersList';
import AddNewClusterPanel from '../../pages/add-new-cluster-page/AddNewClusterPanel';
import './app.scss';

/**
 * Properties
 */
interface Props {
  dataSources: any[];
  path?: string;
  query?: Record<string, any>;
  meta?: any;
}

interface State {
  dataSources: any[];
  loading: boolean;
}

export class App extends PureComponent<Props, State> {
  state: State = {
    loading: true,
    dataSources: [],
  };

  onDeleteDataSource = (id: string) => {};

  async componentDidMount() {
    const dataSources = await getBackendSrv()
      .get('/api/datasources')
      .then((dataSources: any[]) => {
        const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
        return dataSources.filter((dataSource: any) => {
          return dataSource.name.match(regex);
        });
      });

    this.setState({
      dataSources,
      loading: false,
    });
  }

  render() {
    const { path } = this.props;

    if (this.state.loading) {
      return (
        <Alert title="Loading..." severity="info">
          <p>Loading time depends on the number of configured data sources.</p>
        </Alert>
      );
    }

    if (path?.includes('add-new-cluster-page')) {
      return (
        <>
          <AddNewClusterPanel />
        </>
      );
    }

    if (path?.includes('cluster-list-page')) {
      return (
        <>
          <ClustersList />
        </>
      );
    }

    // Default to ClustersList
    return (
      <>
        <ClustersList />
      </>
    );
  }
}
