import React, { FC, useCallback, useEffect, useState } from 'react';
import { DataSourceItem } from '../DataSourceItem/DataSourceItem';
import './data-source-list.scss';
import { useTheme2 } from '@grafana/ui';
import classNames from 'classnames';
import { FormControl, IconButton, MenuItem, Select } from '@mui/material';
import { getBackendSrv } from '@grafana/runtime';
import TextField from '@mui/material/TextField';
import RefreshIcon from '@mui/icons-material/Refresh';
import { DataSource } from '../../types/datasource';
import { Cluster } from '../../types/cluster';

export const DataSourceList: FC = () => {
  const theme = useTheme2();
  const [dataSources, setDataSources] = useState<DataSource[]>([]);
  const [_, setLoading] = useState(true);
  const [selectedTime, setSelectedTime] = useState(60);
  const [filterText, setFilterText] = useState('');

  const fetchDataSources = useCallback(async () => {
    setLoading(true);
    try {
      let sources = await getBackendSrv().get('/api/datasources');
      const regex = /Elasticsearch-direct-prod-.*/g;
      sources = sources.filter((source: DataSource) => source.uid.match(regex));

      const detailedDataPromises = sources.map(async (source: DataSource) => {
        try {
          const detailedInfo = await getBackendSrv().get<Cluster>(
            `/api/datasources/proxy/uid/${source.uid}/_cluster/stats?filter_path=cluster_uuid,cluster_name,status,nodes.versions,indices.count,indices.shards.total,indices.docs.count,indices.store.size_in_bytes,nodes.fs.total_in_bytes,nodes.count.total,nodes.count.data,nodes.count.data_hot,nodes.count.data_warm,nodes.count.data_cold`
          );

          return { ...source, detailedInfo };
        } catch (error) {
          console.error(`Error fetching detailed info for ${source.uid}:`, error);
          return { ...source, detailedInfo: null }; // Ensure detailedInfo is set even if null
        }
      });

      const sourcesWithDetails = await Promise.all(detailedDataPromises);

      console.log(sourcesWithDetails);

      setDataSources(sourcesWithDetails);
    } catch (error) {
      console.error('Error fetching data sources:', error);
    } finally {
      setLoading(false);
    }
  }, []);
  const filteredDataSources = dataSources.filter((dataSource) =>
    dataSource.name.toLowerCase().includes(filterText.toLowerCase())
  );

  console.log(filteredDataSources);

  useEffect(() => {
    fetchDataSources();
  }, [fetchDataSources]);

  const handleFilterText = (value: string) => {
    setFilterText(value);
  };

  const handleRefreshClick = () => {
    fetchDataSources();
  };

  return (
    <div className={classNames({ container: true, isLight: theme.isLight })}>
      <header className="header">
        <h1>Clusters list</h1>
      </header>
      <div className="filters">
        <div className="refresh-input">
          <TextField
            label="Search..."
            variant="outlined"
            fullWidth
            value={filterText}
            onChange={(e) => handleFilterText(e.target.value)}
            margin="normal"
          />
        </div>
        <div className="refresh-select">
          <FormControl fullWidth>
            <Select
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={selectedTime}
              label="Refresh Time"
              onChange={(e) => setSelectedTime(e.target.value as number)}
            >
              <MenuItem value={60}>1min</MenuItem>
              <MenuItem value={120}>2min</MenuItem>
              <MenuItem value={300}>5min</MenuItem>
              <MenuItem value={600}>10min</MenuItem>
            </Select>
          </FormControl>
          <IconButton onClick={handleRefreshClick}>
            <RefreshIcon />
          </IconButton>
        </div>
      </div>
      <section className="card-section card-list-layout-list">
        <ul className="card-list" data-col={filteredDataSources.length}>
          {filteredDataSources.length &&
            filteredDataSources.map((item, index) => (
              <li className="card-item-wrapper" key={index} aria-label="check-card">
                <DataSourceItem
                  refreshTime={selectedTime}
                  onDelete={(id) => {}}
                  theme={theme}
                  dataSourceItem={item.detailedInfo} // Pass detailedInfo, even if null
                />
              </li>
            ))}
        </ul>
      </section>
    </div>
  );
};
