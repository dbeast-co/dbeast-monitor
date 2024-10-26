export interface Cluster {
  cluster_name: string;
  cluster_uuid: string;
  status: string;
  indices: Indices;
  nodes: Nodes;
}

export interface Indices {
  count: number;
  shards: Shards;
  docs: Docs;
  store: Store;
}

export interface Docs {
  count: number;
}

export interface Shards {
  total: number;
}

export interface Store {
  size_in_bytes: number;
}

export interface Nodes {
  count: Count;
  versions: string[];
  fs: FS;
}

export interface Count {
  total: number;
  data: number;
  data_cold: number;
  data_hot: number;
  data_warm: number;
}

export interface FS {
  total_in_bytes: number;
}
