import React, {useState} from 'react';
import {DataSourceList} from '../DataSourceList/DataSourcesList';
import './app.scss';

/**
 * Properties
 */
interface Props {
  dataSources: any[];
}

export const App: React.FC<Props> = ({ dataSources }) => {
  const [currentDataSources, setCurrentDataSources] = useState(dataSources);

  const handleDeleteDataSource = (id: string) => {
    const filteredDataSources = currentDataSources.filter(
        (item) => !item.uid.endsWith(id)
    );
    setCurrentDataSources(filteredDataSources);
  };




  return (
      <DataSourceList
          onDelete={(id) => handleDeleteDataSource(id)}
          dataSources={currentDataSources}
      />
  );
};