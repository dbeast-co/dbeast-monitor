import { AppPlugin } from '@grafana/data';
import { App } from './components/App';
import { AppConfig } from './components/AppConfig';

// @ts-ignore
export const plugin = new AppPlugin<{}>().setRootPage(App).addConfigPage({
    title: 'Configuration',
    icon: 'cog',
    body: AppConfig,
    id: 'configuration',
  });
    // .addConfigPage({
    //     title: 'Clusters List', // Title for the Clusters List page
    //     icon: 'list-ul', // Icon to represent this page
    //     body: ClustersList, // Pass the ClustersList component
    //     id: 'clusters-list', // Unique ID for this page
    // });

