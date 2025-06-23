import pluginJson from '../plugin.json';
export const PLUGIN_BASE_URL = `/a/${pluginJson.id}`;

// Prefixes the route with the base URL of the plugin
export function prefixRoute(route: string): string {
    return `${PLUGIN_BASE_URL}/${route}`;
}