import * as React from 'react';
import { FC } from 'react';
import { DataSourceItem } from '../DataSourceItem/DataSourceItem';
import { getDataSourceListStyles } from './DataSourcesList.styles';
import { useTheme2 } from '@grafana/ui';

interface Props {
  dataSources: any[];
  onDelete: (id: string) => void;
}

export const DataSourceList: FC<Props> = ({ dataSources, onDelete }) => {
  const theme = useTheme2();
  const styles = getDataSourceListStyles(theme);

  const onDeleteItem = (id: string) => {
    onDelete(id);
  };


  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1>Clusters list</h1>
      </header>
      <section className={styles.cardSection}>
        <ul className={styles.cardList} data-col={dataSources.length}>
          {dataSources.map((item, index) => {
            return (
              <li className={styles.cardItemWrapper} key={index} aria-label="check-card">
                <DataSourceItem onDelete={(id) => onDeleteItem(id)} dataSourceItem={item} theme={theme} />
              </li>
            );
          })}
        </ul>
      </section>
    </div>
  );
};
