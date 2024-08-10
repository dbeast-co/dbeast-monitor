import React, { useState } from 'react';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import {CheckboxConfig} from '../models/cluster';
import './AddNewClusterPanel.scss';
import './LogstashConfigurationsPanel.scss';

export const MonitoringClusterInjectionPanel = ({ monitoringClusterInjections }: any) => {
    const [_, setIsChecked] = useState(false);

    console.log(monitoringClusterInjections)

    const onChangeCheckbox = (index: number) => {
        let file = monitoringClusterInjections[index];
        file.is_checked = !file.is_checked;
        setIsChecked(file.is_checked);
    };

    return (
        <div className="monitoringClusterInjectionPanel">
            {monitoringClusterInjections &&
                monitoringClusterInjections.map((config: CheckboxConfig, index: number) => {
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
