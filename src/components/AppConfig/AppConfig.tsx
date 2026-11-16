import * as React from 'react';
import {Alert, Button, Legend, useStyles2} from '@grafana/ui';
import {AppPluginMeta, GrafanaTheme2, PluginConfigPageProps, PluginMeta} from '@grafana/data';
import {getBackendSrv} from '@grafana/runtime';
import {css} from '@emotion/css';
import {lastValueFrom} from 'rxjs';

export type AppPluginSettings = {};

export interface AppConfigProps extends PluginConfigPageProps<AppPluginMeta<AppPluginSettings>> {
}

const GRAFANA_API_DATASOURCE_UID = 'dbeast-grafanaapi-datasource';
const GRAFANA_API_DATASOURCE_TYPE = 'dbeast-grafanaapi-datasource';

interface DataSourceDefinition {
    uid: string;
    type: string;
    name: string;
}

function hasStatus(error: unknown, status: number): boolean {
    return typeof error === 'object' && error !== null && 'status' in error && error.status === status;
}

async function ensureGrafanaApiDataSource(pluginId: string): Promise<void> {
    const backend = getBackendSrv();
    const url = `/api/datasources/uid/${GRAFANA_API_DATASOURCE_UID}`;
    const lookup = () => backend.get<DataSourceDefinition>(url, undefined, undefined, {showErrorAlert: false});
    const verify = (dataSource: DataSourceDefinition) => {
        if (dataSource.type !== GRAFANA_API_DATASOURCE_TYPE) {
            throw new Error(`Datasource UID ${GRAFANA_API_DATASOURCE_UID} is already used by another datasource type.`);
        }
    };

    try {
        verify(await lookup());
        return;
    } catch (error) {
        if (!hasStatus(error, 404)) {
            throw error;
        }
    }

    const definition = await backend.get<DataSourceDefinition>(
        `/api/plugins/${encodeURIComponent(pluginId)}/resources/grafana_api_datasource`,
        undefined,
        undefined,
        {showErrorAlert: false}
    );
    if (definition.uid !== GRAFANA_API_DATASOURCE_UID || definition.type !== GRAFANA_API_DATASOURCE_TYPE || !definition.name) {
        throw new Error('The Grafana API datasource definition is invalid.');
    }
    try {
        await backend.post('/api/datasources', definition, {showErrorAlert: false, showSuccessAlert: false});
    } catch (error) {
        if (!hasStatus(error, 409)) {
            throw error;
        }
        verify(await lookup());
    }
}

export const AppConfig = ({plugin}: AppConfigProps) => {
    const s = useStyles2(getStyles);
    const {enabled, jsonData} = plugin.meta;
    const [dataSourceError, setDataSourceError] = React.useState<string>();
    const [retry, setRetry] = React.useState(0);

    React.useEffect(() => {
        if (!enabled) {
            return;
        }
        let active = true;
        setDataSourceError(undefined);
        ensureGrafanaApiDataSource(plugin.meta.id).catch((error: unknown) => {
            if (active) {
                const message = error instanceof Error ? error.message :
                    typeof error === 'object' && error !== null && 'message' in error && typeof error.message === 'string'
                        ? error.message : 'Check your Grafana permissions and try again.';
                setDataSourceError(message);
            }
        });
        return () => { active = false; };
    }, [enabled, plugin.meta.id, retry]);

    return (
        <div className="gf-form-group">
            {enabled && dataSourceError && (
                <Alert severity="error" title="Failed to configure the Grafana API datasource">
                    <p>{dataSourceError}</p>
                    <Button onClick={() => setRetry((value) => value + 1)}>Retry</Button>
                </Alert>
            )}
            <div>
                {/* Enable the plugin */}
                <Legend>Enable / Disable</Legend>
                {!enabled && (
                    <>
                        <div className={s.colorWeak}>The plugin is currently not enabled.</div>
                        <Button
                            className={s.marginTop}
                            variant="primary"
                            onClick={() =>
                                updatePluginAndReload(plugin.meta.id, {
                                    enabled: true,
                                    pinned: true,
                                    jsonData,
                                })
                            }
                        >
                            Enable plugin
                        </Button>


                    </>
                )}

                {/*Source connection*/}


                {/* Disable the plugin */}
                {enabled && (
                    <>
                        <div className={s.colorWeak}>The plugin is currently enabled.</div>
                        <Button
                            className={s.marginTop}
                            variant="destructive"
                            onClick={() =>
                                updatePluginAndReload(plugin.meta.id, {
                                    enabled: false,
                                    pinned: false,
                                    jsonData,
                                })
                            }
                        >
                            Disable plugin
                        </Button>
                    </>
                )}
            </div>
        </div>
    );
};

const getStyles = (theme: GrafanaTheme2) => ({
    colorWeak: css`
        color: ${theme.colors.text.secondary};
    `,
    marginTop: css`
        margin-top: ${theme.spacing(3)};
    `,
});

const updatePluginAndReload = async (pluginId: string, data: Partial<PluginMeta>) => {
    try {
        await updatePlugin(pluginId, data);
        window.location.reload();
    } catch (e) {
      // eslint-disable-next-line no-console
      console.debug('Error while updating the plugin', e);
    }
};

export const updatePlugin = async (pluginId: string, data: Partial<PluginMeta>) => {
    const response = getBackendSrv().fetch({
        url: `/api/plugins/${pluginId}/settings`,
        method: 'POST',
        data,
    });
    return lastValueFrom(response as any);
};
