import React, {PureComponent} from 'react';
import {getBackendSrv} from '@grafana/runtime';
import {Alert} from '@grafana/ui';
import {DataSourceList} from '../DataSourceList/DataSourcesList';
import './app.scss';

/**
 * Properties
 */
interface Props {
    dataSources: any[];
}

interface State {
    dataSources: any[];

    loading: boolean;
}

export class App extends PureComponent<Props, State> {
    state: State = {
        loading: true,
        dataSources: [],
    };
    onDeleteDataSource = (id: string) => {

        const filteredDataSources = this.state.dataSources.filter((item) => !item.uid.endsWith(id));

        this.setState({dataSources: filteredDataSources});
    };


    async componentDidMount() {
        const dataSources = await getBackendSrv()
            .get('/api/datasources')
            .then((dataSources: any[]) => {
                const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
                return dataSources.filter((dataSource: any) => {
                    return dataSource.uid.match(regex);
                });
            });

        this.setState({
            dataSources,
            loading: false,
        });
    }

    render() {

        if (this.state.loading) {
            return (
                <Alert title="Loading..." severity="info">
                    <p>Loading time depends on the number of configured data sources.</p>
                </Alert>
            );
        }
        return (
            <>
                <div>
                    <h1>Welcome to the Monitoring toolkit for the Elastic Stack</h1>
                </div>
                <DataSourceList onDelete={(id) => this.onDeleteDataSource(id)} dataSources={this.state.dataSources}/>;
                )</>
        );
    }
}
