import React, {useEffect, useState} from "react";
import {Card, Col, Row, Statistic, Table, Tooltip} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";
import {Link} from "react-router-dom";
import {InfoCircleOutlined} from "@ant-design/icons";
import {ColumnsType} from 'antd/es/table';


interface Props {
}

const columns: ColumnsType<any> = [
    {
        title: 'Date',
        dataIndex: 'date',
        // filters: FlightLogStore.tableDataSet.filters['departure'] || null,
        // filterMode: 'menu',
        fixed: 'left',
        width: '120px',
        render: (record: any) => {
            let formattedDate = record.slice(0, 16);
            formattedDate = formattedDate.replace('T', ' - ');
            return (
                <Tooltip placement="topLeft" title={record.date}>
                    {formattedDate}
                </Tooltip>
            );
        },
    },
    {
        title: 'Departure',
        dataIndex: 'departure',
        // filters: FlightLogStore.tableDataSet.data
        //     .filter((elem: any, index: any, self: any) => {
        //         return (
        //             index ===
        //             self.findIndex(
        //                 (t: any) => t.departure.airportId === elem.departure.airportId
        //             )
        //         );
        //     })
        //     .sort((a: any, b: any) =>
        //         a.departure.airportId < b.departure.airportId ? -1 : 1
        //     )
        //     .map((data: any) => {
        //         return {
        //             text: data.departure.airportId,
        //             value: data.departure.airportId,
        //         };
        //     }),
        filterSearch: true,
        // onFilter: (value: string, record: any) =>
        //     record.departure.airportId.startsWith(value),
        width: '90px',
        // render: (record: ModelsFlightInfo) => (
        //     <Tooltip placement="topLeft" title={record.airportName}>
        //         {record.airportId}
        //     </Tooltip>
        // ),
    },
    {
        title: 'Arrival',
        dataIndex: 'arrival',
        // filters: FlightLogStore.tableDataSet.data
        //     .filter((elem: any, index: any, self: any) => {
        //         return (
        //             index ===
        //             self.findIndex(
        //                 (t: any) =>
        //                     t.arrival.airportId === elem.arrival.airportId &&
        //                     t.arrival.airportId.length > 0
        //             )
        //         );
        //     })
        //     .sort((a: any, b: any) =>
        //         a.arrival.airportId < b.arrival.airportId ? -1 : 1
        //     )
        //     .map((data: any) => {
        //         return {
        //             text: data.arrival.airportId,
        //             value: data.arrival.airportId,
        //         };
        //     }),
        filterSearch: true,
        // onFilter: (value: string, record: any) =>
        //     record.arrival.airportId.startsWith(value),
        // fixed: 'left',
        width: '80px',
        // render: (record: ModelsFlightInfo) => (
        //     <Tooltip placement="topLeft" title={record.airportName}>
        //         {record.airportId}
        //     </Tooltip>
        // ),
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
    },
    {
        title: 'Actions',
        key: 'operation',
        fixed: 'right',
        width: 80,
        render: (record: any) =>
            record.hasLocationData ? (
                <>
                    <Link to={`/flight-logs/${record.key}`}>Details</Link>
                </>
            ) : (
                <Tooltip
                    trigger={'click'}
                    title={
                        "This is an imported flight that we don't have enough data to show detailed report"
                    }
                >
                    Not Available <InfoCircleOutlined/>
                </Tooltip>
            ),
    },
];
const Analytics: React.FC<Props> = (props) => {
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const screens = useBreakpoint();
    const [userMetadata, setUserMetadata] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUserMetadata(userInfo);
            }
        })();

    }, [isAuthenticated, fetchUserInfo]);
    const config = {
        appendPadding: 10,
        data: [{type: "B738", value: 2}, {type: "A319", value: 2}],//FlightLogStore.AirplaneStats,
        angleField: 'value',
        colorField: 'type',
        radius: 0.9,
        autoFit: false,
        height: 100,
        // width: 100,
        label: {
            type: 'inner',
            offset: '-30%',
            content: ({percent}: any) => `${(percent * 100).toFixed(0)}%`,
            style: {
                fontSize: 12,
                textAlign: 'center',
            },
        },
        interactions: [
            {
                type: 'element-active',
            },
        ],
    };
    return ((
        <Row style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
        }} gutter={[12, 12]}>
            <Col flex={"auto"}>
                <Card>
                    <Statistic title="Active Users" value={112893}/>
                </Card>
            </Col>
            <Col flex={"auto"}>
                <Card>
                    <Statistic title="Active Users" value={112893}/>
                </Card>
            </Col>
            <Col span={24}>
                <Table
                    columns={columns}
                    scroll={{y: 400}}
                    style={{
                        width: "100%",
                        height: "200px",
                    }}
                    pagination={
                        {
                            pageSize: 10,

                        }
                    }
                    size={"small"}
                />
            </Col>
        </Row>

    ))
        ;
};

export default Analytics;
