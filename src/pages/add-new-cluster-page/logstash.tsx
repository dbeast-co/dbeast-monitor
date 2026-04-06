import React from 'react';
import { Input, Field, useTheme2 } from '@grafana/ui';
import { getLogstashStyles } from './logstash.styles';

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
        <Field label="Server address">
          <Input
            id="standard-basic-1"
            value={logstash.serverAddress ?? ''}
            onChange={onInputServerAddress}
          />
        </Field>
      </div>
      <div className={styles.textField}>
        <Field label="Logstash Api Host">
          <Input
            id="standard-basic-2"
            value={logstash.logstashApiHost ?? ''}
            onChange={onInputHost}
          />
        </Field>
      </div>
      <div className={styles.textField}>
        <Field label="Logstash Logs Folder">
          <Input
            id="standard-basic-3"
            value={logstash.logstashLogsFolder ?? ''}
            onChange={onInputFolder}
          />
        </Field>
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
