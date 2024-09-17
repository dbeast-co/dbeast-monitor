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
