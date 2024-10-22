import React, {FC,} from 'react';
import {DataSourceItem} from '../DataSourceItem/DataSourceItem';
import './data-source-list.scss';
import {useTheme2} from '@grafana/ui';
import classNames from 'classnames';
import {FormControl, InputLabel, MenuItem, Select} from '@mui/material';

/**
 * Properties
 */
interface Props {
  dataSources: any[];
  onDelete: (id: string) => void;
}

export const DataSourceList: FC<Props> = ({ dataSources, onDelete }) => {
  const theme = useTheme2();

  const onDeleteItem = (id: string) => {
    // onDelete(id);

  };
  const [selectedTime, setSelectedTime] = React.useState(0);


  function onSelectAutoRefresh(event: any) {
    console.log('onSelectAutoRefresh');
    setSelectedTime(event?.target!.value);

  }

  return (
    <div
      className={classNames({
        container: true,
        isLight: theme.isLight,
      })}
    >
      <header className="header">
        <h1>Clusters list</h1>
      </header>

      <div className="filters"><FormControl fullWidth>
        <InputLabel id="demo-simple-select-label">Refresh interval</InputLabel>
        <Select
            labelId="demo-simple-select-label"
            id="demo-simple-select"
            value={selectedTime || null}
            label="Age"
            onChange={onSelectAutoRefresh}
        >
          <MenuItem value={10}>10sec</MenuItem>
          <MenuItem value={60}>1min</MenuItem>
          <MenuItem value={120}>2min</MenuItem>
          <MenuItem value={300}>5min</MenuItem>
          <MenuItem value={600}>10min</MenuItem>
        </Select>
      </FormControl></div>
      <section className="card-section card-list-layout-list">


        <ul className="card-list" data-col={dataSources.length}>
          {dataSources.map((item, index) => {
            return (
              <li className="card-item-wrapper" key={index} aria-label="check-card">
                <DataSourceItem refreshTime={selectedTime} onDelete={(id) => onDeleteItem(id)} dataSourceItem={item} theme={theme} />
              </li>
            );
          })}
        </ul>
      </section>
    </div>
  );
};
