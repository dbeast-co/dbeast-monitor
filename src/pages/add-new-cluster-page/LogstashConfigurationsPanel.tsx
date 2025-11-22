import React, { useState, useEffect } from 'react';
import { Checkbox, Field, useTheme2 } from '@grafana/ui';
import { CheckboxConfig } from './models/cluster';
import { getLogstashConfigurationsPanelStyles } from './LogstashConfigurationsPanel.styles';

export const LogstashConfigurationsPanel = ({ files }: any) => {
  const theme = useTheme2();
  const styles = getLogstashConfigurationsPanelStyles(theme);
  const [localFiles, setLocalFiles] = useState<CheckboxConfig[]>(files || []);

  // Update local state when files prop changes
  useEffect(() => {
    setLocalFiles(files || []);
  }, [files]);

  const onChangeCheckbox = (index: number) => {
    setLocalFiles((prevFiles) => {
      const newFiles = [...prevFiles];
      newFiles[index] = {
        ...newFiles[index],
        is_checked: !newFiles[index].is_checked,
      };
      return newFiles;
    });
  };

  return (
    <div className={styles.logstashConfigurationsPanel}>
      {localFiles &&
        localFiles.map((config: CheckboxConfig, index: number) => {
          return (
            <div key={config.id} className={styles.configItem}>
              <Field>
                <Checkbox
                  id={config.id}
                  name={config.id}
                  checked={config.is_checked}
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
