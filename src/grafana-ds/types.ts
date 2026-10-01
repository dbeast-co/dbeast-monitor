import { DataSourceJsonData, FieldType } from '@grafana/data';
import { DataQuery } from "@grafana/schema"

export type QueryLanguage = 'jsonpath' | 'jsonata';

export interface JsonField {
  name?: string;
  jsonPath: string;
  type?: FieldType;
  language?: QueryLanguage;
}

export type Pair<T, K> = [T, K];

export interface JsonApiQuery extends DataQuery {
  fields: JsonField[];
  urlPath: string;
  params: Array<Pair<string, string>>;
  cacheDurationSeconds: number;

  // Keep for backwards compatibility with older version of variables query editor.
  jsonPath?: string;

  // Experimental
  experimentalGroupByField?: string;
  experimentalMetricField?: string;
  experimentalVariableTextField?: string;
  experimentalVariableValueField?: string;
}

export const defaultQuery: Partial<JsonApiQuery> = {
  cacheDurationSeconds: 300,
  urlPath: '',
  params: [],
  fields: [{ jsonPath: '' }],
};

export interface JsonApiDataSourceOptions extends DataSourceJsonData {}
