export interface GrafanaDatasource {
    basicAuth: boolean
    access: string
    basicAuthUser: string
    database: string
    isDefault: boolean
    jsonData: {
        esVersion: string
        includeFrozen: boolean
        logLevelField: string
        logMessageField: string
        maxConcurrentShardRequests: number
        timeField: string
        tlsSkipVerify: boolean
    },
    name: string
    orgId: number
    readOnly: boolean
    secureJsonData: {
        basicAuthPassword: string
    },
    type: string
    typeName: string
    uid: string
    url: string
    withCredentials: boolean
}
