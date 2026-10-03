import { InlineField, InlineFieldRow, RadioButtonGroup, Segment } from '@grafana/ui';

import React, { useState } from 'react';
import { defaultQuery, JsonApiQuery, Pair } from '../types';
import { KeyValueEditor } from './KeyValueEditor';
import { PathEditor } from './PathEditor';
import { FieldEditor } from './FieldEditor';

interface Props {
  onChange: (query: JsonApiQuery) => void;
  onRunQuery: () => void;
  editorContext: string;
  query: JsonApiQuery;

  fieldsTabConfig?: { value: any; onChange: (value: any) => void; limit?: number };
  experimentalTab: React.ReactNode;
}

export const TabbedQueryEditor = ({ query, onChange, onRunQuery, fieldsTabConfig, experimentalTab }: Props) => {
  const [tabIndex, setTabIndex] = useState(0);
  const q = { ...defaultQuery, ...query };

  const onParamsChange = (params: Array<Pair<string, string>>) => {
    onChange({ ...q, params });
    onRunQuery();
  };

  const tabs = [
    {
      title: 'Fields',
      content: fieldsTabConfig ? (
        <FieldEditor
          value={fieldsTabConfig.value}
          onChange={fieldsTabConfig.onChange}
          limit={fieldsTabConfig.limit}
        />
      ) : null,
    },
    {
      title: 'Path',
      content: (
        <PathEditor
          path={q.urlPath ?? ''}
          onPathChange={(path) => {
            onChange({ ...q, urlPath: path });
          }}
        />
      ),
    },
    {
      title: 'Params',
      content: (
        <KeyValueEditor
          addRowLabel={'Add param'}
          columns={['Key', 'Value']}
          values={q.params ?? []}
          onChange={onParamsChange}
          onBlur={() => onRunQuery()}
        />
      ),
    },
    {
      title: 'Experimental',
      content: experimentalTab,
    },
  ];

  return (
    <>
      <InlineFieldRow>
        <InlineField>
          <RadioButtonGroup
            onChange={(e) => setTabIndex(e ?? 0)}
            value={tabIndex}
            options={tabs.map((tab, idx) => ({ label: tab.title, value: idx }))}
          />
        </InlineField>
        <InlineField
          label="Cache Time"
          tooltip="Time in seconds that the response will be cached in Grafana after receiving it."
        >
          <Segment
            value={{ label: formatCacheTimeLabel(q.cacheDurationSeconds), value: q.cacheDurationSeconds }}
            options={[0, 5, 10, 30, 60, 60 * 2, 60 * 5, 60 * 10, 60 * 30, 3600, 3600 * 2, 3600 * 5].map((value) => ({
              label: formatCacheTimeLabel(value),
              value,
              description: value ? '' : 'Response is not cached at all',
            }))}
            onChange={({ value }) => onChange({ ...q, cacheDurationSeconds: value! })}
          />
        </InlineField>
      </InlineFieldRow>
      {tabs[tabIndex].content}
    </>
  );
};

export const formatCacheTimeLabel = (s: number) => {
  if (s < 60) {
    return s + 's';
  } else if (s < 3600) {
    return s / 60 + 'm';
  }

  return s / 3600 + 'h';
};
