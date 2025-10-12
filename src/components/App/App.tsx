import React, { useEffect, useState } from 'react';
import { getBackendSrv } from '@grafana/runtime';
import { Alert, useTheme2 } from '@grafana/ui';
import ClustersList from '../../pages/cluster-list-page/ClustersList';
import AddNewClusterPanel from '../../pages/add-new-cluster-page/AddNewClusterPanel';
import { getAppStyles } from './App.styles';

/**
 * Properties
 */
interface Props {
  dataSources: any[];
  path?: string;
  query?: Record<string, any>;
  meta?: any;
}

export const App: React.FC<Props> = ({ path }) => {
  const theme = useTheme2();
  const styles = getAppStyles(theme);

  const [state, setState] = useState({
    loading: true,
    dataSources: [],
  });

  useEffect(() => {
    const fetchDataSources = async () => {
      const dataSources = await getBackendSrv()
        .get('/api/datasources')
        .then((dataSources: any[]) => {
          const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
          return dataSources.filter((dataSource: any) => {
            return dataSource.name.match(regex);
          });
        });

      setState({
        dataSources,
        loading: false,
      });
    };

    fetchDataSources();
  }, []);

  if (state.loading) {
    return (
      <div className={styles.container}>
        <Alert title="Loading..." severity="info">
          <p>Loading time depends on the number of configured data sources.</p>
        </Alert>
      </div>
    );
  }

  if (path?.includes('add-new-cluster-page')) {
    return (
      <div className={styles.container}>
        <AddNewClusterPanel />
      </div>
    );
  }

  if (path?.includes('cluster-list-page')) {
    return (
      <div className={styles.container}>
        <ClustersList />
      </div>
    );
  }

  // Default to ClustersList
  return (
    <div className={styles.container}>
      <ClustersList />
    </div>
  );
};
