import { Cluster } from '../panels/dbeast-add_new_es_cluster-panel/models/cluster';

export interface DataSource {
  id: number;
  uid: string;
  orgId: number;
  name: string;
  type: string;
  typeName: string;
  typeLogoUrl: string;
  access: string;
  url: string;
  user: string;
  database: string;
  basicAuth: boolean;
  isDefault: boolean;
  jsonData: JSONData;
  readOnly: boolean;
  detailedInfo: Cluster;
}

export interface JSONData {
  tlsSkipVerify: boolean;
}
