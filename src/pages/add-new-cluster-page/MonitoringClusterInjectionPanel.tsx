import React, { useState } from 'react';
import { Checkbox, Field, useTheme2 } from '@grafana/ui';
import { CheckboxConfig } from './models/cluster';
import { getMonitoringClusterInjectionPanelStyles } from './MonitoringClusterInjectionPanel.styles';

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
