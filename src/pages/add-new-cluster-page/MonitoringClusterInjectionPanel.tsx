import React, { useState, useEffect } from 'react';
import { Checkbox, Field, useTheme2 } from '@grafana/ui';
import { CheckboxConfig } from './models/cluster';
import { getMonitoringClusterInjectionPanelStyles } from './MonitoringClusterInjectionPanel.styles';

export const MonitoringClusterInjectionPanel = ({ monitoringClusterInjections }: any) => {
    const theme = useTheme2();
    const styles = getMonitoringClusterInjectionPanelStyles(theme);
    const [localInjections, setLocalInjections] = useState<CheckboxConfig[]>(monitoringClusterInjections || []);

    // Update local state when monitoringClusterInjections prop changes
    useEffect(() => {
        setLocalInjections(monitoringClusterInjections || []);
    }, [monitoringClusterInjections]);

    const onChangeCheckbox = (index: number) => {
        setLocalInjections((prevInjections) => {
            const newInjections = [...prevInjections];
            newInjections[index] = {
                ...newInjections[index],
                is_checked: !newInjections[index].is_checked,
            };
            return newInjections;
        });
    };

    return (
        <div className={styles.monitoringClusterInjectionPanel}>
            {localInjections &&
                localInjections.map((config: CheckboxConfig, index: number) => {
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
