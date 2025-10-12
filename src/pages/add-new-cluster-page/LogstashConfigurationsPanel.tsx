import React, { useState } from 'react';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import { CheckboxConfig } from './models/cluster';
import { getLogstashConfigurationsPanelStyles } from './LogstashConfigurationsPanel.styles';
import { useTheme2 } from '@grafana/ui';

export const LogstashConfigurationsPanel = ({ files }: any) => {
  const theme = useTheme2();
  const styles = getLogstashConfigurationsPanelStyles(theme);
  const [_, setIsChecked] = useState(false);

  const onChangeCheckbox = (index: number) => {
    let file = files[index];
    file.is_checked = !file.is_checked;
    setIsChecked(file.is_checked);
  };

  return (
    <div className={styles.logstashConfigurationsPanel}>
      {files &&
        files.map((config: CheckboxConfig, index: number) => {
          return (
            <div key={config.id} className={styles.configItem}>
              <FormControlLabel
                value={config.label}
                control={
                  <Checkbox
                    id={config.id}
                    name={config.id}
                    checked={config.is_checked}
                    onChange={() => onChangeCheckbox(index)}
                  />
                }
                label={config.label}
              />
            </div>
          );
        })}
    </div>
  );
};
