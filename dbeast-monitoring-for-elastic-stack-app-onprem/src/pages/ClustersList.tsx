import React, {useEffect, useState} from 'react';
import {getBackendSrv, PluginPage} from '@grafana/runtime';
import {App} from "../components/App";
import {Alert} from "@grafana/ui";

export function ClustersList() {
  const [loading, setLoading] = useState(true);
  const [dataSources, setDataSources] = useState([] as any[]);




    useEffect( () =>  {
        setLoading(true);
        const fetchDataSources = async () => {
            const dataSources = await getBackendSrv()
                .get('/api/datasources')
                .then((dataSources: any[]) => {
                    const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
                    setLoading(false);
                    return dataSources.filter((dataSource: any) => {
                        return dataSource.name.match(regex);
                    });
                });

            setDataSources(dataSources);
            setLoading(false);
        };

        fetchDataSources();
    },[loading, dataSources])


    return (
        <PluginPage>
            <div>

                {loading && <Alert title="Loading..." severity="info">
                    <p>Loading time depends on the number of configured data sources.</p>
                </Alert>}

                <App dataSources={dataSources}/>
            </div>
        </PluginPage>
    );
}
