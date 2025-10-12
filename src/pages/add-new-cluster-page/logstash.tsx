import React from 'react';
import TextField from '@mui/material/TextField';
import { getLogstashStyles } from './logstash.styles';
import { useTheme2 } from '@grafana/ui';

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
  const theme = useTheme2();
  const styles = getLogstashStyles(theme);

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

  return (
    <div className={styles.sourcePanel}>
      <div className={styles.textField}>
        <TextField
          id="standard-basic-1"
          label="Server address"
          variant="standard"
          value={logstash.serverAddress ?? ''}
          onChange={onInputServerAddress}
        />
      </div>
      <div className={styles.textField}>
        <TextField
          id="standard-basic-2"
          label="Logstash Api Host"
          variant="standard"
          value={logstash.logstashApiHost ?? ''}
          onChange={onInputHost}
        />
      </div>
      <div className={styles.textField}>
        <TextField
          id="standard-basic-3"
          label="Logstash Logs Folder"
          variant="standard"
          value={logstash.logstashLogsFolder ?? ''}
          onChange={onInputFolder}
        />
      </div>
      <div className={styles.actions}>
        <button className="cancel-btn" onClick={() => onCancel()}>
          Cancel
        </button>
        <button
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
