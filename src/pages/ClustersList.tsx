import React, {useEffect, useState} from 'react';
import {DataSourceList} from "../components/DataSourceList/DataSourcesList";
import {getBackendSrv} from "@grafana/runtime";

function ClustersList(props: any) {
    const [dataSources, setDataSources] = useState([] as any[]);
    const onDeleteDataSource = (id: string) => {
        //
        // const filteredDataSources = this.state.dataSources.filter((item) => !item.uid.endsWith(id));
        //
        // this.setState({ dataSources: filteredDataSources });
    };
    useEffect(() => {
        const fetchDataSources = async () => {
            const dataSourcesFromAPI: any[] = await getBackendSrv()
                .get('/api/datasources')
                .then((dataSources: any[]) => {
                    const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
                    return dataSources.filter((dataSource: any) => {
                        return dataSource.name.match(regex);
                    });
                });
            setDataSources(dataSourcesFromAPI);
        };

        fetchDataSources(); // Call the async function
    }, []); // Add the dependency array


    return (
        <>
            <DataSourceList onDelete={(id) => onDeleteDataSource(id)} dataSources={dataSources}/>
        </>
    );
}

export default ClustersList;

