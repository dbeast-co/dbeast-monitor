import React from 'react';
import TextField from '@mui/material/TextField';
import './logstash.scss';

export interface Props {
    logstash: Logstash;
    onCancel: () => void;
    onSave: (logstash: Logstash) => void;
}

export interface Logstash {
    id: string;
    serverAddress: string;
    logstashApiHost: string;
    logstashLogsFolder: string;
}

const LogstashComponent = (state: Props) => {

    const [logstash, setLogstash] = React.useState<Logstash>({
        id: state.logstash.id,
        serverAddress: state.logstash.serverAddress,
        logstashApiHost: state.logstash.logstashApiHost,
        logstashLogsFolder: state.logstash.logstashLogsFolder,
    });
    const onInputServerAddress = (e: any) => {
        setLogstash({
            ...logstash,
            serverAddress: e.target.value,
        });
    };
    const onInputHost = (e: any) => {
        setLogstash({
            ...logstash,
            logstashApiHost: e.target.value,
        });
    };
    const onInputFolder = (e: any) => {
        setLogstash({
            ...logstash,
            logstashLogsFolder: e.target.value,
        });
    };

    function onCancel() {
        state.onCancel();
    }

    function onSave() {
        state.onSave(logstash);
    }

    // useEffect(() => {
    //   console.log(state)
    // }, [state]);

    return (
        <div className="source_panel">
            <TextField
                id="standard-basic"
                label="Server address"
                variant="standard"
                value={logstash.serverAddress ?? ''}
                onChange={onInputServerAddress}
            />
            <TextField
                id="standard-basic"
                label="Logstash Api Host"
                variant="standard"
                value={logstash.logstashApiHost ?? ''}
                onChange={onInputHost}
            />
            <TextField
                id="standard-basic"
                label="Logstash Logs Folder"
                variant="standard"
                value={logstash.logstashLogsFolder ?? ''}
                onChange={onInputFolder}
            />
            <div className="actions">
                <button className="cancel_btn save_btn" onClick={() => onCancel()}>
                    Cancel
                </button>
                <button
                    className="save_btn"
                    disabled={
                        !logstash.serverAddress.length || !logstash.logstashApiHost.length || !logstash.logstashLogsFolder.length
                    }
                    onClick={() => onSave()}
                    autoFocus
                >
                    Save
                </button>
            </div>
        </div>
    );
};

export default LogstashComponent;
