import React, {FC,} from 'react';
import {DataSourceItem} from '../DataSourceItem/DataSourceItem';
import './data-source-list.scss';
import {useTheme2} from '@grafana/ui';
import classNames from 'classnames';
import {MenuItem, Select} from '@mui/material';
import TextField from '@mui/material/TextField';

/**
 * Properties
 */
interface Props {
  dataSources: any[];
  onDelete: (id: string) => void;
}

export const DataSourceList: FC<Props> = ({ dataSources, onDelete }) => {
  const theme = useTheme2();
  const [selectedSortOption, setSelectedSortOption] = React.useState('');

  const [searchTerm, setSearchTerm] = React.useState('');

  const [dataSourcesState, setDatasourcesState] = React.useState(dataSources);


  const onDeleteItem = (id: string) => {
    // onDelete(id);

  };
  const handleSelectSortOption = (event: string
  ) => {
    setSelectedSortOption(event as string);
    console.log(event);
    const sortedDatasources = dataSources.sort((a, b) => {
        if (event === 'GREEN') {
            return a.status === 'GREEN' ? -1 : 1;
        }
        if (event === 'YELLOW') {
            return a.status === 'YELLOW' ? -1 : 1;
        }
        if (event === 'ERROR' || event === 'RED') {
            return a.status === 'ERROR' ? -1 : 1 || a.status === 'RED' ? -1 : 1;
        }
        return 0;
    });
    setDatasourcesState(sortedDatasources);
  };

  const handleSearchCluster = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
    const filteredDatasources = dataSources.map((item)=> {
        return {
          ...item,
          name: item.name.replace('Elasticsearch-direct-prod-', '').split('--')[0],
        };


    })
        .filter((item) => {

        return item.name.toLowerCase().includes(event.target.value.toLowerCase());
    });
    setDatasourcesState(filteredDatasources);


  };


  return (
    <div
      className={classNames({
        container: true,
        isLight: theme.isLight,
      })}
    >
      <header className="header">

        <h1>Clusters list</h1>

        <div className="sorting-filters-container">
          <div className="search-filter">
            <TextField
                value={searchTerm}
                onChange={handleSearchCluster}
                label="Search"
                variant="outlined"
                fullWidth
            />
          </div>

          <div className="sorting-filters">
            <Select
                className="sorting-select"
                value={selectedSortOption}
                onChange={select => handleSelectSortOption(select.target.value)}
                displayEmpty
            >
              <MenuItem value="" disabled>
                Sort by
              </MenuItem>
              <MenuItem value="GREEN">Green</MenuItem>
              <MenuItem value="YELLOW">Yellow</MenuItem>
              <MenuItem value="ERROR">Error</MenuItem>
            </Select>

          </div>
        </div>
      </header>
      <section className="card-section card-list-layout-list">
        <ul className="card-list" data-col={dataSources.length}>
          {dataSourcesState.map((item, index) => {
            return (
              <li className="card-item-wrapper" key={index} aria-label="check-card">
                <DataSourceItem onDelete={(id) => onDeleteItem(id)} dataSourceItem={item} theme={theme} />
              </li>
            );
          })}
        </ul>
      </section>
    </div>
  );
};
