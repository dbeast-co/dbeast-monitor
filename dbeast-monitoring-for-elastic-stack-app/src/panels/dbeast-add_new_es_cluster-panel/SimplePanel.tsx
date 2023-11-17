import React, { useEffect, useState } from 'react';
import { PanelProps } from '@grafana/data';
import { SimpleOptions } from './types';
import './SimplePanel.scss';
import TextField from '@mui/material/TextField';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { CircularProgress } from '@mui/material';

const settings = require('./config.ts');
// const settings = require('./settings/config');
// const settings = require('./settings/settings.env');
// const config = require('./config/config.js');

interface Response {
  cluster_status: string;
  error: string;
}

interface Project {
  es_host: string;
  kibana_host: string;
  authentication_enabled: boolean;
  username: string | null;
  password: string | null;

  monitoring_host: string;
  monitoring_authentication_enabled: boolean;
  monitoring_username: string;
  monitoring_password: string;
  ssl_enabled?: boolean;
  ssl_file?: string | null;
  status: string;
}

interface Props extends PanelProps<SimpleOptions> {}

export const SimplePanel: React.FC<Props> = ({ options, data, width, height, replaceVariables }) => {
  console.log('INIT SimplePanel in panel')

  const baseUrl = settings.SERVER_URL;

  const hostRegex = new RegExp('(http|https):\\/\\/((\\w|-|\\d|_|\\.)+)\\:\\d{2,5}');
  const [disableControl, setDisableControl] = useState(true);
  const [authChecked, setAuthChecked] = useState(false);
  const [host, setHost] = useState('');
  const [kibanaHost, setKibanaHost] = useState('');
  // const [sslChecked, setSSLChecked] = useState(false);
  const [authUsername, setAuthUsername] = useState('');
  const [authPassword, setAuthPassword] = useState('');
  const [validHost, setValidHost] = useState(false);
  const [validKibanaHost, setValidKibanaHost] = useState(false);
  // const [form, setForm] = useState({});
  const [status, setStatus] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [isDisabled, setIsDisabled] = useState(false);
  const [testDisable, setTestDisable] = useState(true);
  // const [content, setContent] = useState('');

  const [monitoring_host, setMonitoringHost] = useState('');
  const [validMonitoringHost, setValidMonitoringValidHost] = useState(false);
  const [monitoring_is_use_authentication, setMonitoringAuthChecked] = useState(false);
  const [monitoringAuthUsername, setMonitoringAuthUsername] = useState('');
  const [monitoringAuthPassword, setMonitoringAuthPassword] = useState('');
  const [disabledMonitoringControl, setDisableMonitoringControl] = useState(true);
  const onSave = () => {
    setIsLoading(true);
    try {
      const form: Project = {
        es_host: host ?? '',
        kibana_host: kibanaHost ?? '',
        authentication_enabled: authChecked ?? false,
        // ssl_enabled: sslChecked ?? false,
        username: authUsername ?? '',
        password: authPassword ?? '',
        status: status ?? 'UNTESTED',
        monitoring_host: monitoring_host ?? '',
        monitoring_authentication_enabled: monitoring_is_use_authentication ?? false,
        monitoring_username: monitoringAuthUsername ?? '',
        monitoring_password: monitoringAuthPassword ?? '',
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
    setIsLoading(true);
    try {
      const form: Project = {
        es_host: host ?? '',
        kibana_host: kibanaHost ?? '',
        authentication_enabled: authChecked ?? false,
        // ssl_enabled: sslChecked ?? false,
        username: authUsername ?? '',
        password: authPassword ?? '',

        monitoring_host: monitoring_host ?? '',
        monitoring_authentication_enabled: monitoring_is_use_authentication ?? false,
        monitoring_username: monitoringAuthUsername ?? '',
        monitoring_password: monitoringAuthPassword ?? '',
        // ssl_file: content,
        status: status ?? 'UNTESTED',
      };
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

  const onCheckMonitoringAuth = (event: any) => {
    disabledMonitoringControl ? setMonitoringAuthChecked(true) : setMonitoringAuthChecked(false);
    disabledMonitoringControl ? setDisableMonitoringControl(false) : setDisableMonitoringControl(true);
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
  const onInputMonitoringHost = (event: any) => {
    setMonitoringHost(event?.target?.value);

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
    setKibanaHost(event?.target?.value);

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

  // const onCheckSSL = () => {
  //     sslChecked ? setSSLChecked(false) : setSSLChecked(true)
  // };

  // const onSetFile = (file: File) => {
  //     const reader = new FileReader();
  //     reader.readAsText(file);
  //     reader.onload = function (e) {
  //         const content = reader.result;
  //         //Here the content has been read successfuly
  //         setContent(JSON.stringify(content));
  //     }
  //
  //
  // }
  useEffect(() => {
    setHost('');
    setKibanaHost('');
    setIsDisabled(true);
    setValidHost(false);
    setStatus('UNTESTED');
  }, []);

  return (
    <div className="source_panel">
      <div>
        <div className="host_wrapper">
          <TextField
            id="standard-basic"
            label="Elasticsearch Host"
            variant="standard"
            value={host}
            onChange={onInputHost}
          />
          {!validHost && host && <span>Host format is invalid</span>}
        </div>

        <div className="host_wrapper">
          <TextField
            id="standard-basic"
            label="Kibana host"
            variant="standard"
            value={kibanaHost}
            onChange={onInputKibanaHost}
          />
          {!validKibanaHost && kibanaHost && <span>Host format is invalid</span>}
        </div>

        <FormControlLabel
          value="top"
          control={<Checkbox checked={authChecked} onChange={() => onCheckAuth()} />}
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

        {/*<div className="ssl_wrapper">*/}
        {/*    <FormControlLabel*/}
        {/*        value="top"*/}
        {/*        control={<Checkbox checked={sslChecked} onChange={() => onCheckSSL()}/>}*/}
        {/*        label="SSL"*/}
        {/*    />*/}
        {/*    <div className="upload_file_container"><span>SSL Certificate</span>*/}
        {/*        <label htmlFor="file" className="upload_file" >*/}
        {/*            <input type="file" id="file" accept="text" disabled={!sslChecked}*/}
        {/*                   onChange={(event) => onSetFile(event?.target?.files![0])}/>*/}
        {/*            <span>Browse</span>*/}
        {/*        </label></div>*/}
        {/*</div>*/}
      </div>
      {/*<Divider></Divider>*/}

      <div>
        <div className="host_wrapper">
          <TextField
            id="standard-basic"
            label="Monitoring Host"
            variant="standard"
            value={monitoring_host}
            onChange={onInputMonitoringHost}
          />
          {!validMonitoringHost && monitoring_host && <span>Monitoring host format is invalid</span>}
        </div>

        <FormControlLabel
          value="top"
          control={<Checkbox checked={monitoring_is_use_authentication} onChange={(e) => onCheckMonitoringAuth(e)} />}
          label="Use authentication"
        />
        <div className="auth_wrapper">
          <TextField
            id="standard-basic text_1"
            key="text_1"
            label="Username"
            variant="standard"
            disabled={disabledMonitoringControl}
            value={monitoringAuthUsername}
            onChange={(event) => setMonitoringAuthUsername(event.target.value)}
          />
          <TextField
            id="standard-basic text_2"
            type="password"
            key="text_2"
            label="Password"
            variant="standard"
            disabled={disabledMonitoringControl}
            value={monitoringAuthPassword}
            onChange={(event) => setMonitoringAuthPassword(event.target.value)}
          />
        </div>

        {/*<div className="ssl_wrapper">*/}
        {/*    <FormControlLabel*/}
        {/*        value="top"*/}
        {/*        control={<Checkbox checked={sslChecked} onChange={() => onCheckSSL()}/>}*/}
        {/*        label="SSL"*/}
        {/*    />*/}
        {/*    <div className="upload_file_container"><span>SSL Certificate</span>*/}
        {/*        <label htmlFor="file" className="upload_file" >*/}
        {/*            <input type="file" id="file" accept="text" disabled={!sslChecked}*/}
        {/*                   onChange={(event) => onSetFile(event?.target?.files![0])}/>*/}
        {/*            <span>Browse</span>*/}
        {/*        </label></div>*/}
        {/*</div>*/}
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
          <CircularProgress color="primary" />
        </div>
      )}

      <ToastContainer autoClose={3000} position="bottom-right" />
    </div>
  );
};
