import React, {useEffect, useState} from 'react';
import './SimplePanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import {toast, ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {CircularProgress} from '@mui/material';


interface Response {
    cluster_status: string;
    error: string;
}

interface Project {
    es_host: string;
    authentication_enabled: boolean;
    username: string | null;
    password: string | null;

    application_host: string;
    is_replace_keystore?: boolean | true;
    ssl_enabled?: boolean;
    ssl_file?: string | null;
    status: string;
}


// interface Props extends PanelProps<SimpleOptions> {}

export const SimplePanel = (props: any) => {
    // const baseUrl = settings.SERVER_URL;
    const hostRegex = new RegExp('(http|https):\\/\\/((\\w|-|\\d|_|\\.)+)\\:\\d{2,5}');
    const [disableControl, setDisableControl] = useState(true);
    // const [disableControlIsReplaceKeystore, setDisableControlIsReplaceKeystore] = useState(true);
    const [authChecked, setAuthChecked] = useState(false);
    const [host, setHost] = useState('');
    // const [sslChecked, setSSLChecked] = useState(false);
    const [authUsername, setAuthUsername] = useState('');
    const [authPassword, setAuthPassword] = useState('');
    const [validHost, setValidHost] = useState(false);
    // const [form, setForm] = useState({});
    const [status, setStatus] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [isDisabled, setIsDisabled] = useState(false);
    const [testDisable, setTestDisable] = useState(true);
    // const [content, setContent] = useState('');

    const [application_host, setApplicationHost] = useState('');
    const [is_replace_keystore, setIsReplaceKeystore] = useState(false);
    const onSave = () => {
        let baseUrl;
        if (application_host.includes('http')) {
            baseUrl = application_host
        } else {
            baseUrl = 'http://' + application_host
        }
        baseUrl = baseUrl + '/grafana_backend/setup'
        setIsLoading(true);
        try {
            const form: Project = {
                es_host: host ?? '',
                authentication_enabled: authChecked ?? false,
                // ssl_enabled: sslChecked ?? false,
                username: authUsername ?? '',
                password: authPassword ?? '',
                status: status ?? 'UNTESTED',
                application_host: application_host ?? '',
            };
            fetch(`${baseUrl}/save`, {
                method: 'POST',
                body: JSON.stringify(form),
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
        let baseUrl;
        if (application_host.includes('http')) {
            baseUrl = application_host
        } else {
            baseUrl = 'http://' + application_host
        }
        baseUrl = baseUrl + '/grafana_backend/setup'
        setIsLoading(true);
        try {
            const form: Project = {
                es_host: host ?? '',
                authentication_enabled: authChecked ?? false,
                // ssl_enabled: sslChecked ?? false,
                username: authUsername ?? '',
                password: authPassword ?? '',

                application_host: application_host ?? '',
                is_replace_keystore: is_replace_keystore ?? true,
                status: status ?? 'UNTESTED',
            };
            fetch(`${baseUrl}/test_connection`, {
                method: 'POST',
                body: JSON.stringify(form),
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            })
                .then((response) => response.json())
                .then((result: Response) => {
                    if (result.error) {
                        toast.error(`${result.error}`, {
                            position: toast.POSITION.BOTTOM_RIGHT,
                            autoClose: false,
                            closeButton: true,
                            hideProgressBar: true,
                            draggable: false,
                        });
                    }
                    setStatus(result.cluster_status);
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
    const onCheckAuth = () => {
        disableControl ? setAuthChecked(true) : setAuthChecked(false);
        disableControl ? setDisableControl(false) : setDisableControl(true);
    };

    const onInputHost = (event: any) => {
        setHost(event?.target?.value);

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
    const onInputApplicationHost = (event: any) => {
        setApplicationHost(event?.target?.value);

    };
    const onInputIsReplaceKeystore = () => {
        is_replace_keystore ? setIsReplaceKeystore(false) : setIsReplaceKeystore(true);
        // setIsReplaceKeystore(event?.target?.value);

    };

    useEffect(() => {
        setHost('');
        setIsDisabled(true);
        setValidHost(false);
        setStatus('UNTESTED');
    }, []);

    return (
        <>
            <h3 className="title">Application configuration</h3>

            <div className="config">
                <div>
                    <div className="host_wrapper">
                        <TextField
                            id="standard-basic"
                            label="Grafana Host"
                            variant="standard"
                            value={host}
                            onChange={onInputHost}
                        />
                        {!validHost && host && <span>Host format is invalid</span>}
                    </div>

                    <FormControlLabel
                        value="top"
                        control={<Checkbox checked={authChecked} onChange={() => onCheckAuth()}/>}
                        label="Use authentication"
                    />
                    <div className="auth_wrapper">
                        <TextField
                            id="standard-basic text_1"
                            key="text_1"
                            label="Username"
                            variant="standard"
                            disabled={disableControl}
                            value={authUsername}
                            onChange={(event) => setAuthUsername(event.target.value)}
                        />
                        <TextField
                            id="standard-basic text_2"
                            type="password"
                            key="text_2"
                            label="Password"
                            variant="standard"
                            disabled={disableControl}
                            value={authPassword}
                            onChange={(event) => setAuthPassword(event.target.value)}
                        />
                    </div>

                </div>
                {/*<Divider></Divider>*/}

                <div>
                    <div className="host_wrapper">
                        <TextField
                            id="standard-basic"
                            label="Application Host"
                            variant="standard"
                            value={application_host}
                            onChange={onInputApplicationHost}
                        />
                    </div>

                </div>
                <div>
                    <FormControlLabel
                        value="top"
                        control={<Checkbox checked={is_replace_keystore} onChange={() => onInputIsReplaceKeystore()}/>}
                        label="Is replace keystore"
                    />
                </div>

                <div className="actions">
                    <button onClick={() => onTest()} className="btn_test" disabled={testDisable}>
                        {' '}
                        Test
                    </button>
                    <button onClick={() => onSave()} className="btn_save" disabled={testDisable || isDisabled}>
                        Save
                    </button>
                    <span className={status ? status : 'UNTESTED'}>{status ? status : 'UNTESTED'}</span>

                    {/*<button onClick={() => onGetList()}>Get indices</button>*/}
                    {/*<button onClick={() => onNewProject()}>New project</button>*/}
                </div>
                {isLoading && (
                    <div className="spinner_overlay">
                        <CircularProgress color="primary"/>
                    </div>
                )}

                <ToastContainer autoClose={3000} position="bottom-right"/>
            </div>
        </>
    );
};
