import React, {FC} from 'react';
import {DataSourceItem} from '../DataSourceItem/DataSourceItem';
import './data-source-list.scss';
import {useTheme2} from '@grafana/ui';
import classNames from "classnames";

/**
 * Properties
 */
interface Props {
    dataSources: any[];
}

export const DataSourceList: FC<Props> = ({dataSources}) => {
    const theme = useTheme2();
    return (

        <div className={classNames({
            container: true,
            isLight: theme.isLight
        })}>
            <header className="header">
                <h1>Clusters list</h1>
            </header>
            <section className="card-section card-list-layout-list">
                <ul className="card-list" data-col={dataSources.length}>
                    {dataSources.map((item, index) => {
                        return (
                            <li className="card-item-wrapper" key={index} aria-label="check-card">
                                <DataSourceItem dataSourceItem={item} theme={theme}/>
                            </li>
                        );
                    })}
                </ul>
            </section>
        </div>
    );
};
