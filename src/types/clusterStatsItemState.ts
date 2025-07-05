export interface ClusterStatsItemState {
  cluster_name: string;
  cluster_uuid: string;
  status: string;
  versions: string[];
  numberOfIndices: number;
  numberOfShards: number;
  numberOfUnassignedShards: number;
  docsCount: string;
  usedStorage: string;
  totalStorage: string;
  totalNodes: number;
  dataNodes: number;
  dataHotNodes: number;
  dataWarmNodes: number;
  dataColdNodes: number;
  monitorName: string | null | undefined;
  isOpenDialog?: boolean;
  isLoading?: boolean;
}

export interface MonitorState {
  monitorName: string | null | undefined;
}
