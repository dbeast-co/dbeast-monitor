import React, { useEffect, useState } from 'react';
import './AddNewClusterPanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { CircularProgress, Dialog, DialogActions, DialogContent, DialogTitle, Divider } from '@mui/material';

import Button from '@mui/material/Button';
import { useTheme2 } from '@grafana/ui';
import classNames from 'classnames';
import { getBackendSrv } from '@grafana/runtime';
import { ConnectionSettings } from '../models/connection-settings';
import { Datasource } from '../models/datasource';
import { GrafanaDatasource } from '../models/grafana-datasource';
import { BackendResponse } from '../models/backend-response';
import { SERVER_URL } from '../config';
import { Cluster } from '../models/cluster';
import { saveAs } from 'file-saver';

import { DataGrid, GridAddIcon, GridColDef, GridRowsProp } from '@mui/x-data-grid';
import { LogstashConfigurationsPanel } from '../LogstashConfigurationsPanel';

const settings = require('../config.ts');

export enum LogstashFileType {
  ES_MONITORING_CONFIGURATION_FILES = 'download_es_monitoring_configuration_files',
  LOGSTASH_MONITORING_CONFIGURATION_FILES = 'download_logstash_monitoring_configuration_files',
}

export const AddNewClusterPanel = () => {
  const backendSrv = getBackendSrv();
  const theme = useTheme2();
  const baseUrl = settings.SERVER_URL;
  const hostRegex = new RegExp('(http|https):\\/\\/((\\w|-|\\d|_|\\.)+)\\:\\d{2,5}');

  const [validHost, setValidHost] = useState(false);
  const [validKibanaHost, setValidKibanaHost] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isDisabled, setIsDisabled] = useState(false);
  const [testDisable, setTestDisable] = useState(true);

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

  const onSave = () => {
    setIsLoading(true);
    try {
      const formToSave: ConnectionSettings = getCluster();
      if (
        ((formToSave.prod.elasticsearch.status === 'GREEN' || formToSave.prod.elasticsearch.status === 'YELLOW') &&
          formToSave.mon.elasticsearch.status === 'GREEN') ||
        formToSave.mon.elasticsearch.status === 'YELLOW'
      ) {
        const promise2 = backendSrv.post(`${baseUrl}/save`, JSON.stringify(formToSave), {
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
          setIsLoading(false);
          const [value1, value2] = values;
          const dataSourcesFromResponse: Datasource[] = Object.values(value2);
          let isErrorOccurred = false; // Flag to track if an error has occurred
          for (const item of dataSourcesFromResponse) {
            if (!isErrorOccurred) {
              try {
                const isEqualDataSource = value1.find((item1: Datasource) => item1.uid === item.uid);
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
                break; // Exit from the loop on the first error
              }
            } else {
              break; // Exit from the loop if an error has already occurred
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
          }
        });
      }
    } catch (err: any) {
      setIsLoading(false);
    } finally {
      setIsLoading(false);
    }
  };
  const onTest = () => {
    setIsLoading(true);
    try {
      const formToTest: ConnectionSettings = getCluster();
      backendSrv
        .post(`${baseUrl}/test_cluster`, JSON.stringify(formToTest), {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        })

        .then((result: BackendResponse) => {
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

          if (result.prod.elasticsearch.error) {
            toast.error(`${result.prod.elasticsearch.error}`, {
              position: toast.POSITION.BOTTOM_RIGHT,
              autoClose: false,
              closeButton: true,
              hideProgressBar: true,
              draggable: false,
            });
          }
          setIsDisabled(false);
          setIsLoading(false);
        })
        .catch((error: Error) => {
          setIsLoading(false);
          toast.error(`${error.message}`, {
            position: toast.POSITION.BOTTOM_RIGHT,
            autoClose: false,
            closeButton: true,
            hideProgressBar: true,
            draggable: false,
          });
        });
    } catch (error: any) {
      setIsLoading(false);
    } finally {
      setIsLoading(false);
    }
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
    if (isNotEmpty(event.target.value)) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
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
    if (isNotEmpty(event.target.value)) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
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
    if (isNotEmpty(event.target.value)) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
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
    if (isNotEmpty(event.target.value)) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
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
    if (isNotEmpty(event.target.value)) {
      setValidHost(false);
      setTestDisable(true);
    } else {
      setTestDisable(false);
    }
    if (hostRegex.test(event.target.value)) {
      setValidHost(true);
    } else {
      setValidHost(false);
      setTestDisable(true);
    }
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
    if (event.target.value === '') {
      setValidMonitoringValidHost(false);
    }
    if (hostRegex.test(event.target.value)) {
      setValidMonitoringValidHost(true);
    } else {
      setValidMonitoringValidHost(false);
    }
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
    if (event.target.value === '') {
      setValidKibanaHost(false);
      setTestDisable(true);
    } else {
      setTestDisable(false);
    }
    if (hostRegex.test(event.target.value)) {
      setValidKibanaHost(true);
    } else {
      setValidKibanaHost(false);
      setTestDisable(true);
    }
  };

  const getCluster = () => {
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
  const isNotEmpty = (value: string) => {
    return value !== '';
  };

  useEffect(() => {
    setIsDisabled(true);
    setValidHost(false);
  }, []);

  useEffect(() => {
    fetch(`${SERVER_URL}/new_cluster`).then((response) => {
      // fetch(`${MOCK_URL}/project`).then((response) => {
      response.json().then((data: Cluster) => {
        setCluster(data);
      });
    });
  }, []);

  const onDownload = (url: string) => {
    fetch(`${SERVER_URL}/${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ cluster }),
    }).then((response) => {
      const filenameHeader = response.headers.get('Content-Disposition') || 'logstash.zip';
      const filename = filenameHeader.includes('filename=') ? filenameHeader.split('filename=')[1] : 'logstash.zip';
      const formattedFileName = filename.replace(/['"]+/g, '');
      response.blob().then((blob) => {
        saveAs(blob, formattedFileName);
      });
    });
  };

  const onOpenAddDialog = () => {
    setIsOpenAddDialog(true);
  };

  const onSaveAddedLogstash = () => {
    setIsOpenAddDialog(false);
  };

  const onCancel = () => {
    setIsOpenAddDialog(false);
  };

  const columns: GridColDef[] = [
    {
      field: 'server_address',
      headerName: 'Server address',
      editable: true,
      flex: 1,
    },
    {
      field: 'logstash_api_host',
      headerName: 'Logstash Api Host',
      editable: true,
      flex: 1,
    },
    {
      field: 'logstash_config_folder',
      headerName: 'Logstash Config Folder',
      editable: true,
      flex: 1,
    },
    {
      field: 'logstash_logs_folder',
      headerName: 'Logstash Logs Folder',
      editable: true,
      flex: 1,
    },
  ];

  const initialRows: GridRowsProp = [];
  const [rows, setRows] = React.useState(initialRows);

  // const [rowModesModel, setRowModesModel] = React.useState<GridRowModesModel>({});

  const onAddRecord = () => {
    setRows((oldRows) => [
      ...oldRows,
      {
        id: oldRows.length + 1,
        server_address: '',
        logstash_api_host: 'localhost:9600',
        logstash_config_folder: '/etc/logstash/conf.d',
        logstash_logs_folder: 'var/log/logstash',
      },
    ]);
  };

  return (
    <section className="connectionsAndConfig">
      <div
        className={classNames({
          source_panel: true,
          isLight: theme.isLight,
        })}
      >
        <section>
          <div>
            <div className="host_wrapper">
              <TextField
                id="standard-basic"
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
                id="standard-basic"
                label="Kibana host"
                variant="standard"
                value={connectionSettings.prod.kibana.host ?? ''}
                onChange={onInputKibanaHost}
              />
              {!validKibanaHost && connectionSettings.prod.kibana.host && (
                <span className="invalid">Host format is invalid</span>
              )}

              {/*<div className='status'>*/}
              {/*        <span*/}
              {/*            className={connectionSettings.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}>{newProject.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}</span>*/}
              {/*</div>*/}
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
                id="standard-basic text_1"
                key="text_1"
                label="Username"
                variant="standard"
                value={connectionSettings.prod.elasticsearch.username}
                onChange={onUserNameInput}
                disabled={!connectionSettings.prod.elasticsearch.authentication_enabled}
              />
              <TextField
                id="standard-basic text_2"
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

          <div>
            <div className="host_wrapper">
              <TextField
                id="standard-basic"
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
                id="standard-basic text_1"
                key="text_1"
                label="Username"
                variant="standard"
                value={connectionSettings.mon.elasticsearch.username}
                onChange={onInputMonitoringUsername}
                disabled={!connectionSettings.mon.elasticsearch.authentication_enabled}
              />
              <TextField
                id="standard-basic text_2"
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
            <button onClick={() => onTest()} className="btn_test" disabled={testDisable}>
              {' '}
              Test
            </button>
            <button onClick={() => onSave()} className="btn_save" disabled={testDisable || isDisabled}>
              Save
            </button>
          </div>

          {isLoading && (
            <div className="spinner_overlay">
              <CircularProgress color="primary" />
            </div>
          )}
        </section>

        <ToastContainer autoClose={3000} position="bottom-right" />
      </div>
      <div className="config">
        <h3 className="title">Cluster inject configuration</h3>
        <div className="wrapper">
          {cluster && cluster.logstash_configurations && (
            <LogstashConfigurationsPanel files={cluster.logstash_configurations.es_monitoring_configuration_files} />
          )}
          <div className="actions">
            <button onClick={() => onDownload(LogstashFileType.ES_MONITORING_CONFIGURATION_FILES)} className="btn_save">
              Download
            </button>
          </div>
        </div>
        <Divider></Divider>
        <h3 className="title">Logstash inject configurations</h3>
        <div className="wrapper">
          {cluster.logstash_configurations &&
            cluster.logstash_configurations.logstash_monitoring_configuration_files.configurations && (
              <LogstashConfigurationsPanel
                files={cluster.logstash_configurations.logstash_monitoring_configuration_files.configurations}
              />
            )}

          <div className="actions">
            <button className="btn_save" onClick={onOpenAddDialog}>
              Add logstash
            </button>
            yu
            <Dialog
              open={isOpanAddDialog}
              aria-labelledby="alert-dialog-title"
              aria-describedby="alert-dialog-description"
            >
              <DialogTitle id="alert-dialog-title">{'Logstash configurations'}</DialogTitle>
              <div className="header-actions">
                <Button onClick={onAddRecord} className="save_btn" startIcon={<GridAddIcon />}>
                  Add logstash
                </Button>
              </div>

              <DialogContent>
                <div style={{ height: 365, width: '100%' }}>
                  <DataGrid rows={rows} columns={columns} />
                </div>
              </DialogContent>

              <DialogActions>
                <button className="save_btn" onClick={() => onCancel()}>
                  Cancel
                </button>
                <button className="save_btn" onClick={() => onSaveAddedLogstash()} autoFocus>
                  Save
                </button>
              </DialogActions>
            </Dialog>
            <button
              onClick={() => onDownload(LogstashFileType.LOGSTASH_MONITORING_CONFIGURATION_FILES)}
              className="btn_save"
            >
              Download
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};