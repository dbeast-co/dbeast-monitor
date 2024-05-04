import React, { useState } from 'react';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import { CheckboxConfig } from '../models/cluster';
import './AddNewClusterPanel.scss';
import './LogstashConfigurationsPanel.scss';
// interface LogstashConfigurationsPanelProps {
//     files: {
//         es_monitoring_configuration_files: CheckboxConfig[];
//         logstash_monitoring_configuration_files: { configurations: CheckboxConfig[]; hosts: Host[]; };
//     };
// }

export const LogstashConfigurationsPanel = ({ files }: any) => {
  const [_, setIsChecked] = useState(false);

  const onChangeCheckbox = (index: number) => {
    let file = files[index];
    file.is_checked = !file.is_checked;
    setIsChecked(file.is_checked);
  };

  return (
    <div className="logstashConfigurationsPanel">
      {files &&
        files.map((config: CheckboxConfig, index: number) => {
          return (
            <div key={config.id}>
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
