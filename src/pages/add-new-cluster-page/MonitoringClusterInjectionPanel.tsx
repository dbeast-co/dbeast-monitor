import React, { useState } from 'react';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import { CheckboxConfig } from './models/cluster';
import { getMonitoringClusterInjectionPanelStyles } from './MonitoringClusterInjectionPanel.styles';
import { useTheme2 } from '@grafana/ui';

export const MonitoringClusterInjectionPanel = ({ monitoringClusterInjections }: any) => {
    const theme = useTheme2();
    const styles = getMonitoringClusterInjectionPanelStyles(theme);
    const [_, setIsChecked] = useState(false);

    const onChangeCheckbox = (index: number) => {
        let file = monitoringClusterInjections[index];
        file.is_checked = !file.is_checked;
        setIsChecked(file.is_checked);
    };

    return (
        <div className={styles.monitoringClusterInjectionPanel}>
            {monitoringClusterInjections &&
                monitoringClusterInjections.map((config: CheckboxConfig, index: number) => {
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
