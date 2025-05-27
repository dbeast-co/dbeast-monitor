import React, {PureComponent} from 'react';
import {getBackendSrv} from '@grafana/runtime';
import {Alert} from '@grafana/ui';
import './app.scss';
import {DataSourceList} from "../DataSourceList/DataSourcesList";

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
    //
    // const filteredDataSources = this.state.dataSources.filter((item) => !item.uid.endsWith(id));
    //
    // this.setState({ dataSources: filteredDataSources });
  };




  async componentDidMount() {
    const dataSources = await getBackendSrv()
      .get('/api/datasources')
      .then((dataSources: any[]) => {
        const regex = new RegExp(/Elasticsearch-direct-prod-.*/g);
        return dataSources.filter((dataSource: any) => {
          return dataSource.name.match(regex);
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
    // const path = window.location.pathname;

    // if (path.includes('/clusters')) {
    //   return <ClustersList />;
    // }
    return (
    //     <Routes>
    //       <Route path="one" element={<ClustersList />} />
    //       {/*<Route path={`${ROUTES.Three}/:id?`} element={<PageThree />} />*/}
    //
    //       {/*/!* Full-width page (this page will have no side navigation) *!/*/}
    //       {/*<Route path={ROUTES.Four} element={<PageFour />} />*/}
    //
    //       {/*/!* Default page *!/*/}
    //       {/*<Route path="*" element={<PageOne />} />*/}
    //     </Routes>
        <>
          <DataSourceList onDelete={(id) => this.onDeleteDataSource(id)} dataSources={this.state.dataSources}/>
          </>
    );
  }
}
