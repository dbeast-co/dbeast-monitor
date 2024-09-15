import React from 'react';
import {Button, Legend, useStyles2} from '@grafana/ui';
import {AppPluginMeta, GrafanaTheme2, PluginConfigPageProps, PluginMeta} from '@grafana/data';
import {getBackendSrv} from '@grafana/runtime';
import {css} from '@emotion/css';
import {lastValueFrom} from 'rxjs';

export type AppPluginSettings = {};

export interface AppConfigProps extends PluginConfigPageProps<AppPluginMeta<AppPluginSettings>> {
}

export const AppConfig = ({plugin}: AppConfigProps) => {
    const s = useStyles2(getStyles);
    const {enabled, jsonData} = plugin.meta;

    return (
        <div className="gf-form-group">
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
        console.error('Error while updating the plugin', e);
    }
};

export const updatePlugin = async (pluginId: string, data: Partial<PluginMeta>) => {
    const response = getBackendSrv().fetch({
        url: `/api/plugins/${pluginId}/settings`,
        method: 'POST',
        data,
    });
    return lastValueFrom(response);
};
