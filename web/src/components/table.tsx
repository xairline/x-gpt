import {ColumnsType, TableProps} from 'antd/es/table';
import {Card, Table, Tooltip} from 'antd';
import {ModelsFlightInfo} from '../../../../xa-ng/apps/xws/src/store/Api';
import {Link} from 'react-router-dom';
import {useObserver} from 'mobx-react-lite';
import {useStores} from "../stores";

/* eslint-disable-next-line */
export interface TableViewProps {
    // dataSet: TableDataSet;
    height: string;
}

interface DataType {
    key: React.Key;
    name: string;
    age: number;
    address: string;
}

export function TableView(props: TableViewProps) {
    const {FlightLogStore} = useStores();
    const columns: ColumnsType<any> = [
        {
            title: 'Date',
            dataIndex: 'date',
            fixed: 'left',
            width: '160px',
            render: (value, record: any) => {
                let formattedDate = value.slice(0, 16);
                formattedDate = formattedDate.replace('T', ' - ');
                return (
                    <Link to={`/flight-logs/${record.key}`}>
                        <Tooltip placement="topLeft" title={record.date}>
                            {formattedDate}

                        </Tooltip>
                    </Link>
                );
            },
        },
        {
            title: 'Departure',
            dataIndex: 'departure',
            filters: FlightLogStore.tableDataSet.data
                .filter((elem: any, index: any, self: any) => {
                    return (
                        index ===
                        self.findIndex(
                            (t: any) => t.departure.airportId === elem.departure.airportId
                        )
                    );
                })
                .sort((a: any, b: any) =>
                    a.departure.airportId < b.departure.airportId ? -1 : 1
                )
                .map((data: any) => {
                    return {
                        text: data.departure.airportId,
                        value: data.departure.airportId,
                    };
                }),
            filterSearch: true,
            onFilter: (value: any, record: any) =>
                record.departure.airportId.startsWith(value),
            width: '90px',
            render: (record: ModelsFlightInfo) => (
                <Tooltip placement="topLeft" title={record.airportName}>
                    {record.airportId}
                </Tooltip>
            ),
        },
        {
            title: 'Arrival',
            dataIndex: 'arrival',
            filters: FlightLogStore.tableDataSet.data
                .filter((elem: any, index: any, self: any) => {
                    return (
                        index ===
                        self.findIndex(
                            (t: any) =>
                                t.arrival.airportId === elem.arrival.airportId &&
                                t.arrival.airportId.length > 0
                        )
                    );
                })
                .sort((a: any, b: any) =>
                    a.arrival.airportId < b.arrival.airportId ? -1 : 1
                )
                .map((data: any) => {
                    return {
                        text: data.arrival.airportId,
                        value: data.arrival.airportId,
                    };
                }),
            filterSearch: true,
            onFilter: (value: any, record: any) =>
                record.arrival.airportId.startsWith(value),
            // fixed: 'left',
            width: '80px',
            render: (record: ModelsFlightInfo) => (
                <Tooltip placement="topLeft" title={record.airportName}>
                    {record.airportId}
                </Tooltip>
            ),
        },
        {
            title: 'Duration',
            dataIndex: 'duration',
            //onFilter: (value: string, record) => record.arrival.startsWith(value),
            // width: '20%',
            // fixed: 'left',
            width: '100px',
            render: (record: any) =>
                record == '-' ? (
                    '-'
                ) : (
                    <Tooltip placement="topLeft" title={'format: HH:mm'}>
                        {Math.floor(record / 3600) < 10 ? '0' : ''}
                        {Math.floor(record / 3600)}:
                        {Math.floor(record % 3600) / 60 < 10 ? '0' : ''}
                        {Math.floor((record % 3600) / 60)} h
                    </Tooltip>
                ),
            /*
            !flightStatus.departureFlightInfo?.time
                  ? '-'
                  : Math.floor(
                    (flightStatus.arrivalFlightInfo.time -
                      flightStatus.departureFlightInfo?.time || 0) / 3600
                  ) +
                  ':' +
                  Math.floor(
                    ((flightStatus.arrivalFlightInfo.time -
                        flightStatus.departureFlightInfo?.time || 0) %
                      3600) /
                    60
                  ) + " h"
            * */
        }
    ];

    const onChange: TableProps<DataType>['onChange'] = (
        pagination,
        filters,
        sorter,
        extra
    ) => {
        console.log('params', pagination, filters, sorter, extra);
    };
    return useObserver(() => (
        <Card title={'Flight Logs'} size={'small'} headStyle={{background: "#006363", color: "white"}}>
            <Table
                size={'small'}
                columns={columns}
                dataSource={FlightLogStore.tableDataSet.data}
                onChange={onChange}
                scroll={{y: props.height, x: '400px'}}
            />
        </Card>
    ));
}

export default TableView;
