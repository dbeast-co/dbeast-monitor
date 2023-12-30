import React, {useEffect, useState} from 'react';
import {PanelProps} from '@grafana/data';
import {SimpleOptions} from './types';
import './SimplePanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import {toast, ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {CircularProgress} from '@mui/material';
import {useTheme2} from '@grafana/ui';
import classNames from 'classnames';

const settings = require('./config.ts');


interface Response {
    cluster_status: string;
    error: string;
}

// interface Project {
//     es_host: string;
//     kibana_host: string;
//     authentication_enabled: boolean;
//     username: string | null;
//     password: string | null;
//
//     monitoring_host: string;
//     monitoring_authentication_enabled: boolean;
//     monitoring_username: string;
//     monitoring_password: string;
//     ssl_enabled?: boolean;
//     ssl_file?: string | null;
//     status: string;
// }

interface NewProject {
    prod: Prod;
    mon: Mon
}

interface Prod {
    elasticsearch: Project,
    kibana: Project
}

interface Mon {
    elasticsearch: Project,
}

interface Project {
    host: string;
    authentication_enabled: boolean;
    username: string | null;
    password: string | null;
    status: string;
}


interface Props extends PanelProps<SimpleOptions> {
}

export const SimplePanel: React.FC<Props> = ({options, data, width, height, replaceVariables}) => {
    console.log('INIT SimplePanel in panel')
    const theme = useTheme2();

    console.log('theme', theme)

    const baseUrl = settings.SERVER_URL;

    const hostRegex = new RegExp('(http|https):\\/\\/((\\w|-|\\d|_|\\.)+)\\:\\d{2,5}');
    // const [disableControl, setDisableControl] = useState(true);
    // const [authChecked, setAuthChecked] = useState(false);
    // const [host, setHost] = useState('');
    // const [kibanaHost, setKibanaHost] = useState('');
    // const [sslChecked, setSSLChecked] = useState(false);
    // const [authUsername, setAuthUsername] = useState('');
    // const [authPassword, setAuthPassword] = useState('');
    const [validHost, setValidHost] = useState(false);
    const [validKibanaHost, setValidKibanaHost] = useState(false);
    // const [form, setForm] = useState({});
    // const [status, setStatus] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [isDisabled, setIsDisabled] = useState(false);
    const [testDisable, setTestDisable] = useState(true);

    // const [monitoring_host, setMonitoringHost] = useState('');
    const [validMonitoringHost, setValidMonitoringValidHost] = useState(false);
    // const [monitoring_is_use_authentication, setMonitoringAuthChecked] = useState(false);
    // const [monitoringAuthUsername, setMonitoringAuthUsername] = useState('');
    // const [monitoringAuthPassword, setMonitoringAuthPassword] = useState('');
    // const [disabledMonitoringControl, setDisableMonitoringControl] = useState(true);

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

    let form: NewProject = {} as NewProject;
    const onSave = () => {
        setIsLoading(true);
        try {
            // const form: Project = {
            //     es_host: host ?? '',
            //     kibana_host: kibanaHost ?? '',
            //     authentication_enabled: authChecked ?? false,
            //     // ssl_enabled: sslChecked ?? false,
            //     username: authUsername ?? '',
            //     password: authPassword ?? '',
            //     status: status ?? 'UNTESTED',
            //     monitoring_host: monitoring_host ?? '',
            //     monitoring_authentication_enabled: monitoring_is_use_authentication ?? false,
            //     monitoring_username: monitoringAuthUsername ?? '',
            //     monitoring_password: monitoringAuthPassword ?? '',
            // };
            fetch(`${baseUrl}/save`, {
                method: 'POST',
                // body: JSON.stringify(form),
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            })
                .then((response) => {
                    setIsLoading(false);
                    toast.success('Source connection was successfully saved!', {
                        position: toast.POSITION.BOTTOM_RIGHT,
                        autoClose: false,
                        closeButton: true,
                        hideProgressBar: true,
                        draggable: false,
                    });
                })
                .catch((error) => {
                    setIsLoading(false);
                    toast.error(`${error.message}`, {
                        position: toast.POSITION.BOTTOM_RIGHT,
                        autoClose: false,
                        closeButton: true,
                        hideProgressBar: true,
                        draggable: false,
                    });
                });
        } catch (err: any) {
            setIsLoading(false);
        } finally {
            setIsLoading(false);
        }
    };
    const onTest = () => {
        setIsLoading(true);
        try {
            // const form: Project = {
            //     es_host: host ?? '',
            //     kibana_host: kibanaHost ?? '',
            //     authentication_enabled: authChecked ?? false,
            //     // ssl_enabled: sslChecked ?? false,
            //     username: authUsername ?? '',
            //     password: authPassword ?? '',
            //
            //     monitoring_host: monitoring_host ?? '',
            //     monitoring_authentication_enabled: monitoring_is_use_authentication ?? false,
            //     monitoring_username: monitoringAuthUsername ?? '',
            //     monitoring_password: monitoringAuthPassword ?? '',
            //     // ssl_file: content,
            //     status: status ?? 'UNTESTED',
            // };

            form = {
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
            console.log('aaa', form);


            fetch(`${baseUrl}/test_cluster`, {
                method: 'POST',
                body: JSON.stringify(form),
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            })
                .then((response) => response.json())
                .then((result: Response) => {
                    console.log(form)
                    if (result.error) {
                        toast.error(`${result.error}`, {
                            position: toast.POSITION.BOTTOM_RIGHT,
                            autoClose: false,
                            closeButton: true,
                            hideProgressBar: true,
                            draggable: false,
                        });
                    }
                    // setStatus(result.cluster_status);
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
        // disableControl ? setAuthChecked(true) : setAuthChecked(false);
        // disableControl ? setDisableControl(false) : setDisableControl(true);
    };

    const onUserNameInput = (event: any) => {
        // setAuthUsername(event?.target?.value);
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
        if (event.target.value === '') {
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
        if (event.target.value === '') {
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


        //     // disabledMonitoringControl ? setMonitoringAuthChecked(true) : setMonitoringAuthChecked(false);
        //     // disabledMonitoringControl ? setDisableMonitoringControl(false) : setDisableMonitoringControl(true);
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
        if (event.target.value === '') {
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
        if (event.target.value === '') {
            setIsDisabled(true);
        } else {
            setIsDisabled(false);
        }
    }
    const onInputHost = (event: any) => {
        // setHost(event?.target?.value);
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
        //
        //
        if (event.target.value === '') {
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
        // setMonitoringHost(event?.target?.value);
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
        //
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
        // setKibanaHost(event?.target?.value);
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
        //
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


    useEffect(() => {
        // setHost('');
        // setKibanaHost('');
        // setIsDisabled(true);
        // setValidHost(false);
        // setStatus('UNTESTED');
    }, []);

    // @ts-ignore
    return (

        <div  className={classNames({
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
                    {!validHost && newProject.prod.elasticsearch.host && <span className='error'>Host format is invalid</span>}

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
                    {!validKibanaHost && newProject.prod.kibana.host && <span>Host format is invalid</span>}

                    <div className='status'>
                            <span
                                className={newProject.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}>{newProject.prod.kibana.status ? newProject.prod.kibana.status : 'UNTESTED'}</span>
                    </div>

                </div>

                <FormControlLabel
                    value="top"
                    control={<Checkbox onChange={onCheckAuth}
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
                    control={<Checkbox
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
                {/*<span className={status ? status : 'UNTESTED'}>{status ? status : 'UNTESTED'}</span>*/}

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
