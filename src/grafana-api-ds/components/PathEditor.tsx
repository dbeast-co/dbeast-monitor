import { InlineField, InlineFieldRow, Input } from '@grafana/ui';
import React from 'react';

interface Props {
  path: string;
  onPathChange: (path: string) => void;
}

export const PathEditor = ({ path, onPathChange }: Props) => {
  return (
    <InlineFieldRow>
      <InlineField label="API Path" grow tooltip="Grafana API endpoint path (e.g., /api/datasources, /api/dashboards)">
        <Input 
          placeholder="/api/datasources" 
          value={path} 
          onChange={(e) => onPathChange(e.currentTarget.value)}
          autoComplete="off"
        />
      </InlineField>
    </InlineFieldRow>
  );
};
