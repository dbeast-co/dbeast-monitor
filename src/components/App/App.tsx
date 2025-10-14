import React, { useEffect, useState } from 'react';
import { getBackendSrv } from '@grafana/runtime';
import ClustersList from '../../pages/cluster-list-page/ClustersList';
import AddNewClusterPanel from '../../pages/add-new-cluster-page/AddNewClusterPanel';
import { getAppStyles } from './App.styles';
import { Global, css } from '@emotion/react';

import { useTheme2, Alert } from '@grafana/ui';

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

  return (
    <>
      <Global
        styles={css`
          .MuiMenu-paper,
          .MuiPaper-root {
            background: #232733 !important;
            box-shadow: none !important;
            border-radius: 0 !important;
            margin-top: 0 !important;
            padding: 0 !important;
          }
          ul.MuiMenu-list {
            padding: 0 !important;
            background: #232733 !important;
          }
          ul.MuiMenu-list li {
            padding: 10px !important;
            
            &:hover {
              background: #2c3042 !important;
              color: ${theme.colors.text.primary} !important;
            }
          }
        `}
      />
      {path?.includes('add-new-cluster-page') ? (
        <div className={styles.container}>
          <AddNewClusterPanel />
        </div>
      ) : path?.includes('cluster-list-page') ? (
        <div className={styles.container}>
          <ClustersList />
        </div>
      ) : (
        <div className={styles.container}>
          <ClustersList />
        </div>
      )}
    </>
  );
};
