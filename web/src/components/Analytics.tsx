import React, {useEffect, useState} from "react";
import {Col, Row} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";
import Landing from "./landing";
import TableView from "./table";
import Flights from "./flights";


interface Props {
}

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
                <Flights size={"small"}/>
            </Col>
            <Col span={24}>
                <TableView height={'100'}/>
            </Col>
            <Col span={24}>
                <Landing size={"small"}/>
            </Col>
            <Col span={24}>
            </Col>
        </Row>

    ))
        ;
};

export default Analytics;
