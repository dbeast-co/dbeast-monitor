import React, {useEffect, useState} from 'react';
import './AddNewClusterPanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import {toast, ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {CircularProgress} from '@mui/material';
import {useTheme2} from '@grafana/ui';
import classNames from 'classnames';
import {getBackendSrv} from "@grafana/runtime";
import {NewProject} from './models/new-project';
import {Datasource} from './models/datasource';
import {GrafanaDatasource} from './models/grafana-datasource';
import {BackendResponse} from './models/backend-response';

const settings = require('./config.ts');


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
    const [newProject, setNewProject] = useState({
        prod: {
            elasticsearch: {
                host: '',
                authentication_enabled: false,
                username: '',
                password: '',
                status: ''
            },
            kibana: {
                host: '',
                authentication_enabled: false,
                username: '',
                password: '',
                status: ''
            }
        },
        mon: {
            elasticsearch: {
                host: '',
                authentication_enabled: false,
                username: '',
                password: '',
                status: ''
            }
        }
    } as NewProject);

    const [_, __] = useState([] as Datasource[]);
    const [___, ____] = useState([] as GrafanaDatasource[]);


    const onSave = () => {
        setIsLoading(true);
        try {
            const formToSave: NewProject = getNewProject();
            if ((formToSave.prod.elasticsearch.status === 'GREEN' || formToSave.prod.elasticsearch.status === 'YELLOW') && formToSave.mon.elasticsearch.status === 'GREEN' || (formToSave.mon.elasticsearch.status === 'YELLOW')) {
                const promise2 = backendSrv.post(`${baseUrl}/save`,
                    JSON.stringify(formToSave),
                    {
                        headers: {
                            "Content-Type": 'application/json',
                            "Accept": 'application/json'
                        }
                    }
                );
                const promise1 = backendSrv.get(`/api/datasources`, {
                    headers: {
                        "Content-Type": 'application/json',
                        "Accept": 'application/json',
                    },
                })
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

                })
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
            const formToTest: NewProject = getNewProject()
            backendSrv.post(`${baseUrl}/test_cluster`,
                JSON.stringify(formToTest),
                {
                    headers: {
                        "Content-Type": 'application/json',
                        "Accept": 'application/json',
                    }
                }
            )

                .then((result: BackendResponse) => {
                    setNewProject({
                        prod: {
                            elasticsearch: {
                                ...newProject.prod.elasticsearch,
                                status: result.prod.elasticsearch.status.toUpperCase()
                            },
                            kibana: {
                                ...newProject.prod.kibana,
                                status: result.prod.kibana.status.toUpperCase()
                            }
                        },
                        mon: {
                            elasticsearch: {
                                ...newProject.mon.elasticsearch,
                                status: result.mon.elasticsearch.status.toUpperCase()
                            }
                        }
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
        setNewProject({
            ...newProject,
            prod: {
                ...newProject.prod,
                elasticsearch: {
                    ...newProject.prod.elasticsearch,
                    authentication_enabled: event.target.checked
                }
            }
        })
    };

    const onUserNameInput = (event: any) => {
        setNewProject({
            ...newProject,
            prod: {
                ...newProject.prod,
                elasticsearch: {
                    ...newProject.prod.elasticsearch,
                    username: event.target.value
                }
            }
        })
        if (isNotEmpty(event.target.value)) {
            setIsDisabled(true);
        } else {
            setIsDisabled(false);
        }
    }
    const onUserPasswordInput = (event: any) => {
        setNewProject({
            ...newProject,
            prod: {
                ...newProject.prod,
                elasticsearch: {
                    ...newProject.prod.elasticsearch,
                    password: event.target.value
                }
            }
        })
        if (isNotEmpty(event.target.value)) {
            setIsDisabled(true);
        } else {
            setIsDisabled(false);
        }
    }

    const onCheckMonitoringAuth = (event: any) => {
        setNewProject({
            ...newProject,
            mon: {
                ...newProject.mon,
                elasticsearch: {
                    ...newProject.mon.elasticsearch,
                    authentication_enabled: event.target.checked
                }
            }
        })


    };
    const onInputMonitoringUsername = (event: any) => {
        setNewProject({
            ...newProject,
            mon: {
                ...newProject.mon,
                elasticsearch: {
                    ...newProject.mon.elasticsearch,
                    username: event.target.value
                }
            }
        })
        if (isNotEmpty(event.target.value)) {
            setIsDisabled(true);
        } else {
            setIsDisabled(false);
        }
    }

    const onInputMonitoringPassword = (event: any) => {
        setNewProject({
            ...newProject,
            mon: {
                ...newProject.mon,
                elasticsearch: {
                    ...newProject.mon.elasticsearch,
                    password: event.target.value
                }
            }
        })
        if (isNotEmpty(event.target.value)) {
            setIsDisabled(true);
        } else {
            setIsDisabled(false);
        }
    }
    const onInputHost = (event: any) => {
        setNewProject({
            ...newProject,
            prod: {
                ...newProject.prod,
                elasticsearch: {
                    ...newProject.prod.elasticsearch,
                    host: event.target.value
                }
            }
        })
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
        setNewProject({
            ...newProject,
            mon: {
                ...newProject.mon,
                elasticsearch: {
                    ...newProject.mon.elasticsearch,
                    host: event.target.value,

                }
            }
        })
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
        setNewProject({
            ...newProject,
            prod: {
                ...newProject.prod,
                kibana: {
                    ...newProject.prod.kibana,
                    host: event.target.value
                }
            }
        })
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

    const getNewProject = () => {
        return {
            prod: {
                elasticsearch: {
                    host: newProject.prod.elasticsearch.host,
                    authentication_enabled: newProject.prod.elasticsearch.authentication_enabled,
                    username: newProject.prod.elasticsearch.username,
                    password: newProject.prod.elasticsearch.password,
                    status: newProject.prod.elasticsearch.status
                },
                kibana: {
                    host: newProject.prod.kibana.host,
                    authentication_enabled: newProject.prod.elasticsearch.authentication_enabled,
                    username: newProject.prod.elasticsearch.username,
                    password: newProject.prod.elasticsearch.password,
                    status: newProject.prod.kibana.status
                }
            },
            mon: {
                elasticsearch: {
                    host: newProject.mon.elasticsearch.host,
                    authentication_enabled: newProject.mon.elasticsearch.authentication_enabled,
                    username: newProject.mon.elasticsearch.username,
                    password: newProject.mon.elasticsearch.password,
                    status: newProject.mon.elasticsearch.status
                }
            }
        };
    }
    const isNotEmpty = (value: string) => {
        return value !== '';
    }


    useEffect(() => {
        setIsDisabled(true);
        setValidHost(false);
    }, []);


    return (
        <div className={classNames({
            source_panel: true,
            isLight: theme.isLight,
        })}>
            <div>
                <div className="host_wrapper">
                    <TextField
                        id="standard-basic"
                        label="Elasticsearch Host"
                        variant="standard"
                        value={newProject.prod.elasticsearch.host ?? ''}
                        onChange={onInputHost}
                    />
                    {!validHost && newProject.prod.elasticsearch.host &&
                        <span className='invalid'>Host format is invalid</span>}

                    <div className='status'>
                            <span
                                className={newProject.prod.elasticsearch.status ? newProject.prod.elasticsearch.status : 'UNTESTED'}>{newProject.prod.elasticsearch.status ? newProject.prod.elasticsearch.status : 'UNTESTED'}</span>
                    </div>
                </div>
                <div className="host_wrapper">
                    <TextField
                        id="standard-basic"
                        label="Kibana host"
                        variant="standard"
                        value={newProject.prod.kibana.host ?? ''}
                        onChange={onInputKibanaHost}
                    />
                    {!validKibanaHost && newProject.prod.kibana.host &&
                        <span className="invalid">Host format is invalid</span>}

                    {/*<div className='status'>*/}
                    {/*        <span*/}
                    {/*            className={newProject.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}>{newProject.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}</span>*/}
                    {/*</div>*/}
                </div>

                <FormControlLabel
                    value="top"
                    control={<Checkbox id="checkbox1" onChange={onCheckAuth}
                                       checked={newProject.prod.elasticsearch.authentication_enabled}/>}
                    label="Use authentication"
                />
                <div className="auth_wrapper">
                    <TextField
                        id="standard-basic text_1"
                        key="text_1"
                        label="Username"
                        variant="standard"
                        value={newProject.prod.elasticsearch.username}
                        onChange={onUserNameInput}
                        disabled={!newProject.prod.elasticsearch.authentication_enabled}

                    />
                    <TextField
                        id="standard-basic text_2"
                        type="password"
                        key="text_2"
                        label="Password"
                        variant="standard"
                        value={newProject.prod.elasticsearch.password}
                        onChange={onUserPasswordInput}
                        disabled={!newProject.prod.elasticsearch.authentication_enabled}
                    />
                </div>

            </div>

            <div>
                <div className="host_wrapper">
                    <TextField
                        id="standard-basic"
                        label="Monitoring Host"
                        variant="standard"
                        value={newProject.mon.elasticsearch.host}
                        onChange={onInputMonitoringHost}
                    />
                    {!validMonitoringHost && newProject.mon.elasticsearch.host &&
                        <span>Monitoring host format is invalid</span>}

                    <div className='status'>
                            <span
                                className={newProject.mon.elasticsearch.status ? newProject.mon.elasticsearch.status : 'UNTESTED'}>{newProject.mon.elasticsearch.status ? newProject.mon.elasticsearch.status : 'UNTESTED'}</span>
                    </div>
                </div>

                <FormControlLabel
                    value="top"
                    control={<Checkbox id="checkbox2"
                                       checked={newProject.mon.elasticsearch.authentication_enabled}
                                       onChange={onCheckMonitoringAuth}/>}
                    label="Use authentication"

                />
                <div className="auth_wrapper">
                    <TextField
                        id="standard-basic text_1"
                        key="text_1"
                        label="Username"
                        variant="standard"
                        value={newProject.mon.elasticsearch.username}
                        onChange={onInputMonitoringUsername}
                        disabled={!newProject.mon.elasticsearch.authentication_enabled}

                    />
                    <TextField
                        id="standard-basic text_2"
                        type="password"
                        key="text_2"
                        label="Password"
                        variant="standard"
                        value={newProject.mon.elasticsearch.password}
                        onChange={onInputMonitoringPassword}
                        disabled={!newProject.mon.elasticsearch.authentication_enabled}
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
                    <CircularProgress color="primary"/>
                </div>
            )}

            <ToastContainer autoClose={3000} position="bottom-right"/>
        </div>
    );
};
