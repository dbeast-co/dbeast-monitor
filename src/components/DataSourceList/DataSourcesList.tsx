import * as React from 'react';
import { FC } from 'react';
import { DataSourceItem } from '../DataSourceItem/DataSourceItem';
import { getDataSourceListStyles } from './DataSourcesList.styles';
import { RefreshPicker, useTheme2 } from '@grafana/ui';
import { rangeUtil } from '@grafana/data';

interface Props {
  dataSources: any[];
  onDelete: (id: string) => void;
}

export const DataSourceList: FC<Props> = ({ dataSources, onDelete }) => {
  const theme = useTheme2();
  const styles = getDataSourceListStyles(theme);

  const [refreshInterval, setRefreshInterval] = React.useState('1m');
  const [refreshTick, setRefreshTick] = React.useState(0);

  // Items load their stats on mount, so changing the key remounts them and refetches.
  const onRefresh = React.useCallback(() => setRefreshTick((t) => t + 1), []);

  React.useEffect(() => {
    const ms = refreshInterval ? rangeUtil.intervalToMs(refreshInterval) : 0;
    if (!ms) {
      return;
    }
    const id = window.setInterval(onRefresh, ms);
    return () => window.clearInterval(id);
  }, [refreshInterval, onRefresh]);

  const onDeleteItem = (id: string) => {
    onDelete(id);
  };

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1>Clusters list</h1>
        <div className={styles.refreshPicker}>
          <RefreshPicker
            value={refreshInterval}
            intervals={['5s', '10s', '30s', '1m', '5m', '15m', '30m', '1h']}
            onRefresh={onRefresh}
            onIntervalChanged={setRefreshInterval}
            isOnCanvas
          />
        </div>
      </header>
      <section className={styles.cardSection}>
        <ul className={styles.cardList} data-col={dataSources.length}>
          {dataSources.map((item, index) => {
            return (
              <li className={styles.cardItemWrapper} key={`${item.uid ?? index}-${refreshTick}`} aria-label="check-card">
                <DataSourceItem onDelete={(id) => onDeleteItem(id)} dataSourceItem={item} theme={theme} />
              </li>
            );
          })}
        </ul>
      </section>
    </div>
  );
};
