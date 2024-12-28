import { NavModelItem } from '@grafana/data';
import pluginJson from './../plugin.json';

export const PLUGIN_BASE_URL = `/a/${pluginJson.id}`;

export enum ROUTES {
  Update1 = 'update-1',
  Update2 = 'update-2',
  Update3 = 'update-3',
}

export const NAVIGATION_TITLE = 'Basic App Plugin';
export const NAVIGATION_SUBTITLE = 'Some extra description...';

// Add a navigation item for each route you would like to display in the navigation bar
export const NAVIGATION: Record<string, NavModelItem> = {
  [ROUTES.Update1]: {
    id: ROUTES.Update1,
    text: 'Page One',
    icon: 'database',
    url: `${PLUGIN_BASE_URL}/one`,
  },
  [ROUTES.Update2]: {
    id: ROUTES.Update2,
    text: 'Page Two',
    icon: 'key-skeleton-alt',
    url: `${PLUGIN_BASE_URL}/two`,
  },
  [ROUTES.Update3]: {
    id: ROUTES.Update3,
    text: 'Page Three',
    icon: 'chart-line',
    url: `${PLUGIN_BASE_URL}/three`,
  },
};
