import React, { useEffect, useState } from 'react';
import './AddNewClusterPanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { CircularProgress, Dialog, DialogActions, DialogContent, DialogTitle, Divider } from '@mui/material';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';


import { useTheme2 } from '@grafana/ui';
import classNames from 'classnames';
import { getBackendSrv } from '@grafana/runtime';
import { ConnectionSettings } from '../models/connection-settings';
import { Datasource } from '../models/datasource';
import { GrafanaDatasource } from '../models/grafana-datasource';
import { BackendResponse } from '../models/backend-response';
import { SERVER_URL } from '../config';
import { Cluster, Host } from '../models/cluster';
import { saveAs } from 'file-saver';

import { LogstashConfigurationsPanel } from './LogstashConfigurationsPanel';
import LogstashComponent, { Logstash } from './logstash';
import { v4 as uuidv4 } from 'uuid';
import { MonitoringClusterInjectionPanel } from './MonitoringClusterInjectionPanel';

import Tooltip from '@mui/material/Tooltip';

const settings = require('../config.ts');

export enum LogstashFileType {
  ES_MONITORING_CONFIGURATION_FILES = 'download_es_monitoring_configuration_files',
  LOGSTASH_MONITORING_CONFIGURATION_FILES = 'download_logstash_monitoring_configuration_files',
}

export const AddNewClusterPanel = () => {
  const backendSrv = getBackendSrv();
  const theme = useTheme2();
  const baseUrl = settings.SERVER_URL;

  const [validHost, setValidHost] = useState(false);
  const [validKibanaHost, setValidKibanaHost] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isDisabled, setSaveDisabled] = useState(false);
  const [isTestDisabled, setTestDisable] = useState(false);

  const [validMonitoringHost, setValidMonitoringValidHost] = useState(false);
  const [connectionSettings, setConnectionSettings] = useState({
    prod: {
      elasticsearch: {
        host: '',
        authentication_enabled: false,
        username: '',
        password: '',
        status: '',
      },
      kibana: {
        host: '',
        authentication_enabled: false,
        username: '',
        password: '',
        status: '',
      },
    },
    mon: {
      elasticsearch: {
        host: '',
        authentication_enabled: false,
        username: '',
        password: '',
        status: '',
      },
    },
  } as ConnectionSettings);

  const [cluster, setCluster] = useState({} as Cluster);

  const [_, __] = useState([] as Datasource[]);
  const [___, ____] = useState([] as GrafanaDatasource[]);

  const [isOpanAddDialog, setIsOpenAddDialog] = useState(false);

  const initialLogstash: Logstash = {
    id: uuidv4(),
    serverAddress: '',
    logstashApiHost: '',
    logstashLogsFolder: '',
  };

  const [logstash, setLogstash] = useState(initialLogstash);

  const [isShowLogstash, setIsShowLogstash] = useState(false);

  const [logstashList, setLogstashList] = useState([] as Logstash[]);

  const [version, setVersion] = useState('');

  const onSave = () => {
    setIsLoading(true);
    try {
      const formToSave = {
        ...cluster,
        cluster_connection_settings: getConnectionSettings(),
      };

      if (
        ((formToSave.cluster_connection_settings.prod.elasticsearch.status === 'GREEN' ||
          formToSave.cluster_connection_settings.prod.elasticsearch.status === 'YELLOW') &&
          formToSave.cluster_connection_settings.mon.elasticsearch.status === 'GREEN') ||
        formToSave.cluster_connection_settings.mon.elasticsearch.status === 'YELLOW'
      ) {
        const promise2 = backendSrv.post(`${baseUrl}/add_cluster`, JSON.stringify(formToSave), {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        });
        const promise1 = backendSrv.get(`/api/datasources`, {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        });
        Promise.all([promise1, promise2]).then(async (values) => {
          setIsLoading(true);
          const [value1, value2] = values;
          const dataSourcesFromResponse: Datasource[] = Object.values(value2);
          let isErrorOccurred = false; // Flag to track if an error has occurred
          for (const item of dataSourcesFromResponse) {
            if (!isErrorOccurred) {
              try {
                const isEqualDataSource = value1.find((item1: Datasource) => item1.name === item.name);
                if (isEqualDataSource) {
                  await backendSrv.put(`/api/datasources/uid/${item.uid}`, JSON.stringify(item));
                } else {
                  await backendSrv.post('/api/datasources', JSON.stringify(item));
                }
              } catch (error: any) {
                isErrorOccurred = true; // Set the flag to true upon encountering an error
                toast.error(`${error.message}`, {
                  position: toast.POSITION.BOTTOM_RIGHT,
                  autoClose: false,
                  closeButton: true,
                  hideProgressBar: true,
                  draggable: false,
                });
                setIsLoading(false);
                break;
              }
            } else {
              setIsLoading(false);
              break;
            }
          }
          if (!isErrorOccurred) {
            toast.success('Source connections was successfully saved!', {
              position: toast.POSITION.BOTTOM_RIGHT,
              autoClose: false,
              closeButton: true,
              hideProgressBar: true,
              draggable: false,
            });
            setIsLoading(false);
          }
        });
      }
    } catch (err: any) {
      setIsLoading(false);
    } finally {
      setIsLoading(false);
    }
    setIsLoading(false);
    // isSpinnerLoading = false
  };

  const onTest = async () => {
    setIsLoading(true); // Start the loader immediately.

    try {
      const formToTest: ConnectionSettings = getConnectionSettings();

      const result: BackendResponse = await backendSrv.post(
        `${baseUrl}/test_cluster`,
        JSON.stringify(formToTest),
        {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        }
      );

      const { status: prodStatus } = result.prod.elasticsearch;
      const { status: monStatus } = result.mon.elasticsearch;

      // Determine if the configuration is valid.
      if (
        (prodStatus.toLowerCase() === 'green' || prodStatus.toLowerCase() === 'yellow') &&
        (monStatus.toLowerCase() === 'green' || monStatus.toLowerCase() === 'yellow')
      ) {
        setSaveDisabled(false);
      } else {
        setSaveDisabled(true);
      }

      // Update connection settings state.
      setConnectionSettings({
        prod: {
          elasticsearch: {
            ...connectionSettings.prod.elasticsearch,
            status: result.prod.elasticsearch.status.toUpperCase(),
          },
          kibana: {
            ...connectionSettings.prod.kibana,
            status: result.prod.kibana.status.toUpperCase(),
          },
        },
        mon: {
          elasticsearch: {
            ...connectionSettings.mon.elasticsearch,
            status: result.mon.elasticsearch.status.toUpperCase(),
          },
        },
      });

      // Show errors if present.
      if (result.prod.elasticsearch.error) {
        showError(result.prod.elasticsearch.error);
      }
      if (result.mon.elasticsearch.error) {
        showError(result.mon.elasticsearch.error);
      }
    } catch (error: any) {
      // Handle errors and show a toast notification.
      toast.error(`${error.message}`, {
        position: toast.POSITION.BOTTOM_RIGHT,
        autoClose: false,
        closeButton: true,
        hideProgressBar: true,
        draggable: false,
      });
    } finally {
      // Stop the loader after completion, success or failure.
      setIsLoading(false);
    }
  };

  const showError = (message: string) => {
    toast.error(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
      autoClose: false,
      closeButton: true,
      hideProgressBar: true,
      draggable: false,
    });
  };

  const onCheckAuth = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      prod: {
        ...connectionSettings.prod,
        elasticsearch: {
          ...connectionSettings.prod.elasticsearch,
          authentication_enabled: event.target.checked,
        },
      },
    });
  };

  const onUserNameInput = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      prod: {
        ...connectionSettings.prod,
        elasticsearch: {
          ...connectionSettings.prod.elasticsearch,
          username: event.target.value,
        },
      },
    });
  };

  const onUserPasswordInput = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      prod: {
        ...connectionSettings.prod,
        elasticsearch: {
          ...connectionSettings.prod.elasticsearch,
          password: event.target.value,
        },
      },
    });
  };

  const onCheckMonitoringAuth = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      mon: {
        ...connectionSettings.mon,
        elasticsearch: {
          ...connectionSettings.mon.elasticsearch,
          authentication_enabled: event.target.checked,
        },
      },
    });
  };

  const onInputMonitoringUsername = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      mon: {
        ...connectionSettings.mon,
        elasticsearch: {
          ...connectionSettings.mon.elasticsearch,
          username: event.target.value,
        },
      },
    });
  };

  const onInputMonitoringPassword = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      mon: {
        ...connectionSettings.mon,
        elasticsearch: {
          ...connectionSettings.mon.elasticsearch,
          password: event.target.value,
        },
      },
    });
  };

  const onInputHost = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      prod: {
        ...connectionSettings.prod,
        elasticsearch: {
          ...connectionSettings.prod.elasticsearch,
          host: event.target.value,
        },
      },
    });
  };

  const onInputMonitoringHost = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      mon: {
        ...connectionSettings.mon,
        elasticsearch: {
          ...connectionSettings.mon.elasticsearch,
          host: event.target.value,
        },
      },
    });
  };

  const onInputKibanaHost = (event: any) => {
    setConnectionSettings({
      ...connectionSettings,
      prod: {
        ...connectionSettings.prod,
        kibana: {
          ...connectionSettings.prod.kibana,
          host: event.target.value,
        },
      },
    });
  };

  const getConnectionSettings = () => {
    return {
      prod: {
        elasticsearch: {
          host: connectionSettings.prod.elasticsearch.host,
          authentication_enabled: connectionSettings.prod.elasticsearch.authentication_enabled,
          username: connectionSettings.prod.elasticsearch.username,
          password: connectionSettings.prod.elasticsearch.password,
          status: connectionSettings.prod.elasticsearch.status,
        },
        kibana: {
          host: connectionSettings.prod.kibana.host,
          authentication_enabled: connectionSettings.prod.elasticsearch.authentication_enabled,
          username: connectionSettings.prod.elasticsearch.username,
          password: connectionSettings.prod.elasticsearch.password,
          status: connectionSettings.prod.kibana.status,
        },
      },
      mon: {
        elasticsearch: {
          host: connectionSettings.mon.elasticsearch.host,
          authentication_enabled: connectionSettings.mon.elasticsearch.authentication_enabled,
          username: connectionSettings.mon.elasticsearch.username,
          password: connectionSettings.mon.elasticsearch.password,
          status: connectionSettings.mon.elasticsearch.status,
        },
      },
    };
  };

  useEffect(() => {
    fetch(`${SERVER_URL}/version`).then((response: Response) => {
      if (!response) {
        return;
      }

      response
        .json()
        .then((data) => {
          console.log('Fetched version data:', data);
          setVersion(data.version);
        })
        .catch((error) => {
          console.error('Error parsing the response JSON:', error);
        });
    });
  }, []);

  useEffect(() => {
    fetch(`${SERVER_URL}/new_cluster`).then((response) => {
      response.json().then((data: Cluster) => {
        setCluster(data);
        const { status: prodStatus } = data.cluster_connection_settings.prod.elasticsearch;

        const { status: monStatus } = data.cluster_connection_settings.mon.elasticsearch;
        if (prodStatus === 'UNTESTED' || prodStatus === 'ERROR' || monStatus === 'UNTESTED' || monStatus === 'ERROR') {
          setSaveDisabled(true);
          setTestDisable(true);
        } else {
          setSaveDisabled(false);
          setTestDisable(false);
        }
      });
    });
  }, []);

  const onLogstashFilesDeploy = async (url: string) => {
    const cluster_connection_settings = getConnectionSettings();
    const clusterToSend = {
      cluster_connection_settings,
      logstash_configurations: cluster?.logstash_configurations || {},
    };

    setIsLoading(true);

    try {
      const response = await fetch(`${SERVER_URL}/deploy_logstash_configuration`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(clusterToSend),
      });

      if (response.ok) {
        console.log(`Deployment successful. Status: ${response.status}`);
      } else {
        console.error(`Deployment failed. Status: ${response.status}`);
      }
    } catch (err) {
      console.error('There was an error during the logstash deployment:', err);
    } finally {
      setIsLoading(false);
    }
  };


  const onTemplatesDeploy = async (url: string) => {
    const cluster_connection_settings = getConnectionSettings();
    const clusterToSend = {
      cluster_connection_settings,
      monitoring_cluster_injection: cluster?.monitoring_cluster_injection || {},
    };

    setIsLoading(true); // Ensure loader starts before the fetch operation.

    try {
      const response = await fetch(`${SERVER_URL}/deploy_elasticsearch_configuration`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(clusterToSend),
      });

      if (response.ok) {
        console.log(`Templates deployment successful. Status: ${response.status}`);
      } else {
        console.error(`Templates deployment failed. Status: ${response.status}`);
      }
    } catch (err) {
      console.error('There was an error during the Elasticsearch templates deployment:', err);
    } finally {
      // Ensure loader stops after completing the fetch operation.
      setIsLoading(false);
    }
  };

  const onDownload = (url: string) => {
    const cluster_connection_settings = getConnectionSettings();

    if (url === LogstashFileType.ES_MONITORING_CONFIGURATION_FILES) {
      const clusterToSend = {
        cluster_connection_settings,
        logstash_configurations: cluster.logstash_configurations,
      };

      downloadFiles(url, clusterToSend);
    } else {
      if (logstashList.length > 0) {
        const hosts: Host[] = logstashList.map((logstash) => {
          return {
            server_address: logstash?.serverAddress || '',
            logstash_api_host: logstash?.logstashApiHost || '',
            logstash_logs_folder: logstash?.logstashLogsFolder || '',
          };
        });

        cluster.logstash_configurations.logstash_monitoring_configuration_files.hosts = hosts;

        const clusterToSend = {
          cluster_connection_settings,
          logstash_configurations: cluster.logstash_configurations,
        };
        downloadFiles(url, clusterToSend);
      }
    }
  };

  const downloadFiles = (url: string, clusterToSend: any) => {
    fetch(`${SERVER_URL}/${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(clusterToSend),
    }).then((response) => {
      if (response) {
        const filenameHeader = response.headers.get('Content-Disposition') || 'logstash.zip';
        const filename = filenameHeader.includes('filename=') ? filenameHeader.split('filename=')[1] : 'logstash.zip';
        const formattedFileName = filename.replace(/['"]+/g, '');
        response.blob().then((blob: Blob) => {
          saveAs(blob, formattedFileName);
        });
      }
    });
  };

  const onOpenAddDialog = () => {
    setIsOpenAddDialog(true);
    setLogstash(initialLogstash);
  };

  const onSaveAddedLogstash = (item: Logstash) => {
    setIsOpenAddDialog(false);
    setIsShowLogstash(true);
    const existingIndex = logstashList.findIndex((l) => l.id === item.id);
    if (existingIndex > -1) {

      const updatedList = [...logstashList.slice(0, existingIndex), item, ...logstashList.slice(existingIndex + 1)];
      setLogstashList(updatedList);
    } else {
      setLogstashList([...logstashList, item]);
    }
  };

  const onCancel = () => {
    setIsOpenAddDialog(false);
  };

  const handleCardClick = (logstash: Logstash) => {
    setIsOpenAddDialog(true);
    setLogstash(logstash);
  };

  useEffect(() => {
    const regex = new RegExp('(http|https):\\/\\/((\\w|-|\\d|_|\\.)+)\\:\\d{2,5}');
    const updatedState = connectionSettings;

    const isProdElasticHostValid = regex.test(updatedState.prod.elasticsearch.host);
    const isKibanaHostValid = regex.test(updatedState.prod.kibana.host);
    const isMonitoringElasticHostValid = regex.test(updatedState.mon.elasticsearch.host);

    setValidHost(isProdElasticHostValid);
    setValidKibanaHost(isKibanaHostValid);
    setValidMonitoringValidHost(isMonitoringElasticHostValid);

    const isProdAuthenticationEnabled = updatedState.prod.elasticsearch.authentication_enabled;
    const isProdUsernameValid = updatedState.prod.elasticsearch.username;
    const isProdPasswordValid = updatedState.prod.elasticsearch.password;

    const isMonAuthenticationEnabled = updatedState.mon.elasticsearch.authentication_enabled;
    const isMonUsernameValid = updatedState.mon.elasticsearch.username;
    const isMonPasswordValid = updatedState.mon.elasticsearch.password;

    const isProdElasticHostExists = updatedState.prod.elasticsearch.host;
    const isMonElasticHostExists = updatedState.mon.elasticsearch.host;

    const isProdAuthValid =
      !isProdAuthenticationEnabled || (isProdAuthenticationEnabled && isProdUsernameValid && isProdPasswordValid);

    const isMonAuthValid =
      !isMonAuthenticationEnabled || (isMonAuthenticationEnabled && isMonUsernameValid && isMonPasswordValid);

    if (
      isProdElasticHostExists &&
      isMonElasticHostExists &&
      isProdAuthValid &&
      isMonAuthValid &&
      isProdElasticHostValid &&
      isMonitoringElasticHostValid
    ) {
      setTestDisable(false);
    } else {
      setTestDisable(true);
    }
  }, [connectionSettings]);

  return (



    <section className="connectionsAndConfig">

      {isLoading && (
        <div className="spinner_overlay">
          <CircularProgress color="primary" />
        </div>
      )}
      <div
        className={classNames({
          source_panel: true,
          isLight: theme.isLight,
        })}
      >


        <section>
          <div className="source-connection-group">
            <h3 className="title">Source connection</h3>
            <div className="host_wrapper form-wrapper">
              <TextField
                id="standard-basic-1"
                label="Elasticsearch Host"
                variant="standard"
                value={connectionSettings.prod.elasticsearch.host ?? ''}
                onChange={onInputHost}
              />
              {!validHost && connectionSettings.prod.elasticsearch.host && (
                <span className="invalid">Host format is invalid</span>
              )}

              <div className="status">
                <span
                  className={
                    connectionSettings.prod.elasticsearch.status
                      ? connectionSettings.prod.elasticsearch.status
                      : 'UNTESTED'
                  }
                >
                  {connectionSettings.prod.elasticsearch.status
                    ? connectionSettings.prod.elasticsearch.status
                    : 'UNTESTED'}
                </span>
              </div>
            </div>
            <div className="host_wrapper">
              <TextField
                id="standard-basic-2"
                label="Kibana host"
                variant="standard"
                value={connectionSettings.prod.kibana.host ?? ''}
                onChange={onInputKibanaHost}
              />
              {!validKibanaHost && connectionSettings.prod.kibana.host && (
                <span className="invalid">Host format is invalid</span>
              )}
            </div>

            <FormControlLabel
              value="top"
              control={
                <Checkbox
                  id="checkbox1"
                  onChange={onCheckAuth}
                  checked={connectionSettings.prod.elasticsearch.authentication_enabled}
                />
              }
              label="Use authentication"
            />
            <div className="auth_wrapper">
              <TextField
                id="standard-basic-3 text_1"
                key="text_1"
                label="Username"
                variant="standard"
                value={connectionSettings.prod.elasticsearch.username}
                onChange={onUserNameInput}
                disabled={!connectionSettings.prod.elasticsearch.authentication_enabled}
              />
              <TextField
                id="standard-basic-4 text_2"
                type="password"
                key="text_2"
                label="Password"
                variant="standard"
                value={connectionSettings.prod.elasticsearch.password}
                onChange={onUserPasswordInput}
                disabled={!connectionSettings.prod.elasticsearch.authentication_enabled}
              />
            </div>
          </div>
          <Divider className="add-new-cluster-divider"></Divider>
          <div >
            <h3 className="title monitoring-cluster-title">Monitoring cluster</h3>

            <div className="monitoring-cluster-group">
              <div className="host_wrapper">
                <TextField
                  id="standard-basic-5"
                  label="Monitoring Host"
                  variant="standard"
                  value={connectionSettings.mon.elasticsearch.host}
                  onChange={onInputMonitoringHost}
                />
                {!validMonitoringHost && connectionSettings.mon.elasticsearch.host && (
                  <span>Monitoring host format is invalid</span>
                )}

                <div className="status">
                <span
                  className={
                    connectionSettings.mon.elasticsearch.status
                      ? connectionSettings.mon.elasticsearch.status
                      : 'UNTESTED'
                  }
                >
                  {connectionSettings.mon.elasticsearch.status
                    ? connectionSettings.mon.elasticsearch.status
                    : 'UNTESTED'}
                </span>
                </div>
              </div>

              <FormControlLabel
                value="top"
                control={
                  <Checkbox
                    id="checkbox2"
                    checked={connectionSettings.mon.elasticsearch.authentication_enabled}
                    onChange={onCheckMonitoringAuth}
                  />
                }
                label="Use authentication"
              />
              <div className="auth_wrapper">
                <TextField
                  id="standard-basic-6 text_1"
                  key="text_1"
                  label="Username"
                  variant="standard"
                  value={connectionSettings.mon.elasticsearch.username}
                  onChange={onInputMonitoringUsername}
                  disabled={!connectionSettings.mon.elasticsearch.authentication_enabled}
                />
                <TextField
                  id="standard-basic-7 text_2"
                  type="password"
                  key="text_2"
                  label="Password"
                  variant="standard"
                  value={connectionSettings.mon.elasticsearch.password}
                  onChange={onInputMonitoringPassword}
                  disabled={!connectionSettings.mon.elasticsearch.authentication_enabled}
                />
              </div>
            </div>

            <div className="actions">
              <button onClick={() => onTest()} className="btn_test" disabled={isTestDisabled}>
                Test
              </button>
              <button onClick={() => onSave()} className="btn_save" disabled={isDisabled}>
                Add
              </button>
            </div>
          </div>

          {isLoading && (
            <div className="spinner_overlay">
              <CircularProgress color="primary" />
            </div>
          )}
        </section>
        <ToastContainer autoClose={3000} position="bottom-right" />
      </div>
      <div className={isShowLogstash ? 'config not-rounded' : 'config'}>
        <div className="wrapper">
          <div className="section-header">
            <h3 className="title">Monitoring Cluster Injections</h3>
            <Tooltip title="This section displays the configuration files for monitoring cluster injections" arrow>
              <HelpOutlineIcon fontSize="small" />
            </Tooltip></div>
          {cluster && cluster.monitoring_cluster_injection && (
            <MonitoringClusterInjectionPanel monitoringClusterInjections={cluster.monitoring_cluster_injection} />
          )}
          <div className="actions">
            <button
              disabled={isDisabled}
              onClick={() => onTemplatesDeploy(LogstashFileType.ES_MONITORING_CONFIGURATION_FILES)}
              className="btn_save deploy-btn"
            >
              Deploy
            </button>
          </div>
          <Divider className="add-new-cluster-divider-2"></Divider>

          <div className="section-header"><h3 className="title">Production cluster configuration</h3>
            <Tooltip title="This section displays the configuration files for Elasticsearch monitoring configuration"
                     arrow>
              <HelpOutlineIcon fontSize="small" />
            </Tooltip></div>

          {cluster && cluster.logstash_configurations && (
            <LogstashConfigurationsPanel files={cluster.logstash_configurations.es_monitoring_configuration_files} />
          )}
          <div className="actions">
            {version.includes('Container') ? (
              <button
                disabled={isDisabled}
                onClick={() => onLogstashFilesDeploy(LogstashFileType.ES_MONITORING_CONFIGURATION_FILES)}
                className="btn_save deploy-btn"
              >
                Deploy
              </button>
            ) : (
              <button
                disabled={isDisabled}
                onClick={() => onDownload(LogstashFileType.ES_MONITORING_CONFIGURATION_FILES)}
                className="btn_save deploy-btn"
              >
                Download
              </button>
            )}
          </div>
        </div>
        <Divider className="add-new-cluster-divider-2"></Divider>

        <div className="section-header"><h3 className="title">Logstash monitoring </h3>
          <Tooltip title="This section displays the configuration files for Logstash monitoring configuration" arrow>
            <HelpOutlineIcon fontSize="small" />
          </Tooltip></div>

        <div className="wrapper">
          {cluster.logstash_configurations &&
            cluster.logstash_configurations.logstash_monitoring_configuration_files.configurations && (
              <div className="hide">
                <LogstashConfigurationsPanel
                  files={cluster.logstash_configurations.logstash_monitoring_configuration_files.configurations}
                />
              </div>
            )}

          <div className="actions">
            <button className="btn_save" onClick={onOpenAddDialog}>
              Add logstash
            </button>
            <Dialog
              className="source_panel"
              open={isOpanAddDialog}
              aria-labelledby="alert-dialog-title"
              aria-describedby="alert-dialog-description"
            >
              <DialogTitle id="alert-dialog-title">{'Logstash configurations'}</DialogTitle>
              <div className="header-actions"></div>
              <DialogContent>
                <LogstashComponent
                  logstash={logstash}
                  onCancel={onCancel}
                  onSave={(logstash) => onSaveAddedLogstash(logstash)}
                />
              </DialogContent>
              <DialogActions></DialogActions>
            </Dialog>
            {isDisabled}
            <button
              disabled={isDisabled || logstashList.length === 0}
              onClick={() => onDownload(LogstashFileType.LOGSTASH_MONITORING_CONFIGURATION_FILES)}
              className="btn_save"
            >
              Download
            </button>
          </div>
        </div>
      </div>

      {isShowLogstash && (
        <div className="logstash-list">
          {' '}
          <h3 className="title">Logstash connections</h3>{' '}
          <div className="cards-wrapper">
            {logstashList.map((logstash, index) => (
              <div className="logstash-card" key={index} onClick={() => handleCardClick(logstash)}>
                <div className="serverAddress form-group">
                  <div className="logstash-label">Server address</div>
                  <span className="value" key={index}>
                    {logstash.serverAddress}
                  </span>
                </div>
                <div className="logstashApiHost form-group">
                  <div className="logstash-label">Logstash Api Host</div>
                  <span className="value" key={index}>
                    {logstash.logstashApiHost}
                  </span>
                </div>
                <div className="logstashFolder form-group">
                  <div className="logstash-label">Logstash Logs Folder</div>
                  <span className="value" key={index}>
                    {logstash.logstashLogsFolder}
                  </span>
                </div>
              </div>
            ))}{' '}
          </div>
        </div>
      )}
    </section>
  );
};
