import React from 'react';
import {Button, Legend, useStyles2} from '@grafana/ui';
import {AppPluginMeta, GrafanaTheme2, PluginConfigPageProps, PluginMeta} from '@grafana/data';
import {getBackendSrv} from '@grafana/runtime';
import {css} from '@emotion/css';
import {lastValueFrom} from 'rxjs';
import {Project} from '../../panels/dbeast-add_new_es_cluster-panel/models/project';
import PasswordDialog from './password-dialog';
import "./AppConfig.scss"

export type AppPluginSettings = {};

export interface AppConfigProps extends PluginConfigPageProps<AppPluginMeta<AppPluginSettings>> {
}

export const AppConfig = ({plugin}: AppConfigProps) => {
    const s = useStyles2(getStyles);
    const {enabled, jsonData} = plugin.meta;

    const [uniqueProjects, setUniqueProjects] = React.useState<Project[]>([]);

    const [projectIndex, setProjectIndex] = React.useState(0);

    const [showDialog, setShowDialog] = React.useState(false);

    // const [project, setProject] = React.useState<Project>({
    //     host: "",
    //     authentication_enabled: false,
    //     username: "",
    //     status: "",
    //     password: ""
    // });

    const [isShowSpinner, setIsShowSpinner] = React.useState(false);

    const [showError, setShowError] = React.useState(false);

    const settings = require('../../panels/dbeast-add_new_es_cluster-panel/config.ts');







    const fetchDataSources = async () => {
        try {
            const dataSources = await getBackendSrv().get('/api/datasources');
            return dataSources.filter((dataSource: any) => /^Elasticsearch-direct-mon--.*$/.test(dataSource.uid));
        } catch (error) {
            console.error("Error fetching data sources:", error);
            return [];
        }
    };

    const extractUniqueProjects = (dataSources: any[]): Project[] => {
        const urlSet = new Set<string>();
        return dataSources.reduce<Project[]>((acc, dataSource) => {
            if (!urlSet.has(dataSource.url)) {
                urlSet.add(dataSource.url);
                acc.push({
                    host: dataSource.url,
                    authentication_enabled: dataSource.basicAuth,
                    username: dataSource.basicAuthUser,
                    status: "",
                    password: ""
                });
            }
            return acc;
        }, []);
    };

    const onUpgradeAll = async () => {
        const dataSources = await fetchDataSources();
        const projects = extractUniqueProjects(dataSources);
        setUniqueProjects(projects);
        if (projects.length > 0) {

                setProjectIndex(0);
                setShowDialog(true);

            setShowDialog(true);
        }
    };

    const onUpgrade = async (project: Project) => {
        setIsShowSpinner(true);
        const baseUrl = settings.SERVER_URL;
        const response = await getBackendSrv().post(`${baseUrl}/update_cluster`, JSON.stringify(project));
        console.log('Cluster updated successfully:', response);

        if(response === "True"){
            setIsShowSpinner(false);
            setShowDialog(false);
            setShowError(false)
        }else{
            setShowDialog(true);
            setShowError(true)
        }


    };
    const onCloseDialog = () => {
        setShowDialog(false);
    }
    const onSkip = () => {
        console.log("Skip");

        if (projectIndex < uniqueProjects.length - 1) {
            setProjectIndex(prevIndex => prevIndex + 1);
        } else {
            setShowDialog(false);
        }



    }
    return <div className="gf-form-group">
        <div>
            {/* Enable the plugin */}

            <Legend>Enable / Disable</Legend>
            {!enabled && <>
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


            </>}

            {/*Source connection*/}

            <div className="actions">
                <Button variant="primary" onClick={() => onUpgradeAll()}>Upgrade all</Button>
            </div>


            {showDialog && <PasswordDialog isShowError={showError} isShowSpinner={isShowSpinner} handleSkip={onSkip} handleClose={onCloseDialog} project={uniqueProjects[projectIndex]}
                                           handleUpgrade={(project) => onUpgrade(project)}/>}


            {/* Disable the plugin */}
            {enabled && <>
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

            </>}
        </div>
    </div>;
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
