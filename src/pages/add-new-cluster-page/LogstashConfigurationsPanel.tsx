import React, { useState } from 'react';
import { Checkbox, Field, useTheme2 } from '@grafana/ui';
import { CheckboxConfig } from './models/cluster';
import { getLogstashConfigurationsPanelStyles } from './LogstashConfigurationsPanel.styles';

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
              <Field>
                <Checkbox
                  id={config.id}
                  name={config.id}
                  value={config.is_checked}
                  label={config.label}
                  onChange={() => onChangeCheckbox(index)}
                />
              </Field>
            </div>
          );
        })}
    </div>
  );
};
