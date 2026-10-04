import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { Alert } from '@grafana/ui';
import React from 'react';
import { JsonApiDataSourceOptions } from '../types';
import { Divider } from './Divider';

type Props = DataSourcePluginOptionsEditorProps<JsonApiDataSourceOptions>;

/**
 * ConfigEditor for Grafana Internal API Datasource.
 * No configuration needed - this datasource connects only to the local Grafana instance.
 */
export const ConfigEditor: React.FC<Props> = () => {
  return (
    <>
      <Divider />

      <Alert
        severity="info"
        title="Grafana Internal API Datasource"
        style={{ maxWidth: '700px', whiteSpace: 'normal' }}
      >
        This datasource connects to the internal APIs of your Grafana instance. No configuration is needed - enter the
        API path in your query (e.g., /api/datasources, /api/dashboards). Only GET requests are allowed.{' '}
        <a href="https://github.com/dbeast-co/dbeast-monitor/wiki" target="_blank" rel="noreferrer">
          Grafana internal API calls documentation
        </a>
      </Alert>
    </>
  );
};
